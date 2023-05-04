package entity

import "time"

// Generated from Shopware Admin API
// Version 6.5.9999999.9999999-dev at 2023-05-02 19:06:34 UTC

// App data structure
// Added since version: 6.3.1.0
// Required fields: name, path, active, configurable, version, allowDisable, integrationId, aclRoleId, createdAt, label
type App struct {
	AclRole                 *AclRole          `json:"aclRole,omitempty"`
	AclRoleId               string            `json:"aclRoleId,omitempty"`
	ActionButtons           *AppActionButton  `json:"actionButtons,omitempty"`
	Active                  bool              `json:"active,omitempty"`
	AllowDisable            bool              `json:"allowDisable,omitempty"`
	AllowedHosts            *interface{}      `json:"allowedHosts,omitempty"` // map[items:map[type:string] type:array]
	Author                  *string           `json:"author,omitempty"`
	BaseAppUrl              *string           `json:"baseAppUrl,omitempty"`
	CmsBlocks               *AppCmsBlock      `json:"cmsBlocks,omitempty"`
	Configurable            bool              `json:"configurable,omitempty"`
	Cookies                 *interface{}      `json:"cookies,omitempty"` // map[items:map[type:object] type:array]
	Copyright               *string           `json:"copyright,omitempty"`
	CreatedAt               *time.Time        `json:"createdAt,omitempty"`
	CustomFieldSets         *CustomFieldSet   `json:"customFieldSets,omitempty"`
	CustomFields            *[]CustomField    `json:"customFields,omitempty"`
	Description             *string           `json:"description,omitempty"`
	FlowActions             *AppFlowAction    `json:"flowActions,omitempty"`
	Icon                    *string           `json:"icon,omitempty"`
	Id                      string            `json:"id,omitempty"`
	Integration             *Integration      `json:"integration,omitempty"`
	IntegrationId           string            `json:"integrationId,omitempty"`
	Label                   string            `json:"label,omitempty"`
	License                 *string           `json:"license,omitempty"`
	MainModule              *interface{}      `json:"mainModule,omitempty"` // map[type:object]
	Modules                 *interface{}      `json:"modules,omitempty"`    // map[items:map[type:object] type:array]
	Name                    string            `json:"name,omitempty"`
	Path                    string            `json:"path,omitempty"`
	PaymentMethods          *AppPaymentMethod `json:"paymentMethods,omitempty"`
	Privacy                 *string           `json:"privacy,omitempty"`
	PrivacyPolicyExtensions *string           `json:"privacyPolicyExtensions,omitempty"`
	TemplateLoadPriority    *int              `json:"templateLoadPriority,omitempty"`
	Templates               *AppTemplate      `json:"templates,omitempty"`
	Translated              *interface{}      `json:"translated,omitempty"` // map[type:object]
	UpdatedAt               *time.Time        `json:"updatedAt,omitempty"`
	Version                 string            `json:"version,omitempty"`
	Webhooks                *Webhook          `json:"webhooks,omitempty"`
}

// AppActionButton data structure
// Added since version: 6.3.1.0
// Required fields: entity, view, url, action, appId, createdAt, label
type AppActionButton struct {
	Action     string       `json:"action,omitempty"`
	App        *App         `json:"app,omitempty"`
	AppId      string       `json:"appId,omitempty"`
	CreatedAt  *time.Time   `json:"createdAt,omitempty"`
	Entity     string       `json:"entity,omitempty"`
	Id         string       `json:"id,omitempty"`
	Label      string       `json:"label,omitempty"`
	Translated *interface{} `json:"translated,omitempty"` // map[type:object]
	UpdatedAt  *time.Time   `json:"updatedAt,omitempty"`
	Url        string       `json:"url,omitempty"`
	View       string       `json:"view,omitempty"`
}

// AppAdministrationSnippet data structure
// Added since version: 6.4.15.0
// Required fields: value, appId, localeId, createdAt
type AppAdministrationSnippet struct {
	AppId     string     `json:"appId,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Id        string     `json:"id,omitempty"`
	LocaleId  string     `json:"localeId,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	Value     string     `json:"value,omitempty"`
}

