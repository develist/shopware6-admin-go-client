package entity

import "time"

// completed

type Product struct {
	Active             bool      `json:"active,omitempty"`
	AutoIncrement      int       `json:"autoIncrement,omitempty"`
	Available          bool      `json:"available,omitempty"`
	AvailableStock     int       `json:"availableStock,omitempty"`
	CanonicalProduct   *Product  `json:"canonicalProduct,omitempty"`
	CanonicalProductId string    `json:"canonicalProductId,omitempty"`
	Categories         *Category `json:"categories,omitempty"`
	CategoriesRo       *Category `json:"categoriesRo,omitempty"`
	// CategoryIds map[items:map[pattern:^[0-9a-f]{32}$ type:string] readOnly:true type:array] `json:"categoryIds,omitempty"`
	// CategoryTree map[items:map[pattern:^[0-9a-f]{32}$ type:string] readOnly:true type:array] `json:"categoryTree,omitempty"`
	// CheapestPrice map[readOnly:true type:object] `json:"cheapestPrice,omitempty"`
	ChildCount       int      `json:"childCount,omitempty"`
	Children         *Product `json:"children,omitempty"`
	CmsPage          *CmsPage `json:"cmsPage,omitempty"`
	CmsPageId        string   `json:"cmsPageId,omitempty"`
	CmsPageVersionId string   `json:"cmsPageVersionId,omitempty"`
	// ConfiguratorGroupConfig map[type:object] `json:"configuratorGroupConfig,omitempty"`
	ConfiguratorSettings          *ProductConfiguratorSetting          `json:"configuratorSettings,omitempty"`
	Cover                         *ProductMedia                        `json:"cover,omitempty"`
	CoverId                       string                               `json:"coverId,omitempty"`
	CreatedAt                     time.Time                            `json:"createdAt,omitempty"`
	CrossSellingAssignedProducts  *ProductCrossSellingAssignedProducts `json:"crossSellingAssignedProducts,omitempty"`
	CrossSellings                 *ProductCrossSelling                 `json:"crossSellings,omitempty"`
	CustomFieldSetSelectionActive bool                                 `json:"customFieldSetSelectionActive,omitempty"`
	CustomFieldSets               *CustomFieldSet                      `json:"customFieldSets,omitempty"`
	CustomFields                  *[]CustomField                       `json:"customFields,omitempty"`
	// CustomSearchKeywords map[items:map[additionalProperties:false] type:array] `json:"customSearchKeywords,omitempty"`
	DeliveryTime       *DeliveryTime        `json:"deliveryTime,omitempty"`
	DeliveryTimeId     string               `json:"deliveryTimeId,omitempty"`
	Description        string               `json:"description,omitempty"`
	DisplayGroup       string               `json:"displayGroup,omitempty"`
	Ean                string               `json:"ean,omitempty"`
	FeatureSet         *ProductFeatureSet   `json:"featureSet,omitempty"`
	FeatureSetId       string               `json:"featureSetId,omitempty"`
	Height             float64              `json:"height,omitempty"`
	Id                 string               `json:"id,omitempty"`
	IsCloseout         bool                 `json:"isCloseout,omitempty"`
	Keywords           string               `json:"keywords,omitempty"`
	Length             float64              `json:"length,omitempty"`
	MainCategories     *MainCategory        `json:"mainCategories,omitempty"`
	MainVariantId      string               `json:"mainVariantId,omitempty"`
	Manufacturer       *ProductManufacturer `json:"manufacturer,omitempty"`
	ManufacturerId     string               `json:"manufacturerId,omitempty"`
	ManufacturerNumber string               `json:"manufacturerNumber,omitempty"`
	MarkAsTopseller    bool                 `json:"markAsTopseller,omitempty"`
	MaxPurchase        int                  `json:"maxPurchase,omitempty"`
	Media              *ProductMedia        `json:"media,omitempty"`
	MetaDescription    string               `json:"metaDescription,omitempty"`
	MetaTitle          string               `json:"metaTitle,omitempty"`
	MinPurchase        int                  `json:"minPurchase,omitempty"`
	Name               string               `json:"name,omitempty"`
	// OptionIds map[items:map[pattern:^[0-9a-f]{32}$ type:string] readOnly:true type:array] `json:"optionIds,omitempty"`
	Options         *PropertyGroupOption `json:"options,omitempty"`
	OrderLineItems  *OrderLineItem       `json:"orderLineItems,omitempty"`
	PackUnit        string               `json:"packUnit,omitempty"`
	PackUnitPlural  string               `json:"packUnitPlural,omitempty"`
	Parent          *Product             `json:"parent,omitempty"`
	ParentId        string               `json:"parentId,omitempty"`
	ParentVersionId string               `json:"parentVersionId,omitempty"`
	// Price map[type:object] `json:"price,omitempty"`
	Prices                       *ProductPrice        `json:"prices,omitempty"`
	ProductManufacturerVersionId string               `json:"productManufacturerVersionId,omitempty"`
	ProductMediaVersionId        string               `json:"productMediaVersionId,omitempty"`
	ProductNumber                string               `json:"productNumber,omitempty"`
	ProductReviews               *ProductReview       `json:"productReviews,omitempty"`
	Properties                   *PropertyGroupOption `json:"properties,omitempty"`
	// PropertyIds map[items:map[pattern:^[0-9a-f]{32}$ type:string] readOnly:true type:array] `json:"propertyIds,omitempty"`
	// PurchasePrices map[type:object] `json:"purchasePrices,omitempty"`
	PurchaseSteps  int                   `json:"purchaseSteps,omitempty"`
	PurchaseUnit   float64               `json:"purchaseUnit,omitempty"`
	RatingAverage  float64               `json:"ratingAverage,omitempty"`
	ReferenceUnit  float64               `json:"referenceUnit,omitempty"`
	ReleaseDate    time.Time             `json:"releaseDate,omitempty"`
	RestockTime    int                   `json:"restockTime,omitempty"`
	Sales          int                   `json:"sales,omitempty"`
	SearchKeywords *ProductSearchKeyword `json:"searchKeywords,omitempty"`
	SeoUrls        *SeoUrl               `json:"seoUrls,omitempty"`
	ShippingFree   bool                  `json:"shippingFree,omitempty"`
	// SlotConfig map[type:object] `json:"slotConfig,omitempty"`
	Stock   int            `json:"stock,omitempty"`
	Streams *ProductStream `json:"streams,omitempty"`
	// TagIds map[items:map[pattern:^[0-9a-f]{32}$ type:string] readOnly:true type:array] `json:"tagIds,omitempty"`
	Tags  *Tag   `json:"tags,omitempty"`
	Tax   *Tax   `json:"tax,omitempty"`
	TaxId string `json:"taxId,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	Unit      *Unit     `json:"unit,omitempty"`
	UnitId    string    `json:"unitId,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// VariantRestrictions map[type:object] `json:"variantRestrictions,omitempty"`
	// Variation map[items:map[type:string] type:array] `json:"variation,omitempty"`
	VersionId    string                   `json:"versionId,omitempty"`
	Visibilities *ProductVisibility       `json:"visibilities,omitempty"`
	Weight       float64                  `json:"weight,omitempty"`
	Width        float64                  `json:"width,omitempty"`
	Wishlists    *CustomerWishlistProduct `json:"wishlists,omitempty"`
}

