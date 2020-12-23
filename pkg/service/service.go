package service

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/myjupyter/simple-grpc-kv/grpc/kvapi"
	"github.com/myjupyter/simple-grpc-kv/pkg/storage"
)

var (
	null          = &empty.Empty{}
	ErrInvalidReq = fmt.Errorf("invalid request")
)

type KVService struct {
	Storage storage.Storage
	kvapi.UnimplementedKVStorageServer
}

func NewKVService(st storage.Storage) *KVService {
	return &KVService{
		Storage: st,
	}
}

func (kv *KVService) Get(_ context.Context, req *kvapi.KeyRequest) (*kvapi.ValueResponse, error) {
	if req == nil {
		return nil, ErrInvalidReq
	}

	value, err := kv.Storage.Get(req.Key)
	if err != nil {
		return nil, err
	}

	return &kvapi.ValueResponse{
		Value: value,
	}, nil
}

func (kv *KVService) Set(_ context.Context, req *kvapi.KeyValueRequest) (*empty.Empty, error) {
	if req == nil {
		return null, ErrInvalidReq
	}

	if err := kv.Storage.Set(req.Key, req.Value); err != nil {
		return null, err
	}

	return null, nil
}

func (kv *KVService) Update(ctx context.Context, req *kvapi.KeyValueRequest) (*empty.Empty, error) {
	if req == nil {
		return null, ErrInvalidReq
	}

	if err := kv.Storage.Set(req.Key, req.Value); err != nil {
		return null, err
	}

	return null, nil
}

func (kv *KVService) Delete(ctx context.Context, req *kvapi.KeyRequest) (*empty.Empty, error) {
	if req == nil {
		return null, ErrInvalidReq
	}

	if err := kv.Storage.Delete(req.Key); err != nil {
		return null, err
	}

	return null, nil
}
