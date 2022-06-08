package entity

import "time"

// Generated from Shopware Admin API
// Version 6.4.9999999.9999999-dev at 2022-06-06 18:44:04 UTC

// ImportExportFile data structure
// Added since version: 6.0.0.0
// Required fields: originalName, path, expireDate, accessToken, createdAt
type ImportExportFile struct {
	AccessToken  *string          `json:"accessToken,omitempty"`
	CreatedAt    *time.Time       `json:"createdAt,omitempty"`
	ExpireDate   *time.Time       `json:"expireDate,omitempty"`
	Id           *string          `json:"id,omitempty"`
	Log          *ImportExportLog `json:"log,omitempty"`
	OriginalName *string          `json:"originalName,omitempty"`
	Path         *string          `json:"path,omitempty"`
	Size         *int             `json:"size,omitempty"`
	UpdatedAt    *time.Time       `json:"updatedAt,omitempty"`
}

// ImportExportLog data structure
// Added since version: 6.0.0.0
// Required fields: activity, state, records, config, createdAt
type ImportExportLog struct {
	Activity            *string              `json:"activity,omitempty"`
	Config              *interface{}         `json:"config,omitempty"` // map[type:object]
	CreatedAt           *time.Time           `json:"createdAt,omitempty"`
	FailedImportLog     *ImportExportLog     `json:"failedImportLog,omitempty"`
	File                *ImportExportFile    `json:"file,omitempty"`
	FileId              *string              `json:"fileId,omitempty"`
	Id                  *string              `json:"id,omitempty"`
	InvalidRecordsLog   *ImportExportLog     `json:"invalidRecordsLog,omitempty"`
	InvalidRecordsLogId *string              `json:"invalidRecordsLogId,omitempty"`
	Profile             *ImportExportProfile `json:"profile,omitempty"`
	ProfileId           *string              `json:"profileId,omitempty"`
	ProfileName         *string              `json:"profileName,omitempty"`
	Records             *int                 `json:"records,omitempty"`
	State               *string              `json:"state,omitempty"`
	UpdatedAt           *time.Time           `json:"updatedAt,omitempty"`
	User                *User                `json:"user,omitempty"`
	UserId              *string              `json:"userId,omitempty"`
	Username            *string              `json:"username,omitempty"`
}

// ImportExportProfile data structure
// Added since version: 6.0.0.0
// Required fields: label, sourceEntity, fileType, delimiter, enclosure, createdAt
type ImportExportProfile struct {
	Config           *interface{}     `json:"config,omitempty"` // map[type:object]
	CreatedAt        *time.Time       `json:"createdAt,omitempty"`
	Delimiter        *string          `json:"delimiter,omitempty"`
	Enclosure        *string          `json:"enclosure,omitempty"`
	FileType         *string          `json:"fileType,omitempty"`
	Id               *string          `json:"id,omitempty"`
	ImportExportLogs *ImportExportLog `json:"importExportLogs,omitempty"`
	Label            *string          `json:"label,omitempty"`
	Mapping          *interface{}     `json:"mapping,omitempty"` // map[type:object]
	Name             *string          `json:"name,omitempty"`
	SourceEntity     *string          `json:"sourceEntity,omitempty"`
	SystemDefault    *bool            `json:"systemDefault,omitempty"`
	Translated       *interface{}     `json:"translated,omitempty"` // map[type:object]
	UpdatedAt        *time.Time       `json:"updatedAt,omitempty"`
}
