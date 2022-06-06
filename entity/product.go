package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-06 18:44:04 UTC

// Product data structure
// Added since version: 6.0.0.0
// Required fields: taxId, price, productNumber, stock, createdAt, name
type Product struct {
	Active                        bool                                 `json:"active,omitempty"`
	AutoIncrement                 int                                  `json:"autoIncrement,omitempty"`
	Available                     bool                                 `json:"available,omitempty"`
	AvailableStock                int                                  `json:"availableStock,omitempty"`
	CanonicalProduct              *Product                             `json:"canonicalProduct,omitempty"`
	CanonicalProductId            string                               `json:"canonicalProductId,omitempty"`
	Categories                    *Category                            `json:"categories,omitempty"`
	CategoriesRo                  *Category                            `json:"categoriesRo,omitempty"`
	CategoryIds                   *interface{}                         `json:"categoryIds,omitempty"`   // map[items:map[pattern:^[0-9a-f]{32}$ type:string] readOnly:true type:array]
	CategoryTree                  *interface{}                         `json:"categoryTree,omitempty"`  // map[items:map[pattern:^[0-9a-f]{32}$ type:string] readOnly:true type:array]
	CheapestPrice                 *interface{}                         `json:"cheapestPrice,omitempty"` // map[readOnly:true type:object]
	ChildCount                    int                                  `json:"childCount,omitempty"`
	Children                      *Product                             `json:"children,omitempty"`
	CmsPage                       *CmsPage                             `json:"cmsPage,omitempty"`
	CmsPageId                     string                               `json:"cmsPageId,omitempty"`
	CmsPageVersionId              string                               `json:"cmsPageVersionId,omitempty"`
	ConfiguratorGroupConfig       *interface{}                         `json:"configuratorGroupConfig,omitempty"` // map[type:object]
	ConfiguratorSettings          *ProductConfiguratorSetting          `json:"configuratorSettings,omitempty"`
	Cover                         *ProductMedia                        `json:"cover,omitempty"`
	CoverId                       string                               `json:"coverId,omitempty"`
	CreatedAt                     *time.Time                           `json:"createdAt,omitempty"`
	CrossSellingAssignedProducts  *ProductCrossSellingAssignedProducts `json:"crossSellingAssignedProducts,omitempty"`
	CrossSellings                 *ProductCrossSelling                 `json:"crossSellings,omitempty"`
	CustomFieldSetSelectionActive bool                                 `json:"customFieldSetSelectionActive,omitempty"`
	CustomFieldSets               *CustomFieldSet                      `json:"customFieldSets,omitempty"`
	CustomFields                  *[]CustomField                       `json:"customFields,omitempty"`
	CustomSearchKeywords          *interface{}                         `json:"customSearchKeywords,omitempty"` // map[items:map[additionalProperties:false] type:array]
	DeliveryTime                  *DeliveryTime                        `json:"deliveryTime,omitempty"`
	DeliveryTimeId                string                               `json:"deliveryTimeId,omitempty"`
	Description                   string                               `json:"description,omitempty"`
	DisplayGroup                  string                               `json:"displayGroup,omitempty"`
	Ean                           string                               `json:"ean,omitempty"`
	FeatureSet                    *ProductFeatureSet                   `json:"featureSet,omitempty"`
	FeatureSetId                  string                               `json:"featureSetId,omitempty"`
	Height                        float64                              `json:"height,omitempty"`
	Id                            string                               `json:"id,omitempty"`
	IsCloseout                    bool                                 `json:"isCloseout,omitempty"`
	Keywords                      string                               `json:"keywords,omitempty"`
	Length                        float64                              `json:"length,omitempty"`
	MainCategories                *MainCategory                        `json:"mainCategories,omitempty"`
	MainVariantId                 string                               `json:"mainVariantId,omitempty"`
	Manufacturer                  *ProductManufacturer                 `json:"manufacturer,omitempty"`
	ManufacturerId                string                               `json:"manufacturerId,omitempty"`
	ManufacturerNumber            string                               `json:"manufacturerNumber,omitempty"`
	MarkAsTopseller               bool                                 `json:"markAsTopseller,omitempty"`
	MaxPurchase                   int                                  `json:"maxPurchase,omitempty"`
	Media                         *ProductMedia                        `json:"media,omitempty"`
	MetaDescription               string                               `json:"metaDescription,omitempty"`
	MetaTitle                     string                               `json:"metaTitle,omitempty"`
	MinPurchase                   int                                  `json:"minPurchase,omitempty"`
	Name                          string                               `json:"name,omitempty"`
	OptionIds                     *interface{}                         `json:"optionIds,omitempty"` // map[items:map[pattern:^[0-9a-f]{32}$ type:string] readOnly:true type:array]
	Options                       *PropertyGroupOption                 `json:"options,omitempty"`
	OrderLineItems                *OrderLineItem                       `json:"orderLineItems,omitempty"`
	PackUnit                      string                               `json:"packUnit,omitempty"`
	PackUnitPlural                string                               `json:"packUnitPlural,omitempty"`
	Parent                        *Product                             `json:"parent,omitempty"`
	ParentId                      string                               `json:"parentId,omitempty"`
	ParentVersionId               string                               `json:"parentVersionId,omitempty"`
	Price                         *interface{}                         `json:"price,omitempty"` // map[type:object]
	Prices                        *ProductPrice                        `json:"prices,omitempty"`
	ProductManufacturerVersionId  string                               `json:"productManufacturerVersionId,omitempty"`
	ProductMediaVersionId         string                               `json:"productMediaVersionId,omitempty"`
	ProductNumber                 string                               `json:"productNumber,omitempty"`
	ProductReviews                *ProductReview                       `json:"productReviews,omitempty"`
	Properties                    *PropertyGroupOption                 `json:"properties,omitempty"`
	PropertyIds                   *interface{}                         `json:"propertyIds,omitempty"`    // map[items:map[pattern:^[0-9a-f]{32}$ type:string] readOnly:true type:array]
	PurchasePrices                *interface{}                         `json:"purchasePrices,omitempty"` // map[type:object]
	PurchaseSteps                 int                                  `json:"purchaseSteps,omitempty"`
	PurchaseUnit                  float64                              `json:"purchaseUnit,omitempty"`
	RatingAverage                 float64                              `json:"ratingAverage,omitempty"`
	ReferenceUnit                 float64                              `json:"referenceUnit,omitempty"`
	ReleaseDate                   *time.Time                           `json:"releaseDate,omitempty"`
	RestockTime                   int                                  `json:"restockTime,omitempty"`
	Sales                         int                                  `json:"sales,omitempty"`
	SearchKeywords                *ProductSearchKeyword                `json:"searchKeywords,omitempty"`
	SeoUrls                       *SeoUrl                              `json:"seoUrls,omitempty"`
	ShippingFree                  bool                                 `json:"shippingFree,omitempty"`
	SlotConfig                    *interface{}                         `json:"slotConfig,omitempty"` // map[type:object]
	Stock                         int                                  `json:"stock,omitempty"`
	Streams                       *ProductStream                       `json:"streams,omitempty"`
	TagIds                        *interface{}                         `json:"tagIds,omitempty"` // map[items:map[pattern:^[0-9a-f]{32}$ type:string] readOnly:true type:array]
	Tags                          *Tag                                 `json:"tags,omitempty"`
	Tax                           *Tax                                 `json:"tax,omitempty"`
	TaxId                         string                               `json:"taxId,omitempty"`
	Translated                    *interface{}                         `json:"translated,omitempty"` // map[type:object]
	Unit                          *Unit                                `json:"unit,omitempty"`
	UnitId                        string                               `json:"unitId,omitempty"`
	UpdatedAt                     *time.Time                           `json:"updatedAt,omitempty"`
	VariantRestrictions           *interface{}                         `json:"variantRestrictions,omitempty"` // map[type:object]
	Variation                     *interface{}                         `json:"variation,omitempty"`           // map[items:map[type:string] type:array]
	VersionId                     string                               `json:"versionId,omitempty"`
	Visibilities                  *ProductVisibility                   `json:"visibilities,omitempty"`
	Weight                        float64                              `json:"weight,omitempty"`
	Width                         float64                              `json:"width,omitempty"`
	Wishlists                     *CustomerWishlistProduct             `json:"wishlists,omitempty"`
}

