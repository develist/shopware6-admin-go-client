package entity

import (
	"time"
)

// completed

type Order struct {
	Addresses               *OrderAddress  `json:"addresses,omitempty"`
	AffiliateCode           string         `json:"affiliateCode,omitempty"`
	AmountNet               float64        `json:"amountNet,omitempty"`
	AmountTotal             float64        `json:"amountTotal,omitempty"`
	AutoIncrement           int            `json:"autoIncrement,omitempty"`
	BillingAddress          *OrderAddress  `json:"billingAddress,omitempty"`
	BillingAddressId        string         `json:"billingAddressId,omitempty"`
	BillingAddressVersionId string         `json:"billingAddressVersionId,omitempty"`
	CampaignCode            string         `json:"campaignCode,omitempty"`
	CreatedAt               *time.Time     `json:"createdAt,omitempty"`
	CreatedBy               *User          `json:"createdBy,omitempty"`
	CreatedById             string         `json:"createdById,omitempty"`
	Currency                *Currency      `json:"currency,omitempty"`
	CurrencyFactor          float64        `json:"currencyFactor,omitempty"`
	CurrencyId              string         `json:"currencyId,omitempty"`
	CustomFields            *[]CustomField `json:"customFields,omitempty"`
	CustomerComment         string         `json:"customerComment,omitempty"`
	DeepLinkCode            string         `json:"deepLinkCode,omitempty"`
	Deliveries              *OrderDelivery `json:"deliveries,omitempty"`
	Documents               *Document      `json:"documents,omitempty"`
	Id                      string         `json:"id,omitempty"`
	// ItemRounding map[properties:map[decimals:map[format:int64 type:integer] interval:map[format:float type:number] roundForNet:map[type:boolean]] type:object] `json:"itemRounding,omitempty"`
	Language      *Language      `json:"language,omitempty"`
	LanguageId    string         `json:"languageId,omitempty"`
	LineItems     *OrderLineItem `json:"lineItems,omitempty"`
	OrderCustomer *OrderCustomer `json:"orderCustomer,omitempty"`
	OrderDate     string         `json:"orderDate,omitempty"`
	OrderDateTime *time.Time     `json:"orderDateTime,omitempty"`
	OrderNumber   string         `json:"orderNumber,omitempty"`
	PositionPrice float64        `json:"positionPrice,omitempty"`
	// Price map[properties:map[calculatedTaxes:map[type:object] netPrice:map[format:float type:number] positionPrice:map[format:float type:number] rawTotal:map[format:float type:number] taxRules:map[type:object] taxStatus:map[type:string] totalPrice:map[format:float type:number]] required:[netPrice totalPrice positionPrice rawTotal taxStatus] type:object] `json:"price,omitempty"`
	// RuleIds map[items:map[type:string] type:array] `json:"ruleIds,omitempty"`
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId string        `json:"salesChannelId,omitempty"`
	// ShippingCosts map[properties:map[calculatedTaxes:map[type:object] listPrice:map[properties:map[discount:map[format:float type:number] percentage:map[format:float type:number] price:map[format:float type:number]] type:object] quantity:map[format:int64 type:integer] referencePrice:map[type:object] taxRules:map[type:object] totalPrice:map[format:float type:number] unitPrice:map[format:float type:number]] required:[unitPrice totalPrice quantity] type:object] `json:"shippingCosts,omitempty"`
	ShippingTotal     float64            `json:"shippingTotal,omitempty"`
	StateId           string             `json:"stateId,omitempty"`
	StateMachineState *StateMachineState `json:"stateMachineState,omitempty"`
	Tags              *Tag               `json:"tags,omitempty"`
	TaxStatus         string             `json:"taxStatus,omitempty"`
	// TotalRounding map[properties:map[decimals:map[format:int64 type:integer] interval:map[format:float type:number] roundForNet:map[type:boolean]] type:object] `json:"totalRounding,omitempty"`
	Transactions *OrderTransaction `json:"transactions,omitempty"`
	UpdatedAt    *time.Time        `json:"updatedAt,omitempty"`
	UpdatedBy    *User             `json:"updatedBy,omitempty"`
	UpdatedById  string            `json:"updatedById,omitempty"`
	VersionId    string            `json:"versionId,omitempty"`
}

