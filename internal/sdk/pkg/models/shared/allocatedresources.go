// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AllocatedResources struct {
	Shared *AllocatedSharedResources         `json:"Shared,omitempty"`
	Tasks  map[string]AllocatedTaskResources `json:"Tasks,omitempty"`
}