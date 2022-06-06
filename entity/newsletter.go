package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-06 18:44:04 UTC

// NewsletterRecipient data structure
// Added since version: 6.0.0.0
// Required fields: email, status, hash, languageId, salesChannelId, createdAt
type NewsletterRecipient struct {
	City           string         `json:"city,omitempty"`
	ConfirmedAt    *time.Time     `json:"confirmedAt,omitempty"`
	CreatedAt      *time.Time     `json:"createdAt,omitempty"`
	CustomFields   *[]CustomField `json:"customFields,omitempty"`
	Email          string         `json:"email,omitempty"`
	FirstName      string         `json:"firstName,omitempty"`
	Hash           string         `json:"hash,omitempty"`
	Id             string         `json:"id,omitempty"`
	Language       *Language      `json:"language,omitempty"`
	LanguageId     string         `json:"languageId,omitempty"`
	LastName       string         `json:"lastName,omitempty"`
	SalesChannel   *SalesChannel  `json:"salesChannel,omitempty"`
	SalesChannelId string         `json:"salesChannelId,omitempty"`
	Salutation     *Salutation    `json:"salutation,omitempty"`
	SalutationId   string         `json:"salutationId,omitempty"`
	Status         string         `json:"status,omitempty"`
	Street         string         `json:"street,omitempty"`
	Tags           *Tag           `json:"tags,omitempty"`
	Title          string         `json:"title,omitempty"`
	UpdatedAt      *time.Time     `json:"updatedAt,omitempty"`
	ZipCode        string         `json:"zipCode,omitempty"`
}

// NewsletterRecipientTag data structure
// Added since version: 6.0.0.0
// Required fields: newsletterRecipientId, tagId
type NewsletterRecipientTag struct {
	Id                    string               `json:"id,omitempty"`
	NewsletterRecipient   *NewsletterRecipient `json:"newsletterRecipient,omitempty"`
	NewsletterRecipientId string               `json:"newsletterRecipientId,omitempty"`
	Tag                   *Tag                 `json:"tag,omitempty"`
	TagId                 string               `json:"tagId,omitempty"`
}
