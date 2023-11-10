// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CSIVolumeRegisterRequest struct {
	Namespace *string     `json:"Namespace,omitempty"`
	Region    *string     `json:"Region,omitempty"`
	SecretID  *string     `json:"SecretID,omitempty"`
	Volumes   []CSIVolume `json:"Volumes,omitempty"`
}

func (o *CSIVolumeRegisterRequest) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *CSIVolumeRegisterRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *CSIVolumeRegisterRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}

func (o *CSIVolumeRegisterRequest) GetVolumes() []CSIVolume {
	if o == nil {
		return nil
	}
	return o.Volumes
}
