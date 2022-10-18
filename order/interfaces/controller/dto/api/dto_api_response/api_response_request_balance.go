package dto_api_response

import (
	"github.com/high-performance-payment-gateway/order-service/order/application/query_request_balance"
	"github.com/high-performance-payment-gateway/order-service/order/infrastructure/server/web_server"
	"github.com/high-performance-payment-gateway/order-service/order/pkg/external/http/http_status"
)

type (
	ResponseRequestBalanceQuery struct {
		HttpCode int
		Status   string
		Code     int
		Message  string
		Data     any
	}

	ErrorDetailDefault struct {
		ErrorList string //json
	}
)

const (
	STATUS_SUCCESS = "success"
	STATUS_ERROR   = "error"
)

func (r *ResponseRequestBalanceQuery) Response(c web_server.ContextBase) error {
	return c.Status(r.HttpCode).JSON(web_server.MapBase{
		"Status":   r.Status,
		"HttpCode": r.HttpCode,
		"Message":  r.Message,
		"Data":     r.Data,
	})
}

func (r *ResponseRequestBalanceQuery) MappingFrServiceRequestBalanceResponse(response query_request_balance.RequestBalanceResponse) {
	//todo implement mapping error code to error response
	if response.Status == true {
		r.HttpCode = http_status.StatusOK
		r.Status = STATUS_SUCCESS
		r.Code = 200
		r.Message = response.Message
		r.Data = response.Data

		return
	}

	if response.Status == false {
		r.HttpCode = http_status.StatusBadRequest
		r.Status = STATUS_ERROR
		r.Code = 401
		r.Message = response.Message
		r.Data = response.Data
	}
}

func NewResponseRequestBalanceDto() *ResponseRequestBalanceQuery {
	return &ResponseRequestBalanceQuery{}
}
