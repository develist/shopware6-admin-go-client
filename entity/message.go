package entity

import "time"

// completed

type MessageQueueStats struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Id        string     `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Size      int        `json:"size,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}
