package entity

import (
	"time"
)

// Generated from Shopware Admin API
// Version 6.5.9999999.9999999-dev at 2023-05-02 19:06:34 UTC

// Order data structure
// Added since version: 6.0.0.0
// Required fields: billingAddressId, currencyId, languageId, salesChannelId, orderDateTime, currencyFactor, stateId, createdAt
type Order struct {
	Addresses               *OrderAddress      `json:"addresses,omitempty"`
	AffiliateCode           *string            `json:"affiliateCode,omitempty"`
	AmountNet               *float64           `json:"amountNet,omitempty"`
	AmountTotal             *float64           `json:"amountTotal,omitempty"`
	AutoIncrement           *int               `json:"autoIncrement,omitempty"`
	BillingAddress          *OrderAddress      `json:"billingAddress,omitempty"`
	BillingAddressId        string             `json:"billingAddressId,omitempty"`
	BillingAddressVersionId *string            `json:"billingAddressVersionId,omitempty"`
	CampaignCode            *string            `json:"campaignCode,omitempty"`
	CreatedAt               *time.Time         `json:"createdAt,omitempty"`
	CreatedBy               *User              `json:"createdBy,omitempty"`
	CreatedById             *string            `json:"createdById,omitempty"`
	Currency                *Currency          `json:"currency,omitempty"`
	CurrencyFactor          float64            `json:"currencyFactor,omitempty"`
	CurrencyId              string             `json:"currencyId,omitempty"`
	CustomFields            *[]CustomField     `json:"customFields,omitempty"`
	CustomerComment         *string            `json:"customerComment,omitempty"`
	DeepLinkCode            *string            `json:"deepLinkCode,omitempty"`
	Deliveries              *OrderDelivery     `json:"deliveries,omitempty"`
	Documents               *Document          `json:"documents,omitempty"`
	Extensions              *interface{}       `json:"extensions,omitempty"` // map[properties:map[delayActions:map[properties:map[data:map[items:map[properties:map[id:map[example:f2a69391d84240deb6ece2a78b8a6d5b type:string] type:map[example:swag_delay_action type:string]] type:object] type:array] links:map[properties:map[related:map[example:/order/7eeb6a3313d04faf896be34599decfe0/delayActions format:uri-reference type:string]] type:object]] type:object]] type:object]
	Id                      string             `json:"id,omitempty"`
	ItemRounding            *Rounding          `json:"itemRounding,omitempty"` // map[properties:map[decimals:map[format:int64 type:integer] interval:map[format:float type:number] roundForNet:map[type:boolean]] type:object]
	Language                *Language          `json:"language,omitempty"`
	LanguageId              string             `json:"languageId,omitempty"`
	LineItems               *OrderLineItem     `json:"lineItems,omitempty"`
	OrderCustomer           *OrderCustomer     `json:"orderCustomer,omitempty"`
	OrderDate               *string            `json:"orderDate,omitempty"`
	OrderDateTime           *time.Time         `json:"orderDateTime,omitempty"`
	OrderNumber             *string            `json:"orderNumber,omitempty"`
	PositionPrice           *float64           `json:"positionPrice,omitempty"`
	Price                   *interface{}       `json:"price,omitempty"`   // map[properties:map[calculatedTaxes:map[type:object] netPrice:map[format:float type:number] positionPrice:map[format:float type:number] rawTotal:map[format:float type:number] taxRules:map[type:object] taxStatus:map[type:string] totalPrice:map[format:float type:number]] required:[netPrice totalPrice positionPrice rawTotal taxStatus] type:object]
	RuleIds                 *[]string          `json:"ruleIds,omitempty"` // map[items:map[type:string] type:array]
	SalesChannel            *SalesChannel      `json:"salesChannel,omitempty"`
	SalesChannelId          string             `json:"salesChannelId,omitempty"`
	ShippingCosts           *interface{}       `json:"shippingCosts,omitempty"` // map[properties:map[calculatedTaxes:map[type:object] listPrice:map[properties:map[discount:map[format:float type:number] percentage:map[format:float type:number] price:map[format:float type:number]] type:object] quantity:map[format:int64 type:integer] referencePrice:map[type:object] regulationPrice:map[properties:map[price:map[format:float type:number]] type:object] taxRules:map[type:object] totalPrice:map[format:float type:number] unitPrice:map[format:float type:number]] required:[unitPrice totalPrice quantity] type:object]
	ShippingTotal           *float64           `json:"shippingTotal,omitempty"`
	StateId                 string             `json:"stateId,omitempty"`
	StateMachineState       *StateMachineState `json:"stateMachineState,omitempty"`
	Tags                    *Tag               `json:"tags,omitempty"`
	TaxStatus               *string            `json:"taxStatus,omitempty"`
	TotalRounding           *Rounding          `json:"totalRounding,omitempty"` // map[properties:map[decimals:map[format:int64 type:integer] interval:map[format:float type:number] roundForNet:map[type:boolean]] type:object]
	Transactions            *OrderTransaction  `json:"transactions,omitempty"`
	UpdatedAt               *time.Time         `json:"updatedAt,omitempty"`
	UpdatedBy               *User              `json:"updatedBy,omitempty"`
	UpdatedById             *string            `json:"updatedById,omitempty"`
	VersionId               *string            `json:"versionId,omitempty"`
}

