// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type NodeDeviceResource struct {
	Attributes map[string]Attribute `json:"Attributes,omitempty"`
	Instances  []NodeDevice         `json:"Instances,omitempty"`
	Name       *string              `json:"Name,omitempty"`
	Type       *string              `json:"Type,omitempty"`
	Vendor     *string              `json:"Vendor,omitempty"`
}

func (o *NodeDeviceResource) GetAttributes() map[string]Attribute {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *NodeDeviceResource) GetInstances() []NodeDevice {
	if o == nil {
		return nil
	}
	return o.Instances
}

func (o *NodeDeviceResource) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *NodeDeviceResource) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *NodeDeviceResource) GetVendor() *string {
	if o == nil {
		return nil
	}
	return o.Vendor
}
