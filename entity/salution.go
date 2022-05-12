package entity

import "time"

// completed

type Salutation struct {
	CreatedAt            time.Time            `json:"createdAt,omitempty"`
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
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