// OrderAddress data structure
// Added since version: 6.0.0.0
// Required fields: countryId, orderId, salutationId, firstName, lastName, street, city, createdAt
type OrderAddress struct {
	AdditionalAddressLine1 *string        `json:"additionalAddressLine1,omitempty"`
	AdditionalAddressLine2 *string        `json:"additionalAddressLine2,omitempty"`
	City                   string         `json:"city,omitempty"`
	Company                *string        `json:"company,omitempty"`
	Country                *Country       `json:"country,omitempty"`
	CountryId              string         `json:"countryId,omitempty"`
	CountryState           *CountryState  `json:"countryState,omitempty"`
	CountryStateId         *string        `json:"countryStateId,omitempty"`
	CreatedAt              *time.Time     `json:"createdAt,omitempty"`
	CustomFields           *[]CustomField `json:"customFields,omitempty"`
	Department             *string        `json:"department,omitempty"`
	FirstName              string         `json:"firstName,omitempty"`
	Id                     string         `json:"id,omitempty"`
	LastName               string         `json:"lastName,omitempty"`
	Order                  *Order         `json:"order,omitempty"`
	OrderDeliveries        *OrderDelivery `json:"orderDeliveries,omitempty"`
	OrderId                string         `json:"orderId,omitempty"`
	OrderVersionId         *string        `json:"orderVersionId,omitempty"`
	PhoneNumber            *string        `json:"phoneNumber,omitempty"`
	Salutation             *Salutation    `json:"salutation,omitempty"`
	SalutationId           string         `json:"salutationId,omitempty"`
	Street                 string         `json:"street,omitempty"`
	Title                  *string        `json:"title,omitempty"`
	UpdatedAt              *time.Time     `json:"updatedAt,omitempty"`
	VatId                  *string        `json:"vatId,omitempty"`
	VersionId              *string        `json:"versionId,omitempty"`
	Zipcode                *string        `json:"zipcode,omitempty"`
}

