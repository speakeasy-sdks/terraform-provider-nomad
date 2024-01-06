// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/hashicorp/terraform-provider-nomad/v3/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutOperatorAutopilotConfigurationSecurity struct {
	XNomadToken string `security:"scheme,type=apiKey,subtype=header,name=X-Nomad-Token"`
}

func (o *PutOperatorAutopilotConfigurationSecurity) GetXNomadToken() string {
	if o == nil {
		return ""
	}
	return o.XNomadToken
}

type PutOperatorAutopilotConfigurationRequest struct {
	AutopilotConfiguration shared.AutopilotConfiguration `request:"mediaType=application/json"`
	// A Nomad ACL token.
	XNomadToken *string `header:"style=simple,explode=false,name=X-Nomad-Token"`
	// Can be used to ensure operations are only run once.
	IdempotencyToken *string `queryParam:"style=form,explode=true,name=idempotency_token"`
	// Filters results based on the specified namespace.
	Namespace *string `queryParam:"style=form,explode=true,name=namespace"`
	// Filters results based on the specified region.
	Region *string `queryParam:"style=form,explode=true,name=region"`
}

func (o *PutOperatorAutopilotConfigurationRequest) GetAutopilotConfiguration() shared.AutopilotConfiguration {
	if o == nil {
		return shared.AutopilotConfiguration{}
	}
	return o.AutopilotConfiguration
}

func (o *PutOperatorAutopilotConfigurationRequest) GetXNomadToken() *string {
	if o == nil {
		return nil
	}
	return o.XNomadToken
}

func (o *PutOperatorAutopilotConfigurationRequest) GetIdempotencyToken() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyToken
}

func (o *PutOperatorAutopilotConfigurationRequest) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *PutOperatorAutopilotConfigurationRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

type PutOperatorAutopilotConfigurationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	Boolean     *bool
}

func (o *PutOperatorAutopilotConfigurationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutOperatorAutopilotConfigurationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutOperatorAutopilotConfigurationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PutOperatorAutopilotConfigurationResponse) GetBoolean() *bool {
	if o == nil {
		return nil
	}
	return o.Boolean
}
