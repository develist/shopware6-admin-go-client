package entity

import "time"

// completed

type Locale struct {
	Code         string         `json:"code,omitempty"`
	CreatedAt    *time.Time     `json:"createdAt,omitempty"`
	CustomFields *[]CustomField `json:"customFields,omitempty"`
	Id           string         `json:"id,omitempty"`
	Languages    *Language      `json:"languages,omitempty"`
	Name         string         `json:"name,omitempty"`
	Territory    string         `json:"territory,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	Users     *User      `json:"users,omitempty"`
}
