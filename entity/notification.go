package entity

import "time"

// Generated from Shopware Admin API
// Version 6.5.9999999.9999999-dev at 2023-05-02 19:06:34 UTC

// Notification data structure
// Added since version: 6.4.7.0
// Required fields: status, message, createdAt
type Notification struct {
	AdminOnly              *bool        `json:"adminOnly,omitempty"`
	CreatedAt              *time.Time   `json:"createdAt,omitempty"`
	CreatedByIntegration   *Integration `json:"createdByIntegration,omitempty"`
	CreatedByIntegrationId *string      `json:"createdByIntegrationId,omitempty"`
	CreatedByUser          *User        `json:"createdByUser,omitempty"`
	CreatedByUserId        *string      `json:"createdByUserId,omitempty"`
	Id                     string       `json:"id,omitempty"`
	Message                string       `json:"message,omitempty"`
	RequiredPrivileges     *interface{} `json:"requiredPrivileges,omitempty"` // map[items:map[additionalProperties:false] type:array]
	Status                 string       `json:"status,omitempty"`
	UpdatedAt              *time.Time   `json:"updatedAt,omitempty"`
}