type OrderAddress struct {
	AdditionalAddressLine1 string         `json:"additionalAddressLine1,omitempty"`
	AdditionalAddressLine2 string         `json:"additionalAddressLine2,omitempty"`
	City                   string         `json:"city,omitempty"`
	Company                string         `json:"company,omitempty"`
	Country                *Country       `json:"country,omitempty"`
	CountryId              string         `json:"countryId,omitempty"`
	CountryState           *CountryState  `json:"countryState,omitempty"`
	CountryStateId         string         `json:"countryStateId,omitempty"`
	CreatedAt              *time.Time     `json:"createdAt,omitempty"`
	CustomFields           *[]CustomField `json:"customFields,omitempty"`
	Department             string         `json:"department,omitempty"`
	FirstName              string         `json:"firstName,omitempty"`
	Id                     string         `json:"id,omitempty"`
	LastName               string         `json:"lastName,omitempty"`
	Order                  *Order         `json:"order,omitempty"`
	OrderDeliveries        *OrderDelivery `json:"orderDeliveries,omitempty"`
	OrderId                string         `json:"orderId,omitempty"`
	OrderVersionId         string         `json:"orderVersionId,omitempty"`
	PhoneNumber            string         `json:"phoneNumber,omitempty"`
	Salutation             *Salutation    `json:"salutation,omitempty"`
	SalutationId           string         `json:"salutationId,omitempty"`
	Street                 string         `json:"street,omitempty"`
	Title                  string         `json:"title,omitempty"`
	UpdatedAt              *time.Time     `json:"updatedAt,omitempty"`
	VatId                  string         `json:"vatId,omitempty"`
	VersionId              string         `json:"versionId,omitempty"`
	Zipcode                string         `json:"zipcode,omitempty"`
}

type OrderCustomer struct {
	Company        string         `json:"company,omitempty"`
	CreatedAt      *time.Time     `json:"createdAt,omitempty"`
	CustomFields   *[]CustomField `json:"customFields,omitempty"`
	Customer       *Customer      `json:"customer,omitempty"`
	CustomerId     string         `json:"customerId,omitempty"`
	CustomerNumber string         `json:"customerNumber,omitempty"`
	Email          string         `json:"email,omitempty"`
	FirstName      string         `json:"firstName,omitempty"`
	Id             string         `json:"id,omitempty"`
	LastName       string         `json:"lastName,omitempty"`
	Order          *Order         `json:"order,omitempty"`
	OrderId        string         `json:"orderId,omitempty"`
	OrderVersionId string         `json:"orderVersionId,omitempty"`
	RemoteAddress  string         `json:"remoteAddress,omitempty"`
	Salutation     *Salutation    `json:"salutation,omitempty"`
	SalutationId   string         `json:"salutationId,omitempty"`
	Title          string         `json:"title,omitempty"`
	UpdatedAt      *time.Time     `json:"updatedAt,omitempty"`
	// VatIds map[items:map[type:string] type:array] `json:"vatIds,omitempty"`
	VersionId string `json:"versionId,omitempty"`
}

type OrderDelivery struct {
	CreatedAt      *time.Time             `json:"createdAt,omitempty"`
	CustomFields   *[]CustomField         `json:"customFields,omitempty"`
	Id             string                 `json:"id,omitempty"`
	Order          *Order                 `json:"order,omitempty"`
	OrderId        string                 `json:"orderId,omitempty"`
	OrderVersionId string                 `json:"orderVersionId,omitempty"`
	Positions      *OrderDeliveryPosition `json:"positions,omitempty"`
	// ShippingCosts map[properties:map[calculatedTaxes:map[type:object] listPrice:map[properties:map[discount:map[format:float type:number] percentage:map[format:float type:number] price:map[format:float type:number]] type:object] quantity:map[format:int64 type:integer] referencePrice:map[type:object] taxRules:map[type:object] totalPrice:map[format:float type:number] unitPrice:map[format:float type:number]] required:[unitPrice totalPrice quantity] type:object] `json:"shippingCosts,omitempty"`
	ShippingDateEarliest          *time.Time         `json:"shippingDateEarliest,omitempty"`
	ShippingDateLatest            *time.Time         `json:"shippingDateLatest,omitempty"`
	ShippingMethod                *ShippingMethod    `json:"shippingMethod,omitempty"`
	ShippingMethodId              string             `json:"shippingMethodId,omitempty"`
	ShippingOrderAddress          *OrderAddress      `json:"shippingOrderAddress,omitempty"`
	ShippingOrderAddressId        string             `json:"shippingOrderAddressId,omitempty"`
	ShippingOrderAddressVersionId string             `json:"shippingOrderAddressVersionId,omitempty"`
	StateId                       string             `json:"stateId,omitempty"`
	StateMachineState             *StateMachineState `json:"stateMachineState,omitempty"`
	// TrackingCodes map[items:map[type:string] type:array] `json:"trackingCodes,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	VersionId string     `json:"versionId,omitempty"`
}

type OrderDeliveryPosition struct {
	CreatedAt              *time.Time     `json:"createdAt,omitempty"`
	CustomFields           *[]CustomField `json:"customFields,omitempty"`
	Id                     string         `json:"id,omitempty"`
	OrderDelivery          *OrderDelivery `json:"orderDelivery,omitempty"`
	OrderDeliveryId        string         `json:"orderDeliveryId,omitempty"`
	OrderDeliveryVersionId string         `json:"orderDeliveryVersionId,omitempty"`
	OrderLineItem          *OrderLineItem `json:"orderLineItem,omitempty"`
	OrderLineItemId        string         `json:"orderLineItemId,omitempty"`
	OrderLineItemVersionId string         `json:"orderLineItemVersionId,omitempty"`
	// Price map[properties:map[calculatedTaxes:map[type:object] listPrice:map[properties:map[discount:map[format:float type:number] percentage:map[format:float type:number] price:map[format:float type:number]] type:object] quantity:map[format:int64 type:integer] referencePrice:map[type:object] taxRules:map[type:object] totalPrice:map[format:float type:number] unitPrice:map[format:float type:number]] required:[unitPrice totalPrice quantity] type:object] `json:"price,omitempty"`
	Quantity   int        `json:"quantity,omitempty"`
	TotalPrice float64    `json:"totalPrice,omitempty"`
	UnitPrice  float64    `json:"unitPrice,omitempty"`
	UpdatedAt  *time.Time `json:"updatedAt,omitempty"`
	VersionId  string     `json:"versionId,omitempty"`
}

