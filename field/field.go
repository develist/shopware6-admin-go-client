package field

type Field string

const (
	OrderNumber    Field = "orderNumber"
	OrderDateTime        = "orderDateTime"
	BillingAddress       = "billingAddress"
	Deliveries           = "deliveries"
	Transactions         = "transactions"
	PaymentMethod        = "paymentMethod"
)
