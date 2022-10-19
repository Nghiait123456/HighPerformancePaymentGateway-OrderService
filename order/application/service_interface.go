package application

import "github.com/high-performance-payment-gateway/order-service/order/application/query_request_balance"

type (
	ServiceInterface interface {
		Init()
		GetOneOrderInfo(pq query_request_balance.ParamQueryOneBalance) query_request_balance.RequestBalanceResponse
	}
)
