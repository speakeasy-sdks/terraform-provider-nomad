// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CheckRestart struct {
	Grace          *int64 `json:"Grace,omitempty"`
	IgnoreWarnings *bool  `json:"IgnoreWarnings,omitempty"`
	Limit          *int64 `json:"Limit,omitempty"`
}

func (o *CheckRestart) GetGrace() *int64 {
	if o == nil {
		return nil
	}
	return o.Grace
}

func (o *CheckRestart) GetIgnoreWarnings() *bool {
	if o == nil {
		return nil
	}
	return o.IgnoreWarnings
}

func (o *CheckRestart) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}
