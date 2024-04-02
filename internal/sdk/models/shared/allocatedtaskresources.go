// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AllocatedTaskResources struct {
	CPU      *AllocatedCPUResources    `json:"Cpu,omitempty"`
	Devices  []AllocatedDeviceResource `json:"Devices,omitempty"`
	Memory   *AllocatedMemoryResources `json:"Memory,omitempty"`
	Networks []NetworkResource         `json:"Networks,omitempty"`
}

func (o *AllocatedTaskResources) GetCPU() *AllocatedCPUResources {
	if o == nil {
		return nil
	}
	return o.CPU
}

func (o *AllocatedTaskResources) GetDevices() []AllocatedDeviceResource {
	if o == nil {
		return nil
	}
	return o.Devices
}

func (o *AllocatedTaskResources) GetMemory() *AllocatedMemoryResources {
	if o == nil {
		return nil
	}
	return o.Memory
}

func (o *AllocatedTaskResources) GetNetworks() []NetworkResource {
	if o == nil {
		return nil
	}
	return o.Networks
}