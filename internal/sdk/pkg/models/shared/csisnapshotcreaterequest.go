// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CSISnapshotCreateRequest struct {
	Namespace *string       `json:"Namespace,omitempty"`
	Region    *string       `json:"Region,omitempty"`
	SecretID  *string       `json:"SecretID,omitempty"`
	Snapshots []CSISnapshot `json:"Snapshots,omitempty"`
}
