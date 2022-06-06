package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-06 18:44:04 UTC

// App data structure
// Added since version: 6.3.1.0
// Required fields: name, path, active, configurable, version, integrationId, aclRoleId, createdAt, label
type App struct {
	AclRole                 *AclRole          `json:"aclRole,omitempty"`
	AclRoleId               string            `json:"aclRoleId,omitempty"`
	ActionButtons           *AppActionButton  `json:"actionButtons,omitempty"`
	Active                  bool              `json:"active,omitempty"`
	Author                  string            `json:"author,omitempty"`
	Configurable            bool              `json:"configurable,omitempty"`
	Cookies                 *interface{}      `json:"cookies,omitempty"` // map[items:map[type:object] type:array]
	Copyright               string            `json:"copyright,omitempty"`
	CreatedAt               *time.Time        `json:"createdAt,omitempty"`
	CustomFieldSets         *CustomFieldSet   `json:"customFieldSets,omitempty"`
	CustomFields            *[]CustomField    `json:"customFields,omitempty"`
	Description             string            `json:"description,omitempty"`
	Icon                    string            `json:"icon,omitempty"`
	Id                      string            `json:"id,omitempty"`
	Integration             *Integration      `json:"integration,omitempty"`
	IntegrationId           string            `json:"integrationId,omitempty"`
	Label                   string            `json:"label,omitempty"`
	License                 string            `json:"license,omitempty"`
	MainModule              *interface{}      `json:"mainModule,omitempty"` // map[type:object]
	Modules                 *interface{}      `json:"modules,omitempty"`    // map[items:map[type:object] type:array]
	Name                    string            `json:"name,omitempty"`
	Path                    string            `json:"path,omitempty"`
	PaymentMethods          *AppPaymentMethod `json:"paymentMethods,omitempty"`
	Privacy                 string            `json:"privacy,omitempty"`
	PrivacyPolicyExtensions string            `json:"privacyPolicyExtensions,omitempty"`
	Templates               *AppTemplate      `json:"templates,omitempty"`
	Translated              *interface{}      `json:"translated,omitempty"` // map[type:object]
	UpdatedAt               *time.Time        `json:"updatedAt,omitempty"`
	Version                 string            `json:"version,omitempty"`
	Webhooks                *Webhook          `json:"webhooks,omitempty"`
}

// AppActionButton data structure
// Added since version: 6.3.1.0
// Required fields: entity, view, url, action, openNewTab, appId, createdAt, label
type AppActionButton struct {
	Action     string       `json:"action,omitempty"`
	App        *App         `json:"app,omitempty"`
	AppId      string       `json:"appId,omitempty"`
	CreatedAt  *time.Time   `json:"createdAt,omitempty"`
	Entity     string       `json:"entity,omitempty"`
	Id         string       `json:"id,omitempty"`
	Label      string       `json:"label,omitempty"`
	OpenNewTab bool         `json:"openNewTab,omitempty"`
	Translated *interface{} `json:"translated,omitempty"` // map[type:object]
	UpdatedAt  *time.Time   `json:"updatedAt,omitempty"`
	Url        string       `json:"url,omitempty"`
	View       string       `json:"view,omitempty"`
}

// AppPaymentMethod data structure
// Added since version: 6.4.1.0
// Required fields: appName, identifier, paymentMethodId, createdAt
type AppPaymentMethod struct {
	App             *App           `json:"app,omitempty"`
	AppId           string         `json:"appId,omitempty"`
	AppName         string         `json:"appName,omitempty"`
	CreatedAt       *time.Time     `json:"createdAt,omitempty"`
	FinalizeUrl     string         `json:"finalizeUrl,omitempty"`
	Id              string         `json:"id,omitempty"`
	Identifier      string         `json:"identifier,omitempty"`
	OriginalMedia   *Media         `json:"originalMedia,omitempty"`
	OriginalMediaId string         `json:"originalMediaId,omitempty"`
	PayUrl          string         `json:"payUrl,omitempty"`
	PaymentMethod   *PaymentMethod `json:"paymentMethod,omitempty"`
	PaymentMethodId string         `json:"paymentMethodId,omitempty"`
	UpdatedAt       *time.Time     `json:"updatedAt,omitempty"`
}

// AppTemplate data structure
// Added since version: 6.3.1.0
// Required fields: template, path, active, appId, createdAt
type AppTemplate struct {
	Active    bool       `json:"active,omitempty"`
	App       *App       `json:"app,omitempty"`
	AppId     string     `json:"appId,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Id        string     `json:"id,omitempty"`
	Path      string     `json:"path,omitempty"`
	Template  string     `json:"template,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}
