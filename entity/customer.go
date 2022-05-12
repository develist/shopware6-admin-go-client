package entity

import "time"

// completed

type Customer struct {
	Active                   bool              `json:"active,omitempty"`
	Addresses                *CustomerAddress  `json:"addresses,omitempty"`
	AffiliateCode            string            `json:"affiliateCode,omitempty"`
	AutoIncrement            int               `json:"autoIncrement,omitempty"`
	Birthday                 string            `json:"birthday,omitempty"`
	BoundSalesChannel        *SalesChannel     `json:"boundSalesChannel,omitempty"`
	BoundSalesChannelId      string            `json:"boundSalesChannelId,omitempty"`
	CampaignCode             string            `json:"campaignCode,omitempty"`
	Company                  string            `json:"company,omitempty"`
	CreatedAt                time.Time         `json:"createdAt,omitempty"`
	CustomFields             *[]CustomField    `json:"customFields,omitempty"`
	CustomerNumber           string            `json:"customerNumber,omitempty"`
	DefaultBillingAddress    *CustomerAddress  `json:"defaultBillingAddress,omitempty"`
	DefaultBillingAddressId  string            `json:"defaultBillingAddressId,omitempty"`
	DefaultPaymentMethod     *PaymentMethod    `json:"defaultPaymentMethod,omitempty"`
	DefaultPaymentMethodId   string            `json:"defaultPaymentMethodId,omitempty"`
	DefaultShippingAddress   *CustomerAddress  `json:"defaultShippingAddress,omitempty"`
	DefaultShippingAddressId string            `json:"defaultShippingAddressId,omitempty"`
	DoubleOptInConfirmDate   time.Time         `json:"doubleOptInConfirmDate,omitempty"`
	DoubleOptInEmailSentDate time.Time         `json:"doubleOptInEmailSentDate,omitempty"`
	DoubleOptInRegistration  bool              `json:"doubleOptInRegistration,omitempty"`
	Email                    string            `json:"email,omitempty"`
	FirstLogin               time.Time         `json:"firstLogin,omitempty"`
	FirstName                string            `json:"firstName,omitempty"`
	Group                    *CustomerGroup    `json:"group,omitempty"`
	GroupId                  string            `json:"groupId,omitempty"`
	Guest                    bool              `json:"guest,omitempty"`
	Hash                     string            `json:"hash,omitempty"`
	Id                       string            `json:"id,omitempty"`
	Language                 *Language         `json:"language,omitempty"`
	LanguageId               string            `json:"languageId,omitempty"`
	LastLogin                time.Time         `json:"lastLogin,omitempty"`
	LastName                 string            `json:"lastName,omitempty"`
	LastOrderDate            time.Time         `json:"lastOrderDate,omitempty"`
	LastPaymentMethod        *PaymentMethod    `json:"lastPaymentMethod,omitempty"`
	LastPaymentMethodId      string            `json:"lastPaymentMethodId,omitempty"`
	Newsletter               bool              `json:"newsletter,omitempty"`
	OrderCount               int               `json:"orderCount,omitempty"`
	OrderCustomers           *OrderCustomer    `json:"orderCustomers,omitempty"`
	ProductReviews           *ProductReview    `json:"productReviews,omitempty"`
	Promotions               *Promotion        `json:"promotions,omitempty"`
	RecoveryCustomer         *CustomerRecovery `json:"recoveryCustomer,omitempty"`
	RemoteAddress            string            `json:"remoteAddress,omitempty"`
	RequestedGroup           *CustomerGroup    `json:"requestedGroup,omitempty"`
	RequestedGroupId         string            `json:"requestedGroupId,omitempty"`
	SalesChannel             *SalesChannel     `json:"salesChannel,omitempty"`
	SalesChannelId           string            `json:"salesChannelId,omitempty"`
	Salutation               *Salutation       `json:"salutation,omitempty"`
	SalutationId             string            `json:"salutationId,omitempty"`
	// TagIds map[items:map[pattern:^[0-9a-f]{32}$ type:string] readOnly:true type:array] `json:"tagIds,omitempty"`
	Tags      *Tag      `json:"tags,omitempty"`
	Title     string    `json:"title,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// VatIds map[items:map[type:string] type:array] `json:"vatIds,omitempty"`
	Wishlists *CustomerWishlist `json:"wishlists,omitempty"`
}

type CustomerAddress struct {
	AdditionalAddressLine1 string         `json:"additionalAddressLine1,omitempty"`
	AdditionalAddressLine2 string         `json:"additionalAddressLine2,omitempty"`
	City                   string         `json:"city,omitempty"`
	Company                string         `json:"company,omitempty"`
	Country                *Country       `json:"country,omitempty"`
	CountryId              string         `json:"countryId,omitempty"`
	CountryState           *CountryState  `json:"countryState,omitempty"`
	CountryStateId         string         `json:"countryStateId,omitempty"`
	CreatedAt              time.Time      `json:"createdAt,omitempty"`
	CustomFields           *[]CustomField `json:"customFields,omitempty"`
	Customer               *Customer      `json:"customer,omitempty"`
	CustomerId             string         `json:"customerId,omitempty"`
	Department             string         `json:"department,omitempty"`
	FirstName              string         `json:"firstName,omitempty"`
	Id                     string         `json:"id,omitempty"`
	LastName               string         `json:"lastName,omitempty"`
	PhoneNumber            string         `json:"phoneNumber,omitempty"`
	Salutation             *Salutation    `json:"salutation,omitempty"`
	SalutationId           string         `json:"salutationId,omitempty"`
	Street                 string         `json:"street,omitempty"`
	Title                  string         `json:"title,omitempty"`
	UpdatedAt              time.Time      `json:"updatedAt,omitempty"`
	Zipcode                string         `json:"zipcode,omitempty"`
}

type CustomerGroup struct {
	CreatedAt                           time.Time      `json:"createdAt,omitempty"`
	CustomFields                        *[]CustomField `json:"customFields,omitempty"`
	Customers                           *Customer      `json:"customers,omitempty"`
	DisplayGross                        bool           `json:"displayGross,omitempty"`
	Id                                  string         `json:"id,omitempty"`
	Name                                string         `json:"name,omitempty"`
	RegistrationActive                  bool           `json:"registrationActive,omitempty"`
	RegistrationIntroduction            string         `json:"registrationIntroduction,omitempty"`
	RegistrationOnlyCompanyRegistration bool           `json:"registrationOnlyCompanyRegistration,omitempty"`
	RegistrationSalesChannels           *SalesChannel  `json:"registrationSalesChannels,omitempty"`
	RegistrationSeoMetaDescription      string         `json:"registrationSeoMetaDescription,omitempty"`
	RegistrationTitle                   string         `json:"registrationTitle,omitempty"`
	SalesChannels                       *SalesChannel  `json:"salesChannels,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type CustomerGroupRegistrationSalesChannels struct {
	CreatedAt       time.Time      `json:"createdAt,omitempty"`
	CustomerGroup   *CustomerGroup `json:"customerGroup,omitempty"`
	CustomerGroupId string         `json:"customerGroupId,omitempty"`
	Id              string         `json:"id,omitempty"`
	SalesChannel    *SalesChannel  `json:"salesChannel,omitempty"`
	SalesChannelId  string         `json:"salesChannelId,omitempty"`
}