// OrderCustomer data structure
// Added since version: 6.0.0.0
// Required fields: orderId, email, salutationId, firstName, lastName, createdAt
type OrderCustomer struct {
	Company        *string        `json:"company,omitempty"`
	CreatedAt      *time.Time     `json:"createdAt,omitempty"`
	CustomFields   *[]CustomField `json:"customFields,omitempty"`
	Customer       *Customer      `json:"customer,omitempty"`
	CustomerId     *string        `json:"customerId,omitempty"`
	CustomerNumber *string        `json:"customerNumber,omitempty"`
	Email          string         `json:"email,omitempty"`
	FirstName      string         `json:"firstName,omitempty"`
	Id             string         `json:"id,omitempty"`
	LastName       string         `json:"lastName,omitempty"`
	Order          *Order         `json:"order,omitempty"`
	OrderId        string         `json:"orderId,omitempty"`
	OrderVersionId *string        `json:"orderVersionId,omitempty"`
	RemoteAddress  *string        `json:"remoteAddress,omitempty"`
	Salutation     *Salutation    `json:"salutation,omitempty"`
	SalutationId   string         `json:"salutationId,omitempty"`
	Title          *string        `json:"title,omitempty"`
	UpdatedAt      *time.Time     `json:"updatedAt,omitempty"`
	VatIds         *[]string      `json:"vatIds,omitempty"` // map[items:map[type:string] type:array]
	VersionId      *string        `json:"versionId,omitempty"`
}

// OrderDelivery data structure
// Added since version: 6.0.0.0
// Required fields: orderId, shippingOrderAddressId, shippingMethodId, stateId, trackingCodes, shippingDateEarliest, shippingDateLatest, createdAt
type OrderDelivery struct {
	CreatedAt                     *time.Time             `json:"createdAt,omitempty"`
	CustomFields                  *[]CustomField         `json:"customFields,omitempty"`
	Id                            string                 `json:"id,omitempty"`
	Order                         *Order                 `json:"order,omitempty"`
	OrderId                       string                 `json:"orderId,omitempty"`
	OrderVersionId                *string                `json:"orderVersionId,omitempty"`
	Positions                     *OrderDeliveryPosition `json:"positions,omitempty"`
	ShippingCosts                 *interface{}           `json:"shippingCosts,omitempty"` // map[properties:map[calculatedTaxes:map[type:object] listPrice:map[properties:map[discount:map[format:float type:number] percentage:map[format:float type:number] price:map[format:float type:number]] type:object] quantity:map[format:int64 type:integer] referencePrice:map[type:object] regulationPrice:map[properties:map[price:map[format:float type:number]] type:object] taxRules:map[type:object] totalPrice:map[format:float type:number] unitPrice:map[format:float type:number]] required:[unitPrice totalPrice quantity] type:object]
	ShippingDateEarliest          *time.Time             `json:"shippingDateEarliest,omitempty"`
	ShippingDateLatest            *time.Time             `json:"shippingDateLatest,omitempty"`
	ShippingMethod                *ShippingMethod        `json:"shippingMethod,omitempty"`
	ShippingMethodId              string                 `json:"shippingMethodId,omitempty"`
	ShippingOrderAddress          *OrderAddress          `json:"shippingOrderAddress,omitempty"`
	ShippingOrderAddressId        string                 `json:"shippingOrderAddressId,omitempty"`
	ShippingOrderAddressVersionId *string                `json:"shippingOrderAddressVersionId,omitempty"`
	StateId                       string                 `json:"stateId,omitempty"`
	StateMachineState             *StateMachineState     `json:"stateMachineState,omitempty"`
	TrackingCodes                 *[]string              `json:"trackingCodes,omitempty"` // map[items:map[type:string] type:array]
	UpdatedAt                     *time.Time             `json:"updatedAt,omitempty"`
	VersionId                     *string                `json:"versionId,omitempty"`
}

