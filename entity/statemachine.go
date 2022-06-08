package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-06 18:44:04 UTC

// StateMachine data structure
// Added since version: 6.0.0.0
// Required fields: technicalName, createdAt, name
type StateMachine struct {
	CreatedAt      *time.Time              `json:"createdAt,omitempty"`
	CustomFields   *[]CustomField          `json:"customFields,omitempty"`
	HistoryEntries *StateMachineHistory    `json:"historyEntries,omitempty"`
	Id             *string                 `json:"id,omitempty"`
	InitialStateId *string                 `json:"initialStateId,omitempty"`
	Name           *string                 `json:"name,omitempty"`
	States         *StateMachineState      `json:"states,omitempty"`
	TechnicalName  *string                 `json:"technicalName,omitempty"`
	Transitions    *StateMachineTransition `json:"transitions,omitempty"`
	Translated     *interface{}            `json:"translated,omitempty"` // map[type:object]
	UpdatedAt      *time.Time              `json:"updatedAt,omitempty"`
}

// StateMachineHistory data structure
// Added since version: 6.0.0.0
// Required fields: stateMachineId, entityName, entityId, fromStateId, toStateId, createdAt
type StateMachineHistory struct {
	CreatedAt             *time.Time         `json:"createdAt,omitempty"`
	EntityId              *interface{}       `json:"entityId,omitempty"` // map[type:object]
	EntityName            *string            `json:"entityName,omitempty"`
	FromStateId           *string            `json:"fromStateId,omitempty"`
	FromStateMachineState *StateMachineState `json:"fromStateMachineState,omitempty"`
	Id                    *string            `json:"id,omitempty"`
	StateMachine          *StateMachine      `json:"stateMachine,omitempty"`
	StateMachineId        *string            `json:"stateMachineId,omitempty"`
	ToStateId             *string            `json:"toStateId,omitempty"`
	ToStateMachineState   *StateMachineState `json:"toStateMachineState,omitempty"`
	TransitionActionName  *string            `json:"transitionActionName,omitempty"`
	UpdatedAt             *time.Time         `json:"updatedAt,omitempty"`
	User                  *User              `json:"user,omitempty"`
	UserId                *string            `json:"userId,omitempty"`
}

// StateMachineState data structure
// Added since version: 6.0.0.0
// Required fields: technicalName, stateMachineId, createdAt, name
type StateMachineState struct {
	CreatedAt                      *time.Time              `json:"createdAt,omitempty"`
	CustomFields                   *[]CustomField          `json:"customFields,omitempty"`
	FromStateMachineHistoryEntries *StateMachineHistory    `json:"fromStateMachineHistoryEntries,omitempty"`
	FromStateMachineTransitions    *StateMachineTransition `json:"fromStateMachineTransitions,omitempty"`
	Id                             *string                 `json:"id,omitempty"`
	Name                           *string                 `json:"name,omitempty"`
	OrderDeliveries                *OrderDelivery          `json:"orderDeliveries,omitempty"`
	OrderTransactions              *OrderTransaction       `json:"orderTransactions,omitempty"`
	Orders                         *Order                  `json:"orders,omitempty"`
	StateMachine                   *StateMachine           `json:"stateMachine,omitempty"`
	StateMachineId                 *string                 `json:"stateMachineId,omitempty"`
	TechnicalName                  *string                 `json:"technicalName,omitempty"`
	ToStateMachineHistoryEntries   *StateMachineHistory    `json:"toStateMachineHistoryEntries,omitempty"`
	ToStateMachineTransitions      *StateMachineTransition `json:"toStateMachineTransitions,omitempty"`
	Translated                     *interface{}            `json:"translated,omitempty"` // map[type:object]
	UpdatedAt                      *time.Time              `json:"updatedAt,omitempty"`
}

// StateMachineTransition data structure
// Added since version: 6.0.0.0
// Required fields: actionName, stateMachineId, fromStateId, toStateId, createdAt
type StateMachineTransition struct {
	ActionName            *string            `json:"actionName,omitempty"`
	CreatedAt             *time.Time         `json:"createdAt,omitempty"`
	CustomFields          *[]CustomField     `json:"customFields,omitempty"`
	FromStateId           *string            `json:"fromStateId,omitempty"`
	FromStateMachineState *StateMachineState `json:"fromStateMachineState,omitempty"`
	Id                    *string            `json:"id,omitempty"`
	StateMachine          *StateMachine      `json:"stateMachine,omitempty"`
	StateMachineId        *string            `json:"stateMachineId,omitempty"`
	ToStateId             *string            `json:"toStateId,omitempty"`
	ToStateMachineState   *StateMachineState `json:"toStateMachineState,omitempty"`
	UpdatedAt             *time.Time         `json:"updatedAt,omitempty"`
}
