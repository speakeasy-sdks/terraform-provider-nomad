// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Allocation struct {
	AllocModifyIndex      *int64                 `json:"AllocModifyIndex,omitempty"`
	AllocatedResources    *AllocatedResources    `json:"AllocatedResources,omitempty"`
	ClientDescription     *string                `json:"ClientDescription,omitempty"`
	ClientStatus          *string                `json:"ClientStatus,omitempty"`
	CreateIndex           *int64                 `json:"CreateIndex,omitempty"`
	CreateTime            *int64                 `json:"CreateTime,omitempty"`
	DeploymentID          *string                `json:"DeploymentID,omitempty"`
	DeploymentStatus      *AllocDeploymentStatus `json:"DeploymentStatus,omitempty"`
	DesiredDescription    *string                `json:"DesiredDescription,omitempty"`
	DesiredStatus         *string                `json:"DesiredStatus,omitempty"`
	DesiredTransition     *DesiredTransition     `json:"DesiredTransition,omitempty"`
	EvalID                *string                `json:"EvalID,omitempty"`
	FollowupEvalID        *string                `json:"FollowupEvalID,omitempty"`
	ID                    *string                `json:"ID,omitempty"`
	Job                   *Job                   `json:"Job,omitempty"`
	JobID                 *string                `json:"JobID,omitempty"`
	Metrics               *AllocationMetric      `json:"Metrics,omitempty"`
	ModifyIndex           *int64                 `json:"ModifyIndex,omitempty"`
	ModifyTime            *int64                 `json:"ModifyTime,omitempty"`
	Name                  *string                `json:"Name,omitempty"`
	Namespace             *string                `json:"Namespace,omitempty"`
	NextAllocation        *string                `json:"NextAllocation,omitempty"`
	NodeID                *string                `json:"NodeID,omitempty"`
	NodeName              *string                `json:"NodeName,omitempty"`
	PreemptedAllocations  []string               `json:"PreemptedAllocations,omitempty"`
	PreemptedByAllocation *string                `json:"PreemptedByAllocation,omitempty"`
	PreviousAllocation    *string                `json:"PreviousAllocation,omitempty"`
	RescheduleTracker     *RescheduleTracker     `json:"RescheduleTracker,omitempty"`
	Resources             *Resources             `json:"Resources,omitempty"`
	Services              map[string]string      `json:"Services,omitempty"`
	TaskGroup             *string                `json:"TaskGroup,omitempty"`
	TaskResources         map[string]Resources   `json:"TaskResources,omitempty"`
	TaskStates            map[string]TaskState   `json:"TaskStates,omitempty"`
}

func (o *Allocation) GetAllocModifyIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.AllocModifyIndex
}

func (o *Allocation) GetAllocatedResources() *AllocatedResources {
	if o == nil {
		return nil
	}
	return o.AllocatedResources
}

func (o *Allocation) GetClientDescription() *string {
	if o == nil {
		return nil
	}
	return o.ClientDescription
}

func (o *Allocation) GetClientStatus() *string {
	if o == nil {
		return nil
	}
	return o.ClientStatus
}

func (o *Allocation) GetCreateIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.CreateIndex
}

func (o *Allocation) GetCreateTime() *int64 {
	if o == nil {
		return nil
	}
	return o.CreateTime
}

func (o *Allocation) GetDeploymentID() *string {
	if o == nil {
		return nil
	}
	return o.DeploymentID
}

func (o *Allocation) GetDeploymentStatus() *AllocDeploymentStatus {
	if o == nil {
		return nil
	}
	return o.DeploymentStatus
}

func (o *Allocation) GetDesiredDescription() *string {
	if o == nil {
		return nil
	}
	return o.DesiredDescription
}

func (o *Allocation) GetDesiredStatus() *string {
	if o == nil {
		return nil
	}
	return o.DesiredStatus
}

func (o *Allocation) GetDesiredTransition() *DesiredTransition {
	if o == nil {
		return nil
	}
	return o.DesiredTransition
}

func (o *Allocation) GetEvalID() *string {
	if o == nil {
		return nil
	}
	return o.EvalID
}

func (o *Allocation) GetFollowupEvalID() *string {
	if o == nil {
		return nil
	}
	return o.FollowupEvalID
}

func (o *Allocation) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Allocation) GetJob() *Job {
	if o == nil {
		return nil
	}
	return o.Job
}

func (o *Allocation) GetJobID() *string {
	if o == nil {
		return nil
	}
	return o.JobID
}

func (o *Allocation) GetMetrics() *AllocationMetric {
	if o == nil {
		return nil
	}
	return o.Metrics
}

func (o *Allocation) GetModifyIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.ModifyIndex
}

func (o *Allocation) GetModifyTime() *int64 {
	if o == nil {
		return nil
	}
	return o.ModifyTime
}

func (o *Allocation) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Allocation) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *Allocation) GetNextAllocation() *string {
	if o == nil {
		return nil
	}
	return o.NextAllocation
}

func (o *Allocation) GetNodeID() *string {
	if o == nil {
		return nil
	}
	return o.NodeID
}

func (o *Allocation) GetNodeName() *string {
	if o == nil {
		return nil
	}
	return o.NodeName
}

func (o *Allocation) GetPreemptedAllocations() []string {
	if o == nil {
		return nil
	}
	return o.PreemptedAllocations
}

func (o *Allocation) GetPreemptedByAllocation() *string {
	if o == nil {
		return nil
	}
	return o.PreemptedByAllocation
}

func (o *Allocation) GetPreviousAllocation() *string {
	if o == nil {
		return nil
	}
	return o.PreviousAllocation
}

func (o *Allocation) GetRescheduleTracker() *RescheduleTracker {
	if o == nil {
		return nil
	}
	return o.RescheduleTracker
}

func (o *Allocation) GetResources() *Resources {
	if o == nil {
		return nil
	}
	return o.Resources
}

func (o *Allocation) GetServices() map[string]string {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *Allocation) GetTaskGroup() *string {
	if o == nil {
		return nil
	}
	return o.TaskGroup
}

func (o *Allocation) GetTaskResources() map[string]Resources {
	if o == nil {
		return nil
	}
	return o.TaskResources
}

func (o *Allocation) GetTaskStates() map[string]TaskState {
	if o == nil {
		return nil
	}
	return o.TaskStates
}
