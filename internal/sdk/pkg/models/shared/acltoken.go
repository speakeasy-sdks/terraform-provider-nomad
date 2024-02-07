// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/hashicorp/terraform-provider-nomad/v5/internal/sdk/pkg/utils"
	"time"
)

type ACLToken struct {
	AccessorID     *string            `json:"AccessorID,omitempty"`
	CreateIndex    *int64             `json:"CreateIndex,omitempty"`
	CreateTime     *time.Time         `json:"CreateTime,omitempty"`
	ExpirationTTL  *int64             `json:"ExpirationTTL,omitempty"`
	ExpirationTime *time.Time         `json:"ExpirationTime,omitempty"`
	Global         *bool              `json:"Global,omitempty"`
	ModifyIndex    *int64             `json:"ModifyIndex,omitempty"`
	Name           *string            `json:"Name,omitempty"`
	Policies       []string           `json:"Policies,omitempty"`
	Roles          []ACLTokenRoleLink `json:"Roles,omitempty"`
	SecretID       *string            `json:"SecretID,omitempty"`
	Type           *string            `json:"Type,omitempty"`
}

func (a ACLToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *ACLToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ACLToken) GetAccessorID() *string {
	if o == nil {
		return nil
	}
	return o.AccessorID
}

func (o *ACLToken) GetCreateIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.CreateIndex
}

func (o *ACLToken) GetCreateTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreateTime
}

func (o *ACLToken) GetExpirationTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.ExpirationTTL
}

func (o *ACLToken) GetExpirationTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.ExpirationTime
}

func (o *ACLToken) GetGlobal() *bool {
	if o == nil {
		return nil
	}
	return o.Global
}

func (o *ACLToken) GetModifyIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.ModifyIndex
}

func (o *ACLToken) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ACLToken) GetPolicies() []string {
	if o == nil {
		return nil
	}
	return o.Policies
}

func (o *ACLToken) GetRoles() []ACLTokenRoleLink {
	if o == nil {
		return nil
	}
	return o.Roles
}

func (o *ACLToken) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}

func (o *ACLToken) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}
