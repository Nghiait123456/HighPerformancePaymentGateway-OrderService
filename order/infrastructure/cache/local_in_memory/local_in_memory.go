package local_in_memory

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type (
	LocalInMemory struct {
		cache *cache.Cache
	}

	LocalInMemoryInterface interface {
		Set(k string, x any, d time.Duration)
		Get(k string) (any, bool)
		Delete(k string) error
		// Reset resets the storage and delete all keys.
		Reset() error
		Init() error
	}

	Config struct {
		defaultExpiration,
		cleanupInterval time.Duration
	}
)

func (l *LocalInMemory) Set(k string, x any, d time.Duration) {
	l.cache.Set(k, x, d)
}

func (l *LocalInMemory) Get(k string) (any, bool) {
	return l.cache.Get(k)
}

func (l *LocalInMemory) Delete(k string) error {
	l.cache.Delete(k)
	return nil
}

// Reset resets the storage and delete all keys.
func (l *LocalInMemory) Reset() error {
	l.cache.Flush()
	return nil
}

func (l *LocalInMemory) Init() error {
	return nil
}

func NewLocalInMemory(c Config) LocalInMemoryInterface {
	l := LocalInMemory{}
	l.cache = cache.New(c.defaultExpiration, c.cleanupInterval)
	l.Init()

	return &l
}
