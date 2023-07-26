// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type CSIVolumeListStub struct {
	AccessMode          *string       `json:"AccessMode,omitempty"`
	AttachmentMode      *string       `json:"AttachmentMode,omitempty"`
	ControllerRequired  *bool         `json:"ControllerRequired,omitempty"`
	ControllersExpected *int64        `json:"ControllersExpected,omitempty"`
	ControllersHealthy  *int64        `json:"ControllersHealthy,omitempty"`
	CreateIndex         *int64        `json:"CreateIndex,omitempty"`
	CurrentReaders      *int64        `json:"CurrentReaders,omitempty"`
	CurrentWriters      *int64        `json:"CurrentWriters,omitempty"`
	ExternalID          *string       `json:"ExternalID,omitempty"`
	ID                  *string       `json:"ID,omitempty"`
	ModifyIndex         *int64        `json:"ModifyIndex,omitempty"`
	Name                *string       `json:"Name,omitempty"`
	Namespace           *string       `json:"Namespace,omitempty"`
	NodesExpected       *int64        `json:"NodesExpected,omitempty"`
	NodesHealthy        *int64        `json:"NodesHealthy,omitempty"`
	PluginID            *string       `json:"PluginID,omitempty"`
	Provider            *string       `json:"Provider,omitempty"`
	ResourceExhausted   *time.Time    `json:"ResourceExhausted,omitempty"`
	Schedulable         *bool         `json:"Schedulable,omitempty"`
	Topologies          []CSITopology `json:"Topologies,omitempty"`
}