// OrderDeliveryPosition data structure
// Added since version: 6.0.0.0
// Required fields: orderDeliveryId, orderLineItemId, createdAt
type OrderDeliveryPosition struct {
	CreatedAt              *time.Time     `json:"createdAt,omitempty"`
	CustomFields           *[]CustomField `json:"customFields,omitempty"`
	Id                     string         `json:"id,omitempty"`
	OrderDelivery          *OrderDelivery `json:"orderDelivery,omitempty"`
	OrderDeliveryId        string         `json:"orderDeliveryId,omitempty"`
	OrderDeliveryVersionId *string        `json:"orderDeliveryVersionId,omitempty"`
	OrderLineItem          *OrderLineItem `json:"orderLineItem,omitempty"`
	OrderLineItemId        string         `json:"orderLineItemId,omitempty"`
	OrderLineItemVersionId *string        `json:"orderLineItemVersionId,omitempty"`
	Price                  *interface{}   `json:"price,omitempty"` // map[properties:map[calculatedTaxes:map[type:object] listPrice:map[properties:map[discount:map[format:float type:number] percentage:map[format:float type:number] price:map[format:float type:number]] type:object] quantity:map[format:int64 type:integer] referencePrice:map[type:object] regulationPrice:map[properties:map[price:map[format:float type:number]] type:object] taxRules:map[type:object] totalPrice:map[format:float type:number] unitPrice:map[format:float type:number]] required:[unitPrice totalPrice quantity] type:object]
	Quantity               *int           `json:"quantity,omitempty"`
	TotalPrice             *float64       `json:"totalPrice,omitempty"`
	UnitPrice              *float64       `json:"unitPrice,omitempty"`
	UpdatedAt              *time.Time     `json:"updatedAt,omitempty"`
	VersionId              *string        `json:"versionId,omitempty"`
}

// OrderLineItem data structure
// Added since version: 6.0.0.0
// Required fields: orderId, identifier, quantity, label, position, price, children, createdAt
type OrderLineItem struct {
	Children                               *OrderLineItem                         `json:"children,omitempty"`
	Cover                                  *Media                                 `json:"cover,omitempty"`
	CoverId                                *string                                `json:"coverId,omitempty"`
	CreatedAt                              *time.Time                             `json:"createdAt,omitempty"`
	CustomFields                           *[]CustomField                         `json:"customFields,omitempty"`
	Description                            *string                                `json:"description,omitempty"`
	Good                                   *bool                                  `json:"good,omitempty"`
	Id                                     string                                 `json:"id,omitempty"`
	Identifier                             string                                 `json:"identifier,omitempty"`
	Label                                  string                                 `json:"label,omitempty"`
	Order                                  *Order                                 `json:"order,omitempty"`
	OrderDeliveryPositions                 *OrderDeliveryPosition                 `json:"orderDeliveryPositions,omitempty"`
	OrderId                                string                                 `json:"orderId,omitempty"`
	OrderTransactionCaptureRefundPositions *OrderTransactionCaptureRefundPosition `json:"orderTransactionCaptureRefundPositions,omitempty"`
	OrderVersionId                         *string                                `json:"orderVersionId,omitempty"`
	Parent                                 *OrderLineItem                         `json:"parent,omitempty"`
	ParentId                               *string                                `json:"parentId,omitempty"`
	ParentVersionId                        *string                                `json:"parentVersionId,omitempty"`
	Payload                                *interface{}                           `json:"payload,omitempty"` // map[type:object]
	Position                               int                                    `json:"position,omitempty"`
	Price                                  *interface{}                           `json:"price,omitempty"`           // map[properties:map[calculatedTaxes:map[type:object] listPrice:map[properties:map[discount:map[format:float type:number] percentage:map[format:float type:number] price:map[format:float type:number]] type:object] quantity:map[format:int64 type:integer] referencePrice:map[type:object] regulationPrice:map[properties:map[price:map[format:float type:number]] type:object] taxRules:map[type:object] totalPrice:map[format:float type:number] unitPrice:map[format:float type:number]] required:[unitPrice totalPrice quantity] type:object]
	PriceDefinition                        *interface{}                           `json:"priceDefinition,omitempty"` // map[type:object]
	Product                                *Product                               `json:"product,omitempty"`
	ProductId                              *string                                `json:"productId,omitempty"`
	ProductVersionId                       *string                                `json:"productVersionId,omitempty"`
	Promotion                              *Promotion                             `json:"promotion,omitempty"`
	PromotionId                            *string                                `json:"promotionId,omitempty"`
	Quantity                               int                                    `json:"quantity,omitempty"`
	ReferencedId                           *string                                `json:"referencedId,omitempty"`
	Removable                              *bool                                  `json:"removable,omitempty"`
	Stackable                              *bool                                  `json:"stackable,omitempty"`
	TotalPrice                             *float64                               `json:"totalPrice,omitempty"`
	Type                                   *string                                `json:"type,omitempty"`
	UnitPrice                              *float64                               `json:"unitPrice,omitempty"`
	UpdatedAt                              *time.Time                             `json:"updatedAt,omitempty"`
	VersionId                              *string                                `json:"versionId,omitempty"`
}

