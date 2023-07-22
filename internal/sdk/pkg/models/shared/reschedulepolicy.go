// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ReschedulePolicy struct {
	Attempts      *int64  `json:"Attempts,omitempty"`
	Delay         *int64  `json:"Delay,omitempty"`
	DelayFunction *string `json:"DelayFunction,omitempty"`
	Interval      *int64  `json:"Interval,omitempty"`
	MaxDelay      *int64  `json:"MaxDelay,omitempty"`
	Unlimited     *bool   `json:"Unlimited,omitempty"`
}
