// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ConsulConnect struct {
	Gateway        *ConsulGateway        `json:"Gateway,omitempty"`
	Native         *bool                 `json:"Native,omitempty"`
	SidecarService *ConsulSidecarService `json:"SidecarService,omitempty"`
	SidecarTask    *SidecarTask          `json:"SidecarTask,omitempty"`
}

func (o *ConsulConnect) GetGateway() *ConsulGateway {
	if o == nil {
		return nil
	}
	return o.Gateway
}

func (o *ConsulConnect) GetNative() *bool {
	if o == nil {
		return nil
	}
	return o.Native
}

func (o *ConsulConnect) GetSidecarService() *ConsulSidecarService {
	if o == nil {
		return nil
	}
	return o.SidecarService
}

func (o *ConsulConnect) GetSidecarTask() *SidecarTask {
	if o == nil {
		return nil
	}
	return o.SidecarTask
}