// OrderTag data structure
// Added since version: 6.0.0.0
// Required fields: orderId, tagId
type OrderTag struct {
	Id             string  `json:"id,omitempty"`
	Order          *Order  `json:"order,omitempty"`
	OrderId        string  `json:"orderId,omitempty"`
	OrderVersionId *string `json:"orderVersionId,omitempty"`
	Tag            *Tag    `json:"tag,omitempty"`
	TagId          string  `json:"tagId,omitempty"`
}

// OrderTransaction data structure
// Added since version: 6.0.0.0
// Required fields: orderId, paymentMethodId, amount, stateId, createdAt
type OrderTransaction struct {
	Amount            *interface{}             `json:"amount,omitempty"` // map[properties:map[calculatedTaxes:map[type:object] listPrice:map[properties:map[discount:map[format:float type:number] percentage:map[format:float type:number] price:map[format:float type:number]] type:object] quantity:map[format:int64 type:integer] referencePrice:map[type:object] regulationPrice:map[properties:map[price:map[format:float type:number]] type:object] taxRules:map[type:object] totalPrice:map[format:float type:number] unitPrice:map[format:float type:number]] required:[unitPrice totalPrice quantity] type:object]
	Captures          *OrderTransactionCapture `json:"captures,omitempty"`
	CreatedAt         *time.Time               `json:"createdAt,omitempty"`
	CustomFields      *[]CustomField           `json:"customFields,omitempty"`
	Id                string                   `json:"id,omitempty"`
	Order             *Order                   `json:"order,omitempty"`
	OrderId           string                   `json:"orderId,omitempty"`
	OrderVersionId    *string                  `json:"orderVersionId,omitempty"`
	PaymentMethod     *PaymentMethod           `json:"paymentMethod,omitempty"`
	PaymentMethodId   string                   `json:"paymentMethodId,omitempty"`
	StateId           string                   `json:"stateId,omitempty"`
	StateMachineState *StateMachineState       `json:"stateMachineState,omitempty"`
	UpdatedAt         *time.Time               `json:"updatedAt,omitempty"`
	VersionId         *string                  `json:"versionId,omitempty"`
}

// OrderTransactionCapture data structure
// Added since version: 6.4.12.0
// Required fields: orderTransactionId, stateId, amount, createdAt
type OrderTransactionCapture struct {
	Amount                    *interface{}                   `json:"amount,omitempty"` // map[properties:map[calculatedTaxes:map[type:object] listPrice:map[properties:map[discount:map[format:float type:number] percentage:map[format:float type:number] price:map[format:float type:number]] type:object] quantity:map[format:int64 type:integer] referencePrice:map[type:object] regulationPrice:map[properties:map[price:map[format:float type:number]] type:object] taxRules:map[type:object] totalPrice:map[format:float type:number] unitPrice:map[format:float type:number]] required:[unitPrice totalPrice quantity] type:object]
	CreatedAt                 *time.Time                     `json:"createdAt,omitempty"`
	CustomFields              *[]CustomField                 `json:"customFields,omitempty"`
	ExternalReference         *string                        `json:"externalReference,omitempty"`
	Id                        string                         `json:"id,omitempty"`
	OrderTransactionId        string                         `json:"orderTransactionId,omitempty"`
	OrderTransactionVersionId *string                        `json:"orderTransactionVersionId,omitempty"`
	Refunds                   *OrderTransactionCaptureRefund `json:"refunds,omitempty"`
	StateId                   string                         `json:"stateId,omitempty"`
	StateMachineState         *StateMachineState             `json:"stateMachineState,omitempty"`
	TotalAmount               *float64                       `json:"totalAmount,omitempty"`
	Transaction               *OrderTransaction              `json:"transaction,omitempty"`
	UpdatedAt                 *time.Time                     `json:"updatedAt,omitempty"`
}

