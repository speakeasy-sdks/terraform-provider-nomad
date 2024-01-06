// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/hashicorp/terraform-provider-nomad/v3/internal/sdk/pkg/models/shared"
	"net/http"
)

type PostJobRevertSecurity struct {
	XNomadToken string `security:"scheme,type=apiKey,subtype=header,name=X-Nomad-Token"`
}

func (o *PostJobRevertSecurity) GetXNomadToken() string {
	if o == nil {
		return ""
	}
	return o.XNomadToken
}

type PostJobRevertRequest struct {
	JobRevertRequest shared.JobRevertRequest `request:"mediaType=application/json"`
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

func (o *PostJobRevertRequest) GetJobRevertRequest() shared.JobRevertRequest {
	if o == nil {
		return shared.JobRevertRequest{}
	}
	return o.JobRevertRequest
}

func (o *PostJobRevertRequest) GetXNomadToken() *string {
	if o == nil {
		return nil
	}
	return o.XNomadToken
}

func (o *PostJobRevertRequest) GetIdempotencyToken() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyToken
}

func (o *PostJobRevertRequest) GetJobName() string {
	if o == nil {
		return ""
	}
	return o.JobName
}

func (o *PostJobRevertRequest) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *PostJobRevertRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

type PostJobRevertResponse struct {
	// HTTP response content type for this operation
	ContentType         string
	Headers             map[string][]string
	JobRegisterResponse *shared.JobRegisterResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostJobRevertResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostJobRevertResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *PostJobRevertResponse) GetJobRegisterResponse() *shared.JobRegisterResponse {
	if o == nil {
		return nil
	}
	return o.JobRegisterResponse
}

func (o *PostJobRevertResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostJobRevertResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
