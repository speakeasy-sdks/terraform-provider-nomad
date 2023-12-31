// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type JobRevertRequest struct {
	ConsulToken         *string `json:"ConsulToken,omitempty"`
	EnforcePriorVersion *int64  `json:"EnforcePriorVersion,omitempty"`
	JobID               *string `json:"JobID,omitempty"`
	JobVersion          *int64  `json:"JobVersion,omitempty"`
	Namespace           *string `json:"Namespace,omitempty"`
	Region              *string `json:"Region,omitempty"`
	SecretID            *string `json:"SecretID,omitempty"`
	VaultToken          *string `json:"VaultToken,omitempty"`
}

func (o *JobRevertRequest) GetConsulToken() *string {
	if o == nil {
		return nil
	}
	return o.ConsulToken
}

func (o *JobRevertRequest) GetEnforcePriorVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.EnforcePriorVersion
}

func (o *JobRevertRequest) GetJobID() *string {
	if o == nil {
		return nil
	}
	return o.JobID
}

func (o *JobRevertRequest) GetJobVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.JobVersion
}

func (o *JobRevertRequest) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *JobRevertRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *JobRevertRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}

func (o *JobRevertRequest) GetVaultToken() *string {
	if o == nil {
		return nil
	}
	return o.VaultToken
}
