// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type JobListStub struct {
	CreateIndex       *int64      `json:"CreateIndex,omitempty"`
	Datacenters       []string    `json:"Datacenters,omitempty"`
	ID                *string     `json:"ID,omitempty"`
	JobModifyIndex    *int64      `json:"JobModifyIndex,omitempty"`
	JobSummary        *JobSummary `json:"JobSummary,omitempty"`
	ModifyIndex       *int64      `json:"ModifyIndex,omitempty"`
	Name              *string     `json:"Name,omitempty"`
	Namespace         *string     `json:"Namespace,omitempty"`
	ParameterizedJob  *bool       `json:"ParameterizedJob,omitempty"`
	ParentID          *string     `json:"ParentID,omitempty"`
	Periodic          *bool       `json:"Periodic,omitempty"`
	Priority          *int64      `json:"Priority,omitempty"`
	Status            *string     `json:"Status,omitempty"`
	StatusDescription *string     `json:"StatusDescription,omitempty"`
	Stop              *bool       `json:"Stop,omitempty"`
	SubmitTime        *int64      `json:"SubmitTime,omitempty"`
	Type              *string     `json:"Type,omitempty"`
}