type ProductCategory struct {
	Category          *Category `json:"category,omitempty"`
	CategoryId        string    `json:"categoryId,omitempty"`
	CategoryVersionId string    `json:"categoryVersionId,omitempty"`
	Id                string    `json:"id,omitempty"`
	Product           *Product  `json:"product,omitempty"`
	ProductId         string    `json:"productId,omitempty"`
	ProductVersionId  string    `json:"productVersionId,omitempty"`
}

type ProductCategoryTree struct {
	Category          *Category `json:"category,omitempty"`
	CategoryId        string    `json:"categoryId,omitempty"`
	CategoryVersionId string    `json:"categoryVersionId,omitempty"`
	Id                string    `json:"id,omitempty"`
	Product           *Product  `json:"product,omitempty"`
	ProductId         string    `json:"productId,omitempty"`
	ProductVersionId  string    `json:"productVersionId,omitempty"`
}

type ProductConfiguratorSetting struct {
	CreatedAt    time.Time            `json:"createdAt,omitempty"`
	CustomFields *[]CustomField       `json:"customFields,omitempty"`
	Id           string               `json:"id,omitempty"`
	Media        *Media               `json:"media,omitempty"`
	MediaId      string               `json:"mediaId,omitempty"`
	Option       *PropertyGroupOption `json:"option,omitempty"`
	OptionId     string               `json:"optionId,omitempty"`
	Position     int                  `json:"position,omitempty"`
	// Price map[type:object] `json:"price,omitempty"`
	Product          *Product  `json:"product,omitempty"`
	ProductId        string    `json:"productId,omitempty"`
	ProductVersionId string    `json:"productVersionId,omitempty"`
	UpdatedAt        time.Time `json:"updatedAt,omitempty"`
	VersionId        string    `json:"versionId,omitempty"`
}

