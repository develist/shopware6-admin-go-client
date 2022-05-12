package entity

import "time"

// completed

type AclRole struct {
	App          *App         `json:"app,omitempty"`
	CreatedAt    time.Time    `json:"createdAt,omitempty"`
	DeletedAt    time.Time    `json:"deletedAt,omitempty"`
	Description  string       `json:"description,omitempty"`
	Id           string       `json:"id,omitempty"`
	Integrations *Integration `json:"integrations,omitempty"`
	Name         string       `json:"name,omitempty"`
	// Privileges map[items:map[additionalProperties:false] type:array] `json:"privileges,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	Users     *User     `json:"users,omitempty"`
}

type AclUserRole struct {
	AclRole   *AclRole  `json:"aclRole,omitempty"`
	AclRoleId string    `json:"aclRoleId,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Id        string    `json:"id,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	User      *User     `json:"user,omitempty"`
	UserId    string    `json:"userId,omitempty"`
}
