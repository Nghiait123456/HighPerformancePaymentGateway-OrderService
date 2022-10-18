package order_query

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/high-performance-payment-gateway/order-service/order/infrastructure/cache/redis"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type (
	QueryCache struct {
		// todo get cluster from global cf and pass
		rc redis.RedisClusterInterface
	}

	ParamQueryCache    = ParamQuery
	ResponseQueryCache struct {
		Data                      DataQuery
		IsUseDataForResponse      bool
		IsContinueUpdateCacheFrDB bool
		Err                       error
	}

	DataSavedCache struct {
		Data       DataQuery
		IsHaveLock bool
	}

	QueryCacheInterface interface {
		QueryRequestBalance(p ParamQueryCache) ResponseQueryCache
	}
)

func (q *QueryCache) QueryRequestBalance(p ParamQueryCache) ResponseQueryCache {
	keyCache, errK := q.KeyCache(p)
	if errK != nil {
		errM := fmt.Sprintf("ParamQueryCache dont has cacheket valid, Err: %v \n", errK.Error())
		log.Errorln(errM)
		panic(errM)
	}
	var RqBalanceObj DataSavedCache

	rawData, errQC := q.rc.Get(context.Background(), keyCache)
	if q.rc.IsEmptyData(errQC) {
		return ResponseQueryCache{
			Data:                      DataQuery{},
			IsUseDataForResponse:      false,
			IsContinueUpdateCacheFrDB: true,
			Err:                       nil,
		}
	}

	if errQC != nil {
		panic(errQC.Error())
	}

	errU := json.Unmarshal([]byte(rawData), &RqBalanceObj)
	if errU != nil {
		return ResponseQueryCache{
			Data:                      DataQuery{},
			IsUseDataForResponse:      false,
			IsContinueUpdateCacheFrDB: false,
			Err:                       errU,
		}
	}

	if RqBalanceObj.IsHaveLock == false {
		return ResponseQueryCache{
			Data:                      RqBalanceObj.Data,
			IsUseDataForResponse:      true,
			IsContinueUpdateCacheFrDB: false,
			Err:                       nil,
		}
	}

	return ResponseQueryCache{
		Data:                      DataQuery{},
		IsUseDataForResponse:      false,
		IsContinueUpdateCacheFrDB: false,
		Err:                       nil,
	}

}

func (q QueryCache) KeyCache(p ParamQueryCache) (string, error) {
	if p.OrderId == 0 {
		return "", errors.New("empty orderId")
	}

	return strconv.FormatUint(p.OrderId, 10), nil
}

func NewQueryCache() QueryCacheInterface {
	//todo get query  redis.RedisClusterInterface from global cf and map to object
	return &QueryCache{}
}
