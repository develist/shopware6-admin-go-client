package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-06 18:44:04 UTC

// LogEntry data structure
// Added since version: 6.0.0.0
// Required fields: createdAt
type LogEntry struct {
	Channel   *string      `json:"channel,omitempty"`
	Context   *interface{} `json:"context,omitempty"` // map[type:object]
	CreatedAt *time.Time   `json:"createdAt,omitempty"`
	Extra     *interface{} `json:"extra,omitempty"` // map[type:object]
	Id        *string      `json:"id,omitempty"`
	Level     *int         `json:"level,omitempty"`
	Message   *string      `json:"message,omitempty"`
	UpdatedAt *time.Time   `json:"updatedAt,omitempty"`
}
