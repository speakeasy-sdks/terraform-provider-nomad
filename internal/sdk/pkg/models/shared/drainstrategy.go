// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type DrainStrategy struct {
	Deadline         *int64     `json:"Deadline,omitempty"`
	ForceDeadline    *time.Time `json:"ForceDeadline,omitempty"`
	IgnoreSystemJobs *bool      `json:"IgnoreSystemJobs,omitempty"`
	StartedAt        *time.Time `json:"StartedAt,omitempty"`
}
