// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ConsulSidecarService struct {
	DisableDefaultTCPCheck *bool        `json:"DisableDefaultTCPCheck,omitempty"`
	Port                   *string      `json:"Port,omitempty"`
	Proxy                  *ConsulProxy `json:"Proxy,omitempty"`
	Tags                   []string     `json:"Tags,omitempty"`
}

func (o *ConsulSidecarService) GetDisableDefaultTCPCheck() *bool {
	if o == nil {
		return nil
	}
	return o.DisableDefaultTCPCheck
}

func (o *ConsulSidecarService) GetPort() *string {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *ConsulSidecarService) GetProxy() *ConsulProxy {
	if o == nil {
		return nil
	}
	return o.Proxy
}

func (o *ConsulSidecarService) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}