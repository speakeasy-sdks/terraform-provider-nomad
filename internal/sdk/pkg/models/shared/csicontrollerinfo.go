// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CSIControllerInfo struct {
	SupportsAttachDetach             *bool `json:"SupportsAttachDetach,omitempty"`
	SupportsClone                    *bool `json:"SupportsClone,omitempty"`
	SupportsCondition                *bool `json:"SupportsCondition,omitempty"`
	SupportsCreateDelete             *bool `json:"SupportsCreateDelete,omitempty"`
	SupportsCreateDeleteSnapshot     *bool `json:"SupportsCreateDeleteSnapshot,omitempty"`
	SupportsExpand                   *bool `json:"SupportsExpand,omitempty"`
	SupportsGet                      *bool `json:"SupportsGet,omitempty"`
	SupportsGetCapacity              *bool `json:"SupportsGetCapacity,omitempty"`
	SupportsListSnapshots            *bool `json:"SupportsListSnapshots,omitempty"`
	SupportsListVolumes              *bool `json:"SupportsListVolumes,omitempty"`
	SupportsListVolumesAttachedNodes *bool `json:"SupportsListVolumesAttachedNodes,omitempty"`
	SupportsReadOnlyAttach           *bool `json:"SupportsReadOnlyAttach,omitempty"`
}