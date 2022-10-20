package errors_create_order

import "github.com/high-performance-payment-gateway/order-service/order/pkg/external/error_base"

type (
	ErrorInternal     = error_base.Error
	ErrorInternalType = error_base.ErrorBase
)

const (
	ERROR_BANK_CODE_NOT_EXITS_CODE      = 1
	ERROR_BANK_CODE_NOT_EXITS_MESSAGE   = "BankCode not exits"
	ERROR_BANK_CODE_NOT_EXITS_SIGNATURE = "errors_create_order_ERROR_HTTP_INTERNAL_SIGNATURE"
)

func NewErrorBankCodeNotExits() ErrorInternal {
	return &ErrorInternalType{
		Code:      ERROR_BANK_CODE_NOT_EXITS_CODE,
		Message:   ERROR_BANK_CODE_NOT_EXITS_MESSAGE,
		Signature: ERROR_BANK_CODE_NOT_EXITS_SIGNATURE,
	}
}

func IsErrorBankCodeNotExits(e error) bool {
	return error_base.IsErrorOfType(e, ERROR_BANK_CODE_NOT_EXITS_SIGNATURE)
}
