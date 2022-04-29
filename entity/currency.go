package entity

type Currency struct {
	Factor          float64
	Id              string
	IsSystemDefault bool
	IsoCode         string
	Name            string
	Position        int
	ShortName       string
	Symbol          string
}
