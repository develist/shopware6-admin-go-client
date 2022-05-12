package entity

import "time"

// completed

type SalesChannel struct {
	AccessKey      string                 `json:"accessKey,omitempty"`
	Active         bool                   `json:"active,omitempty"`
	Analytics      *SalesChannelAnalytics `json:"analytics,omitempty"`
	AnalyticsId    string                 `json:"analyticsId,omitempty"`
	BoundCustomers *Customer              `json:"boundCustomers,omitempty"`
	// Configuration map[type:object] `json:"configuration,omitempty"`
	Countries                       *Country                        `json:"countries,omitempty"`
	Country                         *Country                        `json:"country,omitempty"`
	CountryId                       string                          `json:"countryId,omitempty"`
	CreatedAt                       time.Time                       `json:"createdAt,omitempty"`
	Currencies                      *Currency                       `json:"currencies,omitempty"`
	Currency                        *Currency                       `json:"currency,omitempty"`
	CurrencyId                      string                          `json:"currencyId,omitempty"`
	CustomFields                    *[]CustomField                  `json:"customFields,omitempty"`
	CustomerGroup                   *CustomerGroup                  `json:"customerGroup,omitempty"`
	CustomerGroupId                 string                          `json:"customerGroupId,omitempty"`
	CustomerGroupsRegistrations     *CustomerGroup                  `json:"customerGroupsRegistrations,omitempty"`
	Customers                       *Customer                       `json:"customers,omitempty"`
	DocumentBaseConfigSalesChannels *DocumentBaseConfigSalesChannel `json:"documentBaseConfigSalesChannels,omitempty"`
	Domains                         *SalesChannelDomain             `json:"domains,omitempty"`
	EventActions                    *EventAction                    `json:"eventActions,omitempty"`
	FooterCategory                  *Category                       `json:"footerCategory,omitempty"`
	FooterCategoryId                string                          `json:"footerCategoryId,omitempty"`
	FooterCategoryVersionId         string                          `json:"footerCategoryVersionId,omitempty"`
	HomeCmsPage                     *CmsPage                        `json:"homeCmsPage,omitempty"`
	HomeCmsPageId                   string                          `json:"homeCmsPageId,omitempty"`
	HomeCmsPageVersionId            string                          `json:"homeCmsPageVersionId,omitempty"`
	HomeEnabled                     bool                            `json:"homeEnabled,omitempty"`
	HomeKeywords                    string                          `json:"homeKeywords,omitempty"`
	HomeMetaDescription             string                          `json:"homeMetaDescription,omitempty"`
	HomeMetaTitle                   string                          `json:"homeMetaTitle,omitempty"`
	HomeName                        string                          `json:"homeName,omitempty"`
	// HomeSlotConfig map[type:object] `json:"homeSlotConfig,omitempty"`
	HreflangActive          bool                `json:"hreflangActive,omitempty"`
	HreflangDefaultDomain   *SalesChannelDomain `json:"hreflangDefaultDomain,omitempty"`
	HreflangDefaultDomainId string              `json:"hreflangDefaultDomainId,omitempty"`
	Id                      string              `json:"id,omitempty"`
	LandingPages            *LandingPage        `json:"landingPages,omitempty"`
	Language                *Language           `json:"language,omitempty"`
	LanguageId              string              `json:"languageId,omitempty"`
	Languages               *Language           `json:"languages,omitempty"`
	MailHeaderFooter        *MailHeaderFooter   `json:"mailHeaderFooter,omitempty"`
	MailHeaderFooterId      string              `json:"mailHeaderFooterId,omitempty"`
	MainCategories          *MainCategory       `json:"mainCategories,omitempty"`
	Maintenance             bool                `json:"maintenance,omitempty"`
	// MaintenanceIpWhitelist map[items:map[additionalProperties:false] type:array] `json:"maintenanceIpWhitelist,omitempty"`
	Name                        string                   `json:"name,omitempty"`
	NavigationCategory          *Category                `json:"navigationCategory,omitempty"`
	NavigationCategoryDepth     int                      `json:"navigationCategoryDepth,omitempty"`
	NavigationCategoryId        string                   `json:"navigationCategoryId,omitempty"`
	NavigationCategoryVersionId string                   `json:"navigationCategoryVersionId,omitempty"`
	NewsletterRecipients        *NewsletterRecipient     `json:"newsletterRecipients,omitempty"`
	NumberRangeSalesChannels    *NumberRangeSalesChannel `json:"numberRangeSalesChannels,omitempty"`
	Orders                      *Order                   `json:"orders,omitempty"`
	PaymentMethod               *PaymentMethod           `json:"paymentMethod,omitempty"`
	PaymentMethodId             string                   `json:"paymentMethodId,omitempty"`
	// PaymentMethodIds map[items:map[pattern:^[0-9a-f]{32}$ type:string] readOnly:true type:array] `json:"paymentMethodIds,omitempty"`
	PaymentMethods           *PaymentMethod         `json:"paymentMethods,omitempty"`
	ProductExports           *ProductExport         `json:"productExports,omitempty"`
	ProductReviews           *ProductReview         `json:"productReviews,omitempty"`
	ProductVisibilities      *ProductVisibility     `json:"productVisibilities,omitempty"`
	PromotionSalesChannels   *PromotionSalesChannel `json:"promotionSalesChannels,omitempty"`
	SeoUrlTemplates          *SeoUrlTemplate        `json:"seoUrlTemplates,omitempty"`
	SeoUrls                  *SeoUrl                `json:"seoUrls,omitempty"`
	ServiceCategory          *Category              `json:"serviceCategory,omitempty"`
	ServiceCategoryId        string                 `json:"serviceCategoryId,omitempty"`
	ServiceCategoryVersionId string                 `json:"serviceCategoryVersionId,omitempty"`
	ShippingMethod           *ShippingMethod        `json:"shippingMethod,omitempty"`
	ShippingMethodId         string                 `json:"shippingMethodId,omitempty"`
	ShippingMethods          *ShippingMethod        `json:"shippingMethods,omitempty"`
	ShortName                string                 `json:"shortName,omitempty"`
	SystemConfigs            *SystemConfig          `json:"systemConfigs,omitempty"`
	TaxCalculationType       string                 `json:"taxCalculationType,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	Type      *SalesChannelType `json:"type,omitempty"`
	TypeId    string            `json:"typeId,omitempty"`
	UpdatedAt time.Time         `json:"updatedAt,omitempty"`
	Wishlists *CustomerWishlist `json:"wishlists,omitempty"`
}

type SalesChannelAnalytics struct {
	Active       bool          `json:"active,omitempty"`
	AnonymizeIp  bool          `json:"anonymizeIp,omitempty"`
	CreatedAt    time.Time     `json:"createdAt,omitempty"`
	Id           string        `json:"id,omitempty"`
	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`
	TrackOrders  bool          `json:"trackOrders,omitempty"`
	TrackingId   string        `json:"trackingId,omitempty"`
	UpdatedAt    time.Time     `json:"updatedAt,omitempty"`
}

