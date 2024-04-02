// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AutopilotConfiguration struct {
	CleanupDeadServers      *bool   `json:"CleanupDeadServers,omitempty"`
	CreateIndex             *int64  `json:"CreateIndex,omitempty"`
	DisableUpgradeMigration *bool   `json:"DisableUpgradeMigration,omitempty"`
	EnableCustomUpgrades    *bool   `json:"EnableCustomUpgrades,omitempty"`
	EnableRedundancyZones   *bool   `json:"EnableRedundancyZones,omitempty"`
	LastContactThreshold    *string `json:"LastContactThreshold,omitempty"`
	MaxTrailingLogs         *int64  `json:"MaxTrailingLogs,omitempty"`
	MinQuorum               *int64  `json:"MinQuorum,omitempty"`
	ModifyIndex             *int64  `json:"ModifyIndex,omitempty"`
	ServerStabilizationTime *string `json:"ServerStabilizationTime,omitempty"`
}

func (o *AutopilotConfiguration) GetCleanupDeadServers() *bool {
	if o == nil {
		return nil
	}
	return o.CleanupDeadServers
}

func (o *AutopilotConfiguration) GetCreateIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.CreateIndex
}

func (o *AutopilotConfiguration) GetDisableUpgradeMigration() *bool {
	if o == nil {
		return nil
	}
	return o.DisableUpgradeMigration
}

func (o *AutopilotConfiguration) GetEnableCustomUpgrades() *bool {
	if o == nil {
		return nil
	}
	return o.EnableCustomUpgrades
}

func (o *AutopilotConfiguration) GetEnableRedundancyZones() *bool {
	if o == nil {
		return nil
	}
	return o.EnableRedundancyZones
}

func (o *AutopilotConfiguration) GetLastContactThreshold() *string {
	if o == nil {
		return nil
	}
	return o.LastContactThreshold
}

func (o *AutopilotConfiguration) GetMaxTrailingLogs() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxTrailingLogs
}

func (o *AutopilotConfiguration) GetMinQuorum() *int64 {
	if o == nil {
		return nil
	}
	return o.MinQuorum
}

func (o *AutopilotConfiguration) GetModifyIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.ModifyIndex
}

func (o *AutopilotConfiguration) GetServerStabilizationTime() *string {
	if o == nil {
		return nil
	}
	return o.ServerStabilizationTime
}