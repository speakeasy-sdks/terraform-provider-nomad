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
