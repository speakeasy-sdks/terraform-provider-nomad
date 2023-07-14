// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type JobPlanResponse struct {
	Annotations        *PlanAnnotations            `json:"Annotations,omitempty"`
	CreatedEvals       []Evaluation                `json:"CreatedEvals,omitempty"`
	Diff               *JobDiff                    `json:"Diff,omitempty"`
	FailedTGAllocs     map[string]AllocationMetric `json:"FailedTGAllocs,omitempty"`
	JobModifyIndex     *int64                      `json:"JobModifyIndex,omitempty"`
	NextPeriodicLaunch *time.Time                  `json:"NextPeriodicLaunch,omitempty"`
	Warnings           *string                     `json:"Warnings,omitempty"`
}
