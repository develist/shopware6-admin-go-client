package entity

import "time"

// completed

type SystemConfig struct {
	ConfigurationKey string `json:"configurationKey,omitempty"`
	// ConfigurationValue map[properties:map[_value:map[type:object]] type:object] `json:"configurationValue,omitempty"`
	CreatedAt      *time.Time    `json:"createdAt,omitempty"`
	Id             string        `json:"id,omitempty"`
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId string        `json:"salesChannelId,omitempty"`
	UpdatedAt      *time.Time    `json:"updatedAt,omitempty"`
}
