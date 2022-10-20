package application

import (
	"github.com/high-performance-payment-gateway/order-service/order/application/query_request_balance"
	"github.com/high-performance-payment-gateway/order-service/order/domain/command/order_cmd/create_order"
	"github.com/high-performance-payment-gateway/order-service/order/domain/command/order_cmd/create_order/global_param_create_order"
	"github.com/high-performance-payment-gateway/order-service/order/domain/command/order_cmd/create_order/response_create_order"
)

type (
	ServiceInterface interface {
		Init()
		GetOneOrderInfo(pq query_request_balance.ParamQueryOneBalance) query_request_balance.RequestBalanceResponse
		CreateNewOrder(PrCreateOrder global_param_create_order.ParamCreateOrder, PrGetInstance create_order.ParamGetInstanceOrder) response_create_order.ResponseCreateOrder
	}
)
