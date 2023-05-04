package entity

import "time"

// Generated from Shopware Admin API
// Version 6.5.9999999.9999999-dev at 2023-05-02 19:06:34 UTC

// Flow data structure
// Added since version: 6.4.6.0
// Required fields: name, eventName, createdAt
type Flow struct {
	Active       *bool          `json:"active,omitempty"`
	CreatedAt    *time.Time     `json:"createdAt,omitempty"`
	CustomFields *[]CustomField `json:"customFields,omitempty"`
	Description  *string        `json:"description,omitempty"`
	EventName    string         `json:"eventName,omitempty"`
	Extensions   *interface{}   `json:"extensions,omitempty"` // map[properties:map[delayActions:map[properties:map[data:map[items:map[properties:map[id:map[example:afaebe0004084c57a4173df11c9bf5e8 type:string] type:map[example:swag_delay_action type:string]] type:object] type:array] links:map[properties:map[related:map[example:/flow/05a775fc095645b6b014238d0448b564/delayActions format:uri-reference type:string]] type:object]] type:object]] type:object]
	Id           string         `json:"id,omitempty"`
	Invalid      *bool          `json:"invalid,omitempty"`
	Name         string         `json:"name,omitempty"`
	Priority     *int           `json:"priority,omitempty"`
	Sequences    *FlowSequence  `json:"sequences,omitempty"`
	UpdatedAt    *time.Time     `json:"updatedAt,omitempty"`
}

// FlowSequence data structure
// Added since version: 6.4.6.0
// Required fields: flowId, createdAt
type FlowSequence struct {
	ActionName      *string        `json:"actionName,omitempty"`
	AppFlowAction   *AppFlowAction `json:"appFlowAction,omitempty"`
	AppFlowActionId *string        `json:"appFlowActionId,omitempty"`
	Children        *FlowSequence  `json:"children,omitempty"`
	Config          *interface{}   `json:"config,omitempty"` // map[type:object]
	CreatedAt       *time.Time     `json:"createdAt,omitempty"`
	CustomFields    *[]CustomField `json:"customFields,omitempty"`
	DisplayGroup    *int           `json:"displayGroup,omitempty"`
	Flow            *Flow          `json:"flow,omitempty"`
	FlowId          string         `json:"flowId,omitempty"`
	Id              string         `json:"id,omitempty"`
	Parent          *FlowSequence  `json:"parent,omitempty"`
	ParentId        *string        `json:"parentId,omitempty"`
	Position        *int           `json:"position,omitempty"`
	Rule            *Rule          `json:"rule,omitempty"`
	RuleId          *string        `json:"ruleId,omitempty"`
	TrueCase        *bool          `json:"trueCase,omitempty"`
	UpdatedAt       *time.Time     `json:"updatedAt,omitempty"`
}

// FlowTemplate data structure
// Added since version: 6.4.18.0
// Required fields: name, createdAt
type FlowTemplate struct {
	Config    *interface{} `json:"config,omitempty"` // map[type:object]
	CreatedAt *time.Time   `json:"createdAt,omitempty"`
	Id        string       `json:"id,omitempty"`
	Name      string       `json:"name,omitempty"`
	UpdatedAt *time.Time   `json:"updatedAt,omitempty"`
}
