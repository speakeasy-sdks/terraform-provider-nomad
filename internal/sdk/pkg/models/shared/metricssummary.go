// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type MetricsSummary struct {
	Counters  []SampledValue `json:"Counters,omitempty"`
	Gauges    []GaugeValue   `json:"Gauges,omitempty"`
	Points    []PointValue   `json:"Points,omitempty"`
	Samples   []SampledValue `json:"Samples,omitempty"`
	Timestamp *string        `json:"Timestamp,omitempty"`
}