// ProductCategory data structure
// Added since version: 6.0.0.0
// Required fields: productId, categoryId
type ProductCategory struct {
	Category          *Category `json:"category,omitempty"`
	CategoryId        string    `json:"categoryId,omitempty"`
	CategoryVersionId string    `json:"categoryVersionId,omitempty"`
	Id                string    `json:"id,omitempty"`
	Product           *Product  `json:"product,omitempty"`
	ProductId         string    `json:"productId,omitempty"`
	ProductVersionId  string    `json:"productVersionId,omitempty"`
}

// ProductCategoryTree data structure
// Added since version: 6.0.0.0
// Required fields: productId, categoryId
type ProductCategoryTree struct {
	Category          *Category `json:"category,omitempty"`
	CategoryId        string    `json:"categoryId,omitempty"`
	CategoryVersionId string    `json:"categoryVersionId,omitempty"`
	Id                string    `json:"id,omitempty"`
	Product           *Product  `json:"product,omitempty"`
	ProductId         string    `json:"productId,omitempty"`
	ProductVersionId  string    `json:"productVersionId,omitempty"`
}

// ProductConfiguratorSetting data structure
// Added since version: 6.0.0.0
// Required fields: productId, optionId, createdAt
type ProductConfiguratorSetting struct {
	CreatedAt        *time.Time           `json:"createdAt,omitempty"`
	CustomFields     *[]CustomField       `json:"customFields,omitempty"`
	Id               string               `json:"id,omitempty"`
	Media            *Media               `json:"media,omitempty"`
	MediaId          string               `json:"mediaId,omitempty"`
	Option           *PropertyGroupOption `json:"option,omitempty"`
	OptionId         string               `json:"optionId,omitempty"`
	Position         int                  `json:"position,omitempty"`
	Price            *interface{}         `json:"price,omitempty"` // map[type:object]
	Product          *Product             `json:"product,omitempty"`
	ProductId        string               `json:"productId,omitempty"`
	ProductVersionId string               `json:"productVersionId,omitempty"`
	UpdatedAt        *time.Time           `json:"updatedAt,omitempty"`
	VersionId        string               `json:"versionId,omitempty"`
}

