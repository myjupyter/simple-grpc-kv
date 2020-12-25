package service

import (
	"context"
	"fmt"
	"sync"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/myjupyter/simple-grpc-kv/grpc/kvapi"
	log "github.com/myjupyter/simple-grpc-kv/pkg/logger"
	"github.com/myjupyter/simple-grpc-kv/pkg/storage"
)

var (
	null          = &empty.Empty{}
	ErrInvalidReq = fmt.Errorf("invalid request")

	getFormat    = "gRPC Get: %s"
	setFormat    = "gRPC Set: %s"
	updateFormat = "gRPC Update: %s"
	deleteFormat = "gRPC Delete: %s"
)

type KVService struct {
	Storage storage.Storage
	Logger  log.Logger

	mu *sync.RWMutex

	kvapi.UnimplementedKVStorageServer
}

func NewKVService(st storage.Storage, logger log.Logger) *KVService {
	return &KVService{
		Storage: st,
		Logger:  logger,
		mu:      &sync.RWMutex{},
	}
}

func (kv *KVService) Get(_ context.Context, req *kvapi.KeyRequest) (*kvapi.ValueResponse, error) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()

	if req == nil {
		kv.Logger.Errorf(getFormat, ErrInvalidReq)
		return nil, ErrInvalidReq
	}

	value, err := kv.Storage.Get(req.Key)
	if err != nil {
		kv.Logger.Warningf(getFormat, err)
		return nil, err
	} else {
		kv.Logger.Infof(getFormat, fmt.Sprintf("request to get key '%s'", req.Key))
	}

	return &kvapi.ValueResponse{
		Value: value,
	}, nil
}

func (kv *KVService) Set(_ context.Context, req *kvapi.KeyValueRequest) (*empty.Empty, error) {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	if req == nil {
		kv.Logger.Errorf(setFormat, ErrInvalidReq)
		return null, ErrInvalidReq
	}

	if err := kv.Storage.Set(req.Key, req.Value); err != nil {
		kv.Logger.Warningf(setFormat, err)
		return null, err
	} else {
		kv.Logger.Infof(setFormat, fmt.Sprintf("request to set key '%s'", req.Key))
	}

	return null, nil
}

func (kv *KVService) Update(ctx context.Context, req *kvapi.KeyValueRequest) (*empty.Empty, error) {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	if req == nil {
		kv.Logger.Errorf(updateFormat, ErrInvalidReq)
		return null, ErrInvalidReq
	}

	if err := kv.Storage.Set(req.Key, req.Value); err != nil {
		kv.Logger.Warningf(updateFormat, err)
		return null, err
	} else {
		kv.Logger.Infof(updateFormat, fmt.Sprintf("request to update key '%s'", req.Key))
	}

	return null, nil
}

func (kv *KVService) Delete(ctx context.Context, req *kvapi.KeyRequest) (*empty.Empty, error) {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	if req == nil {
		kv.Logger.Errorf(deleteFormat, ErrInvalidReq)
		return null, ErrInvalidReq
	}

	if err := kv.Storage.Delete(req.Key); err != nil {
		kv.Logger.Warningf(updateFormat, err)
		return null, err
	} else {
		kv.Logger.Infof(deleteFormat, fmt.Sprintf("request to delete key '%s'", req.Key))
	}

	return null, nil
}
