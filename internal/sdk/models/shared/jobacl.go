// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type JobACL struct {
	Group     *string `json:"Group,omitempty"`
	JobID     *string `json:"JobID,omitempty"`
	Namespace *string `json:"Namespace,omitempty"`
	Task      *string `json:"Task,omitempty"`
}

func (o *JobACL) GetGroup() *string {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *JobACL) GetJobID() *string {
	if o == nil {
		return nil
	}
	return o.JobID
}

func (o *JobACL) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *JobACL) GetTask() *string {
	if o == nil {
		return nil
	}
	return o.Task
}