// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TaskGroup struct {
	Affinities                []Affinity               `json:"Affinities,omitempty"`
	Constraints               []Constraint             `json:"Constraints,omitempty"`
	Consul                    *Consul                  `json:"Consul,omitempty"`
	Count                     *int64                   `json:"Count,omitempty"`
	EphemeralDisk             *EphemeralDisk           `json:"EphemeralDisk,omitempty"`
	MaxClientDisconnect       *int64                   `json:"MaxClientDisconnect,omitempty"`
	Meta                      map[string]string        `json:"Meta,omitempty"`
	Migrate                   *MigrateStrategy         `json:"Migrate,omitempty"`
	Name                      *string                  `json:"Name,omitempty"`
	Networks                  []NetworkResource        `json:"Networks,omitempty"`
	ReschedulePolicy          *ReschedulePolicy        `json:"ReschedulePolicy,omitempty"`
	RestartPolicy             *RestartPolicy           `json:"RestartPolicy,omitempty"`
	Scaling                   *ScalingPolicy           `json:"Scaling,omitempty"`
	Services                  []Service                `json:"Services,omitempty"`
	ShutdownDelay             *int64                   `json:"ShutdownDelay,omitempty"`
	Spreads                   []Spread                 `json:"Spreads,omitempty"`
	StopAfterClientDisconnect *int64                   `json:"StopAfterClientDisconnect,omitempty"`
	Tasks                     []Task                   `json:"Tasks,omitempty"`
	Update                    *UpdateStrategy          `json:"Update,omitempty"`
	Volumes                   map[string]VolumeRequest `json:"Volumes,omitempty"`
}

func (o *TaskGroup) GetAffinities() []Affinity {
	if o == nil {
		return nil
	}
	return o.Affinities
}

func (o *TaskGroup) GetConstraints() []Constraint {
	if o == nil {
		return nil
	}
	return o.Constraints
}

func (o *TaskGroup) GetConsul() *Consul {
	if o == nil {
		return nil
	}
	return o.Consul
}

func (o *TaskGroup) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *TaskGroup) GetEphemeralDisk() *EphemeralDisk {
	if o == nil {
		return nil
	}
	return o.EphemeralDisk
}

func (o *TaskGroup) GetMaxClientDisconnect() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxClientDisconnect
}

func (o *TaskGroup) GetMeta() map[string]string {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *TaskGroup) GetMigrate() *MigrateStrategy {
	if o == nil {
		return nil
	}
	return o.Migrate
}

func (o *TaskGroup) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TaskGroup) GetNetworks() []NetworkResource {
	if o == nil {
		return nil
	}
	return o.Networks
}

func (o *TaskGroup) GetReschedulePolicy() *ReschedulePolicy {
	if o == nil {
		return nil
	}
	return o.ReschedulePolicy
}

func (o *TaskGroup) GetRestartPolicy() *RestartPolicy {
	if o == nil {
		return nil
	}
	return o.RestartPolicy
}

func (o *TaskGroup) GetScaling() *ScalingPolicy {
	if o == nil {
		return nil
	}
	return o.Scaling
}

func (o *TaskGroup) GetServices() []Service {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *TaskGroup) GetShutdownDelay() *int64 {
	if o == nil {
		return nil
	}
	return o.ShutdownDelay
}

func (o *TaskGroup) GetSpreads() []Spread {
	if o == nil {
		return nil
	}
	return o.Spreads
}

func (o *TaskGroup) GetStopAfterClientDisconnect() *int64 {
	if o == nil {
		return nil
	}
	return o.StopAfterClientDisconnect
}

func (o *TaskGroup) GetTasks() []Task {
	if o == nil {
		return nil
	}
	return o.Tasks
}

func (o *TaskGroup) GetUpdate() *UpdateStrategy {
	if o == nil {
		return nil
	}
	return o.Update
}

func (o *TaskGroup) GetVolumes() map[string]VolumeRequest {
	if o == nil {
		return nil
	}
	return o.Volumes
}
