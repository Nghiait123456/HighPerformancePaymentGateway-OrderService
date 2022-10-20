package dto_api_request

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/high-performance-payment-gateway/order-service/order/interfaces/controller/dto/api/dto_api_response"
	"github.com/high-performance-payment-gateway/order-service/order/pkg/external/http/http_status"
	log "github.com/sirupsen/logrus"
)

type (
	CreateOder struct {
		RequestDto CreateOrderDto
	}

	CreateOrderDto struct {
		PaymentType string `json:"PaymentType" ml:"PaymentType" form:"PaymentType" validate:"required,paymentTypeValid"`
		PaymentCode int    `json:"PaymentCode" ml:"PaymentCode" form:"PaymentCode" validate:"required,paymentCodeValid"`
	}
)

func (c *CreateOder) BindDataDto(ctx *fiber.Ctx) (dto_api_response.CreateOrder, error) {
	var o CreateOrderDto
	if errBP := ctx.BodyParser(&o); errBP != nil {
		res := dto_api_response.CreateOrder{
			HttpCode: http_status.StatusBadRequest,
			Status:   dto_api_response.STATUS_ERROR,
			Code:     http_status.StatusBadRequest,
			Message:  "param input not valid, please check doc and try again",
		}

		errML := fmt.Sprintf("param input not valid, please check doc and try again, detail: %s", errBP.Error())
		log.Error(errML)
		return res, errBP
	}

	c.RequestDto = o
	return dto_api_response.CreateOrder{}, nil
}

func NewCreateOder() *CreateOder {
	return &CreateOder{}
}
