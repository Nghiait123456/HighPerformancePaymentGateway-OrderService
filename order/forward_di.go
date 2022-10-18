package balance

import "github.com/high-performance-payment-gateway/order-service/order/application"

/**
forward to wire_gen
*/

func ForwardProviderService() application.ServiceInterface {
	return ProviderService()
}
