package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-06 18:44:04 UTC

// Tax data structure
// Added since version: 6.0.0.0
// Required fields: taxRate, name, position, createdAt
type Tax struct {
	CreatedAt       *time.Time      `json:"createdAt,omitempty"`
	CustomFields    *[]CustomField  `json:"customFields,omitempty"`
	Id              string          `json:"id,omitempty"`
	Name            string          `json:"name,omitempty"`
	Position        int             `json:"position,omitempty"`
	Products        *Product        `json:"products,omitempty"`
	Rules           *TaxRule        `json:"rules,omitempty"`
	ShippingMethods *ShippingMethod `json:"shippingMethods,omitempty"`
	TaxRate         float64         `json:"taxRate,omitempty"`
	UpdatedAt       *time.Time      `json:"updatedAt,omitempty"`
}

// TaxRule data structure
// Added since version: 6.1.0.0
// Required fields: taxRuleTypeId, countryId, taxRate, taxId, createdAt
type TaxRule struct {
	Country       *Country     `json:"country,omitempty"`
	CountryId     string       `json:"countryId,omitempty"`
	CreatedAt     *time.Time   `json:"createdAt,omitempty"`
	Data          *interface{} `json:"data,omitempty"` // map[properties:map[fromZipCode:map[type:string] states:map[items:map[additionalProperties:false] type:array] toZipCode:map[type:string] zipCode:map[type:string]] type:object]
	Id            string       `json:"id,omitempty"`
	Tax           *Tax         `json:"tax,omitempty"`
	TaxId         string       `json:"taxId,omitempty"`
	TaxRate       float64      `json:"taxRate,omitempty"`
	TaxRuleTypeId string       `json:"taxRuleTypeId,omitempty"`
	Type          *TaxRuleType `json:"type,omitempty"`
	UpdatedAt     *time.Time   `json:"updatedAt,omitempty"`
}

// TaxRuleType data structure
// Added since version: 6.1.0.0
// Required fields: technicalName, position, createdAt, typeName
type TaxRuleType struct {
	CreatedAt     *time.Time   `json:"createdAt,omitempty"`
	Id            string       `json:"id,omitempty"`
	Position      int          `json:"position,omitempty"`
	Rules         *TaxRule     `json:"rules,omitempty"`
	TechnicalName string       `json:"technicalName,omitempty"`
	Translated    *interface{} `json:"translated,omitempty"` // map[type:object]
	TypeName      string       `json:"typeName,omitempty"`
	UpdatedAt     *time.Time   `json:"updatedAt,omitempty"`
}
