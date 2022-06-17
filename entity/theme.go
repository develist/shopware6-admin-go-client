package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-17 19:07:53 UTC

// Theme data structure
// Added since version: 6.0.0.0
// Required fields: name, author, active, createdAt
type Theme struct {
	Active         bool           `json:"active,omitempty"`
	Author         string         `json:"author,omitempty"`
	BaseConfig     *interface{}   `json:"baseConfig,omitempty"` // map[type:object]
	ChildThemes    *Theme         `json:"childThemes,omitempty"`
	ConfigValues   *interface{}   `json:"configValues,omitempty"` // map[type:object]
	CreatedAt      *time.Time     `json:"createdAt,omitempty"`
	CustomFields   *[]CustomField `json:"customFields,omitempty"`
	Description    *string        `json:"description,omitempty"`
	HelpTexts      *interface{}   `json:"helpTexts,omitempty"` // map[type:object]
	Id             string         `json:"id,omitempty"`
	Labels         *interface{}   `json:"labels,omitempty"` // map[type:object]
	Media          *Media         `json:"media,omitempty"`
	Name           string         `json:"name,omitempty"`
	ParentThemeId  *string        `json:"parentThemeId,omitempty"`
	PreviewMedia   *Media         `json:"previewMedia,omitempty"`
	PreviewMediaId *string        `json:"previewMediaId,omitempty"`
	SalesChannels  *SalesChannel  `json:"salesChannels,omitempty"`
	TechnicalName  *string        `json:"technicalName,omitempty"`
	Translated     *interface{}   `json:"translated,omitempty"` // map[type:object]
	UpdatedAt      *time.Time     `json:"updatedAt,omitempty"`
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
