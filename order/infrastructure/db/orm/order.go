package orm

type Order struct {
	ID             uint32 `gorm:"<-:create"`
	OrderId        uint64 `gorm:"uniqueIndex"`
	PartnerOrderId uint64 `gorm:"uniqueIndex"`
	PartnerCode    string
	Amount         uint64
	Status         string
	PaymentType    string
	RequestTime    uint32
	ResponseTime   uint32
	CreatedAt      uint32
	UpdatedAt      uint32
}

const (
	ORDER_STATUS_PENDING    = "pending"
	ORDER_STATUS_PROCESSING = "processing"
	ORDER_STATUS_SUCCESS    = "success"
	ORDER_STATUS_CANCEL     = "cancel"
	ORDER_STATUS_TIMEOUT    = "timeout"
	ORDER_STATUS_ERROR      = "error"
)

// TableName overrides
func (b *Order) TableName() string {
	return "order"
}
func (b Order) StatusPending() string {
	return ORDER_STATUS_PENDING
}
func (b Order) StatusSuccess() string {
	return ORDER_STATUS_SUCCESS
}
func (b Order) StatusProcessing() string {
	return ORDER_STATUS_PROCESSING
}
func (b Order) StatusCancel() string {
	return ORDER_STATUS_CANCEL
}
func (b Order) StatusTimeout() string {
	return ORDER_STATUS_TIMEOUT
}

func (b Order) StatusError() string {
	return ORDER_STATUS_ERROR
}
