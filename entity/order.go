package entity

import (
	"time"
)

type Order struct {
	Id                string
	OrderNumber       string
	OrderDateTime     time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
	CurrencyId        string
	CurrencyFactor    float64
	SalesChannelId    string
	OrderCustomer     OrderCustomer
	BillingAddressId  string
	BillingAddress    BillingAddress
	Price             Price
	AmountTotal       float64
	AmountNet         float64
	ShippingCosts     ShippingCosts
	ShippingTotal     float64
	StateId           string
	StateMachineState StateMachineState
	Transactions      []Transaction
	Deliveries        []Delivery
	VersionId         string
}