// ProductCrossSelling data structure
// Added since version: 6.1.0.0
// Required fields: name, position, type, productId, createdAt
type ProductCrossSelling struct {
	Active           bool                                 `json:"active,omitempty"`
	AssignedProducts *ProductCrossSellingAssignedProducts `json:"assignedProducts,omitempty"`
	CreatedAt        *time.Time                           `json:"createdAt,omitempty"`
	Id               string                               `json:"id,omitempty"`
	Limit            int                                  `json:"limit,omitempty"`
	Name             string                               `json:"name,omitempty"`
	Position         int                                  `json:"position,omitempty"`
	Product          *Product                             `json:"product,omitempty"`
	ProductId        string                               `json:"productId,omitempty"`
	ProductStream    *ProductStream                       `json:"productStream,omitempty"`
	ProductStreamId  string                               `json:"productStreamId,omitempty"`
	ProductVersionId string                               `json:"productVersionId,omitempty"`
	SortBy           string                               `json:"sortBy,omitempty"`
	SortDirection    string                               `json:"sortDirection,omitempty"`
	Translated       *interface{}                         `json:"translated,omitempty"` // map[type:object]
	Type             string                               `json:"type,omitempty"`
	UpdatedAt        *time.Time                           `json:"updatedAt,omitempty"`
}

// ProductCrossSellingAssignedProducts data structure
// Added since version: 6.2.0.0
// Required fields: crossSellingId, productId, createdAt
type ProductCrossSellingAssignedProducts struct {
	CreatedAt        *time.Time           `json:"createdAt,omitempty"`
	CrossSelling     *ProductCrossSelling `json:"crossSelling,omitempty"`
	CrossSellingId   string               `json:"crossSellingId,omitempty"`
	Id               string               `json:"id,omitempty"`
	Position         int                  `json:"position,omitempty"`
	Product          *Product             `json:"product,omitempty"`
	ProductId        string               `json:"productId,omitempty"`
	ProductVersionId string               `json:"productVersionId,omitempty"`
	UpdatedAt        *time.Time           `json:"updatedAt,omitempty"`
}

