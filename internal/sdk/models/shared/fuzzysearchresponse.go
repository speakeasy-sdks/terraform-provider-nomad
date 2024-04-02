// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type FuzzySearchResponse struct {
	KnownLeader *bool                   `json:"KnownLeader,omitempty"`
	LastContact *int64                  `json:"LastContact,omitempty"`
	LastIndex   *int64                  `json:"LastIndex,omitempty"`
	Matches     map[string][]FuzzyMatch `json:"Matches,omitempty"`
	NextToken   *string                 `json:"NextToken,omitempty"`
	RequestTime *int64                  `json:"RequestTime,omitempty"`
	Truncations map[string]bool         `json:"Truncations,omitempty"`
}

func (o *FuzzySearchResponse) GetKnownLeader() *bool {
	if o == nil {
		return nil
	}
	return o.KnownLeader
}

func (o *FuzzySearchResponse) GetLastContact() *int64 {
	if o == nil {
		return nil
	}
	return o.LastContact
}

func (o *FuzzySearchResponse) GetLastIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.LastIndex
}

func (o *FuzzySearchResponse) GetMatches() map[string][]FuzzyMatch {
	if o == nil {
		return nil
	}
	return o.Matches
}

func (o *FuzzySearchResponse) GetNextToken() *string {
	if o == nil {
		return nil
	}
	return o.NextToken
}

func (o *FuzzySearchResponse) GetRequestTime() *int64 {
	if o == nil {
		return nil
	}
	return o.RequestTime
}

func (o *FuzzySearchResponse) GetTruncations() map[string]bool {
	if o == nil {
		return nil
	}
	return o.Truncations
}