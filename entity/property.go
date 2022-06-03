package entity

import "time"

// completed

type PropertyGroup struct {
	CreatedAt    *time.Time           `json:"createdAt,omitempty"`
	CustomFields *[]CustomField       `json:"customFields,omitempty"`
	Description  string               `json:"description,omitempty"`
	DisplayType  string               `json:"displayType,omitempty"`
	Filterable   bool                 `json:"filterable,omitempty"`
	Id           string               `json:"id,omitempty"`
	Name         string               `json:"name,omitempty"`
	Options      *PropertyGroupOption `json:"options,omitempty"`
	Position     int                  `json:"position,omitempty"`
	SortingType  string               `json:"sortingType,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt                  *time.Time `json:"updatedAt,omitempty"`
	VisibleOnProductDetailPage bool       `json:"visibleOnProductDetailPage,omitempty"`
}

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
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}
