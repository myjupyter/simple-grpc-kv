package app

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	"github.com/myjupyter/simple-grpc-kv/grpc/kvapi"
	log "github.com/myjupyter/simple-grpc-kv/pkg/logger"
	"github.com/myjupyter/simple-grpc-kv/pkg/service"
	kv "github.com/myjupyter/simple-grpc-kv/pkg/storage"
	"github.com/myjupyter/simple-grpc-kv/pkg/storage/cacher"
)

var ErrAppInterrupted = fmt.Errorf("interrupt signal")

type Options interface {
	SavePath() string
	SaveTime() string
	Host() string
	Port() string
}

type Application struct {
	opts Options
	log  log.Logger

	st        kv.Storage
	kvService kvapi.KVStorageServer
}

func NewApplication(opts Options, logger log.Logger) *Application {
	return &Application{
		opts: opts,
		log:  logger,
	}
}

func (app *Application) Run() error {

	ctx := context.Background()

	address := app.opts.Host() + ":" + app.opts.Port()
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer lis.Close()

	errChanService := handleSignal()

	errChanStorage := make(chan error)

	st := cacher.NewCache(app.opts)

	err = st.Upload(ctx)
	if err != nil {
		app.log.Warningf("Cache upload: %s", err)
	}

	go func(ctx context.Context, s kv.Storage) {
		newCtx, cancel := context.WithCancel(ctx)
		select {
		case errChanStorage <- s.Save(newCtx):
			cancel()
		case <-ctx.Done():
			cancel()
		}
	}(ctx, st)

	go func(lis net.Listener, logger log.Logger) {
		grpcServer := grpc.NewServer()
		kvapi.RegisterKVStorageServer(grpcServer, service.NewKVService(st, logger))
		errChanService <- grpcServer.Serve(lis)
	}(lis, app.log)

	app.log.Infof(
		"Sever successfully started on %s:%s",
		app.opts.Host(),
		app.opts.Port(),
	)

	select {
	case err := <-errChanService:
		return err
	case err = <-errChanStorage:
		return err
	}
}

func (app *Application) Stop() error {
	return nil
}

func handleSignal() chan error {
	errChanService := make(chan error)
	go func() {
		signalChan := make(chan os.Signal)
		signal.Notify(signalChan, os.Interrupt)
		<-signalChan
		errChanService <- ErrAppInterrupted
	}()

	return errChanService
}
