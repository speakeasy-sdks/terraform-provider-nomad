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