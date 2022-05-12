package entity

import "time"

// completed

type MailHeaderFooter struct {
	CreatedAt     time.Time     `json:"createdAt,omitempty"`
	Description   string        `json:"description,omitempty"`
	FooterHtml    string        `json:"footerHtml,omitempty"`
	FooterPlain   string        `json:"footerPlain,omitempty"`
	HeaderHtml    string        `json:"headerHtml,omitempty"`
	HeaderPlain   string        `json:"headerPlain,omitempty"`
	Id            string        `json:"id,omitempty"`
	Name          string        `json:"name,omitempty"`
	SalesChannels *SalesChannel `json:"salesChannels,omitempty"`
	SystemDefault bool          `json:"systemDefault,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type MailTemplate struct {
	ContentHtml        string             `json:"contentHtml,omitempty"`
	ContentPlain       string             `json:"contentPlain,omitempty"`
	CreatedAt          time.Time          `json:"createdAt,omitempty"`
	CustomFields       *[]CustomField     `json:"customFields,omitempty"`
	Description        string             `json:"description,omitempty"`
	Id                 string             `json:"id,omitempty"`
	MailTemplateType   *MailTemplateType  `json:"mailTemplateType,omitempty"`
	MailTemplateTypeId string             `json:"mailTemplateTypeId,omitempty"`
	Media              *MailTemplateMedia `json:"media,omitempty"`
	SenderName         string             `json:"senderName,omitempty"`
	Subject            string             `json:"subject,omitempty"`
	SystemDefault      bool               `json:"systemDefault,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type MailTemplateMedia struct {
	Id             string        `json:"id,omitempty"`
	LanguageId     string        `json:"languageId,omitempty"`
	MailTemplate   *MailTemplate `json:"mailTemplate,omitempty"`
	MailTemplateId string        `json:"mailTemplateId,omitempty"`
	Media          *Media        `json:"media,omitempty"`
	MediaId        string        `json:"mediaId,omitempty"`
	Position       int           `json:"position,omitempty"`
}

type MailTemplateType struct {
	// AvailableEntities map[type:object] `json:"availableEntities,omitempty"`
	CreatedAt     time.Time      `json:"createdAt,omitempty"`
	CustomFields  *[]CustomField `json:"customFields,omitempty"`
	Id            string         `json:"id,omitempty"`
	MailTemplates *MailTemplate  `json:"mailTemplates,omitempty"`
	Name          string         `json:"name,omitempty"`
	TechnicalName string         `json:"technicalName,omitempty"`
	// TemplateData map[type:object] `json:"templateData,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
