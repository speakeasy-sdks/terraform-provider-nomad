// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TaskGroupDiff struct {
	Fields  []FieldDiff       `json:"Fields,omitempty"`
	Name    *string           `json:"Name,omitempty"`
	Objects []ObjectDiffInput `json:"Objects,omitempty"`
	Tasks   []TaskDiff        `json:"Tasks,omitempty"`
	Type    *string           `json:"Type,omitempty"`
	Updates map[string]int64  `json:"Updates,omitempty"`
}

func (o *TaskGroupDiff) GetFields() []FieldDiff {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *TaskGroupDiff) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *TaskGroupDiff) GetObjects() []ObjectDiffInput {
	if o == nil {
		return nil
	}
	return o.Objects
}

func (o *TaskGroupDiff) GetTasks() []TaskDiff {
	if o == nil {
		return nil
	}
	return o.Tasks
}

func (o *TaskGroupDiff) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *TaskGroupDiff) GetUpdates() map[string]int64 {
	if o == nil {
		return nil
	}
	return o.Updates
}