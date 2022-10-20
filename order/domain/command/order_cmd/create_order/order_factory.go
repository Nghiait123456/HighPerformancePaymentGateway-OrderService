package create_order

import (
	"fmt"
	"github.com/high-performance-payment-gateway/order-service/order/domain/command/order_cmd/create_order/base_interface"
	"github.com/high-performance-payment-gateway/order-service/order/domain/command/order_cmd/create_order/pay_direct_with_bank"
	"github.com/high-performance-payment-gateway/order-service/order/infrastructure/cache/redis"
)

type (

	// PaymentType + PaymentCode ==> detect type payment
	ParamGetInstanceOrder struct {
		PaymentType string
		PaymentCode int
	}
)

const (
	PAYMENT_DIRECT_WITH_BANK = "PAYMENT_WITH_BANK"
	PAYMENT_WITH_CREDIT_CARD = "PAYMENT_WITH_CREDIT_CARD"
)

func OrderFactory(p ParamGetInstanceOrder) (base_interface.OrderInterface, error) {
	switch p.PaymentType {
	case PAYMENT_DIRECT_WITH_BANK:
		switch p.PaymentCode {
		case 1:
			//todo get for global value and pass to
			return pay_direct_with_bank.NewPayDirectWithBank(pay_direct_with_bank.BankOrderCf{}, &redis.RedisCluster{}), nil
		}
	}

	return nil, fmt.Errorf("payment type and payment code dont exist")
}