type ProductCrossSelling struct {
	Active           bool                                 `json:"active,omitempty"`
	AssignedProducts *ProductCrossSellingAssignedProducts `json:"assignedProducts,omitempty"`
	CreatedAt        time.Time                            `json:"createdAt,omitempty"`
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
	// Translated map[type:object] `json:"translated,omitempty"`
	Type      string    `json:"type,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type ProductCrossSellingAssignedProducts struct {
	CreatedAt        time.Time            `json:"createdAt,omitempty"`
	CrossSelling     *ProductCrossSelling `json:"crossSelling,omitempty"`
	CrossSellingId   string               `json:"crossSellingId,omitempty"`
	Id               string               `json:"id,omitempty"`
	Position         int                  `json:"position,omitempty"`
	Product          *Product             `json:"product,omitempty"`
	ProductId        string               `json:"productId,omitempty"`
	ProductVersionId string               `json:"productVersionId,omitempty"`
	UpdatedAt        time.Time            `json:"updatedAt,omitempty"`
}

type ProductCustomFieldSet struct {
	CustomFieldSet   *CustomFieldSet `json:"customFieldSet,omitempty"`
	CustomFieldSetId string          `json:"customFieldSetId,omitempty"`
	Id               string          `json:"id,omitempty"`
	Product          *Product        `json:"product,omitempty"`
	ProductId        string          `json:"productId,omitempty"`
	ProductVersionId string          `json:"productVersionId,omitempty"`
}

type ProductExport struct {
	AccessKey                string              `json:"accessKey,omitempty"`
	BodyTemplate             string              `json:"bodyTemplate,omitempty"`
	CreatedAt                time.Time           `json:"createdAt,omitempty"`
	Currency                 *Currency           `json:"currency,omitempty"`
	CurrencyId               string              `json:"currencyId,omitempty"`
	Encoding                 string              `json:"encoding,omitempty"`
	FileFormat               string              `json:"fileFormat,omitempty"`
	FileName                 string              `json:"fileName,omitempty"`
	FooterTemplate           string              `json:"footerTemplate,omitempty"`
	GenerateByCronjob        bool                `json:"generateByCronjob,omitempty"`
	GeneratedAt              time.Time           `json:"generatedAt,omitempty"`
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
	UpdatedAt                time.Time           `json:"updatedAt,omitempty"`
}

type ProductFeatureSet struct {
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	Description string    `json:"description,omitempty"`
	// Features map[type:object] `json:"features,omitempty"`
	Id       string   `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	Products *Product `json:"products,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type ProductKeywordDictionary struct {
	Id         string    `json:"id,omitempty"`
	Keyword    string    `json:"keyword,omitempty"`
	Language   *Language `json:"language,omitempty"`
	LanguageId string    `json:"languageId,omitempty"`
	Reversed   string    `json:"reversed,omitempty"`
}

type ProductManufacturer struct {
	CreatedAt    time.Time      `json:"createdAt,omitempty"`
	CustomFields *[]CustomField `json:"customFields,omitempty"`
	Description  string         `json:"description,omitempty"`
	Id           string         `json:"id,omitempty"`
	Link         string         `json:"link,omitempty"`
	Media        *Media         `json:"media,omitempty"`
	MediaId      string         `json:"mediaId,omitempty"`
	Name         string         `json:"name,omitempty"`
	Products     *Product       `json:"products,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	VersionId string    `json:"versionId,omitempty"`
}

