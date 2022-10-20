package pay_direct_with_bank

import (
	"context"
	"fmt"
	"github.com/high-performance-payment-gateway/order-service/order/domain/command/order_cmd/create_order/base_interface"
	"github.com/high-performance-payment-gateway/order-service/order/domain/command/order_cmd/create_order/global_param_create_order"
	"github.com/high-performance-payment-gateway/order-service/order/domain/command/order_cmd/create_order/response_create_order"
	"github.com/high-performance-payment-gateway/order-service/order/infrastructure/cache/redis"
	"github.com/high-performance-payment-gateway/order-service/order/infrastructure/db/orm"
	"github.com/high-performance-payment-gateway/order-service/order/pkg/external/error_base"
	"github.com/high-performance-payment-gateway/order-service/order/pkg/external/http/http_status"
	validate_base "github.com/high-performance-payment-gateway/order-service/order/pkg/external/validate"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

const (
	STATUS_PENDING = "pending"

	STATUS_CREATE_ORDER_INTERNAL_SUCCESS = "success"
	STATUS_CREATE_ORDER_INTERNAL_ERROR   = "error"

	STATUS_CREATE_ORDER_WITH_BANK_SUCCESS = "success"
	STATUS_CREATE_ORDER_WITH_BANK_ERROR   = "error"

	STATUS_CREATE_ORDER_SUCCESS = "success"
	STATUS_CREATE_ORDER_ERROR   = "error"
)

/**
we connect other bank in the word. We require every bank support api define of our, we connect all bank in the world.
we same master card or visa
*/
type (
	PayDirectWithBank struct {
		bankCf  BankOrderCfInterface
		ParamRq ParamRequestDto
		rc      redis.RedisClusterInterface
		ctx     global_param_create_order.ContextRequest
	}

	DataOrder     = orm.Order
	DataSaveCache struct {
		Data   DataOrder
		IsLock bool
	}

	ResponseCreateOrderInternal struct {
		Status    string
		IsSuccess bool
		Error     error
		Data      DataOrder
	}

	ResponseCreateOrderBank struct {
		Status    string
		IsSuccess bool
		Error     error
	}

	PayDirectWithBankInterface interface {
		base_interface.OrderInterface
		ValidateInput(pr ParamRequestDto) (ResponseValidate, error)
		createOrderInBank() ResponseCreateOrderBank
		createOrderInternal() ResponseCreateOrderInternal
		SetParamRequest(pr ParamRequestDto)
	}
)

func (p PayDirectWithBank) ExpiryTime() time.Duration {
	return time.Hour * 24 * 30
}
func (p PayDirectWithBank) CreateOrderId() uint64 {
	return rand.Uint64()
}

func (p PayDirectWithBank) CreatePartnerOrderId() uint64 {
	return rand.Uint64()
}

func (p *PayDirectWithBank) SetParamRequest(pr ParamRequestDto) {
	p.ParamRq = pr
}

//CreateParamNewOrder: we craeate new PartnerOrderId and OrderId, we want make sure dont have dupplicate. Cost of check dupplicate very large, we dont want action it.
func (p PayDirectWithBank) CreateParamNewOrder() DataOrder {
	return DataOrder{
		OrderId:        p.CreateOrderId(),
		PartnerOrderId: p.CreatePartnerOrderId(),
		PartnerCode:    p.ParamRq.PartnerCode,
		Amount:         p.ParamRq.Amount,
		Status:         STATUS_PENDING,
		PaymentType:    p.ParamRq.PaymentType,
		RequestTime:    time.Now().UTC().Second(),
		ResponseTime:   time.Now().UTC().Second(),
		CreatedAt:      time.Now().UTC().Second(),
		UpdatedAt:      time.Now().UTC().Second(),
	}
}

func (p *PayDirectWithBank) createOrderInternal() ResponseCreateOrderInternal {
	newOrder := p.CreateParamNewOrder()
	// todo save order in mysql cluster and handle status

	dSaveCache := DataSaveCache{}
	key := "....." // todo pass orderId
	errSaveCache := p.rc.SetKeyWithStructValue(context.Background(), key, dSaveCache, p.ExpiryTime())
	if errSaveCache != nil {
		errM := fmt.Sprintf("save Order to Cache error: %v", errSaveCache.Error())
		log.Error(errM)
		panic(errM)
	}

	return ResponseCreateOrderInternal{
		Status:    STATUS_CREATE_ORDER_INTERNAL_SUCCESS,
		IsSuccess: true,
		Error:     nil,
		Data:      newOrder,
	}
}

func (p PayDirectWithBank) UrlCreateOrder(rootDomain string) string {
	return rootDomain + "/create-order"
}

func (p PayDirectWithBank) ValidateInput(pr ParamRequestDto) (ResponseValidate, error) {
	validate := ValidateApiPayWithBank{
		VB:  validate_base.NewBaseValidate(),
		Dto: pr,
	}

	validate.Init()
	return validate.Validate()
}

func (p *PayDirectWithBank) createOrderInBank() ResponseCreateOrderBank {
	_, errorGetRDomain := p.bankCf.GetRootDomainForBank(p.ParamRq.BankCode)
	if errorGetRDomain != nil {
		return ResponseCreateOrderBank{
			Status:    STATUS_CREATE_ORDER_WITH_BANK_SUCCESS,
			IsSuccess: false,
			Error:     errorGetRDomain,
		}
	}

	//todo call api to bank and handle response_create_order
	return ResponseCreateOrderBank{}
}

func (p *PayDirectWithBank) setContextRequest(ctx global_param_create_order.ContextRequest) {
	p.ctx = ctx
}

//CreateOrder
// all call after SetContextRequest
func (p *PayDirectWithBank) CreateOrder(prCreateOrder global_param_create_order.ParamCreateOrder) response_create_order.ResponseCreateOrder {
	p.setContextRequest(prCreateOrder.Ctx)

	dtoRq, errBind := BindDataDtoRequestCreateOrder(p.ctx)
	if errBind != nil {
		return response_create_order.ResponseCreateOrder{
			RequestId:       p.ParamRq.RequestId,
			Status:          STATUS_CREATE_ORDER_ERROR,
			IsSuccess:       false,
			Err:             errBind,
			IsValidateError: true,
			CodeErr:         http_status.StatusBadRequest,
			MessageErr:      "param input not valid, please check doc and try again",
		}
	}
	p.SetParamRequest(dtoRq)

	resVL, errVL := p.ValidateInput(p.ParamRq)
	if errVL != nil {
		return response_create_order.ResponseCreateOrder{
			RequestId:   p.ParamRq.RequestId,
			Status:      STATUS_CREATE_ORDER_ERROR,
			IsSuccess:   false,
			Err:         errVL,
			CodeErr:     resVL.Code,
			MessageErr:  resVL.Message,
			DetailError: resVL.DetailMessageError,
		}
	}

	resCOI := p.createOrderInternal()
	if resCOI.IsSuccess != true || resCOI.Error != nil {
		if error_base.IsErrorBase(resCOI.Error) {
			errBase := error_base.GetErrorBase(resCOI.Error)
			return response_create_order.ResponseCreateOrder{
				RequestId:  p.ParamRq.RequestId,
				Status:     STATUS_CREATE_ORDER_ERROR,
				IsSuccess:  false,
				Err:        errBase,
				CodeErr:    errBase.GetCode(),
				MessageErr: errBase.GetMessage(),
			}
		}
	}

	resCOB := p.createOrderInBank()
	if resCOB.IsSuccess != true || resCOB.Error != nil {
		if error_base.IsErrorBase(resCOI.Error) {
			errBase := error_base.GetErrorBase(resCOI.Error)
			return response_create_order.ResponseCreateOrder{
				RequestId:  p.ParamRq.RequestId,
				Status:     STATUS_CREATE_ORDER_ERROR,
				IsSuccess:  false,
				Err:        errBase,
				CodeErr:    errBase.GetCode(),
				MessageErr: errBase.GetMessage(),
			}
		}
	}

	return response_create_order.ResponseCreateOrder{
		OrderId:        resCOI.Data.OrderId,
		RequestId:      p.ParamRq.RequestId,
		PartnerOrderId: resCOI.Data.PartnerOrderId,
		Status:         STATUS_CREATE_ORDER_SUCCESS,
		IsSuccess:      true,
		Err:            nil,
		Data:           resCOI.Data,
	}
}

func (p PayDirectWithBank) PaymentType() string {
	return "PAY_WITH_BANk"
}

func NewPayDirectWithBank(bankCf BankOrderCfInterface, rc redis.RedisClusterInterface) PayDirectWithBankInterface {
	return &PayDirectWithBank{
		bankCf: bankCf,
		rc:     rc,
	}
}
