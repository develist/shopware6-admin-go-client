package entity

import "time"

type StateMachineState struct {
	Name              string
	TechnicalName     string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	OrderTransactions OrderTransactions
}

type OrderTransactions struct {
	Id                string
	StateMachineState *StateMachineState
}
