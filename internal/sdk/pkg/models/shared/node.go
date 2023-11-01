// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Node struct {
	Attributes            map[string]string          `json:"Attributes,omitempty"`
	CSIControllerPlugins  map[string]CSIInfo         `json:"CSIControllerPlugins,omitempty"`
	CSINodePlugins        map[string]CSIInfo         `json:"CSINodePlugins,omitempty"`
	CgroupParent          *string                    `json:"CgroupParent,omitempty"`
	CreateIndex           *int64                     `json:"CreateIndex,omitempty"`
	Datacenter            *string                    `json:"Datacenter,omitempty"`
	Drain                 *bool                      `json:"Drain,omitempty"`
	DrainStrategy         *DrainStrategy             `json:"DrainStrategy,omitempty"`
	Drivers               map[string]DriverInfo      `json:"Drivers,omitempty"`
	Events                []NodeEvent                `json:"Events,omitempty"`
	HTTPAddr              *string                    `json:"HTTPAddr,omitempty"`
	HostNetworks          map[string]HostNetworkInfo `json:"HostNetworks,omitempty"`
	HostVolumes           map[string]HostVolumeInfo  `json:"HostVolumes,omitempty"`
	ID                    *string                    `json:"ID,omitempty"`
	LastDrain             *DrainMetadata             `json:"LastDrain,omitempty"`
	Links                 map[string]string          `json:"Links,omitempty"`
	Meta                  map[string]string          `json:"Meta,omitempty"`
	ModifyIndex           *int64                     `json:"ModifyIndex,omitempty"`
	Name                  *string                    `json:"Name,omitempty"`
	NodeClass             *string                    `json:"NodeClass,omitempty"`
	NodeResources         *NodeResources             `json:"NodeResources,omitempty"`
	Reserved              *Resources                 `json:"Reserved,omitempty"`
	ReservedResources     *NodeReservedResources     `json:"ReservedResources,omitempty"`
	Resources             *Resources                 `json:"Resources,omitempty"`
	SchedulingEligibility *string                    `json:"SchedulingEligibility,omitempty"`
	Status                *string                    `json:"Status,omitempty"`
	StatusDescription     *string                    `json:"StatusDescription,omitempty"`
	StatusUpdatedAt       *int64                     `json:"StatusUpdatedAt,omitempty"`
	TLSEnabled            *bool                      `json:"TLSEnabled,omitempty"`
}

func (o *Node) GetAttributes() map[string]string {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *Node) GetCSIControllerPlugins() map[string]CSIInfo {
	if o == nil {
		return nil
	}
	return o.CSIControllerPlugins
}

func (o *Node) GetCSINodePlugins() map[string]CSIInfo {
	if o == nil {
		return nil
	}
	return o.CSINodePlugins
}

func (o *Node) GetCgroupParent() *string {
	if o == nil {
		return nil
	}
	return o.CgroupParent
}

func (o *Node) GetCreateIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.CreateIndex
}

func (o *Node) GetDatacenter() *string {
	if o == nil {
		return nil
	}
	return o.Datacenter
}

func (o *Node) GetDrain() *bool {
	if o == nil {
		return nil
	}
	return o.Drain
}

func (o *Node) GetDrainStrategy() *DrainStrategy {
	if o == nil {
		return nil
	}
	return o.DrainStrategy
}

func (o *Node) GetDrivers() map[string]DriverInfo {
	if o == nil {
		return nil
	}
	return o.Drivers
}

func (o *Node) GetEvents() []NodeEvent {
	if o == nil {
		return nil
	}
	return o.Events
}

func (o *Node) GetHTTPAddr() *string {
	if o == nil {
		return nil
	}
	return o.HTTPAddr
}

func (o *Node) GetHostNetworks() map[string]HostNetworkInfo {
	if o == nil {
		return nil
	}
	return o.HostNetworks
}

func (o *Node) GetHostVolumes() map[string]HostVolumeInfo {
	if o == nil {
		return nil
	}
	return o.HostVolumes
}

func (o *Node) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Node) GetLastDrain() *DrainMetadata {
	if o == nil {
		return nil
	}
	return o.LastDrain
}

func (o *Node) GetLinks() map[string]string {
	if o == nil {
		return nil
	}
	return o.Links
}

func (o *Node) GetMeta() map[string]string {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *Node) GetModifyIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.ModifyIndex
}

func (o *Node) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Node) GetNodeClass() *string {
	if o == nil {
		return nil
	}
	return o.NodeClass
}

func (o *Node) GetNodeResources() *NodeResources {
	if o == nil {
		return nil
	}
	return o.NodeResources
}

func (o *Node) GetReserved() *Resources {
	if o == nil {
		return nil
	}
	return o.Reserved
}

func (o *Node) GetReservedResources() *NodeReservedResources {
	if o == nil {
		return nil
	}
	return o.ReservedResources
}

func (o *Node) GetResources() *Resources {
	if o == nil {
		return nil
	}
	return o.Resources
}

func (o *Node) GetSchedulingEligibility() *string {
	if o == nil {
		return nil
	}
	return o.SchedulingEligibility
}

func (o *Node) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *Node) GetStatusDescription() *string {
	if o == nil {
		return nil
	}
	return o.StatusDescription
}

func (o *Node) GetStatusUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.StatusUpdatedAt
}

func (o *Node) GetTLSEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.TLSEnabled
}
