// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DispatchPayloadConfig struct {
	File *string `json:"File,omitempty"`
}

func (o *DispatchPayloadConfig) GetFile() *string {
	if o == nil {
		return nil
	}
	return o.File
}
