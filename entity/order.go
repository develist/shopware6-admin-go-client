package entity

import (
	"time"
)

type Order struct {
	Addresses               OrderAddress `json:"addresses,omitempty"`
	AffiliateCode           string       `json:"affiliateCode,omitempty"`
	AmountNet               float64      `json:"amountNet,omitempty"`
	AmountTotal             float64
	AutoIncrement           int
	BillingAddress          BillingAddress
	BillingAddressId        string
	BillingAddressVersionId string
	CampaignCode            string
	CreatedAt               time.Time
	CreatedBy               User
	CreatedById             string
	Currency                Currency
	CurrencyFactor          float64
	CurrencyId              string
	CustomFields            []CustomField
	CustomerComment         string
	DeepLinkCode            string
	Deliveries              []OrderDelivery
	Documents               []Document
	Id                      string `json:"id,omitempty"`
	ItemRounding            Rounding
	Language                Language
	LanguageId              string
	LineItems               []OrderLineItem
	OrderCustomer           Customer
	OrderDate               string
	OrderDateTime           time.Time
	OrderNumber             string
	PositionPrice           float64
	Price                   Price
	RuleIds                 []string
	SalesChannel            SalesChannel
	SalesChannelId          string
	ShippingCosts           ShippingCosts
	ShippingTotal           float64
	StateId                 string
	StateMachineState       StateMachineState
	Tags                    []OrderTag
	TaxStatus               string
	TotalRounding           Rounding
	Transactions            []OrderTransaction
	UpdatedAt               time.Time
	UpdatedBy               string
	UpdatedById             string
	VersionId               string
}

type OrderAddress struct {
	AdditionalAddressLine1 string `json:"additionalAddressLine1,omitempty"`
	AdditionalAddressLine2 string `json:"additionalAddressLine2,omitempty"`
	City                   string `json:"city,omitempty"`
	Company                string `json:"company,omitempty"`
}

type OrderCustomer struct {
}

type OrderDelivery struct {
	OrderId                string            `json:"orderId,omitempty"`
	ShippingOrderAddressId string            `json:"shippingOrderAddressId,omitempty"`
	ShippingMethodId       string            `json:"shippingMethodId,omitempty"`
	StateMachineState      StateMachineState `json:"stateMachineState,omitempty"`
}

type OrderDeliveryPosition struct {
}

type OrderLineItem struct {
}

type OrderTag struct {
	Id      string `json:"id,omitempty"`
	OrderId string `json:"orderId,omitempty"`
	Tag     Tag    `json:"tag,omitempty"`
}

type OrderTransaction struct {
	Id                string            `json:"id,omitempty"`
	VersionId         string            `json:"versionId,omitempty"`
	OrderId           string            `json:"orderId,omitempty"`
	PaymentMethodId   string            `json:"paymentMethodId,omitempty"`
	CreatedAt         time.Time         `json:"createdAt,omitempty"`
	UpdatedAt         time.Time         `json:"updatedAt,omitempty"`
	PaymentMethod     PaymentMethod     `json:"paymentMethod,omitempty"`
	StateMachineState StateMachineState `json:"stateMachineState,omitempty"`
}

type Rounding struct {
	Decimals    int     `json:"decimals,omitempty"`
	Interval    float64 `json:"interval,omitempty"`
	RoundForNet bool    `json:"roundForNet,omitempty"`
}
