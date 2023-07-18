// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Deployment struct {
	CreateIndex        *int64                     `json:"CreateIndex,omitempty"`
	ID                 *string                    `json:"ID,omitempty"`
	IsMultiregion      *bool                      `json:"IsMultiregion,omitempty"`
	JobCreateIndex     *int64                     `json:"JobCreateIndex,omitempty"`
	JobID              *string                    `json:"JobID,omitempty"`
	JobModifyIndex     *int64                     `json:"JobModifyIndex,omitempty"`
	JobSpecModifyIndex *int64                     `json:"JobSpecModifyIndex,omitempty"`
	JobVersion         *int64                     `json:"JobVersion,omitempty"`
	ModifyIndex        *int64                     `json:"ModifyIndex,omitempty"`
	Namespace          *string                    `json:"Namespace,omitempty"`
	Status             *string                    `json:"Status,omitempty"`
	StatusDescription  *string                    `json:"StatusDescription,omitempty"`
	TaskGroups         map[string]DeploymentState `json:"TaskGroups,omitempty"`
}