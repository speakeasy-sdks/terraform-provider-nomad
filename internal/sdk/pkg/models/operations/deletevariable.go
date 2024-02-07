// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/hashicorp/terraform-provider-nomad/v5/internal/sdk/pkg/models/shared"
	"net/http"
)

type DeleteVariableSecurity struct {
	XNomadToken string `security:"scheme,type=apiKey,subtype=header,name=X-Nomad-Token"`
}

func (o *DeleteVariableSecurity) GetXNomadToken() string {
	if o == nil {
		return ""
	}
	return o.XNomadToken
}

type DeleteVariableRequest struct {
	Variable shared.Variable `request:"mediaType=application/json"`
	// A Nomad ACL token.
	XNomadToken *string `header:"style=simple,explode=false,name=X-Nomad-Token"`
	// A compare-and-set parameter for Nomad Variables
	Cas *int64 `queryParam:"style=form,explode=true,name=cas"`
	// Can be used to ensure operations are only run once.
	IdempotencyToken *string `queryParam:"style=form,explode=true,name=idempotency_token"`
	// Filters results based on the specified namespace.
	Namespace *string `queryParam:"style=form,explode=true,name=namespace"`
	// A path to a Nomad Variable
	Path string `pathParam:"style=simple,explode=false,name=path"`
	// Filters results based on the specified region.
	Region *string `queryParam:"style=form,explode=true,name=region"`
}

func (o *DeleteVariableRequest) GetVariable() shared.Variable {
	if o == nil {
		return shared.Variable{}
	}
	return o.Variable
}

func (o *DeleteVariableRequest) GetXNomadToken() *string {
	if o == nil {
		return nil
	}
	return o.XNomadToken
}

func (o *DeleteVariableRequest) GetCas() *int64 {
	if o == nil {
		return nil
	}
	return o.Cas
}

func (o *DeleteVariableRequest) GetIdempotencyToken() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyToken
}

func (o *DeleteVariableRequest) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *DeleteVariableRequest) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

func (o *DeleteVariableRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

type DeleteVariableResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteVariableResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteVariableResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *DeleteVariableResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteVariableResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
