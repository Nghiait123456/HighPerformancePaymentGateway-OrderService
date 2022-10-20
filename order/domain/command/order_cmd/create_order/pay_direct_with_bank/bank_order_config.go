package pay_direct_with_bank

import (
	"context"
	"errors"
	"fmt"
	"github.com/high-performance-payment-gateway/order-service/order/domain/command/order_cmd/create_order/errors_create_order"
	"github.com/high-performance-payment-gateway/order-service/order/infrastructure/cache/local_in_memory"
	"github.com/high-performance-payment-gateway/order-service/order/infrastructure/cache/redis"
	log "github.com/sirupsen/logrus"
	"time"
)

type (
	BankOrderCf struct {
		rc          redis.RedisClusterInterface
		globalValue local_in_memory.LocalInMemoryInterface
	}

	BankOrderCfInterface interface {
		GetRootDomainForBank(bankCode string) (string, error)
		LoadAllCfAndSave() error
	}

	OneDomainBank struct {
		LinkRootDomain string
	}

	ALlDomainBank struct {
		AllData map[string]OneDomainBank // bankCode ==> OneDomainBank
	}
)

const (
	KEY_ALL_DOMAIN_BANK = "pay_with_bank_DOMAIN_BANK_LIST"
)

func (b BankOrderCf) GetRootDomainForBank(bankCode string) (string, error) {
	allCf, errLCf := b.LoadAllCf()

	if errLCf == nil {
		if val, ok := allCf.AllData[bankCode]; ok {
			return val.LinkRootDomain, nil
		}

		return "", errors_create_order.NewErrorBankCodeNotExits()
	}

	//try load cf from remote
	errLA := b.LoadAllCfAndSave()
	if errLA != nil {
		panicM := fmt.Sprintf("LoadAllCf BankDomain And Save error: %v", errLA.Error())
		log.Error(panicM)
		panic(panicM)
	}

	//try get again
	allCfTry, errTryAgain := b.LoadAllCf()
	if errTryAgain != nil {
		panicM := fmt.Sprintf("LoadAllCf BankDomain And Save error: %v", errTryAgain.Error())
		log.Error(panicM)
		panic(panicM)
	}

	if val, ok := allCfTry.AllData[bankCode]; ok {
		return val.LinkRootDomain, nil
	}

	return "", errors_create_order.NewErrorBankCodeNotExits()

}

func (b BankOrderCf) timeExpiry() time.Duration {
	return time.Hour * 24 * 10
}

func (b BankOrderCf) LoadAllCfAndSave() error {
	var allDomainBank ALlDomainBank
	errGetFrCache := b.rc.GetKeyOfStructValue(context.Background(), KEY_ALL_DOMAIN_BANK, allDomainBank)
	if errGetFrCache != nil {
		errM := fmt.Sprintf("get allDomainBank Fr Cache error: %v", errGetFrCache.Error())
		log.Error(errM)
		panic(errM)
	}

	b.globalValue.Set(KEY_ALL_DOMAIN_BANK, allDomainBank, b.timeExpiry())
	return nil
}

func (b BankOrderCf) LoadAllCf() (ALlDomainBank, error) {
	all, stt := b.globalValue.Get(KEY_ALL_DOMAIN_BANK)
	if stt == false || all == nil {
		return ALlDomainBank{}, errors.New("missing dataCf Bank Domain fr key")
	}

	return all.(ALlDomainBank), nil
}

func NewBankOrderCf(rc redis.RedisClusterInterface, globalValue local_in_memory.LocalInMemoryInterface) BankOrderCfInterface {
	return &BankOrderCf{
		rc:          rc,
		globalValue: globalValue,
	}
}
