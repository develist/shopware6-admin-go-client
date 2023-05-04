package entity

import "time"

// Generated from Shopware Admin API
// Version 6.5.9999999.9999999-dev at 2023-05-02 19:06:34 UTC

// AclRole data structure
// Added since version: 6.0.0.0
// Required fields: name, privileges, createdAt
type AclRole struct {
	App          *App         `json:"app,omitempty"`
	CreatedAt    *time.Time   `json:"createdAt,omitempty"`
	DeletedAt    *time.Time   `json:"deletedAt,omitempty"`
	Description  *string      `json:"description,omitempty"`
	Id           string       `json:"id,omitempty"`
	Integrations *Integration `json:"integrations,omitempty"`
	Name         string       `json:"name,omitempty"`
	Privileges   *interface{} `json:"privileges,omitempty"` // map[items:map[additionalProperties:false] type:array]
	UpdatedAt    *time.Time   `json:"updatedAt,omitempty"`
	Users        *User        `json:"users,omitempty"`
}

// AclUserRole data structure
// Added since version: 6.0.0.0
// Required fields: userId, aclRoleId, createdAt
type AclUserRole struct {
	AclRole   *AclRole   `json:"aclRole,omitempty"`
	AclRoleId string     `json:"aclRoleId,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Id        string     `json:"id,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	User      *User      `json:"user,omitempty"`
	UserId    string     `json:"userId,omitempty"`
}
