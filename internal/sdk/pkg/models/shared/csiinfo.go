// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
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
