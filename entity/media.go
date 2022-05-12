package entity

import "time"

// completed

type Media struct {
	Alt                 string              `json:"alt,omitempty"`
	AppPaymentMethods   *AppPaymentMethod   `json:"appPaymentMethods,omitempty"`
	AvatarUser          *User               `json:"avatarUser,omitempty"`
	Categories          *Category           `json:"categories,omitempty"`
	CmsBlocks           *CmsBlock           `json:"cmsBlocks,omitempty"`
	CmsPages            *CmsPage            `json:"cmsPages,omitempty"`
	CmsSections         *CmsSection         `json:"cmsSections,omitempty"`
	CreatedAt           time.Time           `json:"createdAt,omitempty"`
	CustomFields        *[]CustomField      `json:"customFields,omitempty"`
	DocumentBaseConfigs *DocumentBaseConfig `json:"documentBaseConfigs,omitempty"`
	Documents           *Document           `json:"documents,omitempty"`
	FileExtension       string              `json:"fileExtension,omitempty"`
	FileName            string              `json:"fileName,omitempty"`
	FileSize            int                 `json:"fileSize,omitempty"`
	HasFile             bool                `json:"hasFile,omitempty"`
	Id                  string              `json:"id,omitempty"`
	MailTemplateMedia   *MailTemplateMedia  `json:"mailTemplateMedia,omitempty"`
	MediaFolder         *MediaFolder        `json:"mediaFolder,omitempty"`
	MediaFolderId       string              `json:"mediaFolderId,omitempty"`
	// MediaType map[readOnly:true type:object] `json:"mediaType,omitempty"`
	// MetaData map[readOnly:true type:object] `json:"metaData,omitempty"`
	MimeType                    string                      `json:"mimeType,omitempty"`
	OrderLineItems              *OrderLineItem              `json:"orderLineItems,omitempty"`
	PaymentMethods              *PaymentMethod              `json:"paymentMethods,omitempty"`
	Private                     bool                        `json:"private,omitempty"`
	ProductConfiguratorSettings *ProductConfiguratorSetting `json:"productConfiguratorSettings,omitempty"`
	ProductManufacturers        *ProductManufacturer        `json:"productManufacturers,omitempty"`
	ProductMedia                *ProductMedia               `json:"productMedia,omitempty"`
	PropertyGroupOptions        *PropertyGroupOption        `json:"propertyGroupOptions,omitempty"`
	ShippingMethods             *ShippingMethod             `json:"shippingMethods,omitempty"`
	Tags                        *Tag                        `json:"tags,omitempty"`
	Thumbnails                  *MediaThumbnail             `json:"thumbnails,omitempty"`
	Title                       string                      `json:"title,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
	UploadedAt time.Time `json:"uploadedAt,omitempty"`
	Url        string    `json:"url,omitempty"`
	User       *User     `json:"user,omitempty"`
	UserId     string    `json:"userId,omitempty"`
}

type MediaDefaultFolder struct {
	// AssociationFields map[items:map[type:string] type:array] `json:"associationFields,omitempty"`
	CreatedAt    time.Time      `json:"createdAt,omitempty"`
	CustomFields *[]CustomField `json:"customFields,omitempty"`
	Entity       string         `json:"entity,omitempty"`
	Folder       *MediaFolder   `json:"folder,omitempty"`
	Id           string         `json:"id,omitempty"`
	UpdatedAt    time.Time      `json:"updatedAt,omitempty"`
}

type MediaFolder struct {
	ChildCount             int                       `json:"childCount,omitempty"`
	Children               *MediaFolder              `json:"children,omitempty"`
	Configuration          *MediaFolderConfiguration `json:"configuration,omitempty"`
	ConfigurationId        string                    `json:"configurationId,omitempty"`
	CreatedAt              time.Time                 `json:"createdAt,omitempty"`
	CustomFields           *[]CustomField            `json:"customFields,omitempty"`
	DefaultFolder          *MediaDefaultFolder       `json:"defaultFolder,omitempty"`
	DefaultFolderId        string                    `json:"defaultFolderId,omitempty"`
	Id                     string                    `json:"id,omitempty"`
	Media                  *Media                    `json:"media,omitempty"`
	Name                   string                    `json:"name,omitempty"`
	Parent                 *MediaFolder              `json:"parent,omitempty"`
	ParentId               string                    `json:"parentId,omitempty"`
	UpdatedAt              time.Time                 `json:"updatedAt,omitempty"`
	UseParentConfiguration bool                      `json:"useParentConfiguration,omitempty"`
}

type MediaFolderConfiguration struct {
	CreateThumbnails    bool                `json:"createThumbnails,omitempty"`
	CreatedAt           time.Time           `json:"createdAt,omitempty"`
	CustomFields        *[]CustomField      `json:"customFields,omitempty"`
	Id                  string              `json:"id,omitempty"`
	KeepAspectRatio     bool                `json:"keepAspectRatio,omitempty"`
	MediaFolders        *MediaFolder        `json:"mediaFolders,omitempty"`
	MediaThumbnailSizes *MediaThumbnailSize `json:"mediaThumbnailSizes,omitempty"`
	NoAssociation       bool                `json:"noAssociation,omitempty"`
	Private             bool                `json:"private,omitempty"`
	ThumbnailQuality    int                 `json:"thumbnailQuality,omitempty"`
	UpdatedAt           time.Time           `json:"updatedAt,omitempty"`
}

type MediaFolderConfigurationMediaThumbnailSize struct {
	Id                         string                    `json:"id,omitempty"`
	MediaFolderConfiguration   *MediaFolderConfiguration `json:"mediaFolderConfiguration,omitempty"`
	MediaFolderConfigurationId string                    `json:"mediaFolderConfigurationId,omitempty"`
	MediaThumbnailSize         *MediaThumbnailSize       `json:"mediaThumbnailSize,omitempty"`
	MediaThumbnailSizeId       string                    `json:"mediaThumbnailSizeId,omitempty"`
}

type MediaTag struct {
	Id      string `json:"id,omitempty"`
	Media   *Media `json:"media,omitempty"`
	MediaId string `json:"mediaId,omitempty"`
	Tag     *Tag   `json:"tag,omitempty"`
	TagId   string `json:"tagId,omitempty"`
}

type MediaThumbnail struct {
	CreatedAt    time.Time      `json:"createdAt,omitempty"`
	CustomFields *[]CustomField `json:"customFields,omitempty"`
	Height       int            `json:"height,omitempty"`
	Id           string         `json:"id,omitempty"`
	Media        *Media         `json:"media,omitempty"`
	MediaId      string         `json:"mediaId,omitempty"`
	UpdatedAt    time.Time      `json:"updatedAt,omitempty"`
	Url          string         `json:"url,omitempty"`
	Width        int            `json:"width,omitempty"`
}

type MediaThumbnailSize struct {
	CreatedAt                 time.Time                 `json:"createdAt,omitempty"`
	CustomFields              *[]CustomField            `json:"customFields,omitempty"`
	Height                    int                       `json:"height,omitempty"`
	Id                        string                    `json:"id,omitempty"`
	MediaFolderConfigurations *MediaFolderConfiguration `json:"mediaFolderConfigurations,omitempty"`
	UpdatedAt                 time.Time                 `json:"updatedAt,omitempty"`
	Width                     int                       `json:"width,omitempty"`
}
