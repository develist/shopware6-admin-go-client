package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-17 19:07:53 UTC

// CmsBlock data structure
// Added since version: 6.0.0.0
// Required fields: position, type, sectionId, createdAt
type CmsBlock struct {
	BackgroundColor     *string        `json:"backgroundColor,omitempty"`
	BackgroundMedia     *Media         `json:"backgroundMedia,omitempty"`
	BackgroundMediaId   *string        `json:"backgroundMediaId,omitempty"`
	BackgroundMediaMode *string        `json:"backgroundMediaMode,omitempty"`
	CmsSectionVersionId *string        `json:"cmsSectionVersionId,omitempty"`
	CreatedAt           *time.Time     `json:"createdAt,omitempty"`
	CssClass            *string        `json:"cssClass,omitempty"`
	CustomFields        *[]CustomField `json:"customFields,omitempty"`
	Id                  string         `json:"id,omitempty"`
	Locked              *bool          `json:"locked,omitempty"`
	MarginBottom        *string        `json:"marginBottom,omitempty"`
	MarginLeft          *string        `json:"marginLeft,omitempty"`
	MarginRight         *string        `json:"marginRight,omitempty"`
	MarginTop           *string        `json:"marginTop,omitempty"`
	Name                *string        `json:"name,omitempty"`
	Position            int            `json:"position,omitempty"`
	Section             *CmsSection    `json:"section,omitempty"`
	SectionId           string         `json:"sectionId,omitempty"`
	SectionPosition     *string        `json:"sectionPosition,omitempty"`
	Slots               *CmsSlot       `json:"slots,omitempty"`
	Type                string         `json:"type,omitempty"`
	UpdatedAt           *time.Time     `json:"updatedAt,omitempty"`
	VersionId           *string        `json:"versionId,omitempty"`
}

// CmsPage data structure
// Added since version: 6.0.0.0
// Required fields: type, createdAt, name
type CmsPage struct {
	Categories        *Category      `json:"categories,omitempty"`
	Config            *interface{}   `json:"config,omitempty"` // map[properties:map[backgroundColor:map[type:string]] type:object]
	CreatedAt         *time.Time     `json:"createdAt,omitempty"`
	CustomFields      *[]CustomField `json:"customFields,omitempty"`
	Entity            *string        `json:"entity,omitempty"`
	HomeSalesChannels *SalesChannel  `json:"homeSalesChannels,omitempty"`
	Id                string         `json:"id,omitempty"`
	LandingPages      *LandingPage   `json:"landingPages,omitempty"`
	Locked            *bool          `json:"locked,omitempty"`
	Name              string         `json:"name,omitempty"`
	PreviewMedia      *Media         `json:"previewMedia,omitempty"`
	PreviewMediaId    *string        `json:"previewMediaId,omitempty"`
	Products          *Product       `json:"products,omitempty"`
	Sections          *CmsSection    `json:"sections,omitempty"`
	Translated        *interface{}   `json:"translated,omitempty"` // map[type:object]
	Type              string         `json:"type,omitempty"`
	UpdatedAt         *time.Time     `json:"updatedAt,omitempty"`
	VersionId         *string        `json:"versionId,omitempty"`
}

// CmsSection data structure
// Added since version: 6.0.0.0
// Required fields: position, type, pageId, createdAt
type CmsSection struct {
	BackgroundColor     *string        `json:"backgroundColor,omitempty"`
	BackgroundMedia     *Media         `json:"backgroundMedia,omitempty"`
	BackgroundMediaId   *string        `json:"backgroundMediaId,omitempty"`
	BackgroundMediaMode *string        `json:"backgroundMediaMode,omitempty"`
	Blocks              *CmsBlock      `json:"blocks,omitempty"`
	CmsPageVersionId    *string        `json:"cmsPageVersionId,omitempty"`
	CreatedAt           *time.Time     `json:"createdAt,omitempty"`
	CssClass            *string        `json:"cssClass,omitempty"`
	CustomFields        *[]CustomField `json:"customFields,omitempty"`
	Id                  string         `json:"id,omitempty"`
	Locked              *bool          `json:"locked,omitempty"`
	MobileBehavior      *string        `json:"mobileBehavior,omitempty"`
	Name                *string        `json:"name,omitempty"`
	Page                *CmsPage       `json:"page,omitempty"`
	PageId              string         `json:"pageId,omitempty"`
	Position            int            `json:"position,omitempty"`
	SizingMode          *string        `json:"sizingMode,omitempty"`
	Type                string         `json:"type,omitempty"`
	UpdatedAt           *time.Time     `json:"updatedAt,omitempty"`
	VersionId           *string        `json:"versionId,omitempty"`
}

// CmsSlot data structure
// Added since version: 6.0.0.0
// Required fields: type, slot, blockId, createdAt
type CmsSlot struct {
	Block             *CmsBlock      `json:"block,omitempty"`
	BlockId           string         `json:"blockId,omitempty"`
	CmsBlockVersionId *string        `json:"cmsBlockVersionId,omitempty"`
	Config            *interface{}   `json:"config,omitempty"` // map[type:object]
	CreatedAt         *time.Time     `json:"createdAt,omitempty"`
	CustomFields      *[]CustomField `json:"customFields,omitempty"`
	Data              *interface{}   `json:"data,omitempty"` // map[readOnly:true type:object]
	Id                string         `json:"id,omitempty"`
	Locked            *bool          `json:"locked,omitempty"`
	Slot              string         `json:"slot,omitempty"`
	Translated        *interface{}   `json:"translated,omitempty"` // map[type:object]
	Type              string         `json:"type,omitempty"`
	UpdatedAt         *time.Time     `json:"updatedAt,omitempty"`
	VersionId         *string        `json:"versionId,omitempty"`
}
