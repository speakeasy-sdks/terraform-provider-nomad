// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PortMapping struct {
	HostIP *string `json:"HostIP,omitempty"`
	Label  *string `json:"Label,omitempty"`
	To     *int64  `json:"To,omitempty"`
	Value  *int64  `json:"Value,omitempty"`
}

func (o *PortMapping) GetHostIP() *string {
	if o == nil {
		return nil
	}
	return o.HostIP
}

func (o *PortMapping) GetLabel() *string {
	if o == nil {
		return nil
	}
	return o.Label
}

func (o *PortMapping) GetTo() *int64 {
	if o == nil {
		return nil
	}
	return o.To
}

func (o *PortMapping) GetValue() *int64 {
	if o == nil {
		return nil
	}
	return o.Value
}
