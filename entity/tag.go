package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-06 18:44:04 UTC

// Tag data structure
// Added since version: 6.0.0.0
// Required fields: name, createdAt
type Tag struct {
	Categories           *Category            `json:"categories,omitempty"`
	CreatedAt            *time.Time           `json:"createdAt,omitempty"`
	Customers            *Customer            `json:"customers,omitempty"`
	Id                   *string              `json:"id,omitempty"`
	LandingPages         *LandingPage         `json:"landingPages,omitempty"`
	Media                *Media               `json:"media,omitempty"`
	Name                 *string              `json:"name,omitempty"`
	NewsletterRecipients *NewsletterRecipient `json:"newsletterRecipients,omitempty"`
	Orders               *Order               `json:"orders,omitempty"`
	Products             *Product             `json:"products,omitempty"`
	ShippingMethods      *ShippingMethod      `json:"shippingMethods,omitempty"`
	UpdatedAt            *time.Time           `json:"updatedAt,omitempty"`
}
