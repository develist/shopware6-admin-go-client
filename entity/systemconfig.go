package entity

import "time"

// Generated from Shopware Admin API
// Version 6.5.9999999.9999999-dev at 2023-05-02 19:06:34 UTC

// SystemConfig data structure
// Added since version: 6.0.0.0
// Required fields: configurationKey, configurationValue, createdAt
type SystemConfig struct {
	ConfigurationKey   string        `json:"configurationKey,omitempty"`
	ConfigurationValue *interface{}  `json:"configurationValue,omitempty"` // map[properties:map[_value:map[type:object]] type:object]
	CreatedAt          *time.Time    `json:"createdAt,omitempty"`
	Id                 string        `json:"id,omitempty"`
	SalesChannel       *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId     *string       `json:"salesChannelId,omitempty"`
	UpdatedAt          *time.Time    `json:"updatedAt,omitempty"`
}
