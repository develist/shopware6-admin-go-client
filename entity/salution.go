package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-06 18:44:04 UTC

// Salutation data structure
// Added since version: 6.0.0.0
// Required fields: salutationKey, createdAt, displayName, letterName
type Salutation struct {
	CreatedAt            *time.Time           `json:"createdAt,omitempty"`
	CustomFields         *[]CustomField       `json:"customFields,omitempty"`
	CustomerAddresses    *CustomerAddress     `json:"customerAddresses,omitempty"`
	Customers            *Customer            `json:"customers,omitempty"`
	DisplayName          string               `json:"displayName,omitempty"`
	Id                   string               `json:"id,omitempty"`
	LetterName           string               `json:"letterName,omitempty"`
	NewsletterRecipients *NewsletterRecipient `json:"newsletterRecipients,omitempty"`
	OrderAddresses       *OrderAddress        `json:"orderAddresses,omitempty"`
	OrderCustomers       *OrderCustomer       `json:"orderCustomers,omitempty"`
	SalutationKey        string               `json:"salutationKey,omitempty"`
	Translated           *interface{}         `json:"translated,omitempty"` // map[type:object]
	UpdatedAt            *time.Time           `json:"updatedAt,omitempty"`
}
