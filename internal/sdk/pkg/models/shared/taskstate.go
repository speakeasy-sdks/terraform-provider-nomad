// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
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
