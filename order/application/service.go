package application

import "github.com/high-performance-payment-gateway/order-service/order/application/query_request_balance"

type (
	Service struct {
		allP AllPartnerBalanceQueryInterface
	}
)

func (s *Service) Init() {
	s.allP.Init()
}

func (s *Service) GetOneRequestBalance(pq query_request_balance.ParamQueryOneBalance) query_request_balance.RequestBalanceResponse {
	return s.allP.GetOneRequestBalance(pq)
}

func NewService(allP AllPartnerBalanceQueryInterface) *Service {
	var _ ServiceInterface = (*Service)(nil)
	return &Service{allP: allP}
}
