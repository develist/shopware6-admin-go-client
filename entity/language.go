package entity

import "time"

type Language struct {
	CreatedAt time.Time
	Id        string
	LocaleId  string
	Name      string
}
