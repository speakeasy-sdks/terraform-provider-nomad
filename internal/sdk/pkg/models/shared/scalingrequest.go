// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ScalingRequest struct {
	Count          *int64                 `json:"Count,omitempty"`
	Error          *bool                  `json:"Error,omitempty"`
	Message        *string                `json:"Message,omitempty"`
	Meta           map[string]interface{} `json:"Meta,omitempty"`
	Namespace      *string                `json:"Namespace,omitempty"`
	PolicyOverride *bool                  `json:"PolicyOverride,omitempty"`
	Region         *string                `json:"Region,omitempty"`
	SecretID       *string                `json:"SecretID,omitempty"`
	Target         map[string]string      `json:"Target,omitempty"`
}

func (o *ScalingRequest) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *ScalingRequest) GetError() *bool {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *ScalingRequest) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *ScalingRequest) GetMeta() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *ScalingRequest) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *ScalingRequest) GetPolicyOverride() *bool {
	if o == nil {
		return nil
	}
	return o.PolicyOverride
}

func (o *ScalingRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *ScalingRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}

func (o *ScalingRequest) GetTarget() map[string]string {
	if o == nil {
		return nil
	}
	return o.Target
}
