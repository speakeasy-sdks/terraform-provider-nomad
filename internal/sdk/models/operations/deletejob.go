// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/hashicorp/terraform-provider-nomad/v5/internal/sdk/models/shared"
	"net/http"
)

type DeleteJobRequest struct {
	// A Nomad ACL token.
	XNomadToken *string `header:"style=simple,explode=false,name=X-Nomad-Token"`
	// Boolean flag indicating whether the operation should apply to all instances of the job globally.
	Global *bool `queryParam:"style=form,explode=true,name=global"`
	// Can be used to ensure operations are only run once.
	IdempotencyToken *string `queryParam:"style=form,explode=true,name=idempotency_token"`
	// The job identifier.
	JobName string `pathParam:"style=simple,explode=false,name=jobName"`
	// Filters results based on the specified namespace.
	Namespace *string `queryParam:"style=form,explode=true,name=namespace"`
	// Boolean flag indicating whether to purge allocations of the job after deleting.
	Purge *bool `queryParam:"style=form,explode=true,name=purge"`
	// Filters results based on the specified region.
	Region *string `queryParam:"style=form,explode=true,name=region"`
}

func (o *DeleteJobRequest) GetXNomadToken() *string {
	if o == nil {
		return nil
	}
	return o.XNomadToken
}

func (o *DeleteJobRequest) GetGlobal() *bool {
	if o == nil {
		return nil
	}
	return o.Global
}

func (o *DeleteJobRequest) GetIdempotencyToken() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyToken
}

func (o *DeleteJobRequest) GetJobName() string {
	if o == nil {
		return ""
	}
	return o.JobName
}

func (o *DeleteJobRequest) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *DeleteJobRequest) GetPurge() *bool {
	if o == nil {
		return nil
	}
	return o.Purge
}

func (o *DeleteJobRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

type DeleteJobResponse struct {
	// HTTP response content type for this operation
	ContentType           string
	Headers               map[string][]string
	JobDeregisterResponse *shared.JobDeregisterResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteJobResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteJobResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *DeleteJobResponse) GetJobDeregisterResponse() *shared.JobDeregisterResponse {
	if o == nil {
		return nil
	}
	return o.JobDeregisterResponse
}

func (o *DeleteJobResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteJobResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
