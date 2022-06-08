package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-06 18:44:04 UTC

// SeoUrl data structure
// Added since version: 6.0.0.0
// Required fields: languageId, foreignKey, routeName, pathInfo, seoPathInfo, createdAt
type SeoUrl struct {
	CreatedAt      *time.Time     `json:"createdAt,omitempty"`
	CustomFields   *[]CustomField `json:"customFields,omitempty"`
	ForeignKey     *string        `json:"foreignKey,omitempty"`
	Id             *string        `json:"id,omitempty"`
	IsCanonical    *bool          `json:"isCanonical,omitempty"`
	IsDeleted      *bool          `json:"isDeleted,omitempty"`
	IsModified     *bool          `json:"isModified,omitempty"`
	Language       *Language      `json:"language,omitempty"`
	LanguageId     *string        `json:"languageId,omitempty"`
	PathInfo       *string        `json:"pathInfo,omitempty"`
	RouteName      *string        `json:"routeName,omitempty"`
	SalesChannel   *SalesChannel  `json:"salesChannel,omitempty"`
	SalesChannelId *string        `json:"salesChannelId,omitempty"`
	SeoPathInfo    *string        `json:"seoPathInfo,omitempty"`
	UpdatedAt      *time.Time     `json:"updatedAt,omitempty"`
	Url            *string        `json:"url,omitempty"`
}

// SeoUrlTemplate data structure
// Added since version: 6.0.0.0
// Required fields: entityName, routeName, createdAt
type SeoUrlTemplate struct {
	CreatedAt      *time.Time     `json:"createdAt,omitempty"`
	CustomFields   *[]CustomField `json:"customFields,omitempty"`
	EntityName     *string        `json:"entityName,omitempty"`
	Id             *string        `json:"id,omitempty"`
	IsValid        *bool          `json:"isValid,omitempty"`
	RouteName      *string        `json:"routeName,omitempty"`
	SalesChannel   *SalesChannel  `json:"salesChannel,omitempty"`
	SalesChannelId *string        `json:"salesChannelId,omitempty"`
	Template       *string        `json:"template,omitempty"`
	UpdatedAt      *time.Time     `json:"updatedAt,omitempty"`
}
