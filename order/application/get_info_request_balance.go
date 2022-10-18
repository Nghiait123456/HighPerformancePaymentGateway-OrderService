package application

import (
	"github.com/high-performance-payment-gateway/order-service/order/application/query_request_balance"
	"github.com/high-performance-payment-gateway/order-service/order/domain/query/order_query"
)

type (
	AllPartnerBalanceQuery struct {
		RequestBalanceQuery order_query.OneRequestInterface
	}

	AllPartnerBalanceQueryInterface interface {
		Init() error
		GetOneRequestBalance(p query_request_balance.ParamQueryOneBalance) query_request_balance.RequestBalanceResponse
	}
)

func (a *AllPartnerBalanceQuery) Init() error {
	return nil
}

func (a *AllPartnerBalanceQuery) GetOneRequestBalance(p query_request_balance.ParamQueryOneBalance) query_request_balance.RequestBalanceResponse {
	return a.RequestBalanceQuery.HandleOneRequestQuery(p)
}

func NewAllPartnerBalanceQuery(RequestBalanceQuery order_query.OneRequestInterface) *AllPartnerBalanceQuery {
	var _ AllPartnerBalanceQueryInterface = (*AllPartnerBalanceQuery)(nil)
	return &AllPartnerBalanceQuery{
		RequestBalanceQuery: RequestBalanceQuery,
	}
}
