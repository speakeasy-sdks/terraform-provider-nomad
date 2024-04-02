// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SidecarTask struct {
	Config        map[string]interface{} `json:"Config,omitempty"`
	Driver        *string                `json:"Driver,omitempty"`
	Env           map[string]string      `json:"Env,omitempty"`
	KillSignal    *string                `json:"KillSignal,omitempty"`
	KillTimeout   *int64                 `json:"KillTimeout,omitempty"`
	LogConfig     *LogConfig             `json:"LogConfig,omitempty"`
	Meta          map[string]string      `json:"Meta,omitempty"`
	Name          *string                `json:"Name,omitempty"`
	Resources     *Resources             `json:"Resources,omitempty"`
	ShutdownDelay *int64                 `json:"ShutdownDelay,omitempty"`
	User          *string                `json:"User,omitempty"`
}

func (o *SidecarTask) GetConfig() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *SidecarTask) GetDriver() *string {
	if o == nil {
		return nil
	}
	return o.Driver
}

func (o *SidecarTask) GetEnv() map[string]string {
	if o == nil {
		return nil
	}
	return o.Env
}

func (o *SidecarTask) GetKillSignal() *string {
	if o == nil {
		return nil
	}
	return o.KillSignal
}

func (o *SidecarTask) GetKillTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.KillTimeout
}

func (o *SidecarTask) GetLogConfig() *LogConfig {
	if o == nil {
		return nil
	}
	return o.LogConfig
}

func (o *SidecarTask) GetMeta() map[string]string {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *SidecarTask) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *SidecarTask) GetResources() *Resources {
	if o == nil {
		return nil
	}
	return o.Resources
}

func (o *SidecarTask) GetShutdownDelay() *int64 {
	if o == nil {
		return nil
	}
	return o.ShutdownDelay
}

func (o *SidecarTask) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}