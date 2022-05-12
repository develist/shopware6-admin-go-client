package entity

import "time"

// completed

type DeliveryTime struct {
	CreatedAt       time.Time       `json:"createdAt,omitempty"`
	CustomFields    *[]CustomField  `json:"customFields,omitempty"`
	Id              string          `json:"id,omitempty"`
	Max             int             `json:"max,omitempty"`
	Min             int             `json:"min,omitempty"`
	Name            string          `json:"name,omitempty"`
	Products        *Product        `json:"products,omitempty"`
	ShippingMethods *ShippingMethod `json:"shippingMethods,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	Unit      string    `json:"unit,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
