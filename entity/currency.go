package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-06 18:44:04 UTC

// Currency data structure
// Added since version: 6.0.0.0
// Required fields: factor, symbol, isoCode, itemRounding, totalRounding, createdAt, shortName, name
type Currency struct {
	CountryRoundings               *CurrencyCountryRounding `json:"countryRoundings,omitempty"`
	CreatedAt                      *time.Time               `json:"createdAt,omitempty"`
	CustomFields                   *[]CustomField           `json:"customFields,omitempty"`
	Factor                         *float64                 `json:"factor,omitempty"`
	Id                             *string                  `json:"id,omitempty"`
	IsSystemDefault                *bool                    `json:"isSystemDefault,omitempty"`
	IsoCode                        *string                  `json:"isoCode,omitempty"`
	ItemRounding                   *Rounding                `json:"itemRounding,omitempty"` // map[properties:map[decimals:map[format:*int64 type:*integer] *interval:map[format:*float type:number] roundForNet:map[type:*boolean]] type:object]
	Name                           *string                  `json:"name,omitempty"`
	Orders                         *Order                   `json:"orders,omitempty"`
	Position                       *int                     `json:"position,omitempty"`
	ProductExports                 *ProductExport           `json:"productExports,omitempty"`
	PromotionDiscountPrices        *PromotionDiscountPrices `json:"promotionDiscountPrices,omitempty"`
	SalesChannelDefaultAssignments *SalesChannel            `json:"salesChannelDefaultAssignments,omitempty"`
	SalesChannelDomains            *SalesChannelDomain      `json:"salesChannelDomains,omitempty"`
	SalesChannels                  *SalesChannel            `json:"salesChannels,omitempty"`
	ShortName                      *string                  `json:"shortName,omitempty"`
	Symbol                         *string                  `json:"symbol,omitempty"`
	TotalRounding                  *Rounding                `json:"totalRounding,omitempty"` // map[properties:map[decimals:map[format:*int64 type:*integer] *interval:map[format:*float type:number] roundForNet:map[type:*boolean]] type:object]
	Translated                     *interface{}             `json:"translated,omitempty"`    // map[type:object]
	UpdatedAt                      *time.Time               `json:"updatedAt,omitempty"`
}

// CurrencyCountryRounding data structure
// Added since version: 6.4.0.0
// Required fields: currencyId, countryId, itemRounding, totalRounding, createdAt
type CurrencyCountryRounding struct {
	Country       *Country   `json:"country,omitempty"`
	CountryId     *string    `json:"countryId,omitempty"`
	CreatedAt     *time.Time `json:"createdAt,omitempty"`
	Currency      *Currency  `json:"currency,omitempty"`
	CurrencyId    *string    `json:"currencyId,omitempty"`
	Id            *string    `json:"id,omitempty"`
	ItemRounding  *Rounding  `json:"itemRounding,omitempty"`  // map[properties:map[decimals:map[format:*int64 type:*integer] *interval:map[format:*float type:number] roundForNet:map[type:*boolean]] type:object]
	TotalRounding *Rounding  `json:"totalRounding,omitempty"` // map[properties:map[decimals:map[format:*int64 type:*integer] *interval:map[format:*float type:number] roundForNet:map[type:*boolean]] type:object]
	UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
}

// Rounding helper data structure
type Rounding struct {
	Decimals    *int     `json:"decimals,omitempty"`
	Interval    *float64 `json:"*interval,omitempty"`
	RoundForNet *bool    `json:"roundForNet,omitempty"`
}
