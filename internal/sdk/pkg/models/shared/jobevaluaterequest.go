// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type JobEvaluateRequest struct {
	EvalOptions *EvalOptions `json:"EvalOptions,omitempty"`
	JobID       *string      `json:"JobID,omitempty"`
	Namespace   *string      `json:"Namespace,omitempty"`
	Region      *string      `json:"Region,omitempty"`
	SecretID    *string      `json:"SecretID,omitempty"`
}

func (o *JobEvaluateRequest) GetEvalOptions() *EvalOptions {
	if o == nil {
		return nil
	}
	return o.EvalOptions
}

func (o *JobEvaluateRequest) GetJobID() *string {
	if o == nil {
		return nil
	}
	return o.JobID
}

func (o *JobEvaluateRequest) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *JobEvaluateRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *JobEvaluateRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}
