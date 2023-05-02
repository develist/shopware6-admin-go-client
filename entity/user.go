package entity

import "time"

// Generated from Shopware Admin API
// Version 6.5.9999999.9999999-dev at 2023-05-02 19:06:34 UTC

// User data structure
// Added since version: 6.0.0.0
// Required fields: localeId, username, firstName, lastName, email, timeZone, createdAt
type User struct {
	AccessKeys                 *UserAccessKey       `json:"accessKeys,omitempty"`
	AclRoles                   *AclRole             `json:"aclRoles,omitempty"`
	Active                     *bool                `json:"active,omitempty"`
	Admin                      *bool                `json:"admin,omitempty"`
	AvatarId                   *string              `json:"avatarId,omitempty"`
	AvatarMedia                *Media               `json:"avatarMedia,omitempty"`
	Configs                    *UserConfig          `json:"configs,omitempty"`
	CreatedAt                  *time.Time           `json:"createdAt,omitempty"`
	CreatedCustomers           *Customer            `json:"createdCustomers,omitempty"`
	CreatedOrders              *Order               `json:"createdOrders,omitempty"`
	CustomFields               *[]CustomField       `json:"customFields,omitempty"`
	Email                      string               `json:"email,omitempty"`
	Extensions                 *interface{}         `json:"extensions,omitempty"` // map[properties:map[createdNotifications:map[properties:map[data:map[items:map[properties:map[id:map[example:122aedc28e5e496ba5041f9725147253 type:string] type:map[example:notification type:string]] type:object] type:array] links:map[properties:map[related:map[example:/user/3129263dae954c8bbba1ab4a779295ee/createdNotifications format:uri-reference type:string]] type:object]] type:object]] type:object]
	FirstName                  string               `json:"firstName,omitempty"`
	Id                         string               `json:"id,omitempty"`
	ImportExportLogEntries     *ImportExportLog     `json:"importExportLogEntries,omitempty"`
	LastName                   string               `json:"lastName,omitempty"`
	LastUpdatedPasswordAt      *time.Time           `json:"lastUpdatedPasswordAt,omitempty"`
	Locale                     *Locale              `json:"locale,omitempty"`
	LocaleId                   string               `json:"localeId,omitempty"`
	Media                      *Media               `json:"media,omitempty"`
	RecoveryUser               *UserRecovery        `json:"recoveryUser,omitempty"`
	StateMachineHistoryEntries *StateMachineHistory `json:"stateMachineHistoryEntries,omitempty"`
	TimeZone                   string               `json:"timeZone,omitempty"`
	Title                      *string              `json:"title,omitempty"`
	UpdatedAt                  *time.Time           `json:"updatedAt,omitempty"`
	UpdatedCustomers           *Customer            `json:"updatedCustomers,omitempty"`
	UpdatedOrders              *Order               `json:"updatedOrders,omitempty"`
	Username                   string               `json:"username,omitempty"`
}

// UserAccessKey data structure
// Added since version: 6.0.0.0
// Required fields: userId, accessKey, secretAccessKey, createdAt
type UserAccessKey struct {
	AccessKey       string         `json:"accessKey,omitempty"`
	CreatedAt       *time.Time     `json:"createdAt,omitempty"`
	CustomFields    *[]CustomField `json:"customFields,omitempty"`
	Id              string         `json:"id,omitempty"`
	LastUsageAt     *time.Time     `json:"lastUsageAt,omitempty"`
	SecretAccessKey string         `json:"secretAccessKey,omitempty"`
	UpdatedAt       *time.Time     `json:"updatedAt,omitempty"`
	User            *User          `json:"user,omitempty"`
	UserId          string         `json:"userId,omitempty"`
}

// UserConfig data structure
// Added since version: 6.3.5.0
// Required fields: userId, key, createdAt
type UserConfig struct {
	CreatedAt *time.Time   `json:"createdAt,omitempty"`
	Id        string       `json:"id,omitempty"`
	Key       string       `json:"key,omitempty"`
	UpdatedAt *time.Time   `json:"updatedAt,omitempty"`
	User      *User        `json:"user,omitempty"`
	UserId    string       `json:"userId,omitempty"`
	Value     *interface{} `json:"value,omitempty"` // map[type:object]
}

// UserRecovery data structure
// Added since version: 6.0.0.0
// Required fields: hash, userId, createdAt
type UserRecovery struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Hash      string     `json:"hash,omitempty"`
	Id        string     `json:"id,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	User      *User      `json:"user,omitempty"`
	UserId    string     `json:"userId,omitempty"`
}
