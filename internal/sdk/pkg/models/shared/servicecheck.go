// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ServiceCheck struct {
	AddressMode            *string             `json:"AddressMode,omitempty"`
	Advertise              *string             `json:"Advertise,omitempty"`
	Args                   []string            `json:"Args,omitempty"`
	Body                   *string             `json:"Body,omitempty"`
	CheckRestart           *CheckRestart       `json:"CheckRestart,omitempty"`
	Command                *string             `json:"Command,omitempty"`
	Expose                 *bool               `json:"Expose,omitempty"`
	FailuresBeforeCritical *int64              `json:"FailuresBeforeCritical,omitempty"`
	GRPCService            *string             `json:"GRPCService,omitempty"`
	GRPCUseTLS             *bool               `json:"GRPCUseTLS,omitempty"`
	Header                 map[string][]string `json:"Header,omitempty"`
	InitialStatus          *string             `json:"InitialStatus,omitempty"`
	Interval               *int64              `json:"Interval,omitempty"`
	Method                 *string             `json:"Method,omitempty"`
	Name                   *string             `json:"Name,omitempty"`
	OnUpdate               *string             `json:"OnUpdate,omitempty"`
	Path                   *string             `json:"Path,omitempty"`
	PortLabel              *string             `json:"PortLabel,omitempty"`
	Protocol               *string             `json:"Protocol,omitempty"`
	SuccessBeforePassing   *int64              `json:"SuccessBeforePassing,omitempty"`
	TLSSkipVerify          *bool               `json:"TLSSkipVerify,omitempty"`
	TaskName               *string             `json:"TaskName,omitempty"`
	Timeout                *int64              `json:"Timeout,omitempty"`
	Type                   *string             `json:"Type,omitempty"`
}