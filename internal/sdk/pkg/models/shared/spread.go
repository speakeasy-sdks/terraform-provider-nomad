// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Spread struct {
	Attribute    *string        `json:"Attribute,omitempty"`
	SpreadTarget []SpreadTarget `json:"SpreadTarget,omitempty"`
	Weight       *int64         `json:"Weight,omitempty"`
}

func (o *Spread) GetAttribute() *string {
	if o == nil {
		return nil
	}
	return o.Attribute
}

func (o *Spread) GetSpreadTarget() []SpreadTarget {
	if o == nil {
		return nil
	}
	return o.SpreadTarget
}

func (o *Spread) GetWeight() *int64 {
	if o == nil {
		return nil
	}
	return o.Weight
}
