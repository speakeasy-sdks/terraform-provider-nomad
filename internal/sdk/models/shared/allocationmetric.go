// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AllocationMetric struct {
	AllocationTime     *int64               `json:"AllocationTime,omitempty"`
	ClassExhausted     map[string]int64     `json:"ClassExhausted,omitempty"`
	ClassFiltered      map[string]int64     `json:"ClassFiltered,omitempty"`
	CoalescedFailures  *int64               `json:"CoalescedFailures,omitempty"`
	ConstraintFiltered map[string]int64     `json:"ConstraintFiltered,omitempty"`
	DimensionExhausted map[string]int64     `json:"DimensionExhausted,omitempty"`
	NodesAvailable     map[string]int64     `json:"NodesAvailable,omitempty"`
	NodesEvaluated     *int64               `json:"NodesEvaluated,omitempty"`
	NodesExhausted     *int64               `json:"NodesExhausted,omitempty"`
	NodesFiltered      *int64               `json:"NodesFiltered,omitempty"`
	QuotaExhausted     []string             `json:"QuotaExhausted,omitempty"`
	ResourcesExhausted map[string]Resources `json:"ResourcesExhausted,omitempty"`
	ScoreMetaData      []NodeScoreMeta      `json:"ScoreMetaData,omitempty"`
	Scores             map[string]float64   `json:"Scores,omitempty"`
}

func (o *AllocationMetric) GetAllocationTime() *int64 {
	if o == nil {
		return nil
	}
	return o.AllocationTime
}

func (o *AllocationMetric) GetClassExhausted() map[string]int64 {
	if o == nil {
		return nil
	}
	return o.ClassExhausted
}

func (o *AllocationMetric) GetClassFiltered() map[string]int64 {
	if o == nil {
		return nil
	}
	return o.ClassFiltered
}

func (o *AllocationMetric) GetCoalescedFailures() *int64 {
	if o == nil {
		return nil
	}
	return o.CoalescedFailures
}

func (o *AllocationMetric) GetConstraintFiltered() map[string]int64 {
	if o == nil {
		return nil
	}
	return o.ConstraintFiltered
}

func (o *AllocationMetric) GetDimensionExhausted() map[string]int64 {
	if o == nil {
		return nil
	}
	return o.DimensionExhausted
}

func (o *AllocationMetric) GetNodesAvailable() map[string]int64 {
	if o == nil {
		return nil
	}
	return o.NodesAvailable
}

func (o *AllocationMetric) GetNodesEvaluated() *int64 {
	if o == nil {
		return nil
	}
	return o.NodesEvaluated
}

func (o *AllocationMetric) GetNodesExhausted() *int64 {
	if o == nil {
		return nil
	}
	return o.NodesExhausted
}

func (o *AllocationMetric) GetNodesFiltered() *int64 {
	if o == nil {
		return nil
	}
	return o.NodesFiltered
}

func (o *AllocationMetric) GetQuotaExhausted() []string {
	if o == nil {
		return nil
	}
	return o.QuotaExhausted
}

func (o *AllocationMetric) GetResourcesExhausted() map[string]Resources {
	if o == nil {
		return nil
	}
	return o.ResourcesExhausted
}

func (o *AllocationMetric) GetScoreMetaData() []NodeScoreMeta {
	if o == nil {
		return nil
	}
	return o.ScoreMetaData
}

func (o *AllocationMetric) GetScores() map[string]float64 {
	if o == nil {
		return nil
	}
	return o.Scores
}