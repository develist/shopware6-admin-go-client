package entity

import "time"

// completed

type Rule struct {
	CartPromotions *Promotion     `json:"cartPromotions,omitempty"`
	Conditions     *RuleCondition `json:"conditions,omitempty"`
	CreatedAt      *time.Time     `json:"createdAt,omitempty"`
	CustomFields   *[]CustomField `json:"customFields,omitempty"`
	Description    string         `json:"description,omitempty"`
	EventActions   *EventAction   `json:"eventActions,omitempty"`
	Id             string         `json:"id,omitempty"`
	Invalid        bool           `json:"invalid,omitempty"`
	// ModuleTypes map[type:object] `json:"moduleTypes,omitempty"`
	Name                            string               `json:"name,omitempty"`
	OrderPromotions                 *Promotion           `json:"orderPromotions,omitempty"`
	PaymentMethods                  *PaymentMethod       `json:"paymentMethods,omitempty"`
	PersonaPromotions               *Promotion           `json:"personaPromotions,omitempty"`
	Priority                        int                  `json:"priority,omitempty"`
	ProductPrices                   *ProductPrice        `json:"productPrices,omitempty"`
	PromotionDiscounts              *PromotionDiscount   `json:"promotionDiscounts,omitempty"`
	PromotionSetGroups              *PromotionSetgroup   `json:"promotionSetGroups,omitempty"`
	ShippingMethodPriceCalculations *ShippingMethodPrice `json:"shippingMethodPriceCalculations,omitempty"`
	ShippingMethodPrices            *ShippingMethodPrice `json:"shippingMethodPrices,omitempty"`
	ShippingMethods                 *ShippingMethod      `json:"shippingMethods,omitempty"`
	UpdatedAt                       *time.Time           `json:"updatedAt,omitempty"`
}

type RuleCondition struct {
	Children     *RuleCondition `json:"children,omitempty"`
	CreatedAt    *time.Time     `json:"createdAt,omitempty"`
	CustomFields *[]CustomField `json:"customFields,omitempty"`
	Id           string         `json:"id,omitempty"`
	Parent       *RuleCondition `json:"parent,omitempty"`
	ParentId     string         `json:"parentId,omitempty"`
	Position     int            `json:"position,omitempty"`
	Rule         *Rule          `json:"rule,omitempty"`
	RuleId       string         `json:"ruleId,omitempty"`
	Type         string         `json:"type,omitempty"`
	UpdatedAt    *time.Time     `json:"updatedAt,omitempty"`
	// Value map[type:object] `json:"value,omitempty"`
}
