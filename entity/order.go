package entity

import (
	"time"
)

type Order struct {
	// Addresses
	AffiliateCode           string
	AmountNet               float64
	AmountTotal             float64
	AutoIncrement           int
	BillingAddress          BillingAddress
	BillingAddressId        string
	BillingAddressVersionId string
	CampaignCode            string
	CreatedAt               time.Time
	// CreatedBy
	CreatedById    string
	Currency       Currency
	CurrencyFactor float64
	CurrencyId     string
	// CustomFields
	CustomerComment string
	DeepLinkCode    string
	Deliveries      []Delivery
	// Documents
	Id           string
	ItemRounding ItemRounding
	Language     Language
	LanguageId   string
	// LineItems
	OrderCustomer Customer
	OrderDate     string
	OrderDateTime time.Time
	OrderNumber   string
	PositionPrice float64
	Price         Price
	// RuleIds
	// SalesChannel
	SalesChannelId    string
	ShippingCosts     ShippingCosts
	ShippingTotal     float64
	StateId           string
	StateMachineState StateMachineState
	// Tags
	TaxStatus string
	// TotalRounding
	Transactions []Transaction
	UpdatedAt    time.Time
	UpdatedBy    string
	UpdatedById  string
	VersionId    string
}

type ItemRounding struct {
	Decimals    int
	Interval    float64
	RoundForNet bool
}