type OrderLineItem struct {
	Children               *OrderLineItem         `json:"children,omitempty"`
	Cover                  *Media                 `json:"cover,omitempty"`
	CoverId                string                 `json:"coverId,omitempty"`
	CreatedAt              time.Time              `json:"createdAt,omitempty"`
	CustomFields           *[]CustomField         `json:"customFields,omitempty"`
	Description            string                 `json:"description,omitempty"`
	Good                   bool                   `json:"good,omitempty"`
	Id                     string                 `json:"id,omitempty"`
	Identifier             string                 `json:"identifier,omitempty"`
	Label                  string                 `json:"label,omitempty"`
	Order                  *Order                 `json:"order,omitempty"`
	OrderDeliveryPositions *OrderDeliveryPosition `json:"orderDeliveryPositions,omitempty"`
	OrderId                string                 `json:"orderId,omitempty"`
	OrderVersionId         string                 `json:"orderVersionId,omitempty"`
	Parent                 *OrderLineItem         `json:"parent,omitempty"`
	ParentId               string                 `json:"parentId,omitempty"`
	ParentVersionId        string                 `json:"parentVersionId,omitempty"`
	// Payload map[type:object] `json:"payload,omitempty"`
	Position int `json:"position,omitempty"`
	// Price map[properties:map[calculatedTaxes:map[type:object] listPrice:map[properties:map[discount:map[format:float type:number] percentage:map[format:float type:number] price:map[format:float type:number]] type:object] quantity:map[format:int64 type:integer] referencePrice:map[type:object] taxRules:map[type:object] totalPrice:map[format:float type:number] unitPrice:map[format:float type:number]] required:[unitPrice totalPrice quantity] type:object] `json:"price,omitempty"`
	// PriceDefinition map[type:object] `json:"priceDefinition,omitempty"`
	Product          *Product  `json:"product,omitempty"`
	ProductId        string    `json:"productId,omitempty"`
	ProductVersionId string    `json:"productVersionId,omitempty"`
	Quantity         int       `json:"quantity,omitempty"`
	ReferencedId     string    `json:"referencedId,omitempty"`
	Removable        bool      `json:"removable,omitempty"`
	Stackable        bool      `json:"stackable,omitempty"`
	TotalPrice       float64   `json:"totalPrice,omitempty"`
	Type             string    `json:"type,omitempty"`
	UnitPrice        float64   `json:"unitPrice,omitempty"`
	UpdatedAt        time.Time `json:"updatedAt,omitempty"`
	VersionId        string    `json:"versionId,omitempty"`
}

type OrderTag struct {
	Id             string `json:"id,omitempty"`
	Order          *Order `json:"order,omitempty"`
	OrderId        string `json:"orderId,omitempty"`
	OrderVersionId string `json:"orderVersionId,omitempty"`
	Tag            *Tag   `json:"tag,omitempty"`
	TagId          string `json:"tagId,omitempty"`
}

type OrderTransaction struct {
	// Amount map[properties:map[calculatedTaxes:map[type:object] listPrice:map[properties:map[discount:map[format:float type:number] percentage:map[format:float type:number] price:map[format:float type:number]] type:object] quantity:map[format:int64 type:integer] referencePrice:map[type:object] taxRules:map[type:object] totalPrice:map[format:float type:number] unitPrice:map[format:float type:number]] required:[unitPrice totalPrice quantity] type:object] `json:"amount,omitempty"`
	CreatedAt         time.Time          `json:"createdAt,omitempty"`
	CustomFields      *[]CustomField     `json:"customFields,omitempty"`
	Id                string             `json:"id,omitempty"`
	Order             *Order             `json:"order,omitempty"`
	OrderId           string             `json:"orderId,omitempty"`
	OrderVersionId    string             `json:"orderVersionId,omitempty"`
	PaymentMethod     *PaymentMethod     `json:"paymentMethod,omitempty"`
	PaymentMethodId   string             `json:"paymentMethodId,omitempty"`
	StateId           string             `json:"stateId,omitempty"`
	StateMachineState *StateMachineState `json:"stateMachineState,omitempty"`
	UpdatedAt         time.Time          `json:"updatedAt,omitempty"`
	VersionId         string             `json:"versionId,omitempty"`
}
