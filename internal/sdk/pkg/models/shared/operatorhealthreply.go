// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type OperatorHealthReply struct {
	FailureTolerance *int64         `json:"FailureTolerance,omitempty"`
	Healthy          *bool          `json:"Healthy,omitempty"`
	Servers          []ServerHealth `json:"Servers,omitempty"`
}