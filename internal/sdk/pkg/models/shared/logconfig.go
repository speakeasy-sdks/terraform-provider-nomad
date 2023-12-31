// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type LogConfig struct {
	MaxFileSizeMB *int64 `json:"MaxFileSizeMB,omitempty"`
	MaxFiles      *int64 `json:"MaxFiles,omitempty"`
}

func (o *LogConfig) GetMaxFileSizeMB() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxFileSizeMB
}

func (o *LogConfig) GetMaxFiles() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxFiles
}
