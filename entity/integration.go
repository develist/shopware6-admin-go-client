package entity

import "time"

// completed

type Integration struct {
	AccessKey       string         `json:"accessKey,omitempty"`
	AclRoles        *AclRole       `json:"aclRoles,omitempty"`
	Admin           bool           `json:"admin,omitempty"`
	App             *App           `json:"app,omitempty"`
	CreatedAt       time.Time      `json:"createdAt,omitempty"`
	CustomFields    *[]CustomField `json:"customFields,omitempty"`
	DeletedAt       time.Time      `json:"deletedAt,omitempty"`
	Id              string         `json:"id,omitempty"`
	Label           string         `json:"label,omitempty"`
	LastUsageAt     time.Time      `json:"lastUsageAt,omitempty"`
	SecretAccessKey string         `json:"secretAccessKey,omitempty"`
	UpdatedAt       time.Time      `json:"updatedAt,omitempty"`
	WriteAccess     bool           `json:"writeAccess,omitempty"`
}
type IntegrationRole struct {
	AclRoleId     string       `json:"aclRoleId,omitempty"`
	Id            string       `json:"id,omitempty"`
	Integration   *Integration `json:"integration,omitempty"`
	IntegrationId string       `json:"integrationId,omitempty"`
	Role          *AclRole     `json:"role,omitempty"`
}
