// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type EvalOptions struct {
	ForceReschedule *bool `json:"ForceReschedule,omitempty"`
}

func (o *EvalOptions) GetForceReschedule() *bool {
	if o == nil {
		return nil
	}
	return o.ForceReschedule
}