// AppCmsBlock data structure
// Added since version: 6.4.2.0
// Required fields: name, block, template, styles, appId, createdAt, label
type AppCmsBlock struct {
	App        *App         `json:"app,omitempty"`
	AppId      string       `json:"appId,omitempty"`
	Block      *interface{} `json:"block,omitempty"` // map[type:object]
	CreatedAt  *time.Time   `json:"createdAt,omitempty"`
	Id         string       `json:"id,omitempty"`
	Label      string       `json:"label,omitempty"`
	Name       string       `json:"name,omitempty"`
	Styles     string       `json:"styles,omitempty"`
	Template   string       `json:"template,omitempty"`
	Translated *interface{} `json:"translated,omitempty"` // map[type:object]
	UpdatedAt  *time.Time   `json:"updatedAt,omitempty"`
}

// AppFlowAction data structure
// Added since version: 6.4.10.0
// Required fields: appId, name, url, createdAt, label
type AppFlowAction struct {
	App           *App           `json:"app,omitempty"`
	AppId         string         `json:"appId,omitempty"`
	Badge         *string        `json:"badge,omitempty"`
	Config        *interface{}   `json:"config,omitempty"` // map[type:object]
	CreatedAt     *time.Time     `json:"createdAt,omitempty"`
	CustomFields  *[]CustomField `json:"customFields,omitempty"`
	Delayable     *bool          `json:"delayable,omitempty"`
	Description   *string        `json:"description,omitempty"`
	FlowSequences *FlowSequence  `json:"flowSequences,omitempty"`
	Headers       *interface{}   `json:"headers,omitempty"` // map[type:object]
	Headline      *string        `json:"headline,omitempty"`
	Icon          *string        `json:"icon,omitempty"`
	IconRaw       *string        `json:"iconRaw,omitempty"`
	Id            string         `json:"id,omitempty"`
	Label         string         `json:"label,omitempty"`
	Name          string         `json:"name,omitempty"`
	Parameters    *interface{}   `json:"parameters,omitempty"`   // map[type:object]
	Requirements  *interface{}   `json:"requirements,omitempty"` // map[items:map[type:string] type:array]
	SwIcon        *string        `json:"swIcon,omitempty"`
	Translated    *interface{}   `json:"translated,omitempty"` // map[type:object]
	UpdatedAt     *time.Time     `json:"updatedAt,omitempty"`
	Url           string         `json:"url,omitempty"`
}

// AppPaymentMethod data structure
// Added since version: 6.4.1.0
// Required fields: appName, identifier, paymentMethodId, createdAt
type AppPaymentMethod struct {
	App             *App           `json:"app,omitempty"`
	AppId           *string        `json:"appId,omitempty"`
	AppName         string         `json:"appName,omitempty"`
	CaptureUrl      *string        `json:"captureUrl,omitempty"`
	CreatedAt       *time.Time     `json:"createdAt,omitempty"`
	FinalizeUrl     *string        `json:"finalizeUrl,omitempty"`
	Id              string         `json:"id,omitempty"`
	Identifier      string         `json:"identifier,omitempty"`
	OriginalMedia   *Media         `json:"originalMedia,omitempty"`
	OriginalMediaId *string        `json:"originalMediaId,omitempty"`
	PayUrl          *string        `json:"payUrl,omitempty"`
	PaymentMethod   *PaymentMethod `json:"paymentMethod,omitempty"`
	PaymentMethodId string         `json:"paymentMethodId,omitempty"`
	RefundUrl       *string        `json:"refundUrl,omitempty"`
	UpdatedAt       *time.Time     `json:"updatedAt,omitempty"`
	ValidateUrl     *string        `json:"validateUrl,omitempty"`
}

// AppScriptCondition data structure
// Added since version: 6.4.10.3
// Required fields: identifier, active, appId, createdAt, name
type AppScriptCondition struct {
	Active         bool           `json:"active,omitempty"`
	App            *App           `json:"app,omitempty"`
	AppId          string         `json:"appId,omitempty"`
	Config         *interface{}   `json:"config,omitempty"` // map[type:object]
	CreatedAt      *time.Time     `json:"createdAt,omitempty"`
	Group          *string        `json:"group,omitempty"`
	Id             string         `json:"id,omitempty"`
	Identifier     string         `json:"identifier,omitempty"`
	Name           string         `json:"name,omitempty"`
	RuleConditions *RuleCondition `json:"ruleConditions,omitempty"`
	Script         *string        `json:"script,omitempty"`
	Translated     *interface{}   `json:"translated,omitempty"` // map[type:object]
	UpdatedAt      *time.Time     `json:"updatedAt,omitempty"`
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
