package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-06 18:44:04 UTC

// Integration data structure
// Added since version: 6.0.0.0
// Required fields: label, accessKey, secretAccessKey, createdAt
type Integration struct {
	AccessKey       string         `json:"accessKey,omitempty"`
	AclRoles        *AclRole       `json:"aclRoles,omitempty"`
	Admin           bool           `json:"admin,omitempty"`
	App             *App           `json:"app,omitempty"`
	CreatedAt       *time.Time     `json:"createdAt,omitempty"`
	CustomFields    *[]CustomField `json:"customFields,omitempty"`
	DeletedAt       *time.Time     `json:"deletedAt,omitempty"`
	Id              string         `json:"id,omitempty"`
	Label           string         `json:"label,omitempty"`
	LastUsageAt     *time.Time     `json:"lastUsageAt,omitempty"`
	SecretAccessKey string         `json:"secretAccessKey,omitempty"`
	UpdatedAt       *time.Time     `json:"updatedAt,omitempty"`
	WriteAccess     bool           `json:"writeAccess,omitempty"`
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
