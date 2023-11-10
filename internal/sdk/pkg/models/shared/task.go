// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Task struct {
	Affinities      []Affinity             `json:"Affinities,omitempty"`
	Artifacts       []TaskArtifact         `json:"Artifacts,omitempty"`
	CSIPluginConfig *TaskCSIPluginConfig   `json:"CSIPluginConfig,omitempty"`
	Config          map[string]interface{} `json:"Config,omitempty"`
	Constraints     []Constraint           `json:"Constraints,omitempty"`
	DispatchPayload *DispatchPayloadConfig `json:"DispatchPayload,omitempty"`
	Driver          *string                `json:"Driver,omitempty"`
	Env             map[string]string      `json:"Env,omitempty"`
	KillSignal      *string                `json:"KillSignal,omitempty"`
	KillTimeout     *int64                 `json:"KillTimeout,omitempty"`
	Kind            *string                `json:"Kind,omitempty"`
	Leader          *bool                  `json:"Leader,omitempty"`
	Lifecycle       *TaskLifecycle         `json:"Lifecycle,omitempty"`
	LogConfig       *LogConfig             `json:"LogConfig,omitempty"`
	Meta            map[string]string      `json:"Meta,omitempty"`
	Name            *string                `json:"Name,omitempty"`
	Resources       *Resources             `json:"Resources,omitempty"`
	RestartPolicy   *RestartPolicy         `json:"RestartPolicy,omitempty"`
	ScalingPolicies []ScalingPolicy        `json:"ScalingPolicies,omitempty"`
	Services        []Service              `json:"Services,omitempty"`
	ShutdownDelay   *int64                 `json:"ShutdownDelay,omitempty"`
	Templates       []Template             `json:"Templates,omitempty"`
	User            *string                `json:"User,omitempty"`
	Vault           *Vault                 `json:"Vault,omitempty"`
	VolumeMounts    []VolumeMount          `json:"VolumeMounts,omitempty"`
}

func (o *Task) GetAffinities() []Affinity {
	if o == nil {
		return nil
	}
	return o.Affinities
}

func (o *Task) GetArtifacts() []TaskArtifact {
	if o == nil {
		return nil
	}
	return o.Artifacts
}

func (o *Task) GetCSIPluginConfig() *TaskCSIPluginConfig {
	if o == nil {
		return nil
	}
	return o.CSIPluginConfig
}

func (o *Task) GetConfig() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *Task) GetConstraints() []Constraint {
	if o == nil {
		return nil
	}
	return o.Constraints
}

func (o *Task) GetDispatchPayload() *DispatchPayloadConfig {
	if o == nil {
		return nil
	}
	return o.DispatchPayload
}

func (o *Task) GetDriver() *string {
	if o == nil {
		return nil
	}
	return o.Driver
}

func (o *Task) GetEnv() map[string]string {
	if o == nil {
		return nil
	}
	return o.Env
}

func (o *Task) GetKillSignal() *string {
	if o == nil {
		return nil
	}
	return o.KillSignal
}

func (o *Task) GetKillTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.KillTimeout
}

func (o *Task) GetKind() *string {
	if o == nil {
		return nil
	}
	return o.Kind
}

func (o *Task) GetLeader() *bool {
	if o == nil {
		return nil
	}
	return o.Leader
}

func (o *Task) GetLifecycle() *TaskLifecycle {
	if o == nil {
		return nil
	}
	return o.Lifecycle
}

func (o *Task) GetLogConfig() *LogConfig {
	if o == nil {
		return nil
	}
	return o.LogConfig
}

func (o *Task) GetMeta() map[string]string {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *Task) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Task) GetResources() *Resources {
	if o == nil {
		return nil
	}
	return o.Resources
}

func (o *Task) GetRestartPolicy() *RestartPolicy {
	if o == nil {
		return nil
	}
	return o.RestartPolicy
}

func (o *Task) GetScalingPolicies() []ScalingPolicy {
	if o == nil {
		return nil
	}
	return o.ScalingPolicies
}

func (o *Task) GetServices() []Service {
	if o == nil {
		return nil
	}
	return o.Services
}

func (o *Task) GetShutdownDelay() *int64 {
	if o == nil {
		return nil
	}
	return o.ShutdownDelay
}

func (o *Task) GetTemplates() []Template {
	if o == nil {
		return nil
	}
	return o.Templates
}

func (o *Task) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *Task) GetVault() *Vault {
	if o == nil {
		return nil
	}
	return o.Vault
}

func (o *Task) GetVolumeMounts() []VolumeMount {
	if o == nil {
		return nil
	}
	return o.VolumeMounts
}
