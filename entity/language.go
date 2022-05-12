package entity

import "time"

// completed

type Language struct {
	Children                       *Language                 `json:"children,omitempty"`
	CreatedAt                      time.Time                 `json:"createdAt,omitempty"`
	CustomFields                   *[]CustomField            `json:"customFields,omitempty"`
	Customers                      *Customer                 `json:"customers,omitempty"`
	Id                             string                    `json:"id,omitempty"`
	Locale                         *Locale                   `json:"locale,omitempty"`
	LocaleId                       string                    `json:"localeId,omitempty"`
	Name                           string                    `json:"name,omitempty"`
	NewsletterRecipients           *NewsletterRecipient      `json:"newsletterRecipients,omitempty"`
	Orders                         *Order                    `json:"orders,omitempty"`
	Parent                         *Language                 `json:"parent,omitempty"`
	ParentId                       string                    `json:"parentId,omitempty"`
	ProductKeywordDictionaries     *ProductKeywordDictionary `json:"productKeywordDictionaries,omitempty"`
	ProductReviews                 *ProductReview            `json:"productReviews,omitempty"`
	ProductSearchConfig            *ProductSearchConfig      `json:"productSearchConfig,omitempty"`
	ProductSearchKeywords          *ProductSearchKeyword     `json:"productSearchKeywords,omitempty"`
	SalesChannelDefaultAssignments *SalesChannel             `json:"salesChannelDefaultAssignments,omitempty"`
	SalesChannelDomains            *SalesChannelDomain       `json:"salesChannelDomains,omitempty"`
	SalesChannels                  *SalesChannel             `json:"salesChannels,omitempty"`
	TranslationCode                *Locale                   `json:"translationCode,omitempty"`
	TranslationCodeId              string                    `json:"translationCodeId,omitempty"`
	UpdatedAt                      time.Time                 `json:"updatedAt,omitempty"`
}