// ProductCustomFieldSet data structure
// Added since version: 6.3.0.0
// Required fields: productId, customFieldSetId
type ProductCustomFieldSet struct {
	CustomFieldSet   *CustomFieldSet `json:"customFieldSet,omitempty"`
	CustomFieldSetId string          `json:"customFieldSetId,omitempty"`
	Id               string          `json:"id,omitempty"`
	Product          *Product        `json:"product,omitempty"`
	ProductId        string          `json:"productId,omitempty"`
	ProductVersionId string          `json:"productVersionId,omitempty"`
}

// ProductExport data structure
// Added since version: 6.1.0.0
// Required fields: productStreamId, storefrontSalesChannelId, salesChannelId, salesChannelDomainId, currencyId, fileName, accessKey, encoding, fileFormat, generateByCronjob, interval, createdAt
type ProductExport struct {
	AccessKey                string              `json:"accessKey,omitempty"`
	BodyTemplate             string              `json:"bodyTemplate,omitempty"`
	CreatedAt                *time.Time          `json:"createdAt,omitempty"`
	Currency                 *Currency           `json:"currency,omitempty"`
	CurrencyId               string              `json:"currencyId,omitempty"`
	Encoding                 string              `json:"encoding,omitempty"`
	FileFormat               string              `json:"fileFormat,omitempty"`
	FileName                 string              `json:"fileName,omitempty"`
	FooterTemplate           string              `json:"footerTemplate,omitempty"`
	GenerateByCronjob        bool                `json:"generateByCronjob,omitempty"`
	GeneratedAt              *time.Time          `json:"generatedAt,omitempty"`
	HeaderTemplate           string              `json:"headerTemplate,omitempty"`
	Id                       string              `json:"id,omitempty"`
	IncludeVariants          bool                `json:"includeVariants,omitempty"`
	Interval                 int                 `json:"interval,omitempty"`
	PausedSchedule           bool                `json:"pausedSchedule,omitempty"`
	ProductStream            *ProductStream      `json:"productStream,omitempty"`
	ProductStreamId          string              `json:"productStreamId,omitempty"`
	SalesChannel             *SalesChannel       `json:"salesChannel,omitempty"`
	SalesChannelDomain       *SalesChannelDomain `json:"salesChannelDomain,omitempty"`
	SalesChannelDomainId     string              `json:"salesChannelDomainId,omitempty"`
	SalesChannelId           string              `json:"salesChannelId,omitempty"`
	StorefrontSalesChannel   *SalesChannel       `json:"storefrontSalesChannel,omitempty"`
	StorefrontSalesChannelId string              `json:"storefrontSalesChannelId,omitempty"`
	UpdatedAt                *time.Time          `json:"updatedAt,omitempty"`
}

