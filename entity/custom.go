package entity

import "time"

// Generated from Shopware Admin API
// Version 6.5.9999999.9999999-dev at 2023-05-02 19:06:34 UTC

// CustomEntity data structure
// Added since version: 6.4.9.0
// Required fields: name, fields, createdAt
type CustomEntity struct {
	AppId     *string      `json:"appId,omitempty"`
	CreatedAt *time.Time   `json:"createdAt,omitempty"`
	Fields    *interface{} `json:"fields,omitempty"` // map[type:object]
	Id        string       `json:"id,omitempty"`
	Name      string       `json:"name,omitempty"`
	UpdatedAt *time.Time   `json:"updatedAt,omitempty"`
}

// CustomField data structure
// Added since version: 6.0.0.0
// Required fields: name, type, createdAt
type CustomField struct {
	Active                    *bool                     `json:"active,omitempty"`
	AllowCustomerWrite        *bool                     `json:"allowCustomerWrite,omitempty"`
	Config                    *interface{}              `json:"config,omitempty"` // map[type:object]
	CreatedAt                 *time.Time                `json:"createdAt,omitempty"`
	CustomFieldSet            *CustomFieldSet           `json:"customFieldSet,omitempty"`
	CustomFieldSetId          *string                   `json:"customFieldSetId,omitempty"`
	Id                        string                    `json:"id,omitempty"`
	Name                      string                    `json:"name,omitempty"`
	ProductSearchConfigFields *ProductSearchConfigField `json:"productSearchConfigFields,omitempty"`
	Type                      string                    `json:"type,omitempty"`
	UpdatedAt                 *time.Time                `json:"updatedAt,omitempty"`
}

// CustomFieldSet data structure
// Added since version: 6.0.0.0
// Required fields: name, createdAt
type CustomFieldSet struct {
	Active       *bool                   `json:"active,omitempty"`
	App          *App                    `json:"app,omitempty"`
	AppId        *string                 `json:"appId,omitempty"`
	Config       *interface{}            `json:"config,omitempty"` // map[type:object]
	CreatedAt    *time.Time              `json:"createdAt,omitempty"`
	CustomFields *[]CustomField          `json:"customFields,omitempty"`
	Global       *bool                   `json:"global,omitempty"`
	Id           string                  `json:"id,omitempty"`
	Name         string                  `json:"name,omitempty"`
	Position     *int                    `json:"position,omitempty"`
	Products     *Product                `json:"products,omitempty"`
	Relations    *CustomFieldSetRelation `json:"relations,omitempty"`
	UpdatedAt    *time.Time              `json:"updatedAt,omitempty"`
}

// CustomFieldSetRelation data structure
// Added since version: 6.0.0.0
// Required fields: customFieldSetId, entityName, createdAt
type CustomFieldSetRelation struct {
	CreatedAt        *time.Time      `json:"createdAt,omitempty"`
	CustomFieldSet   *CustomFieldSet `json:"customFieldSet,omitempty"`
	CustomFieldSetId string          `json:"customFieldSetId,omitempty"`
	EntityName       string          `json:"entityName,omitempty"`
	Id               string          `json:"id,omitempty"`
	UpdatedAt        *time.Time      `json:"updatedAt,omitempty"`
}

// CustomPrice data structure
// Added since version:
// Required fields: productId, price, createdAt
type CustomPrice struct {
	CreatedAt        *time.Time     `json:"createdAt,omitempty"`
	Customer         *Customer      `json:"customer,omitempty"`
	CustomerGroup    *CustomerGroup `json:"customerGroup,omitempty"`
	CustomerGroupId  *string        `json:"customerGroupId,omitempty"`
	CustomerId       *string        `json:"customerId,omitempty"`
	Id               string         `json:"id,omitempty"`
	Price            *interface{}   `json:"price,omitempty"` // map[type:object]
	Product          *Product       `json:"product,omitempty"`
	ProductId        string         `json:"productId,omitempty"`
	ProductVersionId *string        `json:"productVersionId,omitempty"`
	UpdatedAt        *time.Time     `json:"updatedAt,omitempty"`
}

// CustomPricingPrice data structure
// An extended schema for the Custom Price 'price' column (as opposed to ProductPrice 'price' column)
// Required fields: quantityStart, quantityEnd, price
type CustomPricingPrice struct {
	Price         *interface{} `json:"price,omitempty"`       // map[items:map[description:This field should house all the normal facets of the `product_price`.`price` column properties:map[currencyId:map[pattern:^[0-9a-f]{32}$ type:string] gross:map[minimum:0 type:number] linked:map[type:boolean] listPrice:map[properties:map[gross:map[minimum:0 type:number] linked:map[type:boolean] net:map[minimum:0 type:number]] type:object] net:map[minimum:0 type:number] regulationPrice:map[properties:map[gross:map[minimum:0 type:number] linked:map[type:boolean] net:map[minimum:0 type:number]] type:object]] required:[currencyId gross net linked] type:object] type:array]
	QuantityEnd   *interface{} `json:"quantityEnd,omitempty"` // map[oneOf:[map[format:int64 minimum:1 type:integer] map[type:null]]]
	QuantityStart int          `json:"quantityStart,omitempty"`
}

// CustomPricingUpsertOperation data structure
// Required fields: action, payload
type CustomPricingUpsertOperation struct {
	Action  string       `json:"action,omitempty"`
	Payload *interface{} `json:"payload,omitempty"` // map[description:Contains a list of changesets for an entity. If the action type is `delete`, a list of identifiers can be provided. items:map[description:A definition almost symmetric with Sync API payload properties:map[customerGroupId:map[description:This parameter should be specified, or 'customerId', not both pattern:^[0-9a-f]{32}$ type:string] customerId:map[description:This parameter should be specified, or 'customerGroupId', not both pattern:^[0-9a-f]{32}$ type:string] id:map[pattern:^[0-9a-f]{32}$ type:string] price:map[items:map[$ref:#/components/schemas/CustomPricingPrice] type:array] productId:map[pattern:^[0-9a-f]{32}$ type:string] productVersionId:map[pattern:^[0-9a-f]{32}$ type:string]] required:[productId prices] type:object] type:array]
}

// CustomPricingResponse data structure
// Required fields: success, data
type CustomPricingResponse struct {
	Data    *interface{} `json:"data,omitempty"` // map[items:map[properties:map[extensions:map[items:map[type:object] type:array] result:map[items:map[properties:map[entities:map[items:map[pattern:^[0-9a-f]{32}$ type:string] maxItems:1 type:array] errors:map[description:A detailed error list addressing specific points in which sync payload does not meet system expectations (data types, structure etc.) items:map[type:string] type:array]] required:[entities errors] type:object] type:array]] type:object] type:array]
	Success bool         `json:"success,omitempty"`
}
