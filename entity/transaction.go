package entity

import "time"

type Transaction struct {
	Id                string
	VersionId         string
	OrderId           string
	PaymentMethodId   string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	PaymentMethod     PaymentMethod
	StateMachineState StateMachineState
}

type PaymentMethod struct {
	PluginId            string
	Name                string
	DistinguishableName string
	Description         string
}
