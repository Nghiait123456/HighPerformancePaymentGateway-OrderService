package update_cache_order_when_invalidate

import (
	"errors"
	"github.com/high-performance-payment-gateway/order-service/order/infrastructure/db/orm"
	log "github.com/sirupsen/logrus"
)

type (
	DataFrDB struct {
		Pr ParamQuery
	}

	DataQuery = orm.BalanceRequestLog

	DataFrDBInterface interface {
		GetFrMysql() (DataQuery, error)
		GetFrCassandra() (DataQuery, error)
		ProcessGet() DataResposne
	}
)

func (d *DataFrDB) GetFrMysql() (DataQuery, error) {
	//todo update get
	return DataQuery{}, nil
}

func (d *DataFrDB) GetFrCassandra() (DataQuery, error) {
	//todo update get
	return DataQuery{}, nil
}

/**
  All data save in casssandra, only data not done just is saved in mysql.
  One Simple smart traffic : we get fr cassadra before, affter get in mysql
*/
func (d *DataFrDB) ProcessGet() DataResposne {
	dataCassandra, errCassandra := d.GetFrCassandra()
	if errCassandra != nil {
		log.WithFields(log.Fields{
			"errM": errCassandra.Error(),
		}).Error("get Fr Cassandra error")
	}

	// exist data
	if dataCassandra.OrderId != 0 {
		return DataResposne{
			Data:           dataCassandra,
			Err:            nil,
			IsOrderIdExist: true,
		}
	}

	//get fr Mysql
	dataMysql, errMysql := d.GetFrMysql()
	if errMysql != nil {
		log.WithFields(log.Fields{
			"errM": errMysql.Error(),
		}).Error("get Fr Mysql error")
	}

	// exist data
	if dataCassandra.OrderId != 0 {
		return DataResposne{
			Data:           dataMysql,
			Err:            nil,
			IsOrderIdExist: true,
		}
	}

	return DataResposne{
		Data:           dataMysql,
		Err:            errors.New("orderId dont exist"),
		IsOrderIdExist: false,
	}
}

func NewDataFrDB(pr ParamQuery) DataFrDBInterface {
	return &DataFrDB{
		pr,
	}
}
