// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type JobPlanRequest struct {
	Diff           *bool   `json:"Diff,omitempty"`
	Job            *Job    `json:"Job,omitempty"`
	Namespace      *string `json:"Namespace,omitempty"`
	PolicyOverride *bool   `json:"PolicyOverride,omitempty"`
	Region         *string `json:"Region,omitempty"`
	SecretID       *string `json:"SecretID,omitempty"`
}

func (o *JobPlanRequest) GetDiff() *bool {
	if o == nil {
		return nil
	}
	return o.Diff
}

func (o *JobPlanRequest) GetJob() *Job {
	if o == nil {
		return nil
	}
	return o.Job
}

func (o *JobPlanRequest) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *JobPlanRequest) GetPolicyOverride() *bool {
	if o == nil {
		return nil
	}
	return o.PolicyOverride
}

func (o *JobPlanRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *JobPlanRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}