// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type NodeReservedMemoryResources struct {
	MemoryMB *int64 `json:"MemoryMB,omitempty"`
}

func (o *NodeReservedMemoryResources) GetMemoryMB() *int64 {
	if o == nil {
		return nil
	}
	return o.MemoryMB
}
