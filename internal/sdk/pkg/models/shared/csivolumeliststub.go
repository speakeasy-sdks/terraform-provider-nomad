// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"nomad/internal/sdk/pkg/utils"
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

func (c CSIVolumeListStub) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CSIVolumeListStub) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CSIVolumeListStub) GetAccessMode() *string {
	if o == nil {
		return nil
	}
	return o.AccessMode
}

func (o *CSIVolumeListStub) GetAttachmentMode() *string {
	if o == nil {
		return nil
	}
	return o.AttachmentMode
}

func (o *CSIVolumeListStub) GetControllerRequired() *bool {
	if o == nil {
		return nil
	}
	return o.ControllerRequired
}

func (o *CSIVolumeListStub) GetControllersExpected() *int64 {
	if o == nil {
		return nil
	}
	return o.ControllersExpected
}

func (o *CSIVolumeListStub) GetControllersHealthy() *int64 {
	if o == nil {
		return nil
	}
	return o.ControllersHealthy
}

func (o *CSIVolumeListStub) GetCreateIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.CreateIndex
}

func (o *CSIVolumeListStub) GetCurrentReaders() *int64 {
	if o == nil {
		return nil
	}
	return o.CurrentReaders
}

func (o *CSIVolumeListStub) GetCurrentWriters() *int64 {
	if o == nil {
		return nil
	}
	return o.CurrentWriters
}

func (o *CSIVolumeListStub) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *CSIVolumeListStub) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CSIVolumeListStub) GetModifyIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.ModifyIndex
}

func (o *CSIVolumeListStub) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CSIVolumeListStub) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *CSIVolumeListStub) GetNodesExpected() *int64 {
	if o == nil {
		return nil
	}
	return o.NodesExpected
}

func (o *CSIVolumeListStub) GetNodesHealthy() *int64 {
	if o == nil {
		return nil
	}
	return o.NodesHealthy
}

func (o *CSIVolumeListStub) GetPluginID() *string {
	if o == nil {
		return nil
	}
	return o.PluginID
}

func (o *CSIVolumeListStub) GetProvider() *string {
	if o == nil {
		return nil
	}
	return o.Provider
}

func (o *CSIVolumeListStub) GetResourceExhausted() *time.Time {
	if o == nil {
		return nil
	}
	return o.ResourceExhausted
}

func (o *CSIVolumeListStub) GetSchedulable() *bool {
	if o == nil {
		return nil
	}
	return o.Schedulable
}

func (o *CSIVolumeListStub) GetTopologies() []CSITopology {
	if o == nil {
		return nil
	}
	return o.Topologies
}