type ProductMedia struct {
	CreatedAt        time.Time      `json:"createdAt,omitempty"`
	CustomFields     *[]CustomField `json:"customFields,omitempty"`
	Id               string         `json:"id,omitempty"`
	Media            *Media         `json:"media,omitempty"`
	MediaId          string         `json:"mediaId,omitempty"`
	Position         int            `json:"position,omitempty"`
	Product          *Product       `json:"product,omitempty"`
	ProductId        string         `json:"productId,omitempty"`
	ProductVersionId string         `json:"productVersionId,omitempty"`
	UpdatedAt        time.Time      `json:"updatedAt,omitempty"`
	VersionId        string         `json:"versionId,omitempty"`
}

type ProductOption struct {
	Id               string               `json:"id,omitempty"`
	Option           *PropertyGroupOption `json:"option,omitempty"`
	OptionId         string               `json:"optionId,omitempty"`
	Product          *Product             `json:"product,omitempty"`
	ProductId        string               `json:"productId,omitempty"`
	ProductVersionId string               `json:"productVersionId,omitempty"`
}

type ProductPrice struct {
	CreatedAt    time.Time      `json:"createdAt,omitempty"`
	CustomFields *[]CustomField `json:"customFields,omitempty"`
	Id           string         `json:"id,omitempty"`
	// Price map[type:object] `json:"price,omitempty"`
	Product          *Product  `json:"product,omitempty"`
	ProductId        string    `json:"productId,omitempty"`
	ProductVersionId string    `json:"productVersionId,omitempty"`
	QuantityEnd      int       `json:"quantityEnd,omitempty"`
	QuantityStart    int       `json:"quantityStart,omitempty"`
	Rule             *Rule     `json:"rule,omitempty"`
	RuleId           string    `json:"ruleId,omitempty"`
	UpdatedAt        time.Time `json:"updatedAt,omitempty"`
	VersionId        string    `json:"versionId,omitempty"`
}

type ProductProperty struct {
	Id               string               `json:"id,omitempty"`
	Option           *PropertyGroupOption `json:"option,omitempty"`
	OptionId         string               `json:"optionId,omitempty"`
	Product          *Product             `json:"product,omitempty"`
	ProductId        string               `json:"productId,omitempty"`
	ProductVersionId string               `json:"productVersionId,omitempty"`
}
type ProductReview struct {
	Comment          string         `json:"comment,omitempty"`
	Content          string         `json:"content,omitempty"`
	CreatedAt        time.Time      `json:"createdAt,omitempty"`
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
	UpdatedAt        time.Time      `json:"updatedAt,omitempty"`
}

type ProductSearchConfig struct {
	AndLogic     bool                      `json:"andLogic,omitempty"`
	ConfigFields *ProductSearchConfigField `json:"configFields,omitempty"`
	CreatedAt    time.Time                 `json:"createdAt,omitempty"`
	// ExcludedTerms map[items:map[type:string] type:array] `json:"excludedTerms,omitempty"`
	Id              string    `json:"id,omitempty"`
	Language        *Language `json:"language,omitempty"`
	LanguageId      string    `json:"languageId,omitempty"`
	MinSearchLength int       `json:"minSearchLength,omitempty"`
	UpdatedAt       time.Time `json:"updatedAt,omitempty"`
}

type ProductSearchConfigField struct {
	CreatedAt      time.Time            `json:"createdAt,omitempty"`
	CustomField    *CustomField         `json:"customField,omitempty"`
	CustomFieldId  string               `json:"customFieldId,omitempty"`
	Field          string               `json:"field,omitempty"`
	Id             string               `json:"id,omitempty"`
	Ranking        int                  `json:"ranking,omitempty"`
	SearchConfig   *ProductSearchConfig `json:"searchConfig,omitempty"`
	SearchConfigId string               `json:"searchConfigId,omitempty"`
	Searchable     bool                 `json:"searchable,omitempty"`
	Tokenize       bool                 `json:"tokenize,omitempty"`
	UpdatedAt      time.Time            `json:"updatedAt,omitempty"`
}