type CustomerRecovery struct {
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	Customer   *Customer `json:"customer,omitempty"`
	CustomerId string    `json:"customerId,omitempty"`
	Hash       string    `json:"hash,omitempty"`
	Id         string    `json:"id,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
}

type CustomerTag struct {
	Customer   *Customer `json:"customer,omitempty"`
	CustomerId string    `json:"customerId,omitempty"`
	Id         string    `json:"id,omitempty"`
	Tag        *Tag      `json:"tag,omitempty"`
	TagId      string    `json:"tagId,omitempty"`
}

type CustomerWishlist struct {
	CreatedAt      time.Time                `json:"createdAt,omitempty"`
	CustomFields   *[]CustomField           `json:"customFields,omitempty"`
	Customer       *Customer                `json:"customer,omitempty"`
	CustomerId     string                   `json:"customerId,omitempty"`
	Id             string                   `json:"id,omitempty"`
	Products       *CustomerWishlistProduct `json:"products,omitempty"`
	SalesChannel   *SalesChannel            `json:"salesChannel,omitempty"`
	SalesChannelId string                   `json:"salesChannelId,omitempty"`
	UpdatedAt      time.Time                `json:"updatedAt,omitempty"`
}

type CustomerWishlistProduct struct {
	CreatedAt        time.Time         `json:"createdAt,omitempty"`
	Id               string            `json:"id,omitempty"`
	Product          *Product          `json:"product,omitempty"`
	ProductId        string            `json:"productId,omitempty"`
	ProductVersionId string            `json:"productVersionId,omitempty"`
	UpdatedAt        time.Time         `json:"updatedAt,omitempty"`
	Wishlist         *CustomerWishlist `json:"wishlist,omitempty"`
	WishlistId       string            `json:"wishlistId,omitempty"`
}
