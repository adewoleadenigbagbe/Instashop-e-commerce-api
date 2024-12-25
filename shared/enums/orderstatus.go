package enums

type OrderStatus int

const (
	Pending   OrderStatus = 1
	Confirmed OrderStatus = 2
	Shipped   OrderStatus = 3
	Delivered OrderStatus = 4
	Cancelled OrderStatus = 5
)
