// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"nomad/v2/internal/sdk/pkg/models/shared"
)

type PostJobSecurity struct {
	XNomadToken string `security:"scheme,type=apiKey,subtype=header,name=X-Nomad-Token"`
}

func (o *PostJobSecurity) GetXNomadToken() string {
	if o == nil {
		return ""
	}
	return o.XNomadToken
}

type PostJobRequest struct {
	JobRegisterRequest shared.JobRegisterRequest `request:"mediaType=application/json"`
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

func (o *PostJobRequest) GetJobRegisterRequest() shared.JobRegisterRequest {
	if o == nil {
		return shared.JobRegisterRequest{}
	}
	return o.JobRegisterRequest
}

func (o *PostJobRequest) GetXNomadToken() *string {
	if o == nil {
		return nil
	}
	return o.XNomadToken
}

func (o *PostJobRequest) GetIdempotencyToken() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyToken
}

func (o *PostJobRequest) GetJobName() string {
	if o == nil {
		return ""
	}
	return o.JobName
}

func (o *PostJobRequest) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *PostJobRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

type PostJobResponse struct {
	// HTTP response content type for this operation
	ContentType         string
	Headers             map[string][]string
	JobRegisterResponse *shared.JobRegisterResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostJobResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostJobResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *PostJobResponse) GetJobRegisterResponse() *shared.JobRegisterResponse {
	if o == nil {
		return nil
	}
	return o.JobRegisterResponse
}

func (o *PostJobResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostJobResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
