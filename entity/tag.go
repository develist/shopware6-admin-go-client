package entity

import "time"

// Generated from Shopware Admin API
// Version 6.5.9999999.9999999-dev at 2023-05-02 19:06:34 UTC

// Tag data structure
// Added since version: 6.0.0.0
// Required fields: name, createdAt
type Tag struct {
	Categories           *Category            `json:"categories,omitempty"`
	CreatedAt            *time.Time           `json:"createdAt,omitempty"`
	Customers            *Customer            `json:"customers,omitempty"`
	Id                   string               `json:"id,omitempty"`
	LandingPages         *LandingPage         `json:"landingPages,omitempty"`
	Media                *Media               `json:"media,omitempty"`
	Name                 string               `json:"name,omitempty"`
	NewsletterRecipients *NewsletterRecipient `json:"newsletterRecipients,omitempty"`
	Orders               *Order               `json:"orders,omitempty"`
	Products             *Product             `json:"products,omitempty"`
	Rules                *Rule                `json:"rules,omitempty"`
	ShippingMethods      *ShippingMethod      `json:"shippingMethods,omitempty"`
	UpdatedAt            *time.Time           `json:"updatedAt,omitempty"`
}
