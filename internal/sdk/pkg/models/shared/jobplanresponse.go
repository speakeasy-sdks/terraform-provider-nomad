// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"nomad/v2/internal/sdk/pkg/utils"
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

func (j JobPlanResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(j, "", false)
}

func (j *JobPlanResponse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &j, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *JobPlanResponse) GetAnnotations() *PlanAnnotations {
	if o == nil {
		return nil
	}
	return o.Annotations
}

func (o *JobPlanResponse) GetCreatedEvals() []Evaluation {
	if o == nil {
		return nil
	}
	return o.CreatedEvals
}

func (o *JobPlanResponse) GetDiff() *JobDiff {
	if o == nil {
		return nil
	}
	return o.Diff
}

func (o *JobPlanResponse) GetFailedTGAllocs() map[string]AllocationMetric {
	if o == nil {
		return nil
	}
	return o.FailedTGAllocs
}

func (o *JobPlanResponse) GetJobModifyIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.JobModifyIndex
}

func (o *JobPlanResponse) GetNextPeriodicLaunch() *time.Time {
	if o == nil {
		return nil
	}
	return o.NextPeriodicLaunch
}

func (o *JobPlanResponse) GetWarnings() *string {
	if o == nil {
		return nil
	}
	return o.Warnings
}
