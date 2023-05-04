package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-17 19:07:53 UTC
// DEPRECATED ???

// EventAction data structure
// Added since version: 6.0.0.0
// Required fields: eventName, actionName, createdAt
type EventAction struct {
	ActionName    string         `json:"actionName,omitempty"`
	Active        *bool          `json:"active,omitempty"`
	Config        *interface{}   `json:"config,omitempty"` // map[type:object]
	CreatedAt     *time.Time     `json:"createdAt,omitempty"`
	CustomFields  *[]CustomField `json:"customFields,omitempty"`
	EventName     string         `json:"eventName,omitempty"`
	Id            string         `json:"id,omitempty"`
	Rules         *Rule          `json:"rules,omitempty"`
	SalesChannels *SalesChannel  `json:"salesChannels,omitempty"`
	Title         *string        `json:"title,omitempty"`
	UpdatedAt     *time.Time     `json:"updatedAt,omitempty"`
}

// EventActionRule data structure
// Added since version: 6.3.3.0
// Required fields: eventActionId, ruleId
type EventActionRule struct {
	EventAction   *EventAction `json:"eventAction,omitempty"`
	EventActionId string       `json:"eventActionId,omitempty"`
	Id            string       `json:"id,omitempty"`
	Rule          *Rule        `json:"rule,omitempty"`
	RuleId        string       `json:"ruleId,omitempty"`
}

// EventActionSalesChannel data structure
// Added since version: 6.3.3.0
// Required fields: eventActionId, salesChannelId
type EventActionSalesChannel struct {
	EventAction    *EventAction  `json:"eventAction,omitempty"`
	EventActionId  string        `json:"eventActionId,omitempty"`
	Id             string        `json:"id,omitempty"`
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId string        `json:"salesChannelId,omitempty"`
}
