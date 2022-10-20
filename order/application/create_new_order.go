package application

import (
	"github.com/high-performance-payment-gateway/order-service/order/domain/command/order_cmd/create_order"
	"github.com/high-performance-payment-gateway/order-service/order/domain/command/order_cmd/create_order/global_param_create_order"
	"github.com/high-performance-payment-gateway/order-service/order/domain/command/order_cmd/create_order/response_create_order"
)

type (
	CreateOrder struct {
		PrCreateOrder global_param_create_order.ParamCreateOrder
		PrGetInstance create_order.ParamGetInstanceOrder
	}

	CreateOrderInterface interface {
		CreateNewOrder() response_create_order.ResponseCreateOrder
	}
)

func (c *CreateOrder) CreateNewOrder() response_create_order.ResponseCreateOrder {
	OrderInstance, err := create_order.OrderFactory(c.PrGetInstance)
	if err != nil {
		panic(err)
	}

	return OrderInstance.CreateOrder(c.PrCreateOrder)
}

func NewCreateOrder(PrCreateOrder global_param_create_order.ParamCreateOrder, PrGetInstance create_order.ParamGetInstanceOrder) CreateOrderInterface {
	return &CreateOrder{
		PrCreateOrder: PrCreateOrder,
		PrGetInstance: PrGetInstance,
	}
}
