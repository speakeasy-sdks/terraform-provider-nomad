// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/hashicorp/terraform-provider-nomad/v5/internal/sdk/internal/utils"
	"time"
)

type NodeEvent struct {
	CreateIndex *int64            `json:"CreateIndex,omitempty"`
	Details     map[string]string `json:"Details,omitempty"`
	Message     *string           `json:"Message,omitempty"`
	Subsystem   *string           `json:"Subsystem,omitempty"`
	Timestamp   *time.Time        `json:"Timestamp,omitempty"`
}

func (n NodeEvent) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(n, "", false)
}

func (n *NodeEvent) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &n, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *NodeEvent) GetCreateIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.CreateIndex
}

func (o *NodeEvent) GetDetails() map[string]string {
	if o == nil {
		return nil
	}
	return o.Details
}

func (o *NodeEvent) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *NodeEvent) GetSubsystem() *string {
	if o == nil {
		return nil
	}
	return o.Subsystem
}

func (o *NodeEvent) GetTimestamp() *time.Time {
	if o == nil {
		return nil
	}
	return o.Timestamp
}
