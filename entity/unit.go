package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-06 18:44:04 UTC

// Unit data structure
// Added since version: 6.0.0.0
// Required fields: createdAt, shortCode, name
type Unit struct {
	CreatedAt    *time.Time     `json:"createdAt,omitempty"`
	CustomFields *[]CustomField `json:"customFields,omitempty"`
	Id           string         `json:"id,omitempty"`
	Name         string         `json:"name,omitempty"`
	Products     *Product       `json:"products,omitempty"`
	ShortCode    string         `json:"shortCode,omitempty"`
	Translated   *interface{}   `json:"translated,omitempty"` // map[type:object]
	UpdatedAt    *time.Time     `json:"updatedAt,omitempty"`
}
