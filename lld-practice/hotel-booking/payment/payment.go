package payment

type PaymentDetails struct {
	paymentId   string
	paymentType PaymentType
	amount      int
	createdAt   int64
}

type Refund struct {
	refundDetails *PaymentDetails
	status        RefundStatus
}
