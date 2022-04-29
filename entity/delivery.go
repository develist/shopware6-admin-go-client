package entity

type Delivery struct {
	OrderId                string
	ShippingOrderAddressId string
	ShippingMethodId       string
	StateMachineState      StateMachineState
}
