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

func (o *DeploymentUpdateResponse) GetDeploymentModifyIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.DeploymentModifyIndex
}

func (o *DeploymentUpdateResponse) GetEvalCreateIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.EvalCreateIndex
}

func (o *DeploymentUpdateResponse) GetEvalID() *string {
	if o == nil {
		return nil
	}
	return o.EvalID
}

func (o *DeploymentUpdateResponse) GetLastIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.LastIndex
}

func (o *DeploymentUpdateResponse) GetRequestTime() *int64 {
	if o == nil {
		return nil
	}
	return o.RequestTime
}

func (o *DeploymentUpdateResponse) GetRevertedJobVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.RevertedJobVersion
}