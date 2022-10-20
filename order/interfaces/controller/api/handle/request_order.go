package handle

import (
	"github.com/gofiber/fiber/v2"
	"github.com/high-performance-payment-gateway/order-service/order/application"
	"github.com/high-performance-payment-gateway/order-service/order/application/query_request_balance"
	"github.com/high-performance-payment-gateway/order-service/order/infrastructure/server/web_server"
	validate_api "github.com/high-performance-payment-gateway/order-service/order/interfaces/controller/api/validate"
	"github.com/high-performance-payment-gateway/order-service/order/interfaces/controller/dto/api/dto_api_request"
	"github.com/high-performance-payment-gateway/order-service/order/interfaces/controller/dto/api/dto_api_response"
	validate_base "github.com/high-performance-payment-gateway/order-service/order/pkg/external/validate"
)

type (
	RequestOrderQuery struct {
		sv application.ServiceInterface
	}

	RequestBalanceQueryResponse struct {
		HttpStatus int
		Status     string
		Code       int
		Message    string
	}
)

func (r *RequestOrderQuery) HealthCheck(c *fiber.Ctx) error {
	return c.Status(200).JSON(web_server.MapBase{
		"status": "ok",
	})
}

func (r *RequestOrderQuery) GetOneOrderInfo(c *fiber.Ctx) error {
	rqDto := dto_api_request.NewRequestBalanceQuery()
	res, errB := rqDto.BindDataDto(c)
	if errB != nil {
		return res.Response(c)
	}

	validate := validate_api.ValidateApiRequestBalance{
		VB:  validate_base.NewBaseValidate(),
		Dto: rqDto.Request,
	}

	validate.Init()
	resV, errV := validate.Validate()
	if errV != nil {
		return resV.Response(c)
	}

	paramSV := query_request_balance.ParamQueryOneBalance{
		OrderId: rqDto.Request.OrderId,
	}

	resRQB := r.sv.GetOneOrderInfo(paramSV)
	resProcess := dto_api_response.NewResponseRequestBalanceDto()
	resProcess.MappingFrServiceRequestBalanceResponse(resRQB)

	return resProcess.Response(c)
}

func NewRequestOrderQuery(sv application.ServiceInterface) *RequestOrderQuery {
	return &RequestOrderQuery{
		sv: sv,
	}
}
