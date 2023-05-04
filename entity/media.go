package entity

import "time"

// Generated from Shopware Admin API
// Version 6.5.9999999.9999999-dev at 2023-05-02 19:06:34 UTC

// Media data structure
// Added since version: 6.0.0.0
// Required fields: createdAt
type Media struct {
	Alt                         *string                     `json:"alt,omitempty"`
	AppPaymentMethods           *AppPaymentMethod           `json:"appPaymentMethods,omitempty"`
	AvatarUser                  *User                       `json:"avatarUser,omitempty"`
	Categories                  *Category                   `json:"categories,omitempty"`
	CmsBlocks                   *CmsBlock                   `json:"cmsBlocks,omitempty"`
	CmsPages                    *CmsPage                    `json:"cmsPages,omitempty"`
	CmsSections                 *CmsSection                 `json:"cmsSections,omitempty"`
	CreatedAt                   *time.Time                  `json:"createdAt,omitempty"`
	CustomFields                *[]CustomField              `json:"customFields,omitempty"`
	DocumentBaseConfigs         *DocumentBaseConfig         `json:"documentBaseConfigs,omitempty"`
	Documents                   *Document                   `json:"documents,omitempty"`
	Extensions                  *interface{}                `json:"extensions,omitempty"` // map[properties:map[themeMedia:map[properties:map[data:map[items:map[properties:map[id:map[example:345e91ec37704ba897929b306c0b943b type:string] type:map[example:theme type:string]] type:object] type:array] links:map[properties:map[related:map[example:/media/518b235332044ad98ca35ebe92d86df8/themeMedia format:uri-reference type:string]] type:object]] type:object] themes:map[properties:map[data:map[items:map[properties:map[id:map[example:443a62a579814c289605a5d9650b0ca1 type:string] type:map[example:theme type:string]] type:object] type:array] links:map[properties:map[related:map[example:/media/518b235332044ad98ca35ebe92d86df8/themes format:uri-reference type:string]] type:object]] type:object]] type:object]
	FileExtension               *string                     `json:"fileExtension,omitempty"`
	FileName                    *string                     `json:"fileName,omitempty"`
	FileSize                    *int                        `json:"fileSize,omitempty"`
	HasFile                     *bool                       `json:"hasFile,omitempty"`
	Id                          string                      `json:"id,omitempty"`
	MailTemplateMedia           *MailTemplateMedia          `json:"mailTemplateMedia,omitempty"`
	MediaFolder                 *MediaFolder                `json:"mediaFolder,omitempty"`
	MediaFolderId               *string                     `json:"mediaFolderId,omitempty"`
	MediaType                   *interface{}                `json:"mediaType,omitempty"` // map[readOnly:true type:object]
	MetaData                    *interface{}                `json:"metaData,omitempty"`  // map[readOnly:true type:object]
	MimeType                    *string                     `json:"mimeType,omitempty"`
	OrderLineItems              *OrderLineItem              `json:"orderLineItems,omitempty"`
	PaymentMethods              *PaymentMethod              `json:"paymentMethods,omitempty"`
	Private                     *bool                       `json:"private,omitempty"`
	ProductConfiguratorSettings *ProductConfiguratorSetting `json:"productConfiguratorSettings,omitempty"`
	ProductManufacturers        *ProductManufacturer        `json:"productManufacturers,omitempty"`
	ProductMedia                *ProductMedia               `json:"productMedia,omitempty"`
	PropertyGroupOptions        *PropertyGroupOption        `json:"propertyGroupOptions,omitempty"`
	ShippingMethods             *ShippingMethod             `json:"shippingMethods,omitempty"`
	Tags                        *Tag                        `json:"tags,omitempty"`
	Thumbnails                  *MediaThumbnail             `json:"thumbnails,omitempty"`
	Title                       *string                     `json:"title,omitempty"`
	Translated                  *interface{}                `json:"translated,omitempty"` // map[type:object]
	UpdatedAt                   *time.Time                  `json:"updatedAt,omitempty"`
	UploadedAt                  *time.Time                  `json:"uploadedAt,omitempty"`
	Url                         *string                     `json:"url,omitempty"`
	User                        *User                       `json:"user,omitempty"`
	UserId                      *string                     `json:"userId,omitempty"`
}

// MediaDefaultFolder data structure
// Added since version: 6.0.0.0
// Required fields: associationFields, entity, createdAt
type MediaDefaultFolder struct {
	AssociationFields *[]string      `json:"associationFields,omitempty"` // map[items:map[type:string] type:array]
	CreatedAt         *time.Time     `json:"createdAt,omitempty"`
	CustomFields      *[]CustomField `json:"customFields,omitempty"`
	Entity            string         `json:"entity,omitempty"`
	Folder            *MediaFolder   `json:"folder,omitempty"`
	Id                string         `json:"id,omitempty"`
	UpdatedAt         *time.Time     `json:"updatedAt,omitempty"`
}

