package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-06 18:44:04 UTC

// MessageQueueStats data structure
// Added since version: 6.0.0.0
// Required fields: name, size, createdAt
type MessageQueueStats struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Id        *string    `json:"id,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Size      *int       `json:"size,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}
