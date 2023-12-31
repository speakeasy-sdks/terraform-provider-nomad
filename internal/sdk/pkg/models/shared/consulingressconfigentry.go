// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ConsulIngressConfigEntry struct {
	Listeners []ConsulIngressListener `json:"Listeners,omitempty"`
	TLS       *ConsulGatewayTLSConfig `json:"TLS,omitempty"`
}

func (o *ConsulIngressConfigEntry) GetListeners() []ConsulIngressListener {
	if o == nil {
		return nil
	}
	return o.Listeners
}

func (o *ConsulIngressConfigEntry) GetTLS() *ConsulGatewayTLSConfig {
	if o == nil {
		return nil
	}
	return o.TLS
}
