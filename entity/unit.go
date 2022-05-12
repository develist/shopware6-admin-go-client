package entity

import "time"

// completed

type Unit struct {
	CreatedAt    time.Time      `json:"createdAt,omitempty"`
	CustomFields *[]CustomField `json:"customFields,omitempty"`
	Id           string         `json:"id,omitempty"`
	Name         string         `json:"name,omitempty"`
	Products     *Product       `json:"products,omitempty"`
	ShortCode    string         `json:"shortCode,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
