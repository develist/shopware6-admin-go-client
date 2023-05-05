package entity

import "time"

// Generated from Shopware Admin API
// Version 6.5.9999999.9999999-dev at 2023-05-02 19:06:34 UTC

// Theme data structure
// Added since version: 6.0.0.0
// Required fields: name, author, active, createdAt
type Theme struct {
	Active          bool           `json:"active,omitempty"`
	Author          string         `json:"author,omitempty"`
	BaseConfig      *interface{}   `json:"baseConfig,omitempty"`   // map[type:object]
	ConfigValues    *interface{}   `json:"configValues,omitempty"` // map[type:object]
	CreatedAt       *time.Time     `json:"createdAt,omitempty"`
	CustomFields    *[]CustomField `json:"customFields,omitempty"`
	DependentThemes *Theme         `json:"dependentThemes,omitempty"`
	Description     *string        `json:"description,omitempty"`
	HelpTexts       *interface{}   `json:"helpTexts,omitempty"` // map[type:object]
	Id              string         `json:"id,omitempty"`
	Labels          *interface{}   `json:"labels,omitempty"` // map[type:object]
	Media           *Media         `json:"media,omitempty"`
	Name            string         `json:"name,omitempty"`
	ParentThemeId   *string        `json:"parentThemeId,omitempty"`
	PreviewMedia    *Media         `json:"previewMedia,omitempty"`
	PreviewMediaId  *string        `json:"previewMediaId,omitempty"`
	SalesChannels   *SalesChannel  `json:"salesChannels,omitempty"`
	TechnicalName   *string        `json:"technicalName,omitempty"`
	Translated      *interface{}   `json:"translated,omitempty"` // map[type:object]
	UpdatedAt       *time.Time     `json:"updatedAt,omitempty"`
}

// ThemeChild data structure
// Added since version: 6.4.8.0
// Required fields: parentId, childId
type ThemeChild struct {
	ChildId     string `json:"childId,omitempty"`
	ChildTheme  *Theme `json:"childTheme,omitempty"`
	Id          string `json:"id,omitempty"`
	ParentId    string `json:"parentId,omitempty"`
	ParentTheme *Theme `json:"parentTheme,omitempty"`
}

// ThemeMedia data structure
// Added since version: 6.0.0.0
// Required fields: themeId, mediaId
type ThemeMedia struct {
	Id      string `json:"id,omitempty"`
	Media   *Media `json:"media,omitempty"`
	MediaId string `json:"mediaId,omitempty"`
	Theme   *Theme `json:"theme,omitempty"`
	ThemeId string `json:"themeId,omitempty"`
}

// ThemeSalesChannel data structure
// Added since version: 6.0.0.0
// Required fields: salesChannelId, themeId
type ThemeSalesChannel struct {
	Id             string        `json:"id,omitempty"`
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId string        `json:"salesChannelId,omitempty"`
	Theme          *Theme        `json:"theme,omitempty"`
	ThemeId        string        `json:"themeId,omitempty"`
}
