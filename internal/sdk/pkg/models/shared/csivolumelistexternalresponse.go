// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CSIVolumeListExternalResponse struct {
	NextToken *string                 `json:"NextToken,omitempty"`
	Volumes   []CSIVolumeExternalStub `json:"Volumes,omitempty"`
}