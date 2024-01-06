// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/hashicorp/terraform-provider-nomad/v3/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetAllocationServicesSecurity struct {
	XNomadToken string `security:"scheme,type=apiKey,subtype=header,name=X-Nomad-Token"`
}

func (o *GetAllocationServicesSecurity) GetXNomadToken() string {
	if o == nil {
		return ""
	}
	return o.XNomadToken
}

type GetAllocationServicesRequest struct {
	// A Nomad ACL token.
	XNomadToken *string `header:"style=simple,explode=false,name=X-Nomad-Token"`
	// Allocation ID.
	AllocID string `pathParam:"style=simple,explode=false,name=allocID"`
	// If set, wait until query exceeds given index. Must be provided with WaitParam.
	Index *int64 `header:"style=simple,explode=false,name=index"`
	// Filters results based on the specified namespace.
	Namespace *string `queryParam:"style=form,explode=true,name=namespace"`
	// Indicates where to start paging for queries that support pagination.
	NextToken *string `queryParam:"style=form,explode=true,name=next_token"`
	// Maximum number of results to return.
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
	// Constrains results to jobs that start with the defined prefix
	Prefix *string `queryParam:"style=form,explode=true,name=prefix"`
	// Filters results based on the specified region.
	Region *string `queryParam:"style=form,explode=true,name=region"`
	// If present, results will include stale reads.
	Stale *string `queryParam:"style=form,explode=true,name=stale"`
	// Provided with IndexParam to wait for change.
	Wait *string `queryParam:"style=form,explode=true,name=wait"`
}

func (o *GetAllocationServicesRequest) GetXNomadToken() *string {
	if o == nil {
		return nil
	}
	return o.XNomadToken
}

func (o *GetAllocationServicesRequest) GetAllocID() string {
	if o == nil {
		return ""
	}
	return o.AllocID
}

func (o *GetAllocationServicesRequest) GetIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.Index
}

func (o *GetAllocationServicesRequest) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *GetAllocationServicesRequest) GetNextToken() *string {
	if o == nil {
		return nil
	}
	return o.NextToken
}

func (o *GetAllocationServicesRequest) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *GetAllocationServicesRequest) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *GetAllocationServicesRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *GetAllocationServicesRequest) GetStale() *string {
	if o == nil {
		return nil
	}
	return o.Stale
}

func (o *GetAllocationServicesRequest) GetWait() *string {
	if o == nil {
		return nil
	}
	return o.Wait
}

type GetAllocationServicesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	Classes     []shared.ServiceRegistration
}

func (o *GetAllocationServicesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAllocationServicesResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *GetAllocationServicesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAllocationServicesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAllocationServicesResponse) GetClasses() []shared.ServiceRegistration {
	if o == nil {
		return nil
	}
	return o.Classes
}
