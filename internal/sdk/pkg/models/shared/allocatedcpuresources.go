// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AllocatedCPUResources struct {
	CPUShares *int64 `json:"CpuShares,omitempty"`
}

func (o *AllocatedCPUResources) GetCPUShares() *int64 {
	if o == nil {
		return nil
	}
	return o.CPUShares
}
