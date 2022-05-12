package entity

import "time"

// completed

type Currency struct {
	CountryRoundings *CurrencyCountryRounding `json:"countryRoundings,omitempty"`
	CreatedAt        time.Time                `json:"createdAt,omitempty"`
	CustomFields     *[]CustomField           `json:"customFields,omitempty"`
	Factor           float64                  `json:"factor,omitempty"`
	Id               string                   `json:"id,omitempty"`
	IsSystemDefault  bool                     `json:"isSystemDefault,omitempty"`
	IsoCode          string                   `json:"isoCode,omitempty"`
	// ItemRounding map[properties:map[decimals:map[format:int64 type:integer] interval:map[format:float type:number] roundForNet:map[type:boolean]] type:object] `json:"itemRounding,omitempty"`
	Name                           string                   `json:"name,omitempty"`
	Orders                         *Order                   `json:"orders,omitempty"`
	Position                       int                      `json:"position,omitempty"`
	ProductExports                 *ProductExport           `json:"productExports,omitempty"`
	PromotionDiscountPrices        *PromotionDiscountPrices `json:"promotionDiscountPrices,omitempty"`
	SalesChannelDefaultAssignments *SalesChannel            `json:"salesChannelDefaultAssignments,omitempty"`
	SalesChannelDomains            *SalesChannelDomain      `json:"salesChannelDomains,omitempty"`
	SalesChannels                  *SalesChannel            `json:"salesChannels,omitempty"`
	ShortName                      string                   `json:"shortName,omitempty"`
	Symbol                         string                   `json:"symbol,omitempty"`
	// TotalRounding map[properties:map[decimals:map[format:int64 type:integer] interval:map[format:float type:number] roundForNet:map[type:boolean]] type:object] `json:"totalRounding,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type CurrencyCountryRounding struct {
	Country    *Country  `json:"country,omitempty"`
	CountryId  string    `json:"countryId,omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	Currency   *Currency `json:"currency,omitempty"`
	CurrencyId string    `json:"currencyId,omitempty"`
	Id         string    `json:"id,omitempty"`
	// ItemRounding map[properties:map[decimals:map[format:int64 type:integer] interval:map[format:float type:number] roundForNet:map[type:boolean]] type:object] `json:"itemRounding,omitempty"`
	// TotalRounding map[properties:map[decimals:map[format:int64 type:integer] interval:map[format:float type:number] roundForNet:map[type:boolean]] type:object] `json:"totalRounding,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
