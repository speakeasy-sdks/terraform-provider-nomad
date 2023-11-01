// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CSIVolumeCapability struct {
	AccessMode     *string `json:"AccessMode,omitempty"`
	AttachmentMode *string `json:"AttachmentMode,omitempty"`
}

func (o *CSIVolumeCapability) GetAccessMode() *string {
	if o == nil {
		return nil
	}
	return o.AccessMode
}

func (o *CSIVolumeCapability) GetAttachmentMode() *string {
	if o == nil {
		return nil
	}
	return o.AttachmentMode
}
