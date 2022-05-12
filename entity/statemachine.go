package entity

import "time"

// completed

type StateMachine struct {
	CreatedAt      time.Time               `json:"createdAt,omitempty"`
	CustomFields   *[]CustomField          `json:"customFields,omitempty"`
	HistoryEntries *StateMachineHistory    `json:"historyEntries,omitempty"`
	Id             string                  `json:"id,omitempty"`
	InitialStateId string                  `json:"initialStateId,omitempty"`
	Name           string                  `json:"name,omitempty"`
	States         *StateMachineState      `json:"states,omitempty"`
	TechnicalName  string                  `json:"technicalName,omitempty"`
	Transitions    *StateMachineTransition `json:"transitions,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type StateMachineHistory struct {
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// EntityId map[type:object] `json:"entityId,omitempty"`
	EntityName            string             `json:"entityName,omitempty"`
	FromStateId           string             `json:"fromStateId,omitempty"`
	FromStateMachineState *StateMachineState `json:"fromStateMachineState,omitempty"`
	Id                    string             `json:"id,omitempty"`
	StateMachine          *StateMachine      `json:"stateMachine,omitempty"`
	StateMachineId        string             `json:"stateMachineId,omitempty"`
	ToStateId             string             `json:"toStateId,omitempty"`
	ToStateMachineState   *StateMachineState `json:"toStateMachineState,omitempty"`
	TransitionActionName  string             `json:"transitionActionName,omitempty"`
	UpdatedAt             time.Time          `json:"updatedAt,omitempty"`
	User                  *User              `json:"user,omitempty"`
	UserId                string             `json:"userId,omitempty"`
}

type StateMachineState struct {
	CreatedAt                      time.Time               `json:"createdAt,omitempty"`
	CustomFields                   *[]CustomField          `json:"customFields,omitempty"`
	FromStateMachineHistoryEntries *StateMachineHistory    `json:"fromStateMachineHistoryEntries,omitempty"`
	FromStateMachineTransitions    *StateMachineTransition `json:"fromStateMachineTransitions,omitempty"`
	Id                             string                  `json:"id,omitempty"`
	Name                           string                  `json:"name,omitempty"`
	OrderDeliveries                *OrderDelivery          `json:"orderDeliveries,omitempty"`
	OrderTransactions              *OrderTransaction       `json:"orderTransactions,omitempty"`
	Orders                         *Order                  `json:"orders,omitempty"`
	StateMachine                   *StateMachine           `json:"stateMachine,omitempty"`
	StateMachineId                 string                  `json:"stateMachineId,omitempty"`
	TechnicalName                  string                  `json:"technicalName,omitempty"`
	ToStateMachineHistoryEntries   *StateMachineHistory    `json:"toStateMachineHistoryEntries,omitempty"`
	ToStateMachineTransitions      *StateMachineTransition `json:"toStateMachineTransitions,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type StateMachineTransition struct {
	ActionName            string             `json:"actionName,omitempty"`
	CreatedAt             time.Time          `json:"createdAt,omitempty"`
	CustomFields          *[]CustomField     `json:"customFields,omitempty"`
	FromStateId           string             `json:"fromStateId,omitempty"`
	FromStateMachineState *StateMachineState `json:"fromStateMachineState,omitempty"`
	Id                    string             `json:"id,omitempty"`
	StateMachine          *StateMachine      `json:"stateMachine,omitempty"`
	StateMachineId        string             `json:"stateMachineId,omitempty"`
	ToStateId             string             `json:"toStateId,omitempty"`
	ToStateMachineState   *StateMachineState `json:"toStateMachineState,omitempty"`
	UpdatedAt             time.Time          `json:"updatedAt,omitempty"`
}
