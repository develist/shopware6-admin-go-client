package entity

import "time"

// completed

type ScheduledTask struct {
	CreatedAt          *time.Time   `json:"createdAt,omitempty"`
	DeadMessages       *DeadMessage `json:"deadMessages,omitempty"`
	Id                 string       `json:"id,omitempty"`
	LastExecutionTime  *time.Time   `json:"lastExecutionTime,omitempty"`
	Name               string       `json:"name,omitempty"`
	NextExecutionTime  *time.Time   `json:"nextExecutionTime,omitempty"`
	RunInterval        int          `json:"runInterval,omitempty"`
	ScheduledTaskClass string       `json:"scheduledTaskClass,omitempty"`
	Status             string       `json:"status,omitempty"`
	UpdatedAt          *time.Time   `json:"updatedAt,omitempty"`
}
