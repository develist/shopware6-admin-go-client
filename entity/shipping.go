package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-06 18:44:04 UTC

// ShippingMethod data structure
// Added since version: 6.0.0.0
// Required fields: availabilityRuleId, deliveryTimeId, taxType, createdAt, name
type ShippingMethod struct {
	Active                         bool                 `json:"active,omitempty"`
	AvailabilityRule               *Rule                `json:"availabilityRule,omitempty"`
	AvailabilityRuleId             string               `json:"availabilityRuleId,omitempty"`
	CreatedAt                      *time.Time           `json:"createdAt,omitempty"`
	CustomFields                   *[]CustomField       `json:"customFields,omitempty"`
	DeliveryTime                   *DeliveryTime        `json:"deliveryTime,omitempty"`
	DeliveryTimeId                 string               `json:"deliveryTimeId,omitempty"`
	Description                    string               `json:"description,omitempty"`
	Id                             string               `json:"id,omitempty"`
	Media                          *Media               `json:"media,omitempty"`
	MediaId                        string               `json:"mediaId,omitempty"`
	Name                           string               `json:"name,omitempty"`
	OrderDeliveries                *OrderDelivery       `json:"orderDeliveries,omitempty"`
	Prices                         *ShippingMethodPrice `json:"prices,omitempty"`
	SalesChannelDefaultAssignments *SalesChannel        `json:"salesChannelDefaultAssignments,omitempty"`
	SalesChannels                  *SalesChannel        `json:"salesChannels,omitempty"`
	Tags                           *Tag                 `json:"tags,omitempty"`
	Tax                            *Tax                 `json:"tax,omitempty"`
	TaxId                          string               `json:"taxId,omitempty"`
	TaxType                        string               `json:"taxType,omitempty"`
	TrackingUrl                    string               `json:"trackingUrl,omitempty"`
	Translated                     *interface{}         `json:"translated,omitempty"` // map[type:object]
	UpdatedAt                      *time.Time           `json:"updatedAt,omitempty"`
}

// ShippingMethodPrice data structure
// Added since version: 6.0.0.0
// Required fields: shippingMethodId, createdAt
type ShippingMethodPrice struct {
	Calculation       int             `json:"calculation,omitempty"`
	CalculationRule   *Rule           `json:"calculationRule,omitempty"`
	CalculationRuleId string          `json:"calculationRuleId,omitempty"`
	CreatedAt         *time.Time      `json:"createdAt,omitempty"`
	CurrencyPrice     *interface{}    `json:"currencyPrice,omitempty"` // map[type:object]
	CustomFields      *[]CustomField  `json:"customFields,omitempty"`
	Id                string          `json:"id,omitempty"`
	QuantityEnd       float64         `json:"quantityEnd,omitempty"`
	QuantityStart     float64         `json:"quantityStart,omitempty"`
	Rule              *Rule           `json:"rule,omitempty"`
	RuleId            string          `json:"ruleId,omitempty"`
	ShippingMethod    *ShippingMethod `json:"shippingMethod,omitempty"`
	ShippingMethodId  string          `json:"shippingMethodId,omitempty"`
	UpdatedAt         *time.Time      `json:"updatedAt,omitempty"`
}

// ShippingMethodTag data structure
// Added since version: 6.0.0.0
// Required fields: shippingMethodId, tagId
type ShippingMethodTag struct {
	Id               string          `json:"id,omitempty"`
	ShippingMethod   *ShippingMethod `json:"shippingMethod,omitempty"`
	ShippingMethodId string          `json:"shippingMethodId,omitempty"`
	Tag              *Tag            `json:"tag,omitempty"`
	TagId            string          `json:"tagId,omitempty"`
}
