//go:build wireinject
// +build wireinject

package balance

import (
	"github.com/google/wire"
	"github.com/high-performance-payment-gateway/order-service/order/application"
)

func ProviderService() application.ServiceInterface {
	wire.Build(
		application.ProviderService,
	)

	return &application.Service{}
}
