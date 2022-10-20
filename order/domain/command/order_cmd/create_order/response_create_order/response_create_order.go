package response_create_order

import (
	"github.com/high-performance-payment-gateway/order-service/order/infrastructure/db/orm"
)

type (
	ResponseCreateOrder struct {
		OrderId         uint64
		RequestId       uint64
		PartnerOrderId  uint64
		Status          string
		IsSuccess       bool
		IsValidateError bool
		Err             error
		CodeErr         int
		MessageErr      string
		DetailError     any
		Data            orm.Order
	}
)
