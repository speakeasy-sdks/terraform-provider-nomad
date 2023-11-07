// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"nomad/v2/internal/sdk/pkg/utils"
	"time"
)

type TaskState struct {
	Events      []TaskEvent `json:"Events,omitempty"`
	Failed      *bool       `json:"Failed,omitempty"`
	FinishedAt  *time.Time  `json:"FinishedAt,omitempty"`
	LastRestart *time.Time  `json:"LastRestart,omitempty"`
	Restarts    *int64      `json:"Restarts,omitempty"`
	StartedAt   *time.Time  `json:"StartedAt,omitempty"`
	State       *string     `json:"State,omitempty"`
	TaskHandle  *TaskHandle `json:"TaskHandle,omitempty"`
}

func (t TaskState) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaskState) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TaskState) GetEvents() []TaskEvent {
	if o == nil {
		return nil
	}
	return o.Events
}

func (o *TaskState) GetFailed() *bool {
	if o == nil {
		return nil
	}
	return o.Failed
}

func (o *TaskState) GetFinishedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.FinishedAt
}

func (o *TaskState) GetLastRestart() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastRestart
}

func (o *TaskState) GetRestarts() *int64 {
	if o == nil {
		return nil
	}
	return o.Restarts
}

func (o *TaskState) GetStartedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartedAt
}

func (o *TaskState) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *TaskState) GetTaskHandle() *TaskHandle {
	if o == nil {
		return nil
	}
	return o.TaskHandle
}
