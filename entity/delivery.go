package entity

import "time"

// Generated from Shopware Admin API
// Version 6.5.9999999.9999999-dev at 2023-05-02 19:06:34 UTC

// DeliveryTime data structure
// Added since version: 6.0.0.0
// Required fields: name, min, max, unit, createdAt
type DeliveryTime struct {
	CreatedAt       *time.Time      `json:"createdAt,omitempty"`
	CustomFields    *[]CustomField  `json:"customFields,omitempty"`
	Id              string          `json:"id,omitempty"`
	Max             int             `json:"max,omitempty"`
	Min             int             `json:"min,omitempty"`
	Name            string          `json:"name,omitempty"`
	Products        *Product        `json:"products,omitempty"`
	ShippingMethods *ShippingMethod `json:"shippingMethods,omitempty"`
	Translated      *interface{}    `json:"translated,omitempty"` // map[type:object]
	Unit            string          `json:"unit,omitempty"`
	UpdatedAt       *time.Time      `json:"updatedAt,omitempty"`
}
