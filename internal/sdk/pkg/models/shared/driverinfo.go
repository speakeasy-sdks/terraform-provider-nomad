// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/hashicorp/terraform-provider-nomad/v4/internal/sdk/pkg/utils"
	"time"
)

type DriverInfo struct {
	Attributes        map[string]string `json:"Attributes,omitempty"`
	Detected          *bool             `json:"Detected,omitempty"`
	HealthDescription *string           `json:"HealthDescription,omitempty"`
	Healthy           *bool             `json:"Healthy,omitempty"`
	UpdateTime        *time.Time        `json:"UpdateTime,omitempty"`
}

func (d DriverInfo) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DriverInfo) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DriverInfo) GetAttributes() map[string]string {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *DriverInfo) GetDetected() *bool {
	if o == nil {
		return nil
	}
	return o.Detected
}

func (o *DriverInfo) GetHealthDescription() *string {
	if o == nil {
		return nil
	}
	return o.HealthDescription
}

func (o *DriverInfo) GetHealthy() *bool {
	if o == nil {
		return nil
	}
	return o.Healthy
}

func (o *DriverInfo) GetUpdateTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdateTime
}
