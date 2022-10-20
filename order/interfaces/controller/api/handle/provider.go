package handle

import (
	"github.com/google/wire"
	"github.com/high-performance-payment-gateway/order-service/order/application"
)

var ProviderRequestBalance = wire.NewSet(
	NewRequestOrderQuery,
	application.ProviderService,
)
