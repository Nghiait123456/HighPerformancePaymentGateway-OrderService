package base_interface

import (
	"github.com/high-performance-payment-gateway/order-service/order/domain/command/order_cmd/create_order/global_param_create_order"
	"github.com/high-performance-payment-gateway/order-service/order/domain/command/order_cmd/create_order/response_create_order"
)

type (
	OrderInterface interface {
		PaymentType() string
		CreateOrder(prCreateOrder global_param_create_order.ParamCreateOrder) response_create_order.ResponseCreateOrder
	}
)
