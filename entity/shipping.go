package entity

type ShippingCosts struct {
	UnitPrice  float64
	Quantity   int
	TotalPrice float64
}

type ShippingMethodResult struct {
	Data ShippingMethod
}

type ShippingMethod struct {
	Id     string
	Name   string
	Active bool
}
