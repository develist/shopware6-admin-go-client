package entity

import "time"

// Generated from Shopware Admin API
// Version 6.5.9999999.9999999-dev at 2023-05-02 19:06:34 UTC

// Integration data structure
// Added since version: 6.0.0.0
// Required fields: label, accessKey, secretAccessKey, createdAt
type Integration struct {
	AccessKey       string         `json:"accessKey,omitempty"`
	AclRoles        *AclRole       `json:"aclRoles,omitempty"`
	Admin           *bool          `json:"admin,omitempty"`
	App             *App           `json:"app,omitempty"`
	CreatedAt       *time.Time     `json:"createdAt,omitempty"`
	CustomFields    *[]CustomField `json:"customFields,omitempty"`
	DeletedAt       *time.Time     `json:"deletedAt,omitempty"`
	Extensions      *interface{}   `json:"extensions,omitempty"` // map[properties:map[createdNotifications:map[properties:map[data:map[items:map[properties:map[id:map[example:6b29a6713fed4b80b4a81f56caf2b623 type:string] type:map[example:notification type:string]] type:object] type:array] links:map[properties:map[related:map[example:/integration/257ee728107d41f5b260762a3ba7e6de/createdNotifications format:uri-reference type:string]] type:object]] type:object]] type:object]
	Id              string         `json:"id,omitempty"`
	Label           string         `json:"label,omitempty"`
	LastUsageAt     *time.Time     `json:"lastUsageAt,omitempty"`
	SecretAccessKey string         `json:"secretAccessKey,omitempty"`
	UpdatedAt       *time.Time     `json:"updatedAt,omitempty"`
	WriteAccess     *bool          `json:"writeAccess,omitempty"`
}

// IntegrationRole data structure
// Added since version: 6.3.3.0
// Required fields: integrationId, aclRoleId
type IntegrationRole struct {
	AclRoleId     string       `json:"aclRoleId,omitempty"`
	Id            string       `json:"id,omitempty"`
	Integration   *Integration `json:"integration,omitempty"`
	IntegrationId string       `json:"integrationId,omitempty"`
	Role          *AclRole     `json:"role,omitempty"`
}
