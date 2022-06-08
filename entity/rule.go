package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-06 18:44:04 UTC

// Rule data structure
// Added since version: 6.0.0.0
// Required fields: name, priority, createdAt
type Rule struct {
	CartPromotions                  *Promotion           `json:"cartPromotions,omitempty"`
	Conditions                      *RuleCondition       `json:"conditions,omitempty"`
	CreatedAt                       *time.Time           `json:"createdAt,omitempty"`
	CustomFields                    *[]CustomField       `json:"customFields,omitempty"`
	Description                     *string              `json:"description,omitempty"`
	EventActions                    *EventAction         `json:"eventActions,omitempty"`
	Id                              *string              `json:"id,omitempty"`
	Invalid                         *bool                `json:"invalid,omitempty"`
	ModuleTypes                     *interface{}         `json:"moduleTypes,omitempty"` // map[type:object]
	Name                            *string              `json:"name,omitempty"`
	OrderPromotions                 *Promotion           `json:"orderPromotions,omitempty"`
	PaymentMethods                  *PaymentMethod       `json:"paymentMethods,omitempty"`
	PersonaPromotions               *Promotion           `json:"personaPromotions,omitempty"`
	Priority                        *int                 `json:"priority,omitempty"`
	ProductPrices                   *ProductPrice        `json:"productPrices,omitempty"`
	PromotionDiscounts              *PromotionDiscount   `json:"promotionDiscounts,omitempty"`
	PromotionSetGroups              *PromotionSetgroup   `json:"promotionSetGroups,omitempty"`
	ShippingMethodPriceCalculations *ShippingMethodPrice `json:"shippingMethodPriceCalculations,omitempty"`
	ShippingMethodPrices            *ShippingMethodPrice `json:"shippingMethodPrices,omitempty"`
	ShippingMethods                 *ShippingMethod      `json:"shippingMethods,omitempty"`
	UpdatedAt                       *time.Time           `json:"updatedAt,omitempty"`
}

// RuleCondition data structure
// Added since version: 6.0.0.0
// Required fields: ruleId, createdAt
type RuleCondition struct {
	Children     *RuleCondition `json:"children,omitempty"`
	CreatedAt    *time.Time     `json:"createdAt,omitempty"`
	CustomFields *[]CustomField `json:"customFields,omitempty"`
	Id           *string        `json:"id,omitempty"`
	Parent       *RuleCondition `json:"parent,omitempty"`
	ParentId     *string        `json:"parentId,omitempty"`
	Position     *int           `json:"position,omitempty"`
	Rule         *Rule          `json:"rule,omitempty"`
	RuleId       *string        `json:"ruleId,omitempty"`
	Type         *string        `json:"type,omitempty"`
	UpdatedAt    *time.Time     `json:"updatedAt,omitempty"`
	Value        *interface{}   `json:"value,omitempty"` // map[type:object]
}
