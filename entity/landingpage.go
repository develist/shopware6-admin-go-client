package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-06 18:44:04 UTC

// LandingPage data structure
// Added since version: 6.4.0.0
// Required fields: createdAt, name, url
type LandingPage struct {
	Active           bool           `json:"active,omitempty"`
	CmsPage          *CmsPage       `json:"cmsPage,omitempty"`
	CmsPageId        string         `json:"cmsPageId,omitempty"`
	CmsPageVersionId string         `json:"cmsPageVersionId,omitempty"`
	CreatedAt        *time.Time     `json:"createdAt,omitempty"`
	CustomFields     *[]CustomField `json:"customFields,omitempty"`
	Id               string         `json:"id,omitempty"`
	Keywords         string         `json:"keywords,omitempty"`
	MetaDescription  string         `json:"metaDescription,omitempty"`
	MetaTitle        string         `json:"metaTitle,omitempty"`
	Name             string         `json:"name,omitempty"`
	SalesChannels    *SalesChannel  `json:"salesChannels,omitempty"`
	SeoUrls          *SeoUrl        `json:"seoUrls,omitempty"`
	SlotConfig       *interface{}   `json:"slotConfig,omitempty"` // map[type:object]
	Tags             *Tag           `json:"tags,omitempty"`
	Translated       *interface{}   `json:"translated,omitempty"` // map[type:object]
	UpdatedAt        *time.Time     `json:"updatedAt,omitempty"`
	Url              string         `json:"url,omitempty"`
	VersionId        string         `json:"versionId,omitempty"`
}

// LandingPageSalesChannel data structure
// Added since version: 6.4.0.0
// Required fields: landingPageId, salesChannelId
type LandingPageSalesChannel struct {
	Id                   string        `json:"id,omitempty"`
	LandingPage          *LandingPage  `json:"landingPage,omitempty"`
	LandingPageId        string        `json:"landingPageId,omitempty"`
	LandingPageVersionId string        `json:"landingPageVersionId,omitempty"`
	SalesChannel         *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId       string        `json:"salesChannelId,omitempty"`
}

// LandingPageTag data structure
// Added since version: 6.4.0.0
// Required fields: landingPageId, tagId
type LandingPageTag struct {
	Id                   string       `json:"id,omitempty"`
	LandingPage          *LandingPage `json:"landingPage,omitempty"`
	LandingPageId        string       `json:"landingPageId,omitempty"`
	LandingPageVersionId string       `json:"landingPageVersionId,omitempty"`
	Tag                  *Tag         `json:"tag,omitempty"`
	TagId                string       `json:"tagId,omitempty"`
}
