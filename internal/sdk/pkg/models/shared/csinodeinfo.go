// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CSINodeInfo struct {
	AccessibleTopology      *CSITopology `json:"AccessibleTopology,omitempty"`
	ID                      *string      `json:"ID,omitempty"`
	MaxVolumes              *int64       `json:"MaxVolumes,omitempty"`
	RequiresNodeStageVolume *bool        `json:"RequiresNodeStageVolume,omitempty"`
	SupportsCondition       *bool        `json:"SupportsCondition,omitempty"`
	SupportsExpand          *bool        `json:"SupportsExpand,omitempty"`
	SupportsStats           *bool        `json:"SupportsStats,omitempty"`
}

func (o *CSINodeInfo) GetAccessibleTopology() *CSITopology {
	if o == nil {
		return nil
	}
	return o.AccessibleTopology
}

func (o *CSINodeInfo) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CSINodeInfo) GetMaxVolumes() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxVolumes
}

func (o *CSINodeInfo) GetRequiresNodeStageVolume() *bool {
	if o == nil {
		return nil
	}
	return o.RequiresNodeStageVolume
}

func (o *CSINodeInfo) GetSupportsCondition() *bool {
	if o == nil {
		return nil
	}
	return o.SupportsCondition
}

func (o *CSINodeInfo) GetSupportsExpand() *bool {
	if o == nil {
		return nil
	}
	return o.SupportsExpand
}

func (o *CSINodeInfo) GetSupportsStats() *bool {
	if o == nil {
		return nil
	}
	return o.SupportsStats
}
