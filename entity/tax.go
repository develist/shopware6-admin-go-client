package entity

import "time"

// completed

type Tax struct {
	CreatedAt       time.Time       `json:"createdAt,omitempty"`
	CustomFields    *[]CustomField  `json:"customFields,omitempty"`
	Id              string          `json:"id,omitempty"`
	Name            string          `json:"name,omitempty"`
	Position        int             `json:"position,omitempty"`
	Products        *Product        `json:"products,omitempty"`
	Rules           *TaxRule        `json:"rules,omitempty"`
	ShippingMethods *ShippingMethod `json:"shippingMethods,omitempty"`
	TaxRate         float64         `json:"taxRate,omitempty"`
	UpdatedAt       time.Time       `json:"updatedAt,omitempty"`
}

type TaxRule struct {
	Country   *Country  `json:"country,omitempty"`
	CountryId string    `json:"countryId,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// Data map[properties:map[fromZipCode:map[type:string] states:map[items:map[additionalProperties:false] type:array] toZipCode:map[type:string] zipCode:map[type:string]] type:object] `json:"data,omitempty"`
	Id            string       `json:"id,omitempty"`
	Tax           *Tax         `json:"tax,omitempty"`
	TaxId         string       `json:"taxId,omitempty"`
	TaxRate       float64      `json:"taxRate,omitempty"`
	TaxRuleTypeId string       `json:"taxRuleTypeId,omitempty"`
	Type          *TaxRuleType `json:"type,omitempty"`
	UpdatedAt     time.Time    `json:"updatedAt,omitempty"`
}

type TaxRuleType struct {
	CreatedAt     time.Time `json:"createdAt,omitempty"`
	Id            string    `json:"id,omitempty"`
	Position      int       `json:"position,omitempty"`
	Rules         *TaxRule  `json:"rules,omitempty"`
	TechnicalName string    `json:"technicalName,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	TypeName  string    `json:"typeName,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
