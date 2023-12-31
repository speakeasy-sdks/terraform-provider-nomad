// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CSISnapshotListResponse struct {
	KnownLeader *bool         `json:"KnownLeader,omitempty"`
	LastContact *int64        `json:"LastContact,omitempty"`
	LastIndex   *int64        `json:"LastIndex,omitempty"`
	NextToken   *string       `json:"NextToken,omitempty"`
	RequestTime *int64        `json:"RequestTime,omitempty"`
	Snapshots   []CSISnapshot `json:"Snapshots,omitempty"`
}

func (o *CSISnapshotListResponse) GetKnownLeader() *bool {
	if o == nil {
		return nil
	}
	return o.KnownLeader
}

func (o *CSISnapshotListResponse) GetLastContact() *int64 {
	if o == nil {
		return nil
	}
	return o.LastContact
}

func (o *CSISnapshotListResponse) GetLastIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.LastIndex
}

func (o *CSISnapshotListResponse) GetNextToken() *string {
	if o == nil {
		return nil
	}
	return o.NextToken
}

func (o *CSISnapshotListResponse) GetRequestTime() *int64 {
	if o == nil {
		return nil
	}
	return o.RequestTime
}

func (o *CSISnapshotListResponse) GetSnapshots() []CSISnapshot {
	if o == nil {
		return nil
	}
	return o.Snapshots
}