// ProductFeatureSet data structure
// Added since version: 6.3.0.0
// Required fields: createdAt, name
type ProductFeatureSet struct {
	CreatedAt   *time.Time   `json:"createdAt,omitempty"`
	Description string       `json:"description,omitempty"`
	Features    *interface{} `json:"features,omitempty"` // map[type:object]
	Id          string       `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	Products    *Product     `json:"products,omitempty"`
	Translated  *interface{} `json:"translated,omitempty"` // map[type:object]
	UpdatedAt   *time.Time   `json:"updatedAt,omitempty"`
}

// ProductKeywordDictionary data structure
// Added since version: 6.0.0.0
// Required fields: languageId, keyword
type ProductKeywordDictionary struct {
	Id         string    `json:"id,omitempty"`
	Keyword    string    `json:"keyword,omitempty"`
	Language   *Language `json:"language,omitempty"`
	LanguageId string    `json:"languageId,omitempty"`
	Reversed   string    `json:"reversed,omitempty"`
}

// ProductManufacturer data structure
// Added since version: 6.0.0.0
// Required fields: createdAt, name
type ProductManufacturer struct {
	CreatedAt    *time.Time     `json:"createdAt,omitempty"`
	CustomFields *[]CustomField `json:"customFields,omitempty"`
	Description  string         `json:"description,omitempty"`
	Id           string         `json:"id,omitempty"`
	Link         string         `json:"link,omitempty"`
	Media        *Media         `json:"media,omitempty"`
	MediaId      string         `json:"mediaId,omitempty"`
	Name         string         `json:"name,omitempty"`
	Products     *Product       `json:"products,omitempty"`
	Translated   *interface{}   `json:"translated,omitempty"` // map[type:object]
	UpdatedAt    *time.Time     `json:"updatedAt,omitempty"`
	VersionId    string         `json:"versionId,omitempty"`
}

// ProductMedia data structure
// Added since version: 6.0.0.0
// Required fields: productId, mediaId, createdAt
type ProductMedia struct {
	CreatedAt        *time.Time     `json:"createdAt,omitempty"`
	CustomFields     *[]CustomField `json:"customFields,omitempty"`
	Id               string         `json:"id,omitempty"`
	Media            *Media         `json:"media,omitempty"`
	MediaId          string         `json:"mediaId,omitempty"`
	Position         int            `json:"position,omitempty"`
	Product          *Product       `json:"product,omitempty"`
	ProductId        string         `json:"productId,omitempty"`
	ProductVersionId string         `json:"productVersionId,omitempty"`
	UpdatedAt        *time.Time     `json:"updatedAt,omitempty"`
	VersionId        string         `json:"versionId,omitempty"`
}

// ProductOption data structure
// Added since version: 6.0.0.0
// Required fields: productId, optionId
type ProductOption struct {
	Id               string               `json:"id,omitempty"`
	Option           *PropertyGroupOption `json:"option,omitempty"`
	OptionId         string               `json:"optionId,omitempty"`
	Product          *Product             `json:"product,omitempty"`
	ProductId        string               `json:"productId,omitempty"`
	ProductVersionId string               `json:"productVersionId,omitempty"`
}

// ProductPrice data structure
// Added since version: 6.0.0.0
// Required fields: productId, ruleId, price, quantityStart, createdAt
type ProductPrice struct {
	CreatedAt        *time.Time     `json:"createdAt,omitempty"`
	CustomFields     *[]CustomField `json:"customFields,omitempty"`
	Id               string         `json:"id,omitempty"`
	Price            *interface{}   `json:"price,omitempty"` // map[type:object]
	Product          *Product       `json:"product,omitempty"`
	ProductId        string         `json:"productId,omitempty"`
	ProductVersionId string         `json:"productVersionId,omitempty"`
	QuantityEnd      int            `json:"quantityEnd,omitempty"`
	QuantityStart    int            `json:"quantityStart,omitempty"`
	Rule             *Rule          `json:"rule,omitempty"`
	RuleId           string         `json:"ruleId,omitempty"`
	UpdatedAt        *time.Time     `json:"updatedAt,omitempty"`
	VersionId        string         `json:"versionId,omitempty"`
}

// ProductProperty data structure
// Added since version: 6.0.0.0
// Required fields: productId, optionId
type ProductProperty struct {
	Id               string               `json:"id,omitempty"`
	Option           *PropertyGroupOption `json:"option,omitempty"`
	OptionId         string               `json:"optionId,omitempty"`
	Product          *Product             `json:"product,omitempty"`
	ProductId        string               `json:"productId,omitempty"`
	ProductVersionId string               `json:"productVersionId,omitempty"`
}

// ProductReview data structure
// Added since version: 6.0.0.0
// Required fields: productId, salesChannelId, languageId, title, content, createdAt
type ProductReview struct {
	Comment          string         `json:"comment,omitempty"`
	Content          string         `json:"content,omitempty"`
	CreatedAt        *time.Time     `json:"createdAt,omitempty"`
	CustomFields     *[]CustomField `json:"customFields,omitempty"`
	Customer         *Customer      `json:"customer,omitempty"`
	CustomerId       string         `json:"customerId,omitempty"`
	ExternalEmail    string         `json:"externalEmail,omitempty"`
	ExternalUser     string         `json:"externalUser,omitempty"`
	Id               string         `json:"id,omitempty"`
	Language         *Language      `json:"language,omitempty"`
	LanguageId       string         `json:"languageId,omitempty"`
	Points           float64        `json:"points,omitempty"`
	Product          *Product       `json:"product,omitempty"`
	ProductId        string         `json:"productId,omitempty"`
	ProductVersionId string         `json:"productVersionId,omitempty"`
	SalesChannel     *SalesChannel  `json:"salesChannel,omitempty"`
	SalesChannelId   string         `json:"salesChannelId,omitempty"`
	Status           bool           `json:"status,omitempty"`
	Title            string         `json:"title,omitempty"`
	UpdatedAt        *time.Time     `json:"updatedAt,omitempty"`
}

// ProductSearchConfig data structure
// Added since version: 6.3.5.0
// Required fields: languageId, andLogic, minSearchLength, createdAt
type ProductSearchConfig struct {
	AndLogic        bool                      `json:"andLogic,omitempty"`
	ConfigFields    *ProductSearchConfigField `json:"configFields,omitempty"`
	CreatedAt       *time.Time                `json:"createdAt,omitempty"`
	ExcludedTerms   *interface{}              `json:"excludedTerms,omitempty"` // map[items:map[type:string] type:array]
	Id              string                    `json:"id,omitempty"`
	Language        *Language                 `json:"language,omitempty"`
	LanguageId      string                    `json:"languageId,omitempty"`
	MinSearchLength int                       `json:"minSearchLength,omitempty"`
	UpdatedAt       *time.Time                `json:"updatedAt,omitempty"`
}

// ProductSearchConfigField data structure
// Added since version: 6.3.5.0
// Required fields: searchConfigId, field, tokenize, searchable, ranking, createdAt
type ProductSearchConfigField struct {
	CreatedAt      *time.Time           `json:"createdAt,omitempty"`
	CustomField    *CustomField         `json:"customField,omitempty"`
	CustomFieldId  string               `json:"customFieldId,omitempty"`
	Field          string               `json:"field,omitempty"`
	Id             string               `json:"id,omitempty"`
	Ranking        int                  `json:"ranking,omitempty"`
	SearchConfig   *ProductSearchConfig `json:"searchConfig,omitempty"`
	SearchConfigId string               `json:"searchConfigId,omitempty"`
	Searchable     bool                 `json:"searchable,omitempty"`
	Tokenize       bool                 `json:"tokenize,omitempty"`
	UpdatedAt      *time.Time           `json:"updatedAt,omitempty"`
}

// ProductSearchKeyword data structure
// Added since version: 6.0.0.0
// Required fields: languageId, productId, keyword, ranking, createdAt
type ProductSearchKeyword struct {
	CreatedAt        *time.Time `json:"createdAt,omitempty"`
	Id               string     `json:"id,omitempty"`
	Keyword          string     `json:"keyword,omitempty"`
	Language         *Language  `json:"language,omitempty"`
	LanguageId       string     `json:"languageId,omitempty"`
	Product          *Product   `json:"product,omitempty"`
	ProductId        string     `json:"productId,omitempty"`
	ProductVersionId string     `json:"productVersionId,omitempty"`
	Ranking          float64    `json:"ranking,omitempty"`
	UpdatedAt        *time.Time `json:"updatedAt,omitempty"`
	VersionId        string     `json:"versionId,omitempty"`
}

// ProductSorting data structure
// Added since version: 6.3.2.0
// Required fields: key, priority, active, fields, createdAt, label
type ProductSorting struct {
	Active     bool         `json:"active,omitempty"`
	CreatedAt  *time.Time   `json:"createdAt,omitempty"`
	Fields     *interface{} `json:"fields,omitempty"` // map[type:object]
	Id         string       `json:"id,omitempty"`
	Key        string       `json:"key,omitempty"`
	Label      string       `json:"label,omitempty"`
	Locked     bool         `json:"locked,omitempty"`
	Priority   int          `json:"priority,omitempty"`
	Translated *interface{} `json:"translated,omitempty"` // map[type:object]
	UpdatedAt  *time.Time   `json:"updatedAt,omitempty"`
}

// ProductStream data structure
// Added since version: 6.0.0.0
// Required fields: createdAt, name
type ProductStream struct {
	ApiFilter            *interface{}         `json:"apiFilter,omitempty"` // map[readOnly:true type:object]
	Categories           *Category            `json:"categories,omitempty"`
	CreatedAt            *time.Time           `json:"createdAt,omitempty"`
	CustomFields         *[]CustomField       `json:"customFields,omitempty"`
	Description          string               `json:"description,omitempty"`
	Filters              *ProductStreamFilter `json:"filters,omitempty"`
	Id                   string               `json:"id,omitempty"`
	Invalid              bool                 `json:"invalid,omitempty"`
	Name                 string               `json:"name,omitempty"`
	ProductCrossSellings *ProductCrossSelling `json:"productCrossSellings,omitempty"`
	ProductExports       *ProductExport       `json:"productExports,omitempty"`
	Translated           *interface{}         `json:"translated,omitempty"` // map[type:object]
	UpdatedAt            *time.Time           `json:"updatedAt,omitempty"`
}

// ProductStreamFilter data structure
// Added since version: 6.0.0.0
// Required fields: productStreamId, type, createdAt
type ProductStreamFilter struct {
	CreatedAt       *time.Time           `json:"createdAt,omitempty"`
	CustomFields    *[]CustomField       `json:"customFields,omitempty"`
	Field           string               `json:"field,omitempty"`
	Id              string               `json:"id,omitempty"`
	Operator        string               `json:"operator,omitempty"`
	Parameters      *interface{}         `json:"parameters,omitempty"` // map[type:object]
	Parent          *ProductStreamFilter `json:"parent,omitempty"`
	ParentId        string               `json:"parentId,omitempty"`
	Position        int                  `json:"position,omitempty"`
	ProductStream   *ProductStream       `json:"productStream,omitempty"`
	ProductStreamId string               `json:"productStreamId,omitempty"`
	Queries         *ProductStreamFilter `json:"queries,omitempty"`
	Type            string               `json:"type,omitempty"`
	UpdatedAt       *time.Time           `json:"updatedAt,omitempty"`
	Value           string               `json:"value,omitempty"`
}

// ProductStreamMapping data structure
// Added since version: 6.4.0.0
// Required fields: productId, productStreamId
type ProductStreamMapping struct {
	Id               string         `json:"id,omitempty"`
	Product          *Product       `json:"product,omitempty"`
	ProductId        string         `json:"productId,omitempty"`
	ProductStream    *ProductStream `json:"productStream,omitempty"`
	ProductStreamId  string         `json:"productStreamId,omitempty"`
	ProductVersionId string         `json:"productVersionId,omitempty"`
}

// ProductTag data structure
// Added since version: 6.0.0.0
// Required fields: productId, tagId
type ProductTag struct {
	Id               string   `json:"id,omitempty"`
	Product          *Product `json:"product,omitempty"`
	ProductId        string   `json:"productId,omitempty"`
	ProductVersionId string   `json:"productVersionId,omitempty"`
	Tag              *Tag     `json:"tag,omitempty"`
	TagId            string   `json:"tagId,omitempty"`
}

// ProductVisibility data structure
// Added since version: 6.0.0.0
// Required fields: productId, salesChannelId, visibility, createdAt
type ProductVisibility struct {
	CreatedAt        *time.Time    `json:"createdAt,omitempty"`
	Id               string        `json:"id,omitempty"`
	Product          *Product      `json:"product,omitempty"`
	ProductId        string        `json:"productId,omitempty"`
	ProductVersionId string        `json:"productVersionId,omitempty"`
	SalesChannel     *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId   string        `json:"salesChannelId,omitempty"`
	UpdatedAt        *time.Time    `json:"updatedAt,omitempty"`
	Visibility       int           `json:"visibility,omitempty"`
}
