package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-06 18:44:04 UTC

// PropertyGroup data structure
// Added since version: 6.0.0.0
// Required fields: displayType, sortingType, createdAt, name
type PropertyGroup struct {
	CreatedAt                  *time.Time           `json:"createdAt,omitempty"`
	CustomFields               *[]CustomField       `json:"customFields,omitempty"`
	Description                string               `json:"description,omitempty"`
	DisplayType                string               `json:"displayType,omitempty"`
	Filterable                 bool                 `json:"filterable,omitempty"`
	Id                         string               `json:"id,omitempty"`
	Name                       string               `json:"name,omitempty"`
	Options                    *PropertyGroupOption `json:"options,omitempty"`
	Position                   int                  `json:"position,omitempty"`
	SortingType                string               `json:"sortingType,omitempty"`
	Translated                 *interface{}         `json:"translated,omitempty"` // map[type:object]
	UpdatedAt                  *time.Time           `json:"updatedAt,omitempty"`
	VisibleOnProductDetailPage bool                 `json:"visibleOnProductDetailPage,omitempty"`
}

// PropertyGroupOption data structure
// Added since version: 6.0.0.0
// Required fields: groupId, createdAt, name
type PropertyGroupOption struct {
	ColorHexCode                string                      `json:"colorHexCode,omitempty"`
	CreatedAt                   *time.Time                  `json:"createdAt,omitempty"`
	CustomFields                *[]CustomField              `json:"customFields,omitempty"`
	Group                       *PropertyGroup              `json:"group,omitempty"`
	GroupId                     string                      `json:"groupId,omitempty"`
	Id                          string                      `json:"id,omitempty"`
	Media                       *Media                      `json:"media,omitempty"`
	MediaId                     string                      `json:"mediaId,omitempty"`
	Name                        string                      `json:"name,omitempty"`
	Position                    int                         `json:"position,omitempty"`
	ProductConfiguratorSettings *ProductConfiguratorSetting `json:"productConfiguratorSettings,omitempty"`
	ProductOptions              *Product                    `json:"productOptions,omitempty"`
	ProductProperties           *Product                    `json:"productProperties,omitempty"`
	Translated                  *interface{}                `json:"translated,omitempty"` // map[type:object]
	UpdatedAt                   *time.Time                  `json:"updatedAt,omitempty"`
}