// OrderTransactionCaptureRefund data structure
// Added since version: 6.4.12.0
// Required fields: captureId, stateId, amount, createdAt
type OrderTransactionCaptureRefund struct {
	Amount             *interface{}                           `json:"amount,omitempty"` // map[properties:map[calculatedTaxes:map[type:object] listPrice:map[properties:map[discount:map[format:float type:number] percentage:map[format:float type:number] price:map[format:float type:number]] type:object] quantity:map[format:int64 type:integer] referencePrice:map[type:object] regulationPrice:map[properties:map[price:map[format:float type:number]] type:object] taxRules:map[type:object] totalPrice:map[format:float type:number] unitPrice:map[format:float type:number]] required:[unitPrice totalPrice quantity] type:object]
	CaptureId          string                                 `json:"captureId,omitempty"`
	CreatedAt          *time.Time                             `json:"createdAt,omitempty"`
	CustomFields       *[]CustomField                         `json:"customFields,omitempty"`
	ExternalReference  *string                                `json:"externalReference,omitempty"`
	Id                 string                                 `json:"id,omitempty"`
	Positions          *OrderTransactionCaptureRefundPosition `json:"positions,omitempty"`
	Reason             *string                                `json:"reason,omitempty"`
	StateId            string                                 `json:"stateId,omitempty"`
	StateMachineState  *StateMachineState                     `json:"stateMachineState,omitempty"`
	TotalAmount        *float64                               `json:"totalAmount,omitempty"`
	TransactionCapture *OrderTransactionCapture               `json:"transactionCapture,omitempty"`
	UpdatedAt          *time.Time                             `json:"updatedAt,omitempty"`
}

// OrderTransactionCaptureRefundPosition data structure
// Added since version: 6.4.12.0
// Required fields: refundId, orderLineItemId, amount, createdAt
type OrderTransactionCaptureRefundPosition struct {
	Amount                        *interface{}                   `json:"amount,omitempty"` // map[properties:map[calculatedTaxes:map[type:object] listPrice:map[properties:map[discount:map[format:float type:number] percentage:map[format:float type:number] price:map[format:float type:number]] type:object] quantity:map[format:int64 type:integer] referencePrice:map[type:object] regulationPrice:map[properties:map[price:map[format:float type:number]] type:object] taxRules:map[type:object] totalPrice:map[format:float type:number] unitPrice:map[format:float type:number]] required:[unitPrice totalPrice quantity] type:object]
	CreatedAt                     *time.Time                     `json:"createdAt,omitempty"`
	CustomFields                  *[]CustomField                 `json:"customFields,omitempty"`
	ExternalReference             *string                        `json:"externalReference,omitempty"`
	Id                            string                         `json:"id,omitempty"`
	OrderLineItem                 *OrderLineItem                 `json:"orderLineItem,omitempty"`
	OrderLineItemId               string                         `json:"orderLineItemId,omitempty"`
	OrderLineItemVersionId        *string                        `json:"orderLineItemVersionId,omitempty"`
	OrderTransactionCaptureRefund *OrderTransactionCaptureRefund `json:"orderTransactionCaptureRefund,omitempty"`
	Quantity                      *int                           `json:"quantity,omitempty"`
	Reason                        *string                        `json:"reason,omitempty"`
	RefundId                      string                         `json:"refundId,omitempty"`
	RefundPrice                   *float64                       `json:"refundPrice,omitempty"`
	UpdatedAt                     *time.Time                     `json:"updatedAt,omitempty"`
}
