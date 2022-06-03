package entity

import "time"

// completed

type App struct {
	AclRole       *AclRole         `json:"aclRole,omitempty"`
	AclRoleId     string           `json:"aclRoleId,omitempty"`
	ActionButtons *AppActionButton `json:"actionButtons,omitempty"`
	Active        bool             `json:"active,omitempty"`
	Author        string           `json:"author,omitempty"`
	Configurable  bool             `json:"configurable,omitempty"`
	// Cookies map[items:map[type:object] type:array] `json:"cookies,omitempty"`
	Copyright       string          `json:"copyright,omitempty"`
	CreatedAt       *time.Time      `json:"createdAt,omitempty"`
	CustomFieldSets *CustomFieldSet `json:"customFieldSets,omitempty"`
	CustomFields    *[]CustomField  `json:"customFields,omitempty"`
	Description     string          `json:"description,omitempty"`
	Icon            string          `json:"icon,omitempty"`
	Id              string          `json:"id,omitempty"`
	Integration     *Integration    `json:"integration,omitempty"`
	IntegrationId   string          `json:"integrationId,omitempty"`
	Label           string          `json:"label,omitempty"`
	License         string          `json:"license,omitempty"`
	// MainModule map[type:object] `json:"mainModule,omitempty"`
	// Modules map[items:map[type:object] type:array] `json:"modules,omitempty"`
	Name                    string            `json:"name,omitempty"`
	Path                    string            `json:"path,omitempty"`
	PaymentMethods          *AppPaymentMethod `json:"paymentMethods,omitempty"`
	Privacy                 string            `json:"privacy,omitempty"`
	PrivacyPolicyExtensions string            `json:"privacyPolicyExtensions,omitempty"`
	Templates               *AppTemplate      `json:"templates,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	Version   string     `json:"version,omitempty"`
	Webhooks  *Webhook   `json:"webhooks,omitempty"`
}

type AppActionButton struct {
	Action     string     `json:"action,omitempty"`
	App        *App       `json:"app,omitempty"`
	AppId      string     `json:"appId,omitempty"`
	CreatedAt  *time.Time `json:"createdAt,omitempty"`
	Entity     string     `json:"entity,omitempty"`
	Id         string     `json:"id,omitempty"`
	Label      string     `json:"label,omitempty"`
	OpenNewTab bool       `json:"openNewTab,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	Url       string     `json:"url,omitempty"`
	View      string     `json:"view,omitempty"`
}

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
