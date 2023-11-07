// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"nomad/v2/internal/sdk/pkg/models/shared"
)

type PostJobValidateRequestSecurity struct {
	XNomadToken string `security:"scheme,type=apiKey,subtype=header,name=X-Nomad-Token"`
}

func (o *PostJobValidateRequestSecurity) GetXNomadToken() string {
	if o == nil {
		return ""
	}
	return o.XNomadToken
}

type PostJobValidateRequestRequest struct {
	JobValidateRequest shared.JobValidateRequest `request:"mediaType=application/json"`
	// A Nomad ACL token.
	XNomadToken *string `header:"style=simple,explode=false,name=X-Nomad-Token"`
	// Can be used to ensure operations are only run once.
	IdempotencyToken *string `queryParam:"style=form,explode=true,name=idempotency_token"`
	// Filters results based on the specified namespace.
	Namespace *string `queryParam:"style=form,explode=true,name=namespace"`
	// Filters results based on the specified region.
	Region *string `queryParam:"style=form,explode=true,name=region"`
}

func (o *PostJobValidateRequestRequest) GetJobValidateRequest() shared.JobValidateRequest {
	if o == nil {
		return shared.JobValidateRequest{}
	}
	return o.JobValidateRequest
}

func (o *PostJobValidateRequestRequest) GetXNomadToken() *string {
	if o == nil {
		return nil
	}
	return o.XNomadToken
}

func (o *PostJobValidateRequestRequest) GetIdempotencyToken() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyToken
}

func (o *PostJobValidateRequestRequest) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *PostJobValidateRequestRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

type PostJobValidateRequestResponse struct {
	// HTTP response content type for this operation
	ContentType         string
	Headers             map[string][]string
	JobValidateResponse *shared.JobValidateResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostJobValidateRequestResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostJobValidateRequestResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *PostJobValidateRequestResponse) GetJobValidateResponse() *shared.JobValidateResponse {
	if o == nil {
		return nil
	}
	return o.JobValidateResponse
}

func (o *PostJobValidateRequestResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostJobValidateRequestResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
