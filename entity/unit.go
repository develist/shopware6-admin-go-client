package entity

import "time"

// Generated from Shopware Admin API
// Version 6.5.9999999.9999999-dev at 2023-05-02 19:06:34 UTC

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
