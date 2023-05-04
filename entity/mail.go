package entity

import "time"

// Generated from Shopware Admin API
// Version 6.5.9999999.9999999-dev at 2023-05-02 19:06:34 UTC

// MailHeaderFooter data structure
// Added since version: 6.0.0.0
// Required fields: createdAt, name
type MailHeaderFooter struct {
	CreatedAt     *time.Time    `json:"createdAt,omitempty"`
	Description   *string       `json:"description,omitempty"`
	FooterHtml    *string       `json:"footerHtml,omitempty"`
	FooterPlain   *string       `json:"footerPlain,omitempty"`
	HeaderHtml    *string       `json:"headerHtml,omitempty"`
	HeaderPlain   *string       `json:"headerPlain,omitempty"`
	Id            string        `json:"id,omitempty"`
	Name          string        `json:"name,omitempty"`
	SalesChannels *SalesChannel `json:"salesChannels,omitempty"`
	SystemDefault *bool         `json:"systemDefault,omitempty"`
	Translated    *interface{}  `json:"translated,omitempty"` // map[type:object]
	UpdatedAt     *time.Time    `json:"updatedAt,omitempty"`
}

// MailTemplate data structure
// Added since version: 6.0.0.0
// Required fields: mailTemplateTypeId, createdAt, subject, contentHtml, contentPlain
type MailTemplate struct {
	ContentHtml        string             `json:"contentHtml,omitempty"`
	ContentPlain       string             `json:"contentPlain,omitempty"`
	CreatedAt          *time.Time         `json:"createdAt,omitempty"`
	CustomFields       *[]CustomField     `json:"customFields,omitempty"`
	Description        *string            `json:"description,omitempty"`
	Id                 string             `json:"id,omitempty"`
	MailTemplateType   *MailTemplateType  `json:"mailTemplateType,omitempty"`
	MailTemplateTypeId string             `json:"mailTemplateTypeId,omitempty"`
	Media              *MailTemplateMedia `json:"media,omitempty"`
	SenderName         *string            `json:"senderName,omitempty"`
	Subject            string             `json:"subject,omitempty"`
	SystemDefault      *bool              `json:"systemDefault,omitempty"`
	Translated         *interface{}       `json:"translated,omitempty"` // map[type:object]
	UpdatedAt          *time.Time         `json:"updatedAt,omitempty"`
}

// MailTemplateMedia data structure
// Added since version: 6.0.0.0
// Required fields: mailTemplateId, languageId, mediaId
type MailTemplateMedia struct {
	Id             string        `json:"id,omitempty"`
	LanguageId     string        `json:"languageId,omitempty"`
	MailTemplate   *MailTemplate `json:"mailTemplate,omitempty"`
	MailTemplateId string        `json:"mailTemplateId,omitempty"`
	Media          *Media        `json:"media,omitempty"`
	MediaId        string        `json:"mediaId,omitempty"`
	Position       *int          `json:"position,omitempty"`
}

// MailTemplateType data structure
// Added since version: 6.0.0.0
// Required fields: technicalName, createdAt, name
type MailTemplateType struct {
	AvailableEntities *interface{}   `json:"availableEntities,omitempty"` // map[type:object]
	CreatedAt         *time.Time     `json:"createdAt,omitempty"`
	CustomFields      *[]CustomField `json:"customFields,omitempty"`
	Id                string         `json:"id,omitempty"`
	MailTemplates     *MailTemplate  `json:"mailTemplates,omitempty"`
	Name              string         `json:"name,omitempty"`
	TechnicalName     string         `json:"technicalName,omitempty"`
	TemplateData      *interface{}   `json:"templateData,omitempty"` // map[type:object]
	Translated        *interface{}   `json:"translated,omitempty"`   // map[type:object]
	UpdatedAt         *time.Time     `json:"updatedAt,omitempty"`
}
