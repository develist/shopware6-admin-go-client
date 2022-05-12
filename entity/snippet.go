package entity

import "time"

// completed

type Snippet struct {
	Author         string         `json:"author,omitempty"`
	CreatedAt      time.Time      `json:"createdAt,omitempty"`
	CustomFields   *[]CustomField `json:"customFields,omitempty"`
	Id             string         `json:"id,omitempty"`
	Set            *SnippetSet    `json:"set,omitempty"`
	SetId          string         `json:"setId,omitempty"`
	TranslationKey string         `json:"translationKey,omitempty"`
	UpdatedAt      time.Time      `json:"updatedAt,omitempty"`
	Value          string         `json:"value,omitempty"`
}

type SnippetSet struct {
	BaseFile            string              `json:"baseFile,omitempty"`
	CreatedAt           time.Time           `json:"createdAt,omitempty"`
	CustomFields        *[]CustomField      `json:"customFields,omitempty"`
	Id                  string              `json:"id,omitempty"`
	Iso                 string              `json:"iso,omitempty"`
	Name                string              `json:"name,omitempty"`
	SalesChannelDomains *SalesChannelDomain `json:"salesChannelDomains,omitempty"`
	Snippets            *Snippet            `json:"snippets,omitempty"`
	UpdatedAt           time.Time           `json:"updatedAt,omitempty"`
}
