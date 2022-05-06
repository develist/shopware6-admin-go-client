package entity

type User struct {
	Active bool   `json:"active,omitempty"`
	Admin  bool   `json:"admin,omitempty"`
	Email  string `json:"email,omitempty"`
}
