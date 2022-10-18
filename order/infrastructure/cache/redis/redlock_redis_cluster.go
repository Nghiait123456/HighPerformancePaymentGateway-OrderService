package redis

import (
	redlock "github.com/Nghiait123456/redlock"
	"github.com/Nghiait123456/redlock/redis/goredis/v8"
	"github.com/go-redis/redis/v8"
)

type (
	RedLockCluster struct {
		cc      *redis.ClusterClient
		RedLock *redlock.Redsync
	}

	Option = redlock.Option

	RedLockClusterInterface interface {
		NewMutex(name string, options ...Option) *redlock.Mutex
	}
)

func (r *RedLockCluster) newRedLock() *redlock.Redsync {
	pool := goredis.NewPool(r.cc)
	return redlock.New(pool)
}

func (r RedLockCluster) NewMutex(name string, options ...Option) *redlock.Mutex {
	return r.RedLock.NewMutex(name, options...)
}

func NewRedLockCluster(cc *redis.ClusterClient) RedLockClusterInterface {
	//var _ RedLockClusterInterface = (*RedLockCluster)(nil)
	rl := RedLockCluster{
		cc: cc,
	}

	rl.RedLock = rl.newRedLock()
	return &rl
}