type SalesChannelCountry struct {
	Country        *Country      `json:"country,omitempty"`
	CountryId      string        `json:"countryId,omitempty"`
	Id             string        `json:"id,omitempty"`
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId string        `json:"salesChannelId,omitempty"`
}

type SalesChannelCurrency struct {
	Currency       *Currency     `json:"currency,omitempty"`
	CurrencyId     string        `json:"currencyId,omitempty"`
	Id             string        `json:"id,omitempty"`
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId string        `json:"salesChannelId,omitempty"`
}

type SalesChannelDomain struct {
	CreatedAt                   time.Time      `json:"createdAt,omitempty"`
	Currency                    *Currency      `json:"currency,omitempty"`
	CurrencyId                  string         `json:"currencyId,omitempty"`
	CustomFields                *[]CustomField `json:"customFields,omitempty"`
	HreflangUseOnlyLocale       bool           `json:"hreflangUseOnlyLocale,omitempty"`
	Id                          string         `json:"id,omitempty"`
	Language                    *Language      `json:"language,omitempty"`
	LanguageId                  string         `json:"languageId,omitempty"`
	ProductExports              *ProductExport `json:"productExports,omitempty"`
	SalesChannel                *SalesChannel  `json:"salesChannel,omitempty"`
	SalesChannelDefaultHreflang *SalesChannel  `json:"salesChannelDefaultHreflang,omitempty"`
	SalesChannelId              string         `json:"salesChannelId,omitempty"`
	SnippetSet                  *SnippetSet    `json:"snippetSet,omitempty"`
	SnippetSetId                string         `json:"snippetSetId,omitempty"`
	UpdatedAt                   time.Time      `json:"updatedAt,omitempty"`
	Url                         string         `json:"url,omitempty"`
}

type SalesChannelLanguage struct {
	Id             string        `json:"id,omitempty"`
	Language       *Language     `json:"language,omitempty"`
	LanguageId     string        `json:"languageId,omitempty"`
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId string        `json:"salesChannelId,omitempty"`
}

type SalesChannelPaymentMethod struct {
	Id              string         `json:"id,omitempty"`
	PaymentMethod   *PaymentMethod `json:"paymentMethod,omitempty"`
	PaymentMethodId string         `json:"paymentMethodId,omitempty"`
	SalesChannel    *SalesChannel  `json:"salesChannel,omitempty"`
	SalesChannelId  string         `json:"salesChannelId,omitempty"`
}

type SalesChannelShippingMethod struct {
	Id               string          `json:"id,omitempty"`
	SalesChannel     *SalesChannel   `json:"salesChannel,omitempty"`
	SalesChannelId   string          `json:"salesChannelId,omitempty"`
	ShippingMethod   *ShippingMethod `json:"shippingMethod,omitempty"`
	ShippingMethodId string          `json:"shippingMethodId,omitempty"`
}

type SalesChannelType struct {
	CoverUrl        string         `json:"coverUrl,omitempty"`
	CreatedAt       time.Time      `json:"createdAt,omitempty"`
	CustomFields    *[]CustomField `json:"customFields,omitempty"`
	Description     string         `json:"description,omitempty"`
	DescriptionLong string         `json:"descriptionLong,omitempty"`
	IconName        string         `json:"iconName,omitempty"`
	Id              string         `json:"id,omitempty"`
	Manufacturer    string         `json:"manufacturer,omitempty"`
	Name            string         `json:"name,omitempty"`
	SalesChannels   *SalesChannel  `json:"salesChannels,omitempty"`
	// ScreenshotUrls map[items:map[type:string] type:array] `json:"screenshotUrls,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
