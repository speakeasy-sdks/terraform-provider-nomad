// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type HostVolumeInfo struct {
	Path     *string `json:"Path,omitempty"`
	ReadOnly *bool   `json:"ReadOnly,omitempty"`
}

func (o *HostVolumeInfo) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *HostVolumeInfo) GetReadOnly() *bool {
	if o == nil {
		return nil
	}
	return o.ReadOnly
}