// MediaFolder data structure
// Added since version: 6.0.0.0
// Required fields: configurationId, name, createdAt
type MediaFolder struct {
	ChildCount             *int                      `json:"childCount,omitempty"`
	Children               *MediaFolder              `json:"children,omitempty"`
	Configuration          *MediaFolderConfiguration `json:"configuration,omitempty"`
	ConfigurationId        string                    `json:"configurationId,omitempty"`
	CreatedAt              *time.Time                `json:"createdAt,omitempty"`
	CustomFields           *[]CustomField            `json:"customFields,omitempty"`
	DefaultFolder          *MediaDefaultFolder       `json:"defaultFolder,omitempty"`
	DefaultFolderId        *string                   `json:"defaultFolderId,omitempty"`
	Id                     string                    `json:"id,omitempty"`
	Media                  *Media                    `json:"media,omitempty"`
	Name                   string                    `json:"name,omitempty"`
	Parent                 *MediaFolder              `json:"parent,omitempty"`
	ParentId               *string                   `json:"parentId,omitempty"`
	Path                   *string                   `json:"path,omitempty"`
	UpdatedAt              *time.Time                `json:"updatedAt,omitempty"`
	UseParentConfiguration *bool                     `json:"useParentConfiguration,omitempty"`
}

// MediaFolderConfiguration data structure
// Added since version: 6.0.0.0
// Required fields: createdAt
type MediaFolderConfiguration struct {
	CreateThumbnails    *bool               `json:"createThumbnails,omitempty"`
	CreatedAt           *time.Time          `json:"createdAt,omitempty"`
	CustomFields        *[]CustomField      `json:"customFields,omitempty"`
	Id                  string              `json:"id,omitempty"`
	KeepAspectRatio     *bool               `json:"keepAspectRatio,omitempty"`
	MediaFolders        *MediaFolder        `json:"mediaFolders,omitempty"`
	MediaThumbnailSizes *MediaThumbnailSize `json:"mediaThumbnailSizes,omitempty"`
	NoAssociation       *bool               `json:"noAssociation,omitempty"`
	Private             *bool               `json:"private,omitempty"`
	ThumbnailQuality    *int                `json:"thumbnailQuality,omitempty"`
	UpdatedAt           *time.Time          `json:"updatedAt,omitempty"`
}

// MediaFolderConfigurationMediaThumbnailSize data structure
// Added since version: 6.0.0.0
// Required fields: mediaFolderConfigurationId, mediaThumbnailSizeId
type MediaFolderConfigurationMediaThumbnailSize struct {
	Id                         string                    `json:"id,omitempty"`
	MediaFolderConfiguration   *MediaFolderConfiguration `json:"mediaFolderConfiguration,omitempty"`
	MediaFolderConfigurationId string                    `json:"mediaFolderConfigurationId,omitempty"`
	MediaThumbnailSize         *MediaThumbnailSize       `json:"mediaThumbnailSize,omitempty"`
	MediaThumbnailSizeId       string                    `json:"mediaThumbnailSizeId,omitempty"`
}

// MediaTag data structure
// Added since version: 6.0.0.0
// Required fields: mediaId, tagId
type MediaTag struct {
	Id      string `json:"id,omitempty"`
	Media   *Media `json:"media,omitempty"`
	MediaId string `json:"mediaId,omitempty"`
	Tag     *Tag   `json:"tag,omitempty"`
	TagId   string `json:"tagId,omitempty"`
}

// MediaThumbnail data structure
// Added since version: 6.0.0.0
// Required fields: mediaId, width, height, createdAt
type MediaThumbnail struct {
	CreatedAt    *time.Time     `json:"createdAt,omitempty"`
	CustomFields *[]CustomField `json:"customFields,omitempty"`
	Height       int            `json:"height,omitempty"`
	Id           string         `json:"id,omitempty"`
	Media        *Media         `json:"media,omitempty"`
	MediaId      string         `json:"mediaId,omitempty"`
	UpdatedAt    *time.Time     `json:"updatedAt,omitempty"`
	Url          *string        `json:"url,omitempty"`
	Width        int            `json:"width,omitempty"`
}

// MediaThumbnailSize data structure
// Added since version: 6.0.0.0
// Required fields: width, height, createdAt
type MediaThumbnailSize struct {
	CreatedAt                 *time.Time                `json:"createdAt,omitempty"`
	CustomFields              *[]CustomField            `json:"customFields,omitempty"`
	Height                    int                       `json:"height,omitempty"`
	Id                        string                    `json:"id,omitempty"`
	MediaFolderConfigurations *MediaFolderConfiguration `json:"mediaFolderConfigurations,omitempty"`
	UpdatedAt                 *time.Time                `json:"updatedAt,omitempty"`
	Width                     int                       `json:"width,omitempty"`
}
