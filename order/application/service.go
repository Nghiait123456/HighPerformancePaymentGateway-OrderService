package application

import (
	"github.com/high-performance-payment-gateway/order-service/order/application/query_request_balance"
	"github.com/high-performance-payment-gateway/order-service/order/domain/command/order_cmd/create_order"
	"github.com/high-performance-payment-gateway/order-service/order/domain/command/order_cmd/create_order/global_param_create_order"
	"github.com/high-performance-payment-gateway/order-service/order/domain/command/order_cmd/create_order/response_create_order"
)

type (
	Service struct {
		allP AllPartnerBalanceQueryInterface
	}
)

func (s *Service) Init() {
	s.allP.Init()
}

func (s *Service) GetOneOrderInfo(pq query_request_balance.ParamQueryOneBalance) query_request_balance.RequestBalanceResponse {
	return s.allP.GetOneOrderInfo(pq)
}

func (s *Service) CreateNewOrder(PrCreateOrder global_param_create_order.ParamCreateOrder, PrGetInstance create_order.ParamGetInstanceOrder) response_create_order.ResponseCreateOrder {
	createOrderSv := NewCreateOrder(PrCreateOrder, PrGetInstance)
	return createOrderSv.CreateNewOrder()
}

func NewService(allP AllPartnerBalanceQueryInterface) *Service {
	var _ ServiceInterface = (*Service)(nil)
	return &Service{
		allP: allP,
	}
}
