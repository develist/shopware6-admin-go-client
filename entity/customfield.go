package entity

import "time"

// completed

type CustomField struct {
	Active bool `json:"active,omitempty"`
	// Config map[type:object] `json:"config,omitempty"`
	CreatedAt                 *time.Time                `json:"createdAt,omitempty"`
	CustomFieldSet            *CustomFieldSet           `json:"customFieldSet,omitempty"`
	CustomFieldSetId          string                    `json:"customFieldSetId,omitempty"`
	Id                        string                    `json:"id,omitempty"`
	Name                      string                    `json:"name,omitempty"`
	ProductSearchConfigFields *ProductSearchConfigField `json:"productSearchConfigFields,omitempty"`
	Type                      string                    `json:"type,omitempty"`
	UpdatedAt                 *time.Time                `json:"updatedAt,omitempty"`
}

type CustomFieldSet struct {
	Active bool   `json:"active,omitempty"`
	App    *App   `json:"app,omitempty"`
	AppId  string `json:"appId,omitempty"`
	// Config map[type:object] `json:"config,omitempty"`
	CreatedAt    *time.Time              `json:"createdAt,omitempty"`
	CustomFields *[]CustomField          `json:"customFields,omitempty"`
	Global       bool                    `json:"global,omitempty"`
	Id           string                  `json:"id,omitempty"`
	Name         string                  `json:"name,omitempty"`
	Position     int                     `json:"position,omitempty"`
	Products     *Product                `json:"products,omitempty"`
	Relations    *CustomFieldSetRelation `json:"relations,omitempty"`
	UpdatedAt    *time.Time              `json:"updatedAt,omitempty"`
}

type CustomFieldSetRelation struct {
	CreatedAt        *time.Time      `json:"createdAt,omitempty"`
	CustomFieldSet   *CustomFieldSet `json:"customFieldSet,omitempty"`
	CustomFieldSetId string          `json:"customFieldSetId,omitempty"`
	EntityName       string          `json:"entityName,omitempty"`
	Id               string          `json:"id,omitempty"`
	UpdatedAt        *time.Time      `json:"updatedAt,omitempty"`
}
