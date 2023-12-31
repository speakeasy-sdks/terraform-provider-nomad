// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type MultiregionStrategy struct {
	MaxParallel *int64  `json:"MaxParallel,omitempty"`
	OnFailure   *string `json:"OnFailure,omitempty"`
}

func (o *MultiregionStrategy) GetMaxParallel() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxParallel
}

func (o *MultiregionStrategy) GetOnFailure() *string {
	if o == nil {
		return nil
	}
	return o.OnFailure
}
