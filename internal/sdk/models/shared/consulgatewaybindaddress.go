// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ConsulGatewayBindAddress struct {
	Address *string `json:"Address,omitempty"`
	Name    *string `json:"Name,omitempty"`
	Port    *int64  `json:"Port,omitempty"`
}

func (o *ConsulGatewayBindAddress) GetAddress() *string {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *ConsulGatewayBindAddress) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConsulGatewayBindAddress) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}