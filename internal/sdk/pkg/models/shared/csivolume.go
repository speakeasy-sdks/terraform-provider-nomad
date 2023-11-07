// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"nomad/v2/internal/sdk/pkg/utils"
	"time"
)

type CSIVolume struct {
	AccessMode            *string               `json:"AccessMode,omitempty"`
	Allocations           []AllocationListStub  `json:"Allocations,omitempty"`
	AttachmentMode        *string               `json:"AttachmentMode,omitempty"`
	Capacity              *int64                `json:"Capacity,omitempty"`
	CloneID               *string               `json:"CloneID,omitempty"`
	Context               map[string]string     `json:"Context,omitempty"`
	ControllerRequired    *bool                 `json:"ControllerRequired,omitempty"`
	ControllersExpected   *int64                `json:"ControllersExpected,omitempty"`
	ControllersHealthy    *int64                `json:"ControllersHealthy,omitempty"`
	CreateIndex           *int64                `json:"CreateIndex,omitempty"`
	ExternalID            *string               `json:"ExternalID,omitempty"`
	ID                    *string               `json:"ID,omitempty"`
	ModifyIndex           *int64                `json:"ModifyIndex,omitempty"`
	MountOptions          *CSIMountOptions      `json:"MountOptions,omitempty"`
	Name                  *string               `json:"Name,omitempty"`
	Namespace             *string               `json:"Namespace,omitempty"`
	NodesExpected         *int64                `json:"NodesExpected,omitempty"`
	NodesHealthy          *int64                `json:"NodesHealthy,omitempty"`
	Parameters            map[string]string     `json:"Parameters,omitempty"`
	PluginID              *string               `json:"PluginID,omitempty"`
	Provider              *string               `json:"Provider,omitempty"`
	ProviderVersion       *string               `json:"ProviderVersion,omitempty"`
	ReadAllocs            map[string]Allocation `json:"ReadAllocs,omitempty"`
	RequestedCapabilities []CSIVolumeCapability `json:"RequestedCapabilities,omitempty"`
	RequestedCapacityMax  *int64                `json:"RequestedCapacityMax,omitempty"`
	RequestedCapacityMin  *int64                `json:"RequestedCapacityMin,omitempty"`
	RequestedTopologies   *CSITopologyRequest   `json:"RequestedTopologies,omitempty"`
	ResourceExhausted     *time.Time            `json:"ResourceExhausted,omitempty"`
	Schedulable           *bool                 `json:"Schedulable,omitempty"`
	Secrets               map[string]string     `json:"Secrets,omitempty"`
	SnapshotID            *string               `json:"SnapshotID,omitempty"`
	Topologies            []CSITopology         `json:"Topologies,omitempty"`
	WriteAllocs           map[string]Allocation `json:"WriteAllocs,omitempty"`
}

func (c CSIVolume) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CSIVolume) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CSIVolume) GetAccessMode() *string {
	if o == nil {
		return nil
	}
	return o.AccessMode
}

func (o *CSIVolume) GetAllocations() []AllocationListStub {
	if o == nil {
		return nil
	}
	return o.Allocations
}

func (o *CSIVolume) GetAttachmentMode() *string {
	if o == nil {
		return nil
	}
	return o.AttachmentMode
}

func (o *CSIVolume) GetCapacity() *int64 {
	if o == nil {
		return nil
	}
	return o.Capacity
}

func (o *CSIVolume) GetCloneID() *string {
	if o == nil {
		return nil
	}
	return o.CloneID
}

func (o *CSIVolume) GetContext() map[string]string {
	if o == nil {
		return nil
	}
	return o.Context
}

func (o *CSIVolume) GetControllerRequired() *bool {
	if o == nil {
		return nil
	}
	return o.ControllerRequired
}

func (o *CSIVolume) GetControllersExpected() *int64 {
	if o == nil {
		return nil
	}
	return o.ControllersExpected
}

func (o *CSIVolume) GetControllersHealthy() *int64 {
	if o == nil {
		return nil
	}
	return o.ControllersHealthy
}

func (o *CSIVolume) GetCreateIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.CreateIndex
}

func (o *CSIVolume) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *CSIVolume) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CSIVolume) GetModifyIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.ModifyIndex
}

func (o *CSIVolume) GetMountOptions() *CSIMountOptions {
	if o == nil {
		return nil
	}
	return o.MountOptions
}

func (o *CSIVolume) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CSIVolume) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *CSIVolume) GetNodesExpected() *int64 {
	if o == nil {
		return nil
	}
	return o.NodesExpected
}

func (o *CSIVolume) GetNodesHealthy() *int64 {
	if o == nil {
		return nil
	}
	return o.NodesHealthy
}

func (o *CSIVolume) GetParameters() map[string]string {
	if o == nil {
		return nil
	}
	return o.Parameters
}

func (o *CSIVolume) GetPluginID() *string {
	if o == nil {
		return nil
	}
	return o.PluginID
}

func (o *CSIVolume) GetProvider() *string {
	if o == nil {
		return nil
	}
	return o.Provider
}

func (o *CSIVolume) GetProviderVersion() *string {
	if o == nil {
		return nil
	}
	return o.ProviderVersion
}

func (o *CSIVolume) GetReadAllocs() map[string]Allocation {
	if o == nil {
		return nil
	}
	return o.ReadAllocs
}

func (o *CSIVolume) GetRequestedCapabilities() []CSIVolumeCapability {
	if o == nil {
		return nil
	}
	return o.RequestedCapabilities
}

func (o *CSIVolume) GetRequestedCapacityMax() *int64 {
	if o == nil {
		return nil
	}
	return o.RequestedCapacityMax
}

func (o *CSIVolume) GetRequestedCapacityMin() *int64 {
	if o == nil {
		return nil
	}
	return o.RequestedCapacityMin
}

func (o *CSIVolume) GetRequestedTopologies() *CSITopologyRequest {
	if o == nil {
		return nil
	}
	return o.RequestedTopologies
}

func (o *CSIVolume) GetResourceExhausted() *time.Time {
	if o == nil {
		return nil
	}
	return o.ResourceExhausted
}

func (o *CSIVolume) GetSchedulable() *bool {
	if o == nil {
		return nil
	}
	return o.Schedulable
}

func (o *CSIVolume) GetSecrets() map[string]string {
	if o == nil {
		return nil
	}
	return o.Secrets
}

func (o *CSIVolume) GetSnapshotID() *string {
	if o == nil {
		return nil
	}
	return o.SnapshotID
}

func (o *CSIVolume) GetTopologies() []CSITopology {
	if o == nil {
		return nil
	}
	return o.Topologies
}

func (o *CSIVolume) GetWriteAllocs() map[string]Allocation {
	if o == nil {
		return nil
	}
	return o.WriteAllocs
}
