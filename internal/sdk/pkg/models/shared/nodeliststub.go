// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type NodeListStub struct {
	Address               *string                `json:"Address,omitempty"`
	Attributes            map[string]string      `json:"Attributes,omitempty"`
	CreateIndex           *int64                 `json:"CreateIndex,omitempty"`
	Datacenter            *string                `json:"Datacenter,omitempty"`
	Drain                 *bool                  `json:"Drain,omitempty"`
	Drivers               map[string]DriverInfo  `json:"Drivers,omitempty"`
	ID                    *string                `json:"ID,omitempty"`
	LastDrain             *DrainMetadata         `json:"LastDrain,omitempty"`
	ModifyIndex           *int64                 `json:"ModifyIndex,omitempty"`
	Name                  *string                `json:"Name,omitempty"`
	NodeClass             *string                `json:"NodeClass,omitempty"`
	NodeResources         *NodeResources         `json:"NodeResources,omitempty"`
	ReservedResources     *NodeReservedResources `json:"ReservedResources,omitempty"`
	SchedulingEligibility *string                `json:"SchedulingEligibility,omitempty"`
	Status                *string                `json:"Status,omitempty"`
	StatusDescription     *string                `json:"StatusDescription,omitempty"`
	Version               *string                `json:"Version,omitempty"`
}
