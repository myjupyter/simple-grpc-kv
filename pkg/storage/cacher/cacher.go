package cacher

import (
	"context"
	"fmt"
	"time"

	fast "github.com/VictoriaMetrics/fastcache"
)

const MAX_SIZE = 1 << 25

type Options interface {
	Path() string
}

type Cache struct {
	db   *fast.Cache
	opts Options
}

func New(opts Options) *Cache {
	return &Cache{
		db:   fast.New(MAX_SIZE),
		opts: opts,
	}
}

func (c *Cache) Get(key string) (string, error) {
	value, has := c.db.HasGet(nil, []byte(key))
	if !has {
		return "", fmt.Errorf("no key '%s'", key)
	}
	return string(value), nil
}

func (c *Cache) Set(key, value string) error {
	c.db.Set([]byte(key), []byte(value))
	return nil
}

func (c *Cache) Update(key, value string) error {
	return c.Set(key, value)
}

func (c *Cache) Delete(key string) error {
	_, has := c.db.HasGet(nil, []byte(key))
	if !has {
		return fmt.Errorf("no key to delete '%s'", key)
	}
	c.db.Del([]byte(key))
	return nil
}

func (c *Cache) Upload(_ context.Context) error {
	db, err := fast.LoadFromFile(c.opts.Path())
	if err != nil {
		return err
	}
	c.db = db
	return nil
}

func (c *Cache) Save() error {
	return c.db.SaveToFile(c.opts.Path())
}

func (c *Cache) SaveEvery(ctx context.Context, t time.Duration) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(t):
			if err := c.Save(); err != nil {
				return err
			}
		}
	}
}
