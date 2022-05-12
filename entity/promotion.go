package entity

import "time"

// completed

type Promotion struct {
	Active              bool               `json:"active,omitempty"`
	CartRules           *Rule              `json:"cartRules,omitempty"`
	Code                string             `json:"code,omitempty"`
	CreatedAt           time.Time          `json:"createdAt,omitempty"`
	CustomFields        *[]CustomField     `json:"customFields,omitempty"`
	CustomerRestriction bool               `json:"customerRestriction,omitempty"`
	Discounts           *PromotionDiscount `json:"discounts,omitempty"`
	// ExclusionIds map[items:map[pattern:^[0-9a-f]{32}$ type:string] type:array] `json:"exclusionIds,omitempty"`
	Exclusive                 bool                     `json:"exclusive,omitempty"`
	Id                        string                   `json:"id,omitempty"`
	IndividualCodePattern     string                   `json:"individualCodePattern,omitempty"`
	IndividualCodes           *PromotionIndividualCode `json:"individualCodes,omitempty"`
	MaxRedemptionsGlobal      int                      `json:"maxRedemptionsGlobal,omitempty"`
	MaxRedemptionsPerCustomer int                      `json:"maxRedemptionsPerCustomer,omitempty"`
	Name                      string                   `json:"name,omitempty"`
	OrderCount                int                      `json:"orderCount,omitempty"`
	OrderRules                *Rule                    `json:"orderRules,omitempty"`
	// OrdersPerCustomerCount map[readOnly:true type:object] `json:"ordersPerCustomerCount,omitempty"`
	PersonaCustomers *Customer              `json:"personaCustomers,omitempty"`
	PersonaRules     *Rule                  `json:"personaRules,omitempty"`
	SalesChannels    *PromotionSalesChannel `json:"salesChannels,omitempty"`
	Setgroups        *PromotionSetgroup     `json:"setgroups,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt          time.Time `json:"updatedAt,omitempty"`
	UseCodes           bool      `json:"useCodes,omitempty"`
	UseIndividualCodes bool      `json:"useIndividualCodes,omitempty"`
	UseSetGroups       bool      `json:"useSetGroups,omitempty"`
	ValidFrom          time.Time `json:"validFrom,omitempty"`
	ValidUntil         time.Time `json:"validUntil,omitempty"`
}

type PromotionCartRule struct {
	Id          string     `json:"id,omitempty"`
	Promotion   *Promotion `json:"promotion,omitempty"`
	PromotionId string     `json:"promotionId,omitempty"`
	Rule        *Rule      `json:"rule,omitempty"`
	RuleId      string     `json:"ruleId,omitempty"`
}

type PromotionDiscount struct {
	ApplierKey              string                   `json:"applierKey,omitempty"`
	ConsiderAdvancedRules   bool                     `json:"considerAdvancedRules,omitempty"`
	CreatedAt               time.Time                `json:"createdAt,omitempty"`
	DiscountRules           *Rule                    `json:"discountRules,omitempty"`
	Id                      string                   `json:"id,omitempty"`
	MaxValue                float64                  `json:"maxValue,omitempty"`
	PickerKey               string                   `json:"pickerKey,omitempty"`
	Promotion               *Promotion               `json:"promotion,omitempty"`
	PromotionDiscountPrices *PromotionDiscountPrices `json:"promotionDiscountPrices,omitempty"`
	PromotionId             string                   `json:"promotionId,omitempty"`
	Scope                   string                   `json:"scope,omitempty"`
	SorterKey               string                   `json:"sorterKey,omitempty"`
	Type                    string                   `json:"type,omitempty"`
	UpdatedAt               time.Time                `json:"updatedAt,omitempty"`
	UsageKey                string                   `json:"usageKey,omitempty"`
	Value                   float64                  `json:"value,omitempty"`
}

type PromotionDiscountPrices struct {
	CreatedAt         time.Time          `json:"createdAt,omitempty"`
	Currency          *Currency          `json:"currency,omitempty"`
	CurrencyId        string             `json:"currencyId,omitempty"`
	DiscountId        string             `json:"discountId,omitempty"`
	Id                string             `json:"id,omitempty"`
	Price             float64            `json:"price,omitempty"`
	PromotionDiscount *PromotionDiscount `json:"promotionDiscount,omitempty"`
	UpdatedAt         time.Time          `json:"updatedAt,omitempty"`
}

type PromotionDiscountRule struct {
	Discount   *PromotionDiscount `json:"discount,omitempty"`
	DiscountId string             `json:"discountId,omitempty"`
	Id         string             `json:"id,omitempty"`
	Rule       *Rule              `json:"rule,omitempty"`
	RuleId     string             `json:"ruleId,omitempty"`
}

type PromotionIndividualCode struct {
	Code      string    `json:"code,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Id        string    `json:"id,omitempty"`
	// Payload map[type:object] `json:"payload,omitempty"`
	Promotion   *Promotion `json:"promotion,omitempty"`
	PromotionId string     `json:"promotionId,omitempty"`
	UpdatedAt   time.Time  `json:"updatedAt,omitempty"`
}

type PromotionOrderRule struct {
	Id          string     `json:"id,omitempty"`
	Promotion   *Promotion `json:"promotion,omitempty"`
	PromotionId string     `json:"promotionId,omitempty"`
	Rule        *Rule      `json:"rule,omitempty"`
	RuleId      string     `json:"ruleId,omitempty"`
}

type PromotionPersonaCustomer struct {
	Customer    *Customer  `json:"customer,omitempty"`
	CustomerId  string     `json:"customerId,omitempty"`
	Id          string     `json:"id,omitempty"`
	Promotion   *Promotion `json:"promotion,omitempty"`
	PromotionId string     `json:"promotionId,omitempty"`
}

type PromotionPersonaRule struct {
	Id          string     `json:"id,omitempty"`
	Promotion   *Promotion `json:"promotion,omitempty"`
	PromotionId string     `json:"promotionId,omitempty"`
	Rule        *Rule      `json:"rule,omitempty"`
	RuleId      string     `json:"ruleId,omitempty"`
}

type PromotionSalesChannel struct {
	CreatedAt      time.Time     `json:"createdAt,omitempty"`
	Id             string        `json:"id,omitempty"`
	Priority       int           `json:"priority,omitempty"`
	Promotion      *Promotion    `json:"promotion,omitempty"`
	PromotionId    string        `json:"promotionId,omitempty"`
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId string        `json:"salesChannelId,omitempty"`
	UpdatedAt      time.Time     `json:"updatedAt,omitempty"`
}

type PromotionSetgroup struct {
	CreatedAt     time.Time  `json:"createdAt,omitempty"`
	Id            string     `json:"id,omitempty"`
	PackagerKey   string     `json:"packagerKey,omitempty"`
	Promotion     *Promotion `json:"promotion,omitempty"`
	PromotionId   string     `json:"promotionId,omitempty"`
	SetGroupRules *Rule      `json:"setGroupRules,omitempty"`
	SorterKey     string     `json:"sorterKey,omitempty"`
	UpdatedAt     time.Time  `json:"updatedAt,omitempty"`
	Value         float64    `json:"value,omitempty"`
}

type PromotionSetgroupRule struct {
	Id         string             `json:"id,omitempty"`
	Rule       *Rule              `json:"rule,omitempty"`
	RuleId     string             `json:"ruleId,omitempty"`
	Setgroup   *PromotionSetgroup `json:"setgroup,omitempty"`
	SetgroupId string             `json:"setgroupId,omitempty"`
}
