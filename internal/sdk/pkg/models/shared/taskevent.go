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

func (o *TaskEvent) GetDetails() map[string]string {
	if o == nil {
		return nil
	}
	return o.Details
}

func (o *TaskEvent) GetDiskLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.DiskLimit
}

func (o *TaskEvent) GetDiskSize() *int64 {
	if o == nil {
		return nil
	}
	return o.DiskSize
}

func (o *TaskEvent) GetDisplayMessage() *string {
	if o == nil {
		return nil
	}
	return o.DisplayMessage
}

func (o *TaskEvent) GetDownloadError() *string {
	if o == nil {
		return nil
	}
	return o.DownloadError
}

func (o *TaskEvent) GetDriverError() *string {
	if o == nil {
		return nil
	}
	return o.DriverError
}

func (o *TaskEvent) GetDriverMessage() *string {
	if o == nil {
		return nil
	}
	return o.DriverMessage
}

func (o *TaskEvent) GetExitCode() *int64 {
	if o == nil {
		return nil
	}
	return o.ExitCode
}

func (o *TaskEvent) GetFailedSibling() *string {
	if o == nil {
		return nil
	}
	return o.FailedSibling
}

func (o *TaskEvent) GetFailsTask() *bool {
	if o == nil {
		return nil
	}
	return o.FailsTask
}

func (o *TaskEvent) GetGenericSource() *string {
	if o == nil {
		return nil
	}
	return o.GenericSource
}

func (o *TaskEvent) GetKillError() *string {
	if o == nil {
		return nil
	}
	return o.KillError
}

func (o *TaskEvent) GetKillReason() *string {
	if o == nil {
		return nil
	}
	return o.KillReason
}

func (o *TaskEvent) GetKillTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.KillTimeout
}

func (o *TaskEvent) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *TaskEvent) GetRestartReason() *string {
	if o == nil {
		return nil
	}
	return o.RestartReason
}

func (o *TaskEvent) GetSetupError() *string {
	if o == nil {
		return nil
	}
	return o.SetupError
}

func (o *TaskEvent) GetSignal() *int64 {
	if o == nil {
		return nil
	}
	return o.Signal
}

func (o *TaskEvent) GetStartDelay() *int64 {
	if o == nil {
		return nil
	}
	return o.StartDelay
}

func (o *TaskEvent) GetTaskSignal() *string {
	if o == nil {
		return nil
	}
	return o.TaskSignal
}

func (o *TaskEvent) GetTaskSignalReason() *string {
	if o == nil {
		return nil
	}
	return o.TaskSignalReason
}

func (o *TaskEvent) GetTime() *int64 {
	if o == nil {
		return nil
	}
	return o.Time
}

func (o *TaskEvent) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *TaskEvent) GetValidationError() *string {
	if o == nil {
		return nil
	}
	return o.ValidationError
}

func (o *TaskEvent) GetVaultError() *string {
	if o == nil {
		return nil
	}
	return o.VaultError
}
