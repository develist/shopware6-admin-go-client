package entity

import "time"

// completed

type LandingPage struct {
	Active           bool           `json:"active,omitempty"`
	CmsPage          *CmsPage       `json:"cmsPage,omitempty"`
	CmsPageId        string         `json:"cmsPageId,omitempty"`
	CmsPageVersionId string         `json:"cmsPageVersionId,omitempty"`
	CreatedAt        time.Time      `json:"createdAt,omitempty"`
	CustomFields     *[]CustomField `json:"customFields,omitempty"`
	Id               string         `json:"id,omitempty"`
	Keywords         string         `json:"keywords,omitempty"`
	MetaDescription  string         `json:"metaDescription,omitempty"`
	MetaTitle        string         `json:"metaTitle,omitempty"`
	Name             string         `json:"name,omitempty"`
	SalesChannels    *SalesChannel  `json:"salesChannels,omitempty"`
	SeoUrls          *SeoUrl        `json:"seoUrls,omitempty"`
	// SlotConfig map[type:object] `json:"slotConfig,omitempty"`
	Tags *Tag `json:"tags,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	Url       string    `json:"url,omitempty"`
	VersionId string    `json:"versionId,omitempty"`
}

type LandingPageSalesChannel struct {
	Id                   string        `json:"id,omitempty"`
	LandingPage          *LandingPage  `json:"landingPage,omitempty"`
	LandingPageId        string        `json:"landingPageId,omitempty"`
	LandingPageVersionId string        `json:"landingPageVersionId,omitempty"`
	SalesChannel         *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId       string        `json:"salesChannelId,omitempty"`
}

type LandingPageTag struct {
	Id                   string       `json:"id,omitempty"`
	LandingPage          *LandingPage `json:"landingPage,omitempty"`
	LandingPageId        string       `json:"landingPageId,omitempty"`
	LandingPageVersionId string       `json:"landingPageVersionId,omitempty"`
	Tag                  *Tag         `json:"tag,omitempty"`
	TagId                string       `json:"tagId,omitempty"`
}
