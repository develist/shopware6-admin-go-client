package entity

type CustomerAddress struct {
	AdditionalAddressLine1 string `json:"additionalAddressLine1,omitempty"`
	AdditionalAddressLine2 string `json:"additionalAddressLine2,omitempty"`
	City                   string `json:"city,omitempty"`
	Company                string `json:"company,omitempty"`
}

type Customer struct {
	Active        bool   `json:"active,omitempty"`
	AutoIncrement int    `json:"autoIncrement,omitempty"`
	Email         string `json:"email,omitempty"`
	FirstName     string `json:"firstName,omitempty"`
	Id            string `json:"id,omitempty"`
	LastName      string `json:"lastName,omitempty"`
}
