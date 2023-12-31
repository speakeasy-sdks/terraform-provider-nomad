// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Port struct {
	HostNetwork *string `json:"HostNetwork,omitempty"`
	Label       *string `json:"Label,omitempty"`
	To          *int64  `json:"To,omitempty"`
	Value       *int64  `json:"Value,omitempty"`
}

func (o *Port) GetHostNetwork() *string {
	if o == nil {
		return nil
	}
	return o.HostNetwork
}

func (o *Port) GetLabel() *string {
	if o == nil {
		return nil
	}
	return o.Label
}

func (o *Port) GetTo() *int64 {
	if o == nil {
		return nil
	}
	return o.To
}

func (o *Port) GetValue() *int64 {
	if o == nil {
		return nil
	}
	return o.Value
}
