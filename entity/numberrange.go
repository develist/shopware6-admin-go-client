package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-06 18:44:04 UTC

// NumberRange data structure
// Added since version: 6.0.0.0
// Required fields: typeId, global, pattern, start, createdAt, name
type NumberRange struct {
	CreatedAt                *time.Time               `json:"createdAt,omitempty"`
	CustomFields             *[]CustomField           `json:"customFields,omitempty"`
	Description              string                   `json:"description,omitempty"`
	Global                   bool                     `json:"global,omitempty"`
	Id                       string                   `json:"id,omitempty"`
	Name                     string                   `json:"name,omitempty"`
	NumberRangeSalesChannels *NumberRangeSalesChannel `json:"numberRangeSalesChannels,omitempty"`
	Pattern                  string                   `json:"pattern,omitempty"`
	Start                    int                      `json:"start,omitempty"`
	State                    *NumberRangeState        `json:"state,omitempty"`
	Translated               *interface{}             `json:"translated,omitempty"` // map[type:object]
	Type                     *NumberRangeType         `json:"type,omitempty"`
	TypeId                   string                   `json:"typeId,omitempty"`
	UpdatedAt                *time.Time               `json:"updatedAt,omitempty"`
}

// NumberRangeSalesChannel data structure
// Added since version: 6.0.0.0
// Required fields: numberRangeId, salesChannelId, createdAt
type NumberRangeSalesChannel struct {
	CreatedAt         *time.Time       `json:"createdAt,omitempty"`
	Id                string           `json:"id,omitempty"`
	NumberRange       *NumberRange     `json:"numberRange,omitempty"`
	NumberRangeId     string           `json:"numberRangeId,omitempty"`
	NumberRangeType   *NumberRangeType `json:"numberRangeType,omitempty"`
	NumberRangeTypeId string           `json:"numberRangeTypeId,omitempty"`
	SalesChannel      *SalesChannel    `json:"salesChannel,omitempty"`
	SalesChannelId    string           `json:"salesChannelId,omitempty"`
	UpdatedAt         *time.Time       `json:"updatedAt,omitempty"`
}

// NumberRangeState data structure
// Added since version: 6.0.0.0
// Required fields: numberRangeId, lastValue, createdAt
type NumberRangeState struct {
	CreatedAt     *time.Time   `json:"createdAt,omitempty"`
	Id            string       `json:"id,omitempty"`
	LastValue     int          `json:"lastValue,omitempty"`
	NumberRange   *NumberRange `json:"numberRange,omitempty"`
	NumberRangeId string       `json:"numberRangeId,omitempty"`
	UpdatedAt     *time.Time   `json:"updatedAt,omitempty"`
}

// NumberRangeType data structure
// Added since version: 6.0.0.0
// Required fields: global, createdAt, typeName
type NumberRangeType struct {
	CreatedAt                *time.Time               `json:"createdAt,omitempty"`
	CustomFields             *[]CustomField           `json:"customFields,omitempty"`
	Global                   bool                     `json:"global,omitempty"`
	Id                       string                   `json:"id,omitempty"`
	NumberRangeSalesChannels *NumberRangeSalesChannel `json:"numberRangeSalesChannels,omitempty"`
	NumberRanges             *NumberRange             `json:"numberRanges,omitempty"`
	TechnicalName            string                   `json:"technicalName,omitempty"`
	Translated               *interface{}             `json:"translated,omitempty"` // map[type:object]
	TypeName                 string                   `json:"typeName,omitempty"`
	UpdatedAt                *time.Time               `json:"updatedAt,omitempty"`
}
