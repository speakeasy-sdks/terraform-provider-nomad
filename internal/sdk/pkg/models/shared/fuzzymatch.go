// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type FuzzyMatch struct {
	ID    *string  `json:"ID,omitempty"`
	Scope []string `json:"Scope,omitempty"`
}

func (o *FuzzyMatch) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *FuzzyMatch) GetScope() []string {
	if o == nil {
		return nil
	}
	return o.Scope
}
