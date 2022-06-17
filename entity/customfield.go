package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-17 19:07:53 UTC

// CustomField data structure
// Added since version: 6.0.0.0
// Required fields: name, type, createdAt
type CustomField struct {
	Active                    *bool                     `json:"active,omitempty"`
	Config                    *interface{}              `json:"config,omitempty"` // map[type:object]
	CreatedAt                 *time.Time                `json:"createdAt,omitempty"`
	CustomFieldSet            *CustomFieldSet           `json:"customFieldSet,omitempty"`
	CustomFieldSetId          *string                   `json:"customFieldSetId,omitempty"`
	Id                        string                    `json:"id,omitempty"`
	Name                      string                    `json:"name,omitempty"`
	ProductSearchConfigFields *ProductSearchConfigField `json:"productSearchConfigFields,omitempty"`
	Type                      string                    `json:"type,omitempty"`
	UpdatedAt                 *time.Time                `json:"updatedAt,omitempty"`
}

// CustomFieldSet data structure
// Added since version: 6.0.0.0
// Required fields: name, createdAt
type CustomFieldSet struct {
	Active       *bool                   `json:"active,omitempty"`
	App          *App                    `json:"app,omitempty"`
	AppId        *string                 `json:"appId,omitempty"`
	Config       *interface{}            `json:"config,omitempty"` // map[type:object]
	CreatedAt    *time.Time              `json:"createdAt,omitempty"`
	CustomFields *[]CustomField          `json:"customFields,omitempty"`
	Global       *bool                   `json:"global,omitempty"`
	Id           string                  `json:"id,omitempty"`
	Name         string                  `json:"name,omitempty"`
	Position     *int                    `json:"position,omitempty"`
	Products     *Product                `json:"products,omitempty"`
	Relations    *CustomFieldSetRelation `json:"relations,omitempty"`
	UpdatedAt    *time.Time              `json:"updatedAt,omitempty"`
}

// CustomFieldSetRelation data structure
// Added since version: 6.0.0.0
// Required fields: customFieldSetId, entityName, createdAt
type CustomFieldSetRelation struct {
	CreatedAt        *time.Time      `json:"createdAt,omitempty"`
	CustomFieldSet   *CustomFieldSet `json:"customFieldSet,omitempty"`
	CustomFieldSetId string          `json:"customFieldSetId,omitempty"`
	EntityName       string          `json:"entityName,omitempty"`
	Id               string          `json:"id,omitempty"`
	UpdatedAt        *time.Time      `json:"updatedAt,omitempty"`
}
