// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CSISnapshotCreateResponse struct {
	KnownLeader *bool         `json:"KnownLeader,omitempty"`
	LastContact *int64        `json:"LastContact,omitempty"`
	LastIndex   *int64        `json:"LastIndex,omitempty"`
	NextToken   *string       `json:"NextToken,omitempty"`
	RequestTime *int64        `json:"RequestTime,omitempty"`
	Snapshots   []CSISnapshot `json:"Snapshots,omitempty"`
}

func (o *CSISnapshotCreateResponse) GetKnownLeader() *bool {
	if o == nil {
		return nil
	}
	return o.KnownLeader
}

func (o *CSISnapshotCreateResponse) GetLastContact() *int64 {
	if o == nil {
		return nil
	}
	return o.LastContact
}

func (o *CSISnapshotCreateResponse) GetLastIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.LastIndex
}

func (o *CSISnapshotCreateResponse) GetNextToken() *string {
	if o == nil {
		return nil
	}
	return o.NextToken
}

func (o *CSISnapshotCreateResponse) GetRequestTime() *int64 {
	if o == nil {
		return nil
	}
	return o.RequestTime
}

func (o *CSISnapshotCreateResponse) GetSnapshots() []CSISnapshot {
	if o == nil {
		return nil
	}
	return o.Snapshots
}
