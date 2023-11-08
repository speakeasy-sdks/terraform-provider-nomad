// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/hashicorp/terraform-provider-nomad/v2/internal/sdk/pkg/utils"
	"time"
)

type CSIInfo struct {
	AllocID                  *string            `json:"AllocID,omitempty"`
	ControllerInfo           *CSIControllerInfo `json:"ControllerInfo,omitempty"`
	HealthDescription        *string            `json:"HealthDescription,omitempty"`
	Healthy                  *bool              `json:"Healthy,omitempty"`
	NodeInfo                 *CSINodeInfo       `json:"NodeInfo,omitempty"`
	PluginID                 *string            `json:"PluginID,omitempty"`
	RequiresControllerPlugin *bool              `json:"RequiresControllerPlugin,omitempty"`
	RequiresTopologies       *bool              `json:"RequiresTopologies,omitempty"`
	UpdateTime               *time.Time         `json:"UpdateTime,omitempty"`
}

func (c CSIInfo) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CSIInfo) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CSIInfo) GetAllocID() *string {
	if o == nil {
		return nil
	}
	return o.AllocID
}

func (o *CSIInfo) GetControllerInfo() *CSIControllerInfo {
	if o == nil {
		return nil
	}
	return o.ControllerInfo
}

func (o *CSIInfo) GetHealthDescription() *string {
	if o == nil {
		return nil
	}
	return o.HealthDescription
}

func (o *CSIInfo) GetHealthy() *bool {
	if o == nil {
		return nil
	}
	return o.Healthy
}

func (o *CSIInfo) GetNodeInfo() *CSINodeInfo {
	if o == nil {
		return nil
	}
	return o.NodeInfo
}

func (o *CSIInfo) GetPluginID() *string {
	if o == nil {
		return nil
	}
	return o.PluginID
}

func (o *CSIInfo) GetRequiresControllerPlugin() *bool {
	if o == nil {
		return nil
	}
	return o.RequiresControllerPlugin
}

func (o *CSIInfo) GetRequiresTopologies() *bool {
	if o == nil {
		return nil
	}
	return o.RequiresTopologies
}

func (o *CSIInfo) GetUpdateTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdateTime
}
