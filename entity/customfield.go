package entity

type CustomField struct {
	Active bool   `json:"active,omitempty"`
	Id     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Type   string `json:"type,omitempty"`
}
