// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type UpdateStrategy struct {
	AutoPromote      *bool   `json:"AutoPromote,omitempty"`
	AutoRevert       *bool   `json:"AutoRevert,omitempty"`
	Canary           *int64  `json:"Canary,omitempty"`
	HealthCheck      *string `json:"HealthCheck,omitempty"`
	HealthyDeadline  *int64  `json:"HealthyDeadline,omitempty"`
	MaxParallel      *int64  `json:"MaxParallel,omitempty"`
	MinHealthyTime   *int64  `json:"MinHealthyTime,omitempty"`
	ProgressDeadline *int64  `json:"ProgressDeadline,omitempty"`
	Stagger          *int64  `json:"Stagger,omitempty"`
}