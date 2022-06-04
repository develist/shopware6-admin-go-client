package entity

import "time"

// completed

type Document struct {
	// Config map[type:object] `json:"config,omitempty"`
	CreatedAt            *time.Time     `json:"createdAt,omitempty"`
	CustomFields         *[]CustomField `json:"customFields,omitempty"`
	DeepLinkCode         string         `json:"deepLinkCode,omitempty"`
	DependentDocuments   *Document      `json:"dependentDocuments,omitempty"`
	DocumentMediaFile    *Media         `json:"documentMediaFile,omitempty"`
	DocumentMediaFileId  string         `json:"documentMediaFileId,omitempty"`
	DocumentType         *DocumentType  `json:"documentType,omitempty"`
	DocumentTypeId       string         `json:"documentTypeId,omitempty"`
	FileType             string         `json:"fileType,omitempty"`
	Id                   string         `json:"id,omitempty"`
	Order                *Order         `json:"order,omitempty"`
	OrderId              string         `json:"orderId,omitempty"`
	OrderVersionId       string         `json:"orderVersionId,omitempty"`
	ReferencedDocument   *Document      `json:"referencedDocument,omitempty"`
	ReferencedDocumentId string         `json:"referencedDocumentId,omitempty"`
	Sent                 bool           `json:"sent,omitempty"`
	Static               bool           `json:"static,omitempty"`
	UpdatedAt            *time.Time     `json:"updatedAt,omitempty"`
}

type DocumentBaseConfig struct {
	// Config map[type:object] `json:"config,omitempty"`
	CreatedAt      *time.Time                      `json:"createdAt,omitempty"`
	CustomFields   *[]CustomField                  `json:"customFields,omitempty"`
	DocumentNumber string                          `json:"documentNumber,omitempty"`
	DocumentType   *DocumentType                   `json:"documentType,omitempty"`
	DocumentTypeId string                          `json:"documentTypeId,omitempty"`
	FilenamePrefix string                          `json:"filenamePrefix,omitempty"`
	FilenameSuffix string                          `json:"filenameSuffix,omitempty"`
	Global         bool                            `json:"global,omitempty"`
	Id             string                          `json:"id,omitempty"`
	Logo           *Media                          `json:"logo,omitempty"`
	LogoId         string                          `json:"logoId,omitempty"`
	Name           string                          `json:"name,omitempty"`
	SalesChannels  *DocumentBaseConfigSalesChannel `json:"salesChannels,omitempty"`
	UpdatedAt      *time.Time                      `json:"updatedAt,omitempty"`
}

type DocumentBaseConfigSalesChannel struct {
	CreatedAt            *time.Time          `json:"createdAt,omitempty"`
	DocumentBaseConfig   *DocumentBaseConfig `json:"documentBaseConfig,omitempty"`
	DocumentBaseConfigId string              `json:"documentBaseConfigId,omitempty"`
	DocumentType         *DocumentType       `json:"documentType,omitempty"`
	DocumentTypeId       string              `json:"documentTypeId,omitempty"`
	Id                   string              `json:"id,omitempty"`
	SalesChannel         *SalesChannel       `json:"salesChannel,omitempty"`
	SalesChannelId       string              `json:"salesChannelId,omitempty"`
	UpdatedAt            *time.Time          `json:"updatedAt,omitempty"`
}

type DocumentType struct {
	CreatedAt                       *time.Time                      `json:"createdAt,omitempty"`
	CustomFields                    *[]CustomField                  `json:"customFields,omitempty"`
	DocumentBaseConfigSalesChannels *DocumentBaseConfigSalesChannel `json:"documentBaseConfigSalesChannels,omitempty"`
	DocumentBaseConfigs             *DocumentBaseConfig             `json:"documentBaseConfigs,omitempty"`
	Documents                       *Document                       `json:"documents,omitempty"`
	Id                              string                          `json:"id,omitempty"`
	Name                            string                          `json:"name,omitempty"`
	TechnicalName                   string                          `json:"technicalName,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}