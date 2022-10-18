package dto_api_request

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/high-performance-payment-gateway/order-service/order/interfaces/controller/dto/api/dto_api_response"
	"github.com/high-performance-payment-gateway/order-service/order/pkg/external/http/http_status"
	log "github.com/sirupsen/logrus"
)

type (
	OneRequestBalanceQueryDto struct {
		OrderId uint64 `json:"OrderId" ml:"OrderId" form:"OrderId" validate:"required,minAmount,maxAmount"`
	}

	RequestBalanceQuery struct {
		Request OneRequestBalanceQueryDto
	}

	RequestBalanceQueryInterface interface {
	}
)

func (one *RequestBalanceQuery) BindDataDto(c *fiber.Ctx) (dto_api_response.ResponseRequestBalanceQuery, error) {
	var o OneRequestBalanceQueryDto
	if errBP := c.BodyParser(&o); errBP != nil {
		res := dto_api_response.ResponseRequestBalanceQuery{
			HttpCode: http_status.StatusBadRequest,
			Status:   dto_api_response.STATUS_ERROR,
			Code:     http_status.StatusBadRequest,
			Message:  "param input not valid, please check doc and try again",
			Data:     "",
		}

		errML := fmt.Sprintf("param input not valid, please check doc and try again, detail: %s", errBP.Error())
		log.Error(errML)
		return res, errBP
	}

	one.Request = o
	return dto_api_response.ResponseRequestBalanceQuery{}, nil
}

func NewRequestBalanceQuery() *RequestBalanceQuery {
	return &RequestBalanceQuery{}
}
