package balance

import (
	"github.com/google/wire"
	"github.com/high-performance-payment-gateway/order-service/order/interfaces/controller/api/handle"
)

func NewViewController() any {
	return nil
}

var ProviderRouters = wire.NewSet(
	NewRouter,
	NewViewController,
	handle.ProviderRequestBalance,
)
