// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DeploymentUpdateResponse struct {
	DeploymentModifyIndex *int64  `json:"DeploymentModifyIndex,omitempty"`
	EvalCreateIndex       *int64  `json:"EvalCreateIndex,omitempty"`
	EvalID                *string `json:"EvalID,omitempty"`
	LastIndex             *int64  `json:"LastIndex,omitempty"`
	RequestTime           *int64  `json:"RequestTime,omitempty"`
	RevertedJobVersion    *int64  `json:"RevertedJobVersion,omitempty"`
}
