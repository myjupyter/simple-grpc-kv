package storage

import "context"

type Storage interface {
	Get(string) (string, error)
	Set(string, string) error
	Update(string, string) error
	Delete(string) error
	Upload(context.Context) error
	Save(context.Context) error
}
