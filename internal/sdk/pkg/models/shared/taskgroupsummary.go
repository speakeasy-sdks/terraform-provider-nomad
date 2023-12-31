// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TaskGroupSummary struct {
	Complete *int64 `json:"Complete,omitempty"`
	Failed   *int64 `json:"Failed,omitempty"`
	Lost     *int64 `json:"Lost,omitempty"`
	Queued   *int64 `json:"Queued,omitempty"`
	Running  *int64 `json:"Running,omitempty"`
	Starting *int64 `json:"Starting,omitempty"`
	Unknown  *int64 `json:"Unknown,omitempty"`
}

func (o *TaskGroupSummary) GetComplete() *int64 {
	if o == nil {
		return nil
	}
	return o.Complete
}

func (o *TaskGroupSummary) GetFailed() *int64 {
	if o == nil {
		return nil
	}
	return o.Failed
}

func (o *TaskGroupSummary) GetLost() *int64 {
	if o == nil {
		return nil
	}
	return o.Lost
}

func (o *TaskGroupSummary) GetQueued() *int64 {
	if o == nil {
		return nil
	}
	return o.Queued
}

func (o *TaskGroupSummary) GetRunning() *int64 {
	if o == nil {
		return nil
	}
	return o.Running
}

func (o *TaskGroupSummary) GetStarting() *int64 {
	if o == nil {
		return nil
	}
	return o.Starting
}

func (o *TaskGroupSummary) GetUnknown() *int64 {
	if o == nil {
		return nil
	}
	return o.Unknown
}
