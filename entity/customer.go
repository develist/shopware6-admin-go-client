package entity

type Customer struct {
	Active        bool   `json:"active,omitempty"`
	AutoIncrement int    `json:"autoIncrement,omitempty"`
	Email         string `json:"email,omitempty"`
	FirstName     string `json:"firstName,omitempty"`
	Id            string `json:"id,omitempty"`
	LastName      string `json:"lastName,omitempty"`
}
