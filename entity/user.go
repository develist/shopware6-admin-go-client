package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-06 18:44:04 UTC

// User data structure
// Added since version: 6.0.0.0
// Required fields: localeId, username, firstName, lastName, email, createdAt
type User struct {
	AccessKeys                 *UserAccessKey       `json:"accessKeys,omitempty"`
	AclRoles                   *AclRole             `json:"aclRoles,omitempty"`
	Active                     *bool                `json:"active,omitempty"`
	Admin                      *bool                `json:"admin,omitempty"`
	AvatarId                   *string              `json:"avatarId,omitempty"`
	AvatarMedia                *Media               `json:"avatarMedia,omitempty"`
	Configs                    *UserConfig          `json:"configs,omitempty"`
	CreatedAt                  *time.Time           `json:"createdAt,omitempty"`
	CreatedOrders              *Order               `json:"createdOrders,omitempty"`
	CustomFields               *[]CustomField       `json:"customFields,omitempty"`
	Email                      *string              `json:"email,omitempty"`
	FirstName                  *string              `json:"firstName,omitempty"`
	Id                         *string              `json:"id,omitempty"`
	ImportExportLogEntries     *ImportExportLog     `json:"importExportLogEntries,omitempty"`
	LastName                   *string              `json:"lastName,omitempty"`
	LastUpdatedPasswordAt      *time.Time           `json:"lastUpdatedPasswordAt,omitempty"`
	Locale                     *Locale              `json:"locale,omitempty"`
	LocaleId                   *string              `json:"localeId,omitempty"`
	Media                      *Media               `json:"media,omitempty"`
	RecoveryUser               *UserRecovery        `json:"recoveryUser,omitempty"`
	StateMachineHistoryEntries *StateMachineHistory `json:"stateMachineHistoryEntries,omitempty"`
	Title                      *string              `json:"title,omitempty"`
	UpdatedAt                  *time.Time           `json:"updatedAt,omitempty"`
	UpdatedOrders              *Order               `json:"updatedOrders,omitempty"`
	Username                   *string              `json:"username,omitempty"`
}

// UserAccessKey data structure
// Added since version: 6.0.0.0
// Required fields: userId, accessKey, secretAccessKey, createdAt
type UserAccessKey struct {
	AccessKey       *string        `json:"accessKey,omitempty"`
	CreatedAt       *time.Time     `json:"createdAt,omitempty"`
	CustomFields    *[]CustomField `json:"customFields,omitempty"`
	Id              *string        `json:"id,omitempty"`
	LastUsageAt     *time.Time     `json:"lastUsageAt,omitempty"`
	SecretAccessKey *string        `json:"secretAccessKey,omitempty"`
	UpdatedAt       *time.Time     `json:"updatedAt,omitempty"`
	User            *User          `json:"user,omitempty"`
	UserId          *string        `json:"userId,omitempty"`
	WriteAccess     *bool          `json:"writeAccess,omitempty"`
}

// UserConfig data structure
// Added since version: 6.3.5.0
// Required fields: userId, key, createdAt
type UserConfig struct {
	CreatedAt *time.Time   `json:"createdAt,omitempty"`
	Id        *string      `json:"id,omitempty"`
	Key       *string      `json:"key,omitempty"`
	UpdatedAt *time.Time   `json:"updatedAt,omitempty"`
	User      *User        `json:"user,omitempty"`
	UserId    *string      `json:"userId,omitempty"`
	Value     *interface{} `json:"value,omitempty"` // map[type:object]
}

// UserRecovery data structure
// Added since version: 6.0.0.0
// Required fields: hash, userId, createdAt
type UserRecovery struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Hash      *string    `json:"hash,omitempty"`
	Id        *string    `json:"id,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	User      *User      `json:"user,omitempty"`
	UserId    *string    `json:"userId,omitempty"`
}
