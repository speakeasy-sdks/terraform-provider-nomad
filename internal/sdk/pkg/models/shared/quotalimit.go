// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type QuotaLimit struct {
	Hash           *string    `json:"Hash,omitempty"`
	Region         *string    `json:"Region,omitempty"`
	RegionLimit    *Resources `json:"RegionLimit,omitempty"`
	VariablesLimit *int64     `json:"VariablesLimit,omitempty"`
}
