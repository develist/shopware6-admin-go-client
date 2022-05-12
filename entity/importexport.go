package entity

import "time"

// completed

type ImportExportFile struct {
	AccessToken  string           `json:"accessToken,omitempty"`
	CreatedAt    time.Time        `json:"createdAt,omitempty"`
	ExpireDate   time.Time        `json:"expireDate,omitempty"`
	Id           string           `json:"id,omitempty"`
	Log          *ImportExportLog `json:"log,omitempty"`
	OriginalName string           `json:"originalName,omitempty"`
	Path         string           `json:"path,omitempty"`
	Size         int              `json:"size,omitempty"`
	UpdatedAt    time.Time        `json:"updatedAt,omitempty"`
}

type ImportExportLog struct {
	Activity string `json:"activity,omitempty"`
	// Config map[type:object] `json:"config,omitempty"`
	CreatedAt           time.Time            `json:"createdAt,omitempty"`
	FailedImportLog     *ImportExportLog     `json:"failedImportLog,omitempty"`
	File                *ImportExportFile    `json:"file,omitempty"`
	FileId              string               `json:"fileId,omitempty"`
	Id                  string               `json:"id,omitempty"`
	InvalidRecordsLog   *ImportExportLog     `json:"invalidRecordsLog,omitempty"`
	InvalidRecordsLogId string               `json:"invalidRecordsLogId,omitempty"`
	Profile             *ImportExportProfile `json:"profile,omitempty"`
	ProfileId           string               `json:"profileId,omitempty"`
	ProfileName         string               `json:"profileName,omitempty"`
	Records             int                  `json:"records,omitempty"`
	State               string               `json:"state,omitempty"`
	UpdatedAt           time.Time            `json:"updatedAt,omitempty"`
	User                *User                `json:"user,omitempty"`
	UserId              string               `json:"userId,omitempty"`
	Username            string               `json:"username,omitempty"`
}

type ImportExportProfile struct {
	// Config map[type:object] `json:"config,omitempty"`
	CreatedAt        time.Time        `json:"createdAt,omitempty"`
	Delimiter        string           `json:"delimiter,omitempty"`
	Enclosure        string           `json:"enclosure,omitempty"`
	FileType         string           `json:"fileType,omitempty"`
	Id               string           `json:"id,omitempty"`
	ImportExportLogs *ImportExportLog `json:"importExportLogs,omitempty"`
	Label            string           `json:"label,omitempty"`
	// Mapping map[type:object] `json:"mapping,omitempty"`
	Name          string `json:"name,omitempty"`
	SourceEntity  string `json:"sourceEntity,omitempty"`
	SystemDefault bool   `json:"systemDefault,omitempty"`
	// Translated map[type:object] `json:"translated,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
