package entity

import "time"

// completed

type Country struct {
	Active                         bool                     `json:"active,omitempty"`
	CheckVatIdPattern              bool                     `json:"checkVatIdPattern,omitempty"`
	CompanyTaxFree                 bool                     `json:"companyTaxFree,omitempty"`
	CreatedAt                      *time.Time               `json:"createdAt,omitempty"`
	CurrencyCountryRoundings       *CurrencyCountryRounding `json:"currencyCountryRoundings,omitempty"`
	CustomFields                   *[]CustomField           `json:"customFields,omitempty"`
	CustomerAddresses              *CustomerAddress         `json:"customerAddresses,omitempty"`
	DisplayStateInRegistration     bool                     `json:"displayStateInRegistration,omitempty"`
	ForceStateInRegistration       bool                     `json:"forceStateInRegistration,omitempty"`
	Id                             string                   `json:"id,omitempty"`
	Iso                            string                   `json:"iso,omitempty"`
	Iso3                           string                   `json:"iso3,omitempty"`
	Name                           string                   `json:"name,omitempty"`
	OrderAddresses                 *OrderAddress            `json:"orderAddresses,omitempty"`
	Position                       int                      `json:"position,omitempty"`
	SalesChannelDefaultAssignments *SalesChannel            `json:"salesChannelDefaultAssignments,omitempty"`
	SalesChannels                  *SalesChannel            `json:"salesChannels,omitempty"`
	ShippingAvailable              bool                     `json:"shippingAvailable,omitempty"`
	States                         *CountryState            `json:"states,omitempty"`
	TaxFree                        bool                     `json:"taxFree,omitempty"`
	TaxRules                       *TaxRule                 `json:"taxRules,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt    *time.Time `json:"updatedAt,omitempty"`
	VatIdPattern string     `json:"vatIdPattern,omitempty"`
}

type CountryState struct {
	Active            bool             `json:"active,omitempty"`
	Country           *Country         `json:"country,omitempty"`
	CountryId         string           `json:"countryId,omitempty"`
	CreatedAt         *time.Time       `json:"createdAt,omitempty"`
	CustomFields      *[]CustomField   `json:"customFields,omitempty"`
	CustomerAddresses *CustomerAddress `json:"customerAddresses,omitempty"`
	Id                string           `json:"id,omitempty"`
	Name              string           `json:"name,omitempty"`
	OrderAddresses    *OrderAddress    `json:"orderAddresses,omitempty"`
	Position          int              `json:"position,omitempty"`
	ShortCode         string           `json:"shortCode,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}
