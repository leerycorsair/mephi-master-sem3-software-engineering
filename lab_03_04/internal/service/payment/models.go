package payment

const (
	PaymentPending   = "Pending"
	PaymentCompleted = "Completed"
	PaymentRefunded  = "Refunded"
)

type Payment struct {
	OrderID int
	Amount  float64
	Method  string
	Status  string
}

type OrderItem struct {
	MenuItemID int
	Quantity   int
}
