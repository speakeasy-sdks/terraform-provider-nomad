// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/hashicorp/terraform-provider-nomad/v5/internal/sdk/pkg/models/shared"
	"net/http"
)

type PostJobScalingRequestSecurity struct {
	XNomadToken string `security:"scheme,type=apiKey,subtype=header,name=X-Nomad-Token"`
}

func (o *PostJobScalingRequestSecurity) GetXNomadToken() string {
	if o == nil {
		return ""
	}
	return o.XNomadToken
}

type PostJobScalingRequestRequest struct {
	ScalingRequest shared.ScalingRequest `request:"mediaType=application/json"`
	// A Nomad ACL token.
	XNomadToken *string `header:"style=simple,explode=false,name=X-Nomad-Token"`
	// Can be used to ensure operations are only run once.
	IdempotencyToken *string `queryParam:"style=form,explode=true,name=idempotency_token"`
	// The job identifier.
	JobName string `pathParam:"style=simple,explode=false,name=jobName"`
	// Filters results based on the specified namespace.
	Namespace *string `queryParam:"style=form,explode=true,name=namespace"`
	// Filters results based on the specified region.
	Region *string `queryParam:"style=form,explode=true,name=region"`
}

func (o *PostJobScalingRequestRequest) GetScalingRequest() shared.ScalingRequest {
	if o == nil {
		return shared.ScalingRequest{}
	}
	return o.ScalingRequest
}

func (o *PostJobScalingRequestRequest) GetXNomadToken() *string {
	if o == nil {
		return nil
	}
	return o.XNomadToken
}

func (o *PostJobScalingRequestRequest) GetIdempotencyToken() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyToken
}

func (o *PostJobScalingRequestRequest) GetJobName() string {
	if o == nil {
		return ""
	}
	return o.JobName
}

func (o *PostJobScalingRequestRequest) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *PostJobScalingRequestRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

type PostJobScalingRequestResponse struct {
	// HTTP response content type for this operation
	ContentType         string
	Headers             map[string][]string
	JobRegisterResponse *shared.JobRegisterResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostJobScalingRequestResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostJobScalingRequestResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *PostJobScalingRequestResponse) GetJobRegisterResponse() *shared.JobRegisterResponse {
	if o == nil {
		return nil
	}
	return o.JobRegisterResponse
}

func (o *PostJobScalingRequestResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostJobScalingRequestResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
