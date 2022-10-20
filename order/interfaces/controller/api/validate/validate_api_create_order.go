package validate

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/high-performance-payment-gateway/order-service/order/interfaces/controller/dto/api/dto_api_request"
	"github.com/high-performance-payment-gateway/order-service/order/interfaces/controller/dto/api/dto_api_response"
	"github.com/high-performance-payment-gateway/order-service/order/pkg/external/http/http_status"
	"github.com/high-performance-payment-gateway/order-service/order/pkg/external/validate"
	log "github.com/sirupsen/logrus"
)

type (
	ValidateApiCreateOrder struct {
		VB  validate.ValidateBaseInterface
		Dto dto_api_request.CreateOrderDto
	}
)

const (
	MIN_AMOUNT_CREATE_ORDER = 10000
	MAX_AMOUNT_CREATE_ORDER = 5000000
)

func (v *ValidateApiCreateOrder) Init() {
	v.VB.ResignValidateCustom("paymentTypeValid", v.paymentTypeValid)
	v.VB.ResignValidateCustom("paymentCodeValid", v.paymentCodeValid)

	message := make(validate.MapMessage)
	message["paymentTypeValid"] = "paymentType is not valid"
	message["paymentCodeValid"] = "paymentCodeValid is not valid"
	v.VB.SetMessageForRule(message)
}

// return struct response_create_order, error
func (v *ValidateApiCreateOrder) Validate() (dto_api_response.ResponseRequestBalanceQuery, error) {
	errV := v.VB.Validate().Struct(v.Dto)
	if errV != nil {
		message, errCE := v.VB.ConvertErrorValidate(errV)
		if errCE != nil {
			fmt.Println("invalidate error")
			res := dto_api_response.ResponseRequestBalanceQuery{
				HttpCode: http_status.StatusBadRequest,
				Status:   dto_api_response.STATUS_ERROR,
				Code:     http_status.StatusBadRequest,
				Message:  "Param is invalid format, please check and try again",
				Data:     "",
			}
			return res, errV
		}

		fmt.Println("message =", message)

		// show message
		errSE, detail := v.VB.ShowErrors(message, v.CustomShowError)
		fmt.Println("detail =", detail)
		if errSE != nil {
			messageErr := fmt.Sprintf("ShowErrors validate error: %s", errSE.Error())
			log.WithFields(log.Fields{
				"errMessage": errSE.Error(),
			}).Error("")
			panic(messageErr)
		}

		res := dto_api_response.ResponseRequestBalanceQuery{
			HttpCode: http_status.StatusBadRequest,
			Status:   dto_api_response.STATUS_ERROR,
			Code:     http_status.StatusBadRequest,
			Message:  "Param missing or invalid format, please check and try again",
			Data:     "",
		}
		fmt.Println("res", res)

		return res, errors.New("Validate has errors_create_order")
	}

	return dto_api_response.ResponseRequestBalanceQuery{}, nil
}

func (v ValidateApiCreateOrder) CustomShowError(mE validate.MessageErrors) (error, interface{}) {
	ListES := validate.ListErrorsDefaultShow{}

	for _, v := range mE {
		oneErr := validate.OneErrorDefaultShow{
			Field:      v.Field,
			Rule:       v.Rule,
			Message:    v.Message,
			ParamRule:  v.ParamRule,
			ValueError: v.ValueError,
		}

		ListES[v.Field] = oneErr
	}

	showE := validate.DefaultShowError{
		ListError: ListES,
	}

	return nil, showE
}

func (v ValidateApiCreateOrder) paymentTypeValid(fl validator.FieldLevel) bool {
	//todo check
	return true
}

func (v ValidateApiCreateOrder) paymentCodeValid(fl validator.FieldLevel) bool {
	//todo check
	return true
}
