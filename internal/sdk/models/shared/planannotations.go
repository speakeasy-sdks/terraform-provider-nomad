// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PlanAnnotations struct {
	DesiredTGUpdates map[string]DesiredUpdates `json:"DesiredTGUpdates,omitempty"`
	PreemptedAllocs  []AllocationListStub      `json:"PreemptedAllocs,omitempty"`
}

func (o *PlanAnnotations) GetDesiredTGUpdates() map[string]DesiredUpdates {
	if o == nil {
		return nil
	}
	return o.DesiredTGUpdates
}

func (o *PlanAnnotations) GetPreemptedAllocs() []AllocationListStub {
	if o == nil {
		return nil
	}
	return o.PreemptedAllocs
}