package entity

import "time"

// completed

type NumberRange struct {
	CreatedAt                time.Time                `json:"createdAt,omitempty"`
	CustomFields             *[]CustomField           `json:"customFields,omitempty"`
	Description              string                   `json:"description,omitempty"`
	Global                   bool                     `json:"global,omitempty"`
	Id                       string                   `json:"id,omitempty"`
	Name                     string                   `json:"name,omitempty"`
	NumberRangeSalesChannels *NumberRangeSalesChannel `json:"numberRangeSalesChannels,omitempty"`
	Pattern                  string                   `json:"pattern,omitempty"`
	Start                    int                      `json:"start,omitempty"`
	State                    *NumberRangeState        `json:"state,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	Type      *NumberRangeType `json:"type,omitempty"`
	TypeId    string           `json:"typeId,omitempty"`
	UpdatedAt time.Time        `json:"updatedAt,omitempty"`
}

type NumberRangeSalesChannel struct {
	CreatedAt         time.Time        `json:"createdAt,omitempty"`
	Id                string           `json:"id,omitempty"`
	NumberRange       *NumberRange     `json:"numberRange,omitempty"`
	NumberRangeId     string           `json:"numberRangeId,omitempty"`
	NumberRangeType   *NumberRangeType `json:"numberRangeType,omitempty"`
	NumberRangeTypeId string           `json:"numberRangeTypeId,omitempty"`
	SalesChannel      *SalesChannel    `json:"salesChannel,omitempty"`
	SalesChannelId    string           `json:"salesChannelId,omitempty"`
	UpdatedAt         time.Time        `json:"updatedAt,omitempty"`
}

type NumberRangeState struct {
	CreatedAt     time.Time    `json:"createdAt,omitempty"`
	Id            string       `json:"id,omitempty"`
	LastValue     int          `json:"lastValue,omitempty"`
	NumberRange   *NumberRange `json:"numberRange,omitempty"`
	NumberRangeId string       `json:"numberRangeId,omitempty"`
	UpdatedAt     time.Time    `json:"updatedAt,omitempty"`
}

type NumberRangeType struct {
	CreatedAt                time.Time                `json:"createdAt,omitempty"`
	CustomFields             *[]CustomField           `json:"customFields,omitempty"`
	Global                   bool                     `json:"global,omitempty"`
	Id                       string                   `json:"id,omitempty"`
	NumberRangeSalesChannels *NumberRangeSalesChannel `json:"numberRangeSalesChannels,omitempty"`
	NumberRanges             *NumberRange             `json:"numberRanges,omitempty"`
	TechnicalName            string                   `json:"technicalName,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	TypeName  string    `json:"typeName,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
