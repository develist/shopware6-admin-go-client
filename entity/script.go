package entity

import "time"

// Generated from Shopware Admin API
// Version 6.5.9999999.9999999-dev at 2023-05-02 19:06:34 UTC

// Script data structure
// Added since version: 6.4.7.0
// Required fields: script, hook, name, active, createdAt
type Script struct {
	Active    bool       `json:"active,omitempty"`
	App       *App       `json:"app,omitempty"`
	AppId     *string    `json:"appId,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Hook      string     `json:"hook,omitempty"`
	Id        string     `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Script    string     `json:"script,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}
