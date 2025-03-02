package payment

type PaymentType string

const (
	CREDIT_CARD PaymentType = "credit_card"
	DEBIT_CARD  PaymentType = "debit_card"
	UPI         PaymentType = "upi"
)

type RefundStatus string

const (
	PENDING   RefundStatus = "pending"
	PROCESSED RefundStatus = "processed"
	CANCELLED RefundStatus = "cancelled"
)
