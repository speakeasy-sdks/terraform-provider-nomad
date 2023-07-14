// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TaskEvent struct {
	Details          map[string]string `json:"Details,omitempty"`
	DiskLimit        *int64            `json:"DiskLimit,omitempty"`
	DiskSize         *int64            `json:"DiskSize,omitempty"`
	DisplayMessage   *string           `json:"DisplayMessage,omitempty"`
	DownloadError    *string           `json:"DownloadError,omitempty"`
	DriverError      *string           `json:"DriverError,omitempty"`
	DriverMessage    *string           `json:"DriverMessage,omitempty"`
	ExitCode         *int64            `json:"ExitCode,omitempty"`
	FailedSibling    *string           `json:"FailedSibling,omitempty"`
	FailsTask        *bool             `json:"FailsTask,omitempty"`
	GenericSource    *string           `json:"GenericSource,omitempty"`
	KillError        *string           `json:"KillError,omitempty"`
	KillReason       *string           `json:"KillReason,omitempty"`
	KillTimeout      *int64            `json:"KillTimeout,omitempty"`
	Message          *string           `json:"Message,omitempty"`
	RestartReason    *string           `json:"RestartReason,omitempty"`
	SetupError       *string           `json:"SetupError,omitempty"`
	Signal           *int64            `json:"Signal,omitempty"`
	StartDelay       *int64            `json:"StartDelay,omitempty"`
	TaskSignal       *string           `json:"TaskSignal,omitempty"`
	TaskSignalReason *string           `json:"TaskSignalReason,omitempty"`
	Time             *int64            `json:"Time,omitempty"`
	Type             *string           `json:"Type,omitempty"`
	ValidationError  *string           `json:"ValidationError,omitempty"`
	VaultError       *string           `json:"VaultError,omitempty"`
}
