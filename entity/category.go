package entity

import "time"

// completed

type Category struct {
	Active                 bool   `json:"active,omitempty"`
	AfterCategoryId        string `json:"afterCategoryId,omitempty"`
	AfterCategoryVersionId string `json:"afterCategoryVersionId,omitempty"`
	AutoIncrement          int    `json:"autoIncrement,omitempty"`
	// Breadcrumb map[items:map[type:string] readOnly:true type:array] `json:"breadcrumb,omitempty"`
	ChildCount              int            `json:"childCount,omitempty"`
	Children                *Category      `json:"children,omitempty"`
	CmsPage                 *CmsPage       `json:"cmsPage,omitempty"`
	CmsPageId               string         `json:"cmsPageId,omitempty"`
	CmsPageVersionId        string         `json:"cmsPageVersionId,omitempty"`
	CreatedAt               *time.Time     `json:"createdAt,omitempty"`
	CustomFields            *[]CustomField `json:"customFields,omitempty"`
	Description             string         `json:"description,omitempty"`
	DisplayNestedProducts   bool           `json:"displayNestedProducts,omitempty"`
	ExternalLink            string         `json:"externalLink,omitempty"`
	FooterSalesChannels     *SalesChannel  `json:"footerSalesChannels,omitempty"`
	Id                      string         `json:"id,omitempty"`
	InternalLink            string         `json:"internalLink,omitempty"`
	Keywords                string         `json:"keywords,omitempty"`
	Level                   int            `json:"level,omitempty"`
	LinkNewTab              bool           `json:"linkNewTab,omitempty"`
	LinkType                string         `json:"linkType,omitempty"`
	MainCategories          *MainCategory  `json:"mainCategories,omitempty"`
	Media                   *Media         `json:"media,omitempty"`
	MediaId                 string         `json:"mediaId,omitempty"`
	MetaDescription         string         `json:"metaDescription,omitempty"`
	MetaTitle               string         `json:"metaTitle,omitempty"`
	Name                    string         `json:"name,omitempty"`
	NavigationSalesChannels *SalesChannel  `json:"navigationSalesChannels,omitempty"`
	NestedProducts          *Product       `json:"nestedProducts,omitempty"`
	Parent                  *Category      `json:"parent,omitempty"`
	ParentId                string         `json:"parentId,omitempty"`
	ParentVersionId         string         `json:"parentVersionId,omitempty"`
	Path                    string         `json:"path,omitempty"`
	ProductAssignmentType   string         `json:"productAssignmentType,omitempty"`
	ProductStream           *ProductStream `json:"productStream,omitempty"`
	ProductStreamId         string         `json:"productStreamId,omitempty"`
	Products                *Product       `json:"products,omitempty"`
	SeoUrls                 *SeoUrl        `json:"seoUrls,omitempty"`
	ServiceSalesChannels    *SalesChannel  `json:"serviceSalesChannels,omitempty"`
	// SlotConfig map[type:object] `json:"slotConfig,omitempty"`
	Tags *Tag `json:"tags,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	Type      string     `json:"type,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	VersionId string     `json:"versionId,omitempty"`
	Visible   bool       `json:"visible,omitempty"`
}

type CategoryTag struct {
	Category          *Category `json:"category,omitempty"`
	CategoryId        string    `json:"categoryId,omitempty"`
	CategoryVersionId string    `json:"categoryVersionId,omitempty"`
	Id                string    `json:"id,omitempty"`
	Tag               *Tag      `json:"tag,omitempty"`
	TagId             string    `json:"tagId,omitempty"`
}

type MainCategory struct {
	Category          *Category     `json:"category,omitempty"`
	CategoryId        string        `json:"categoryId,omitempty"`
	CategoryVersionId string        `json:"categoryVersionId,omitempty"`
	CreatedAt         *time.Time    `json:"createdAt,omitempty"`
	Id                string        `json:"id,omitempty"`
	Product           *Product      `json:"product,omitempty"`
	ProductId         string        `json:"productId,omitempty"`
	ProductVersionId  string        `json:"productVersionId,omitempty"`
	SalesChannel      *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId    string        `json:"salesChannelId,omitempty"`
	UpdatedAt         *time.Time    `json:"updatedAt,omitempty"`
}
