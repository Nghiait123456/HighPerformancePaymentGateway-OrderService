package update_cache_order_when_invalidate

import (
	"context"
	"fmt"
	redlock "github.com/Nghiait123456/redlock"
	"github.com/high-performance-payment-gateway/order-service/order/infrastructure/cache/redis"
	"github.com/high-performance-payment-gateway/order-service/order/infrastructure/db/orm"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type (
	UpdateCache struct {
		pr           ParamQuery
		clusterLock  redis.RedLockClusterInterface
		redisCluster redis.RedisClusterInterface
	}

	ParamQuery struct {
		OrderId uint64
	}

	DataResposne struct {
		Err            error
		IsOrderIdExist bool
		Data           orm.Order
	}

	DataSavedCache struct {
		Data       DataQuery
		IsHaveLock bool
	}

	UpdateCacheInterface interface {
		createNewMutexAndLock() (*redlock.Mutex, error)
		saveRecordLock() error
		refreshLock(r *redlock.Mutex) (bool, error)
		getDataFrDB() DataResposne
		updateCache(cache DataSavedCache) error
		HandleRequestUpdateCacheFrDB() DataResposne
	}
)

func (u UpdateCache) createNewMutexAndLock() (*redlock.Mutex, error) {
	k := strconv.FormatUint(u.pr.OrderId, 10)
	customExpiry := redlock.OptionFunc(func(mutex *redlock.Mutex) {
		mutex.SetExpiry(8 * time.Second)
	})

	customTries := redlock.OptionFunc(func(mutex *redlock.Mutex) {
		mutex.SetTries(3)
	})

	customDelayFc := redlock.OptionFunc(func(mutex *redlock.Mutex) {
		mutex.SetDelayFunc(func(tries int) time.Duration {
			return 200 * time.Microsecond
		})
	})

	mutex := u.clusterLock.NewMutex(k, customExpiry, customTries, customDelayFc)
	if err := mutex.Lock(); err != nil {
		return &redlock.Mutex{}, err
	}

	return mutex, nil
}

func (u *UpdateCache) refreshLock(r *redlock.Mutex) (bool, error) {
	return r.Unlock()
}

func (u *UpdateCache) saveRecordLock() error {
	ctx := context.Background()
	k := strconv.FormatUint(u.pr.OrderId, 10)
	recordLock := CreateRecordLock()
	exp := time.Second * 9000
	return u.redisCluster.SetKeyWithStructValue(ctx, k, recordLock, exp)
}

func (u UpdateCache) getDataFrDB() DataResposne {
	qr := NewDataFrDB(u.pr)
	return qr.ProcessGet()
}
func (u UpdateCache) getExpira() time.Duration {
	return time.Hour * 24 * 30
}

func (u *UpdateCache) updateCache(cache DataSavedCache) error {
	k := strconv.FormatUint(u.pr.OrderId, 10)
	err := u.redisCluster.SetKeyWithStructValue(context.Background(), k, cache, u.getExpira())
	if err != nil {
		log.WithFields(log.Fields{
			"errM": err.Error(),
		}).Error("save orderId data to cache error")

		return err
	}

	return nil
}

func (u *UpdateCache) HandleRequestUpdateCacheFrDB() DataResposne {
	lock, errCM := u.createNewMutexAndLock()
	if errCM != nil {
		return DataResposne{
			Data:           DataQuery{},
			IsOrderIdExist: false,
			Err:            errCM,
		}
	}

	errSaveRocordLock := u.saveRecordLock()
	if errSaveRocordLock != nil {
		log.Errorf("save record lock in pc update cache requets balance is error, errorM : %v", errSaveRocordLock.Error())
		return DataResposne{
			Data: DataQuery{},
			Err:  errSaveRocordLock,
		}
	}

	dataFrDB := u.getDataFrDB()
	//update cache and rewrite IsHaveLock
	if dataFrDB.IsOrderIdExist == true && dataFrDB.Err != nil {
		dataSCache := DataSavedCache{
			Data:       dataFrDB.Data,
			IsHaveLock: false,
		}

		errUC := u.updateCache(dataSCache)
		if errUC != nil {
			panicM := fmt.Sprintf("update cachae of request balance is  err: %v", errUC.Error())
			log.Error(panicM)
			panic(panicM)
		}
	} else {
		dataSCache := DataSavedCache{
			Data:       DataQuery{},
			IsHaveLock: false,
		}

		errUC := u.updateCache(dataSCache)
		if errUC != nil {
			panicM := fmt.Sprintf("update cachae of request balance is  err: %v", errUC.Error())
			log.Error(panicM)
			panic(panicM)
		}
	}

	//free mutex
	status, errRL := u.refreshLock(lock)
	if status != true && errRL != nil {
		panicM := fmt.Sprintf("dont refresh Lock, err: %v", errRL.Error())
		log.Error(panicM)
		panic(panicM)
	}

	return dataFrDB
}

func NewCacheUpdate(pr ParamQuery, clusterLock redis.RedLockClusterInterface, redisCluster redis.RedisClusterInterface) UpdateCacheInterface {
	return &UpdateCache{
		pr:           pr,
		clusterLock:  clusterLock,
		redisCluster: redisCluster,
	}
}
