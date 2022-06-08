package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-06 18:44:04 UTC

// Promotion data structure
// Added since version: 6.0.0.0
// Required fields: active, exclusive, useCodes, useIndividualCodes, useSetGroups, createdAt, name
type Promotion struct {
	Active                    *bool                    `json:"active,omitempty"`
	CartRules                 *Rule                    `json:"cartRules,omitempty"`
	Code                      *string                  `json:"code,omitempty"`
	CreatedAt                 *time.Time               `json:"createdAt,omitempty"`
	CustomFields              *[]CustomField           `json:"customFields,omitempty"`
	CustomerRestriction       *bool                    `json:"customerRestriction,omitempty"`
	Discounts                 *PromotionDiscount       `json:"discounts,omitempty"`
	ExclusionIds              *[]string                `json:"exclusionIds,omitempty"` // map[items:map[pattern:^[0-9a-f]{32}$ type:string] type:array]
	Exclusive                 *bool                    `json:"exclusive,omitempty"`
	Id                        *string                  `json:"id,omitempty"`
	IndividualCodePattern     *string                  `json:"individualCodePattern,omitempty"`
	IndividualCodes           *PromotionIndividualCode `json:"individualCodes,omitempty"`
	MaxRedemptionsGlobal      *int                     `json:"maxRedemptionsGlobal,omitempty"`
	MaxRedemptionsPerCustomer *int                     `json:"maxRedemptionsPerCustomer,omitempty"`
	Name                      *string                  `json:"name,omitempty"`
	OrderCount                *int                     `json:"orderCount,omitempty"`
	OrderRules                *Rule                    `json:"orderRules,omitempty"`
	OrdersPerCustomerCount    *interface{}             `json:"ordersPerCustomerCount,omitempty"` // map[readOnly:true type:object]
	PersonaCustomers          *Customer                `json:"personaCustomers,omitempty"`
	PersonaRules              *Rule                    `json:"personaRules,omitempty"`
	SalesChannels             *PromotionSalesChannel   `json:"salesChannels,omitempty"`
	Setgroups                 *PromotionSetgroup       `json:"setgroups,omitempty"`
	Translated                *interface{}             `json:"translated,omitempty"` // map[type:object]
	UpdatedAt                 *time.Time               `json:"updatedAt,omitempty"`
	UseCodes                  *bool                    `json:"useCodes,omitempty"`
	UseIndividualCodes        *bool                    `json:"useIndividualCodes,omitempty"`
	UseSetGroups              *bool                    `json:"useSetGroups,omitempty"`
	ValidFrom                 *time.Time               `json:"validFrom,omitempty"`
	ValidUntil                *time.Time               `json:"validUntil,omitempty"`
}

// PromotionCartRule data structure
// Added since version: 6.0.0.0
// Required fields: promotionId, ruleId
type PromotionCartRule struct {
	Id          *string    `json:"id,omitempty"`
	Promotion   *Promotion `json:"promotion,omitempty"`
	PromotionId *string    `json:"promotionId,omitempty"`
	Rule        *Rule      `json:"rule,omitempty"`
	RuleId      *string    `json:"ruleId,omitempty"`
}

// PromotionDiscount data structure
// Added since version: 6.0.0.0
// Required fields: promotionId, scope, type, value, considerAdvancedRules, createdAt
type PromotionDiscount struct {
	ApplierKey              *string                  `json:"applierKey,omitempty"`
	ConsiderAdvancedRules   *bool                    `json:"considerAdvancedRules,omitempty"`
	CreatedAt               *time.Time               `json:"createdAt,omitempty"`
	DiscountRules           *Rule                    `json:"discountRules,omitempty"`
	Id                      *string                  `json:"id,omitempty"`
	MaxValue                *float64                 `json:"maxValue,omitempty"`
	PickerKey               *string                  `json:"pickerKey,omitempty"`
	Promotion               *Promotion               `json:"promotion,omitempty"`
	PromotionDiscountPrices *PromotionDiscountPrices `json:"promotionDiscountPrices,omitempty"`
	PromotionId             *string                  `json:"promotionId,omitempty"`
	Scope                   *string                  `json:"scope,omitempty"`
	SorterKey               *string                  `json:"sorterKey,omitempty"`
	Type                    *string                  `json:"type,omitempty"`
	UpdatedAt               *time.Time               `json:"updatedAt,omitempty"`
	UsageKey                *string                  `json:"usageKey,omitempty"`
	Value                   *float64                 `json:"value,omitempty"`
}

