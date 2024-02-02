// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/hashicorp/terraform-provider-nomad/v4/internal/sdk/pkg/models/shared"
	"net/http"
)

type RegisterJobSecurity struct {
	XNomadToken string `security:"scheme,type=apiKey,subtype=header,name=X-Nomad-Token"`
}

func (o *RegisterJobSecurity) GetXNomadToken() string {
	if o == nil {
		return ""
	}
	return o.XNomadToken
}

type RegisterJobRequest struct {
	JobRegisterRequest shared.JobRegisterRequest `request:"mediaType=application/json"`
	// A Nomad ACL token.
	XNomadToken *string `header:"style=simple,explode=false,name=X-Nomad-Token"`
	// Can be used to ensure operations are only run once.
	IdempotencyToken *string `queryParam:"style=form,explode=true,name=idempotency_token"`
	// Filters results based on the specified namespace.
	Namespace *string `queryParam:"style=form,explode=true,name=namespace"`
	// Filters results based on the specified region.
	Region *string `queryParam:"style=form,explode=true,name=region"`
}

func (o *RegisterJobRequest) GetJobRegisterRequest() shared.JobRegisterRequest {
	if o == nil {
		return shared.JobRegisterRequest{}
	}
	return o.JobRegisterRequest
}

func (o *RegisterJobRequest) GetXNomadToken() *string {
	if o == nil {
		return nil
	}
	return o.XNomadToken
}

func (o *RegisterJobRequest) GetIdempotencyToken() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyToken
}

func (o *RegisterJobRequest) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *RegisterJobRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

type RegisterJobResponse struct {
	// HTTP response content type for this operation
	ContentType         string
	Headers             map[string][]string
	JobRegisterResponse *shared.JobRegisterResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RegisterJobResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RegisterJobResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *RegisterJobResponse) GetJobRegisterResponse() *shared.JobRegisterResponse {
	if o == nil {
		return nil
	}
	return o.JobRegisterResponse
}

func (o *RegisterJobResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RegisterJobResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
