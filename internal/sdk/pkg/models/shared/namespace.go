// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Namespace struct {
	Capabilities *NamespaceCapabilities `json:"Capabilities,omitempty"`
	CreateIndex  *int64                 `json:"CreateIndex,omitempty"`
	Description  *string                `json:"Description,omitempty"`
	Meta         map[string]string      `json:"Meta,omitempty"`
	ModifyIndex  *int64                 `json:"ModifyIndex,omitempty"`
	Name         *string                `json:"Name,omitempty"`
	Quota        *string                `json:"Quota,omitempty"`
}