// PromotionDiscountPrices data structure
// Added since version: 6.0.0.0
// Required fields: discountId, currencyId, price, createdAt
type PromotionDiscountPrices struct {
	CreatedAt         *time.Time         `json:"createdAt,omitempty"`
	Currency          *Currency          `json:"currency,omitempty"`
	CurrencyId        *string            `json:"currencyId,omitempty"`
	DiscountId        *string            `json:"discountId,omitempty"`
	Id                *string            `json:"id,omitempty"`
	Price             *float64           `json:"price,omitempty"`
	PromotionDiscount *PromotionDiscount `json:"promotionDiscount,omitempty"`
	UpdatedAt         *time.Time         `json:"updatedAt,omitempty"`
}

// PromotionDiscountRule data structure
// Added since version: 6.0.0.0
// Required fields: discountId, ruleId
type PromotionDiscountRule struct {
	Discount   *PromotionDiscount `json:"discount,omitempty"`
	DiscountId *string            `json:"discountId,omitempty"`
	Id         *string            `json:"id,omitempty"`
	Rule       *Rule              `json:"rule,omitempty"`
	RuleId     *string            `json:"ruleId,omitempty"`
}

// PromotionIndividualCode data structure
// Added since version: 6.0.0.0
// Required fields: promotionId, code, createdAt
type PromotionIndividualCode struct {
	Code        *string      `json:"code,omitempty"`
	CreatedAt   *time.Time   `json:"createdAt,omitempty"`
	Id          *string      `json:"id,omitempty"`
	Payload     *interface{} `json:"payload,omitempty"` // map[type:object]
	Promotion   *Promotion   `json:"promotion,omitempty"`
	PromotionId *string      `json:"promotionId,omitempty"`
	UpdatedAt   *time.Time   `json:"updatedAt,omitempty"`
}

// PromotionOrderRule data structure
// Added since version: 6.0.0.0
// Required fields: promotionId, ruleId
type PromotionOrderRule struct {
	Id          *string    `json:"id,omitempty"`
	Promotion   *Promotion `json:"promotion,omitempty"`
	PromotionId *string    `json:"promotionId,omitempty"`
	Rule        *Rule      `json:"rule,omitempty"`
	RuleId      *string    `json:"ruleId,omitempty"`
}

// PromotionPersonaCustomer data structure
// Added since version: 6.0.0.0
// Required fields: promotionId, customerId
type PromotionPersonaCustomer struct {
	Customer    *Customer  `json:"customer,omitempty"`
	CustomerId  *string    `json:"customerId,omitempty"`
	Id          *string    `json:"id,omitempty"`
	Promotion   *Promotion `json:"promotion,omitempty"`
	PromotionId *string    `json:"promotionId,omitempty"`
}

// PromotionPersonaRule data structure
// Added since version: 6.0.0.0
// Required fields: promotionId, ruleId
type PromotionPersonaRule struct {
	Id          *string    `json:"id,omitempty"`
	Promotion   *Promotion `json:"promotion,omitempty"`
	PromotionId *string    `json:"promotionId,omitempty"`
	Rule        *Rule      `json:"rule,omitempty"`
	RuleId      *string    `json:"ruleId,omitempty"`
}

// PromotionSalesChannel data structure
// Added since version: 6.0.0.0
// Required fields: promotionId, salesChannelId, priority, createdAt
type PromotionSalesChannel struct {
	CreatedAt      *time.Time    `json:"createdAt,omitempty"`
	Id             *string       `json:"id,omitempty"`
	Priority       *int          `json:"priority,omitempty"`
	Promotion      *Promotion    `json:"promotion,omitempty"`
	PromotionId    *string       `json:"promotionId,omitempty"`
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId *string       `json:"salesChannelId,omitempty"`
	UpdatedAt      *time.Time    `json:"updatedAt,omitempty"`
}

// PromotionSetgroup data structure
// Added since version: 6.0.0.0
// Required fields: promotionId, packagerKey, sorterKey, value, createdAt
type PromotionSetgroup struct {
	CreatedAt     *time.Time `json:"createdAt,omitempty"`
	Id            *string    `json:"id,omitempty"`
	PackagerKey   *string    `json:"packagerKey,omitempty"`
	Promotion     *Promotion `json:"promotion,omitempty"`
	PromotionId   *string    `json:"promotionId,omitempty"`
	SetGroupRules *Rule      `json:"setGroupRules,omitempty"`
	SorterKey     *string    `json:"sorterKey,omitempty"`
	UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
	Value         *float64   `json:"value,omitempty"`
}

// PromotionSetgroupRule data structure
// Added since version: 6.0.0.0
// Required fields: setgroupId, ruleId
type PromotionSetgroupRule struct {
	Id         *string            `json:"id,omitempty"`
	Rule       *Rule              `json:"rule,omitempty"`
	RuleId     *string            `json:"ruleId,omitempty"`
	Setgroup   *PromotionSetgroup `json:"setgroup,omitempty"`
	SetgroupId *string            `json:"setgroupId,omitempty"`
}
