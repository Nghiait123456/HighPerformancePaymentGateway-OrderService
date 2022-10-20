package pay_direct_with_bank

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/high-performance-payment-gateway/order-service/order/pkg/external/http/http_status"
	"github.com/high-performance-payment-gateway/order-service/order/pkg/external/validate"
	log "github.com/sirupsen/logrus"
)

type (
	ValidateApiPayWithBank struct {
		VB  validate.ValidateBaseInterface
		Dto ParamRequestDto
	}

	ParamRequestDto struct {
		RequestId       uint64 `json:"RequestId" ml:"RequestId" form:"RequestId" validate:"required"`
		BankCode        string `json:"BankCode" ml:"BankCode" form:"BankCode" validate:"required"`
		TypePayWithBank string `json:"TypePayWithBank" ml:"TypePayWithBank" form:"TypePayWithBank" validate:"required,typePayWithBankValid"`
		Amount          uint64 `json:"Amount" ml:"Amount" form:"Amount" validate:"required,minAmount,maxAmount"`
		PartnerCode     string `json:"PartnerCode" ml:"PartnerCode" form:"PartnerCode" validate:"required,partnerCodeValid"`
		PaymentType     string `json:"PaymentType" ml:"PaymentType" form:"PaymentType" validate:"required,paymentTypeValid"`
	}

	ResponseValidate struct {
		IsSuccess          bool
		Code               int
		Message            string
		DetailMessageError validate.MessageErrors
	}
)

const (
	MIN_AMOUNT = 10000
	MAX_AMOUNT = 5000000
)

func (v *ValidateApiPayWithBank) Init() {
	v.VB.ResignValidateCustom("minAmount", v.minAmount)
	v.VB.ResignValidateCustom("maxAmount", v.maxAmount)
	v.VB.ResignValidateCustom("typePayWithBankValid", v.typePayWithBankValid)
	v.VB.ResignValidateCustom("partnerCodeValid", v.partnerCodeValid)
	v.VB.ResignValidateCustom("paymentTypeValid", v.paymentTypeValid)

	message := make(validate.MapMessage)
	message["minAmount"] = "Amount is less than min allow"
	message["maxAmount"] = "Amount is  greater max allow"
	message["typePayWithBankValid"] = "Type payment direct with bank is not valid"
	message["partnerCodeValid"] = "PartnerCode is not valid"
	message["paymentTypeValid"] = "PaymentType is not valid"
	v.VB.SetMessageForRule(message)
}

// return struct response_create_order, error
func (v *ValidateApiPayWithBank) Validate() (ResponseValidate, error) {
	errV := v.VB.Validate().Struct(v.Dto)
	if errV != nil {
		message, errCE := v.VB.ConvertErrorValidate(errV)
		if errCE != nil {
			fmt.Println("invalidate error")
			res := ResponseValidate{
				IsSuccess: false,
				Code:      http_status.StatusBadRequest,
				Message:   "Param is invalid format, please check and try again",
			}
			return res, errV
		}

		// show message
		errSE, detail := v.VB.ShowErrors(message, v.CustomShowError)
		if errSE != nil {
			messageErr := fmt.Sprintf("ShowErrors validate error: %s", errSE.Error())
			log.WithFields(log.Fields{
				"errMessage": errSE.Error(),
			}).Error("")
			panic(messageErr)
		}

		res := ResponseValidate{
			IsSuccess:          false,
			Code:               http_status.StatusBadRequest,
			Message:            "Param is invalid format, please check and try again",
			DetailMessageError: detail.(validate.MessageErrors),
		}

		return res, errors.New("Validate has errors_create_order")
	}

	return ResponseValidate{}, nil
}

func (v ValidateApiPayWithBank) CustomShowError(mE validate.MessageErrors) (error, interface{}) {
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

func (v ValidateApiPayWithBank) minAmount(fl validator.FieldLevel) bool {
	return fl.Field().Uint() > MIN_AMOUNT
}

func (v ValidateApiPayWithBank) maxAmount(fl validator.FieldLevel) bool {
	return fl.Field().Uint() < MAX_AMOUNT
}

func (v ValidateApiPayWithBank) typePayWithBankValid(fl validator.FieldLevel) bool {
	//todo check
	return true
}

func (v ValidateApiPayWithBank) partnerCodeValid(fl validator.FieldLevel) bool {
	//todo check
	return true
}

func (v ValidateApiPayWithBank) paymentTypeValid(fl validator.FieldLevel) bool {
	//todo check
	return true
}