type ProductSearchKeyword struct {
	CreatedAt        time.Time `json:"createdAt,omitempty"`
	Id               string    `json:"id,omitempty"`
	Keyword          string    `json:"keyword,omitempty"`
	Language         *Language `json:"language,omitempty"`
	LanguageId       string    `json:"languageId,omitempty"`
	Product          *Product  `json:"product,omitempty"`
	ProductId        string    `json:"productId,omitempty"`
	ProductVersionId string    `json:"productVersionId,omitempty"`
	Ranking          float64   `json:"ranking,omitempty"`
	UpdatedAt        time.Time `json:"updatedAt,omitempty"`
	VersionId        string    `json:"versionId,omitempty"`
}

type ProductSorting struct {
	Active    bool      `json:"active,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// Fields map[type:object] `json:"fields,omitempty"`
	Id       string `json:"id,omitempty"`
	Key      string `json:"key,omitempty"`
	Label    string `json:"label,omitempty"`
	Locked   bool   `json:"locked,omitempty"`
	Priority int    `json:"priority,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type ProductStream struct {
	// ApiFilter map[readOnly:true type:object] `json:"apiFilter,omitempty"`
	Categories           *Category            `json:"categories,omitempty"`
	CreatedAt            time.Time            `json:"createdAt,omitempty"`
	CustomFields         *[]CustomField       `json:"customFields,omitempty"`
	Description          string               `json:"description,omitempty"`
	Filters              *ProductStreamFilter `json:"filters,omitempty"`
	Id                   string               `json:"id,omitempty"`
	Invalid              bool                 `json:"invalid,omitempty"`
	Name                 string               `json:"name,omitempty"`
	ProductCrossSellings *ProductCrossSelling `json:"productCrossSellings,omitempty"`
	ProductExports       *ProductExport       `json:"productExports,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type ProductStreamFilter struct {
	CreatedAt    time.Time      `json:"createdAt,omitempty"`
	CustomFields *[]CustomField `json:"customFields,omitempty"`
	Field        string         `json:"field,omitempty"`
	Id           string         `json:"id,omitempty"`
	Operator     string         `json:"operator,omitempty"`
	// Parameters map[type:object] `json:"parameters,omitempty"`
	Parent          *ProductStreamFilter `json:"parent,omitempty"`
	ParentId        string               `json:"parentId,omitempty"`
	Position        int                  `json:"position,omitempty"`
	ProductStream   *ProductStream       `json:"productStream,omitempty"`
	ProductStreamId string               `json:"productStreamId,omitempty"`
	Queries         *ProductStreamFilter `json:"queries,omitempty"`
	Type            string               `json:"type,omitempty"`
	UpdatedAt       time.Time            `json:"updatedAt,omitempty"`
	Value           string               `json:"value,omitempty"`
}

type ProductStreamMapping struct {
	Id               string         `json:"id,omitempty"`
	Product          *Product       `json:"product,omitempty"`
	ProductId        string         `json:"productId,omitempty"`
	ProductStream    *ProductStream `json:"productStream,omitempty"`
	ProductStreamId  string         `json:"productStreamId,omitempty"`
	ProductVersionId string         `json:"productVersionId,omitempty"`
}

type ProductTag struct {
	Id               string   `json:"id,omitempty"`
	Product          *Product `json:"product,omitempty"`
	ProductId        string   `json:"productId,omitempty"`
	ProductVersionId string   `json:"productVersionId,omitempty"`
	Tag              *Tag     `json:"tag,omitempty"`
	TagId            string   `json:"tagId,omitempty"`
}

type ProductVisibility struct {
	CreatedAt        time.Time     `json:"createdAt,omitempty"`
	Id               string        `json:"id,omitempty"`
	Product          *Product      `json:"product,omitempty"`
	ProductId        string        `json:"productId,omitempty"`
	ProductVersionId string        `json:"productVersionId,omitempty"`
	SalesChannel     *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId   string        `json:"salesChannelId,omitempty"`
	UpdatedAt        time.Time     `json:"updatedAt,omitempty"`
	Visibility       int           `json:"visibility,omitempty"`
}
