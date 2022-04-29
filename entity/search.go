package entity

import "github.com/develist/shopware6-admin-go-client/field"

type Criteria struct {
	Page         int                      `json:"page,omitempty"`
	Limit        int                      `json:"limit,omitempty"`
	Filter       []FilterCriteria         `json:"filter,omitempty"`
	Sort         []SortCriteria           `json:"sort,omitempty"`
	Associations map[field.Field]Criteria `json:"associations,omitempty"`
}

type FilterCriteria struct {
	Type       FilterType                 `json:"type"`
	Field      field.Field                `json:"field"`
	Parameters map[FilterParameter]string `json:"parameters"`
}

type FilterType string

const (
	Equals    FilterType = "equals"
	EqualsAny            = "equalsAny"
	Contains             = "contains"
	Range                = "range"
	Not                  = "not"
	Multi                = "multi"
	Prefix               = "prefix"
	Suffix               = "suffix"
)

type FilterParameter string

const (
	Gte FilterParameter = "gte"
	Lte                 = "lte"
	Gt                  = "gt"
	Lt                  = "lt"
)

type SortCriteria struct {
	Field          field.Field `json:"field"`
	Order          SortOrder   `json:"order"`
	NaturalSorting bool        `json:"naturalSorting"`
}

type SortOrder string

const (
	Asc  SortOrder = "ASC"
	Desc           = "DESC"
)

type ResultTypes[T any] interface {
	Result[T] | Results[T]
}

type Results[T any] struct {
	Total int `json:"total"`
	Data  []T `json:"data"`
}

type Result[T any] struct {
	Data T `json:"data"`
}
