package storage

import (
	"context"
	"time"
)

type Storage interface {
	Get(string) (string, error)
	Set(string, string) error
	Update(string, string) error
	Delete(string) error
	Upload(context.Context) error
	Save() error
	SaveEvery(context.Context, time.Duration) error
}
