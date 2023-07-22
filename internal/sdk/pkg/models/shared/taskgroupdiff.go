// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TaskGroupDiff struct {
	Fields  []FieldDiff      `json:"Fields,omitempty"`
	Name    *string          `json:"Name,omitempty"`
	Objects []ObjectDiff     `json:"Objects,omitempty"`
	Tasks   []TaskDiff       `json:"Tasks,omitempty"`
	Type    *string          `json:"Type,omitempty"`
	Updates map[string]int64 `json:"Updates,omitempty"`
}
