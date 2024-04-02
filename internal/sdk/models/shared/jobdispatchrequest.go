// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type JobDispatchRequest struct {
	JobID   *string           `json:"JobID,omitempty"`
	Meta    map[string]string `json:"Meta,omitempty"`
	Payload *string           `json:"Payload,omitempty"`
}

func (o *JobDispatchRequest) GetJobID() *string {
	if o == nil {
		return nil
	}
	return o.JobID
}

func (o *JobDispatchRequest) GetMeta() map[string]string {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *JobDispatchRequest) GetPayload() *string {
	if o == nil {
		return nil
	}
	return o.Payload
}