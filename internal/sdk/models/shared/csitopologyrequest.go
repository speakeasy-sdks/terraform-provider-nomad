// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CSITopologyRequest struct {
	Preferred []CSITopology `json:"Preferred,omitempty"`
	Required  []CSITopology `json:"Required,omitempty"`
}

func (o *CSITopologyRequest) GetPreferred() []CSITopology {
	if o == nil {
		return nil
	}
	return o.Preferred
}

func (o *CSITopologyRequest) GetRequired() []CSITopology {
	if o == nil {
		return nil
	}
	return o.Required
}