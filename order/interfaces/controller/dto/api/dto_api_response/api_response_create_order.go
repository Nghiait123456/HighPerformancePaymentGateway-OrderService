package dto_api_response

import (
	"github.com/high-performance-payment-gateway/order-service/order/domain/command/order_cmd/create_order/response_create_order"
	"github.com/high-performance-payment-gateway/order-service/order/infrastructure/server/web_server"
	"github.com/high-performance-payment-gateway/order-service/order/pkg/external/http/http_status"
)

type (
	CreateOrder struct {
		HttpCode int
		Status   string
		Code     int
		Message  string
		Data     any
	}
)

func (r *CreateOrder) Response(c web_server.ContextBase) error {
	return c.Status(r.HttpCode).JSON(web_server.MapBase{
		"Status":   r.Status,
		"HttpCode": r.HttpCode,
		"Message":  r.Message,
		"Data":     r.Data,
	})
}

func (r *CreateOrder) MappingFrServiceRequestBalanceResponse(response response_create_order.ResponseCreateOrder) {
	//todo implement mapping error code to error response_create_order
	if response.IsSuccess == true {
		r.HttpCode = http_status.StatusOK
		r.Status = STATUS_SUCCESS
		r.Code = 200
		r.Data = response.Data
	}

	if response.IsSuccess == false {
		r.HttpCode = http_status.StatusBadRequest
		r.Status = STATUS_ERROR
		r.Code = 401
		r.Message = response.MessageErr
	}
}

func NewResponseCreateOrder() *CreateOrder {
	return &CreateOrder{}
}
