package entity

import "time"

// completed

type LogEntry struct {
	Channel string `json:"channel,omitempty"`
	// Context map[type:object] `json:"context,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Extra map[type:object] `json:"extra,omitempty"`
	Id        string     `json:"id,omitempty"`
	Level     int        `json:"level,omitempty"`
	Message   string     `json:"message,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}
