// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type RescheduleTracker struct {
	Events []RescheduleEvent `json:"Events,omitempty"`
}

func (o *RescheduleTracker) GetEvents() []RescheduleEvent {
	if o == nil {
		return nil
	}
	return o.Events
}
