// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SampledValue struct {
	Count  *int64            `json:"Count,omitempty"`
	Labels map[string]string `json:"Labels,omitempty"`
	Max    *float64          `json:"Max,omitempty"`
	Mean   *float64          `json:"Mean,omitempty"`
	Min    *float64          `json:"Min,omitempty"`
	Name   *string           `json:"Name,omitempty"`
	Rate   *float64          `json:"Rate,omitempty"`
	Stddev *float64          `json:"Stddev,omitempty"`
	Sum    *float64          `json:"Sum,omitempty"`
}

func (o *SampledValue) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *SampledValue) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *SampledValue) GetMax() *float64 {
	if o == nil {
		return nil
	}
	return o.Max
}

func (o *SampledValue) GetMean() *float64 {
	if o == nil {
		return nil
	}
	return o.Mean
}

func (o *SampledValue) GetMin() *float64 {
	if o == nil {
		return nil
	}
	return o.Min
}

func (o *SampledValue) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *SampledValue) GetRate() *float64 {
	if o == nil {
		return nil
	}
	return o.Rate
}

func (o *SampledValue) GetStddev() *float64 {
	if o == nil {
		return nil
	}
	return o.Stddev
}

func (o *SampledValue) GetSum() *float64 {
	if o == nil {
		return nil
	}
	return o.Sum
}