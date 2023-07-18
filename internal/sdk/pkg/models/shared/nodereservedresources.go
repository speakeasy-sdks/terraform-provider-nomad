// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type NodeReservedResources struct {
	CPU      *NodeReservedCPUResources     `json:"Cpu,omitempty"`
	Disk     *NodeReservedDiskResources    `json:"Disk,omitempty"`
	Memory   *NodeReservedMemoryResources  `json:"Memory,omitempty"`
	Networks *NodeReservedNetworkResources `json:"Networks,omitempty"`
}