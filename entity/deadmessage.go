package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-17 19:07:53 UTC
// DEPRECATED ???

// DeadMessage data structure
// Added since version: 6.0.0.0
// Required fields: originalMessageClass, handlerClass, encrypted, errorCount, nextExecutionTime, exception, exceptionMessage, exceptionFile, exceptionLine, createdAt
type DeadMessage struct {
	CreatedAt            *time.Time     `json:"createdAt,omitempty"`
	Encrypted            bool           `json:"encrypted,omitempty"`
	ErrorCount           int            `json:"errorCount,omitempty"`
	Exception            string         `json:"exception,omitempty"`
	ExceptionFile        string         `json:"exceptionFile,omitempty"`
	ExceptionLine        int            `json:"exceptionLine,omitempty"`
	ExceptionMessage     string         `json:"exceptionMessage,omitempty"`
	HandlerClass         string         `json:"handlerClass,omitempty"`
	Id                   string         `json:"id,omitempty"`
	NextExecutionTime    *time.Time     `json:"nextExecutionTime,omitempty"`
	OriginalMessageClass string         `json:"originalMessageClass,omitempty"`
	ScheduledTask        *ScheduledTask `json:"scheduledTask,omitempty"`
	ScheduledTaskId      *string        `json:"scheduledTaskId,omitempty"`
	UpdatedAt            *time.Time     `json:"updatedAt,omitempty"`
}
