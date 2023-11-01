// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CSIVolumeListExternalResponse struct {
	NextToken *string                 `json:"NextToken,omitempty"`
	Volumes   []CSIVolumeExternalStub `json:"Volumes,omitempty"`
}

func (o *CSIVolumeListExternalResponse) GetNextToken() *string {
	if o == nil {
		return nil
	}
	return o.NextToken
}

func (o *CSIVolumeListExternalResponse) GetVolumes() []CSIVolumeExternalStub {
	if o == nil {
		return nil
	}
	return o.Volumes
}
