package application

import (
	"github.com/google/wire"
	"github.com/high-performance-payment-gateway/order-service/order/domain/query/order_query"
)

var ProviderAllPartnerBalanceQuery = wire.NewSet(
	NewAllPartnerBalanceQuery,
	order_query.NewOneRequest,
	wire.Bind(new(AllPartnerBalanceQueryInterface), new(*AllPartnerBalanceQuery)),
)

var ProviderService = wire.NewSet(
	NewService, ProviderAllPartnerBalanceQuery, wire.Bind(new(ServiceInterface), new(*Service)),
)
