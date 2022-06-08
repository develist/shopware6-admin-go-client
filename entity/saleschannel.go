package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-06 18:44:04 UTC

// SalesChannel data structure
// Added since version: 6.0.0.0
// Required fields: typeId, languageId, customerGroupId, currencyId, paymentMethodId, shippingMethodId, countryId, navigationCategoryId, accessKey, createdAt, name, homeEnabled
type SalesChannel struct {
	AccessKey                       *string                         `json:"accessKey,omitempty"`
	Active                          *bool                           `json:"active,omitempty"`
	Analytics                       *SalesChannelAnalytics          `json:"analytics,omitempty"`
	AnalyticsId                     *string                         `json:"analyticsId,omitempty"`
	BoundCustomers                  *Customer                       `json:"boundCustomers,omitempty"`
	Configuration                   *interface{}                    `json:"configuration,omitempty"` // map[type:object]
	Countries                       *Country                        `json:"countries,omitempty"`
	Country                         *Country                        `json:"country,omitempty"`
	CountryId                       *string                         `json:"countryId,omitempty"`
	CreatedAt                       *time.Time                      `json:"createdAt,omitempty"`
	Currencies                      *Currency                       `json:"currencies,omitempty"`
	Currency                        *Currency                       `json:"currency,omitempty"`
	CurrencyId                      *string                         `json:"currencyId,omitempty"`
	CustomFields                    *[]CustomField                  `json:"customFields,omitempty"`
	CustomerGroup                   *CustomerGroup                  `json:"customerGroup,omitempty"`
	CustomerGroupId                 *string                         `json:"customerGroupId,omitempty"`
	CustomerGroupsRegistrations     *CustomerGroup                  `json:"customerGroupsRegistrations,omitempty"`
	Customers                       *Customer                       `json:"customers,omitempty"`
	DocumentBaseConfigSalesChannels *DocumentBaseConfigSalesChannel `json:"documentBaseConfigSalesChannels,omitempty"`
	Domains                         *SalesChannelDomain             `json:"domains,omitempty"`
	EventActions                    *EventAction                    `json:"eventActions,omitempty"`
	FooterCategory                  *Category                       `json:"footerCategory,omitempty"`
	FooterCategoryId                *string                         `json:"footerCategoryId,omitempty"`
	FooterCategoryVersionId         *string                         `json:"footerCategoryVersionId,omitempty"`
	HomeCmsPage                     *CmsPage                        `json:"homeCmsPage,omitempty"`
	HomeCmsPageId                   *string                         `json:"homeCmsPageId,omitempty"`
	HomeCmsPageVersionId            *string                         `json:"homeCmsPageVersionId,omitempty"`
	HomeEnabled                     *bool                           `json:"homeEnabled,omitempty"`
	HomeKeywords                    *string                         `json:"homeKeywords,omitempty"`
	HomeMetaDescription             *string                         `json:"homeMetaDescription,omitempty"`
	HomeMetaTitle                   *string                         `json:"homeMetaTitle,omitempty"`
	HomeName                        *string                         `json:"homeName,omitempty"`
	HomeSlotConfig                  *interface{}                    `json:"homeSlotConfig,omitempty"` // map[type:object]
	HreflangActive                  *bool                           `json:"hreflangActive,omitempty"`
	HreflangDefaultDomain           *SalesChannelDomain             `json:"hreflangDefaultDomain,omitempty"`
	HreflangDefaultDomainId         *string                         `json:"hreflangDefaultDomainId,omitempty"`
	Id                              *string                         `json:"id,omitempty"`
	LandingPages                    *LandingPage                    `json:"landingPages,omitempty"`
	Language                        *Language                       `json:"language,omitempty"`
	LanguageId                      *string                         `json:"languageId,omitempty"`
	Languages                       *Language                       `json:"languages,omitempty"`
	MailHeaderFooter                *MailHeaderFooter               `json:"mailHeaderFooter,omitempty"`
	MailHeaderFooterId              *string                         `json:"mailHeaderFooterId,omitempty"`
	MainCategories                  *MainCategory                   `json:"mainCategories,omitempty"`
	Maintenance                     *bool                           `json:"maintenance,omitempty"`
	MaintenanceIpWhitelist          *interface{}                    `json:"maintenanceIpWhitelist,omitempty"` // map[items:map[additionalProperties:false] type:array]
	Name                            *string                         `json:"name,omitempty"`
	NavigationCategory              *Category                       `json:"navigationCategory,omitempty"`
	NavigationCategoryDepth         *int                            `json:"navigationCategoryDepth,omitempty"`
	NavigationCategoryId            *string                         `json:"navigationCategoryId,omitempty"`
	NavigationCategoryVersionId     *string                         `json:"navigationCategoryVersionId,omitempty"`
	NewsletterRecipients            *NewsletterRecipient            `json:"newsletterRecipients,omitempty"`
	NumberRangeSalesChannels        *NumberRangeSalesChannel        `json:"numberRangeSalesChannels,omitempty"`
	Orders                          *Order                          `json:"orders,omitempty"`
	PaymentMethod                   *PaymentMethod                  `json:"paymentMethod,omitempty"`
	PaymentMethodId                 *string                         `json:"paymentMethodId,omitempty"`
	PaymentMethodIds                *interface{}                    `json:"paymentMethodIds,omitempty"` // map[items:map[pattern:^[0-9a-f]{32}$ type:string] readOnly:true type:array]
	PaymentMethods                  *PaymentMethod                  `json:"paymentMethods,omitempty"`
	ProductExports                  *ProductExport                  `json:"productExports,omitempty"`
	ProductReviews                  *ProductReview                  `json:"productReviews,omitempty"`
	ProductVisibilities             *ProductVisibility              `json:"productVisibilities,omitempty"`
	PromotionSalesChannels          *PromotionSalesChannel          `json:"promotionSalesChannels,omitempty"`
	SeoUrlTemplates                 *SeoUrlTemplate                 `json:"seoUrlTemplates,omitempty"`
	SeoUrls                         *SeoUrl                         `json:"seoUrls,omitempty"`
	ServiceCategory                 *Category                       `json:"serviceCategory,omitempty"`
	ServiceCategoryId               *string                         `json:"serviceCategoryId,omitempty"`
	ServiceCategoryVersionId        *string                         `json:"serviceCategoryVersionId,omitempty"`
	ShippingMethod                  *ShippingMethod                 `json:"shippingMethod,omitempty"`
	ShippingMethodId                *string                         `json:"shippingMethodId,omitempty"`
	ShippingMethods                 *ShippingMethod                 `json:"shippingMethods,omitempty"`
	ShortName                       *string                         `json:"shortName,omitempty"`
	SystemConfigs                   *SystemConfig                   `json:"systemConfigs,omitempty"`
	TaxCalculationType              *string                         `json:"taxCalculationType,omitempty"`
	Translated                      *interface{}                    `json:"translated,omitempty"` // map[type:object]
	Type                            *SalesChannelType               `json:"type,omitempty"`
	TypeId                          *string                         `json:"typeId,omitempty"`
	UpdatedAt                       *time.Time                      `json:"updatedAt,omitempty"`
	Wishlists                       *CustomerWishlist               `json:"wishlists,omitempty"`
}

