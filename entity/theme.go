package entity

import "time"

// completed

type Theme struct {
	Active bool   `json:"active,omitempty"`
	Author string `json:"author,omitempty"`
	// BaseConfig map[type:object] `json:"baseConfig,omitempty"`
	ChildThemes *Theme `json:"childThemes,omitempty"`
	// ConfigValues map[type:object] `json:"configValues,omitempty"`
	CreatedAt    time.Time      `json:"createdAt,omitempty"`
	CustomFields *[]CustomField `json:"customFields,omitempty"`
	Description  string         `json:"description,omitempty"`
	// HelpTexts map[type:object] `json:"helpTexts,omitempty"`
	Id string `json:"id,omitempty"`
	// Labels map[type:object] `json:"labels,omitempty"`
	Media          *Media        `json:"media,omitempty"`
	Name           string        `json:"name,omitempty"`
	ParentThemeId  string        `json:"parentThemeId,omitempty"`
	PreviewMedia   *Media        `json:"previewMedia,omitempty"`
	PreviewMediaId string        `json:"previewMediaId,omitempty"`
	SalesChannels  *SalesChannel `json:"salesChannels,omitempty"`
	TechnicalName  string        `json:"technicalName,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type ThemeMedia struct {
	Id      string `json:"id,omitempty"`
	Media   *Media `json:"media,omitempty"`
	MediaId string `json:"mediaId,omitempty"`
	Theme   *Theme `json:"theme,omitempty"`
	ThemeId string `json:"themeId,omitempty"`
}

type ThemeSalesChannel struct {
	Id             string        `json:"id,omitempty"`
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId string        `json:"salesChannelId,omitempty"`
	Theme          *Theme        `json:"theme,omitempty"`
	ThemeId        string        `json:"themeId,omitempty"`
}
