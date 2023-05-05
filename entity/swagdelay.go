package entity

import "time"

// Generated from Shopware Admin API
// Version 6.5.9999999.9999999-dev at 2023-05-02 19:06:34 UTC

// SwagDelayAction data structure
// Added since version:
// Required fields: flowId, executionTime, stored, createdAt
type SwagDelayAction struct {
	CreatedAt       *time.Time    `json:"createdAt,omitempty"`
	Customer        *Customer     `json:"customer,omitempty"`
	CustomerId      *string       `json:"customerId,omitempty"`
	DelaySequenceId *string       `json:"delaySequenceId,omitempty"`
	EventName       *string       `json:"eventName,omitempty"`
	ExecutionTime   *time.Time    `json:"executionTime,omitempty"`
	Expired         *bool         `json:"expired,omitempty"`
	Flow            *Flow         `json:"flow,omitempty"`
	FlowId          string        `json:"flowId,omitempty"`
	Id              string        `json:"id,omitempty"`
	Order           *Order        `json:"order,omitempty"`
	OrderId         *string       `json:"orderId,omitempty"`
	OrderVersionId  *string       `json:"orderVersionId,omitempty"`
	Sequence        *FlowSequence `json:"sequence,omitempty"`
	Stored          *interface{}  `json:"stored,omitempty"` // map[type:object]
	UpdatedAt       *time.Time    `json:"updatedAt,omitempty"`
}
