package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-06 18:44:04 UTC

// ScheduledTask data structure
// Added since version: 6.0.0.0
// Required fields: name, scheduledTaskClass, runInterval, status, nextExecutionTime, createdAt
type ScheduledTask struct {
	CreatedAt          *time.Time   `json:"createdAt,omitempty"`
	DeadMessages       *DeadMessage `json:"deadMessages,omitempty"`
	Id                 *string      `json:"id,omitempty"`
	LastExecutionTime  *time.Time   `json:"lastExecutionTime,omitempty"`
	Name               *string      `json:"name,omitempty"`
	NextExecutionTime  *time.Time   `json:"nextExecutionTime,omitempty"`
	RunInterval        *int         `json:"runInterval,omitempty"`
	ScheduledTaskClass *string      `json:"scheduledTaskClass,omitempty"`
	Status             *string      `json:"status,omitempty"`
	UpdatedAt          *time.Time   `json:"updatedAt,omitempty"`
}
