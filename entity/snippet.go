package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-17 19:07:53 UTC

// Snippet data structure
// Added since version: 6.0.0.0
// Required fields: setId, translationKey, value, author, createdAt
type Snippet struct {
	Author         string         `json:"author,omitempty"`
	CreatedAt      *time.Time     `json:"createdAt,omitempty"`
	CustomFields   *[]CustomField `json:"customFields,omitempty"`
	Id             string         `json:"id,omitempty"`
	Set            *SnippetSet    `json:"set,omitempty"`
	SetId          string         `json:"setId,omitempty"`
	TranslationKey string         `json:"translationKey,omitempty"`
	UpdatedAt      *time.Time     `json:"updatedAt,omitempty"`
	Value          string         `json:"value,omitempty"`
}

// SnippetSet data structure
// Added since version: 6.0.0.0
// Required fields: name, baseFile, iso, createdAt
type SnippetSet struct {
	BaseFile            string              `json:"baseFile,omitempty"`
	CreatedAt           *time.Time          `json:"createdAt,omitempty"`
	CustomFields        *[]CustomField      `json:"customFields,omitempty"`
	Id                  string              `json:"id,omitempty"`
	Iso                 string              `json:"iso,omitempty"`
	Name                string              `json:"name,omitempty"`
	SalesChannelDomains *SalesChannelDomain `json:"salesChannelDomains,omitempty"`
	Snippets            *Snippet            `json:"snippets,omitempty"`
	UpdatedAt           *time.Time          `json:"updatedAt,omitempty"`
}