// SalesChannelAnalytics data structure
// Added since version: 6.2.0.0
// Required fields: createdAt
type SalesChannelAnalytics struct {
	Active       *bool         `json:"active,omitempty"`
	AnonymizeIp  *bool         `json:"anonymizeIp,omitempty"`
	CreatedAt    *time.Time    `json:"createdAt,omitempty"`
	Id           *string       `json:"id,omitempty"`
	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`
	TrackOrders  *bool         `json:"trackOrders,omitempty"`
	TrackingId   *string       `json:"trackingId,omitempty"`
	UpdatedAt    *time.Time    `json:"updatedAt,omitempty"`
}

// SalesChannelCountry data structure
// Added since version: 6.0.0.0
// Required fields: salesChannelId, countryId
type SalesChannelCountry struct {
	Country        *Country      `json:"country,omitempty"`
	CountryId      *string       `json:"countryId,omitempty"`
	Id             *string       `json:"id,omitempty"`
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId *string       `json:"salesChannelId,omitempty"`
}

// SalesChannelCurrency data structure
// Added since version: 6.0.0.0
// Required fields: salesChannelId, currencyId
type SalesChannelCurrency struct {
	Currency       *Currency     `json:"currency,omitempty"`
	CurrencyId     *string       `json:"currencyId,omitempty"`
	Id             *string       `json:"id,omitempty"`
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId *string       `json:"salesChannelId,omitempty"`
}

// SalesChannelDomain data structure
// Added since version: 6.0.0.0
// Required fields: url, salesChannelId, languageId, currencyId, snippetSetId, createdAt
type SalesChannelDomain struct {
	CreatedAt                   *time.Time     `json:"createdAt,omitempty"`
	Currency                    *Currency      `json:"currency,omitempty"`
	CurrencyId                  *string        `json:"currencyId,omitempty"`
	CustomFields                *[]CustomField `json:"customFields,omitempty"`
	HreflangUseOnlyLocale       *bool          `json:"hreflangUseOnlyLocale,omitempty"`
	Id                          *string        `json:"id,omitempty"`
	Language                    *Language      `json:"language,omitempty"`
	LanguageId                  *string        `json:"languageId,omitempty"`
	ProductExports              *ProductExport `json:"productExports,omitempty"`
	SalesChannel                *SalesChannel  `json:"salesChannel,omitempty"`
	SalesChannelDefaultHreflang *SalesChannel  `json:"salesChannelDefaultHreflang,omitempty"`
	SalesChannelId              *string        `json:"salesChannelId,omitempty"`
	SnippetSet                  *SnippetSet    `json:"snippetSet,omitempty"`
	SnippetSetId                *string        `json:"snippetSetId,omitempty"`
	UpdatedAt                   *time.Time     `json:"updatedAt,omitempty"`
	Url                         *string        `json:"url,omitempty"`
}

// SalesChannelLanguage data structure
// Added since version: 6.0.0.0
// Required fields: salesChannelId, languageId
type SalesChannelLanguage struct {
	Id             *string       `json:"id,omitempty"`
	Language       *Language     `json:"language,omitempty"`
	LanguageId     *string       `json:"languageId,omitempty"`
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId *string       `json:"salesChannelId,omitempty"`
}

// SalesChannelPaymentMethod data structure
// Added since version: 6.0.0.0
// Required fields: salesChannelId, paymentMethodId
type SalesChannelPaymentMethod struct {
	Id              *string        `json:"id,omitempty"`
	PaymentMethod   *PaymentMethod `json:"paymentMethod,omitempty"`
	PaymentMethodId *string        `json:"paymentMethodId,omitempty"`
	SalesChannel    *SalesChannel  `json:"salesChannel,omitempty"`
	SalesChannelId  *string        `json:"salesChannelId,omitempty"`
}

// SalesChannelShippingMethod data structure
// Added since version: 6.0.0.0
// Required fields: salesChannelId, shippingMethodId
type SalesChannelShippingMethod struct {
	Id               *string         `json:"id,omitempty"`
	SalesChannel     *SalesChannel   `json:"salesChannel,omitempty"`
	SalesChannelId   *string         `json:"salesChannelId,omitempty"`
	ShippingMethod   *ShippingMethod `json:"shippingMethod,omitempty"`
	ShippingMethodId *string         `json:"shippingMethodId,omitempty"`
}

// SalesChannelType data structure
// Added since version: 6.0.0.0
// Required fields: createdAt, name
type SalesChannelType struct {
	CoverUrl        *string        `json:"coverUrl,omitempty"`
	CreatedAt       *time.Time     `json:"createdAt,omitempty"`
	CustomFields    *[]CustomField `json:"customFields,omitempty"`
	Description     *string        `json:"description,omitempty"`
	DescriptionLong *string        `json:"descriptionLong,omitempty"`
	IconName        *string        `json:"iconName,omitempty"`
	Id              *string        `json:"id,omitempty"`
	Manufacturer    *string        `json:"manufacturer,omitempty"`
	Name            *string        `json:"name,omitempty"`
	SalesChannels   *SalesChannel  `json:"salesChannels,omitempty"`
	ScreenshotUrls  *[]string      `json:"screenshotUrls,omitempty"` // map[items:map[type:string] type:array]
	Translated      *interface{}   `json:"translated,omitempty"`     // map[type:object]
	UpdatedAt       *time.Time     `json:"updatedAt,omitempty"`
}
