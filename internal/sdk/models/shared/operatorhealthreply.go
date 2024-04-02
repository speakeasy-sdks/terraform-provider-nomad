// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type OperatorHealthReply struct {
	FailureTolerance *int64         `json:"FailureTolerance,omitempty"`
	Healthy          *bool          `json:"Healthy,omitempty"`
	Servers          []ServerHealth `json:"Servers,omitempty"`
}

func (o *OperatorHealthReply) GetFailureTolerance() *int64 {
	if o == nil {
		return nil
	}
	return o.FailureTolerance
}

func (o *OperatorHealthReply) GetHealthy() *bool {
	if o == nil {
		return nil
	}
	return o.Healthy
}

func (o *OperatorHealthReply) GetServers() []ServerHealth {
	if o == nil {
		return nil
	}
	return o.Servers
}