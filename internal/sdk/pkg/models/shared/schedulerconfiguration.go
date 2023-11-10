// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SchedulerConfiguration struct {
	CreateIndex                   *int64            `json:"CreateIndex,omitempty"`
	MemoryOversubscriptionEnabled *bool             `json:"MemoryOversubscriptionEnabled,omitempty"`
	ModifyIndex                   *int64            `json:"ModifyIndex,omitempty"`
	PauseEvalBroker               *bool             `json:"PauseEvalBroker,omitempty"`
	PreemptionConfig              *PreemptionConfig `json:"PreemptionConfig,omitempty"`
	RejectJobRegistration         *bool             `json:"RejectJobRegistration,omitempty"`
	SchedulerAlgorithm            *string           `json:"SchedulerAlgorithm,omitempty"`
}

func (o *SchedulerConfiguration) GetCreateIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.CreateIndex
}

func (o *SchedulerConfiguration) GetMemoryOversubscriptionEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.MemoryOversubscriptionEnabled
}

func (o *SchedulerConfiguration) GetModifyIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.ModifyIndex
}

func (o *SchedulerConfiguration) GetPauseEvalBroker() *bool {
	if o == nil {
		return nil
	}
	return o.PauseEvalBroker
}

func (o *SchedulerConfiguration) GetPreemptionConfig() *PreemptionConfig {
	if o == nil {
		return nil
	}
	return o.PreemptionConfig
}

func (o *SchedulerConfiguration) GetRejectJobRegistration() *bool {
	if o == nil {
		return nil
	}
	return o.RejectJobRegistration
}

func (o *SchedulerConfiguration) GetSchedulerAlgorithm() *string {
	if o == nil {
		return nil
	}
	return o.SchedulerAlgorithm
}
