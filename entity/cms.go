package entity

import "time"

// completed

type CmsBlock struct {
	BackgroundColor     string         `json:"backgroundColor,omitempty"`
	BackgroundMedia     *Media         `json:"backgroundMedia,omitempty"`
	BackgroundMediaId   string         `json:"backgroundMediaId,omitempty"`
	BackgroundMediaMode string         `json:"backgroundMediaMode,omitempty"`
	CmsSectionVersionId string         `json:"cmsSectionVersionId,omitempty"`
	CreatedAt           time.Time      `json:"createdAt,omitempty"`
	CssClass            string         `json:"cssClass,omitempty"`
	CustomFields        *[]CustomField `json:"customFields,omitempty"`
	Id                  string         `json:"id,omitempty"`
	Locked              bool           `json:"locked,omitempty"`
	MarginBottom        string         `json:"marginBottom,omitempty"`
	MarginLeft          string         `json:"marginLeft,omitempty"`
	MarginRight         string         `json:"marginRight,omitempty"`
	MarginTop           string         `json:"marginTop,omitempty"`
	Name                string         `json:"name,omitempty"`
	Position            int            `json:"position,omitempty"`
	Section             *CmsSection    `json:"section,omitempty"`
	SectionId           string         `json:"sectionId,omitempty"`
	SectionPosition     string         `json:"sectionPosition,omitempty"`
	Slots               *CmsSlot       `json:"slots,omitempty"`
	Type                string         `json:"type,omitempty"`
	UpdatedAt           time.Time      `json:"updatedAt,omitempty"`
	VersionId           string         `json:"versionId,omitempty"`
}

type CmsPage struct {
	Categories *Category `json:"categories,omitempty"`
	// Config map[properties:map[backgroundColor:map[type:string]] type:object] `json:"config,omitempty"`
	CreatedAt         time.Time      `json:"createdAt,omitempty"`
	CustomFields      *[]CustomField `json:"customFields,omitempty"`
	Entity            string         `json:"entity,omitempty"`
	HomeSalesChannels *SalesChannel  `json:"homeSalesChannels,omitempty"`
	Id                string         `json:"id,omitempty"`
	LandingPages      *LandingPage   `json:"landingPages,omitempty"`
	Locked            bool           `json:"locked,omitempty"`
	Name              string         `json:"name,omitempty"`
	PreviewMedia      *Media         `json:"previewMedia,omitempty"`
	PreviewMediaId    string         `json:"previewMediaId,omitempty"`
	Products          *Product       `json:"products,omitempty"`
	Sections          *CmsSection    `json:"sections,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	Type      string    `json:"type,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	VersionId string    `json:"versionId,omitempty"`
}

type CmsSection struct {
	BackgroundColor     string         `json:"backgroundColor,omitempty"`
	BackgroundMedia     *Media         `json:"backgroundMedia,omitempty"`
	BackgroundMediaId   string         `json:"backgroundMediaId,omitempty"`
	BackgroundMediaMode string         `json:"backgroundMediaMode,omitempty"`
	Blocks              *CmsBlock      `json:"blocks,omitempty"`
	CmsPageVersionId    string         `json:"cmsPageVersionId,omitempty"`
	CreatedAt           time.Time      `json:"createdAt,omitempty"`
	CssClass            string         `json:"cssClass,omitempty"`
	CustomFields        *[]CustomField `json:"customFields,omitempty"`
	Id                  string         `json:"id,omitempty"`
	Locked              bool           `json:"locked,omitempty"`
	MobileBehavior      string         `json:"mobileBehavior,omitempty"`
	Name                string         `json:"name,omitempty"`
	Page                *CmsPage       `json:"page,omitempty"`
	PageId              string         `json:"pageId,omitempty"`
	Position            int            `json:"position,omitempty"`
	SizingMode          string         `json:"sizingMode,omitempty"`
	Type                string         `json:"type,omitempty"`
	UpdatedAt           time.Time      `json:"updatedAt,omitempty"`
	VersionId           string         `json:"versionId,omitempty"`
}

type CmsSlot struct {
	Block             *CmsBlock `json:"block,omitempty"`
	BlockId           string    `json:"blockId,omitempty"`
	CmsBlockVersionId string    `json:"cmsBlockVersionId,omitempty"`
	// Config map[type:object] `json:"config,omitempty"`
	CreatedAt    time.Time      `json:"createdAt,omitempty"`
	CustomFields *[]CustomField `json:"customFields,omitempty"`
	// Data map[readOnly:true type:object] `json:"data,omitempty"`
	Id     string `json:"id,omitempty"`
	Locked bool   `json:"locked,omitempty"`
	Slot   string `json:"slot,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	Type      string    `json:"type,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	VersionId string    `json:"versionId,omitempty"`
}
