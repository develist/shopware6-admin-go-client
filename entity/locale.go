package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-06 18:44:04 UTC

// Locale data structure
// Added since version: 6.0.0.0
// Required fields: code, createdAt, name, territory
type Locale struct {
	Code         string         `json:"code,omitempty"`
	CreatedAt    *time.Time     `json:"createdAt,omitempty"`
	CustomFields *[]CustomField `json:"customFields,omitempty"`
	Id           string         `json:"id,omitempty"`
	Languages    *Language      `json:"languages,omitempty"`
	Name         string         `json:"name,omitempty"`
	Territory    string         `json:"territory,omitempty"`
	Translated   *interface{}   `json:"translated,omitempty"` // map[type:object]
	UpdatedAt    *time.Time     `json:"updatedAt,omitempty"`
	Users        *User          `json:"users,omitempty"`
}
