package handle

import (
	"github.com/gofiber/fiber/v2"
	"github.com/high-performance-payment-gateway/order-service/order/application"
	"github.com/high-performance-payment-gateway/order-service/order/domain/command/order_cmd/create_order"
	"github.com/high-performance-payment-gateway/order-service/order/domain/command/order_cmd/create_order/global_param_create_order"
	validate_api "github.com/high-performance-payment-gateway/order-service/order/interfaces/controller/api/validate"
	"github.com/high-performance-payment-gateway/order-service/order/interfaces/controller/dto/api/dto_api_request"
	"github.com/high-performance-payment-gateway/order-service/order/interfaces/controller/dto/api/dto_api_response"
	validate_base "github.com/high-performance-payment-gateway/order-service/order/pkg/external/validate"
)

type (
	CreateOrder struct {
		sv application.ServiceInterface
	}
)

func (c *CreateOrder) CreateNewOrder(ctx *fiber.Ctx) error {
	rqDto := dto_api_request.NewCreateOder()
	res, errB := rqDto.BindDataDto(ctx)
	if errB != nil {
		return res.Response(ctx)
	}

	validate := validate_api.ValidateApiCreateOrder{
		VB:  validate_base.NewBaseValidate(),
		Dto: rqDto.RequestDto,
	}

	validate.Init()
	resV, errV := validate.Validate()
	if errV != nil {
		return resV.Response(ctx)
	}

	PrCreateOrder := global_param_create_order.ParamCreateOrder{
		Ctx: ctx,
	}
	PrGetInstance := create_order.ParamGetInstanceOrder{
		PaymentType: rqDto.RequestDto.PaymentType,
		PaymentCode: rqDto.RequestDto.PaymentCode,
	}
	resCO := c.sv.CreateNewOrder(PrCreateOrder, PrGetInstance)

	newResCO := dto_api_response.NewResponseCreateOrder()
	newResCO.MappingFrServiceRequestBalanceResponse(resCO)
	return newResCO.Response(ctx)
}

func NewCreateOrder(sv application.ServiceInterface) *CreateOrder {
	return &CreateOrder{
		sv: sv,
	}
}
