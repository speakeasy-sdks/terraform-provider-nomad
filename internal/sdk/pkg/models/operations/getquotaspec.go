// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/hashicorp/terraform-provider-nomad/v2/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetQuotaSpecSecurity struct {
	XNomadToken string `security:"scheme,type=apiKey,subtype=header,name=X-Nomad-Token"`
}

func (o *GetQuotaSpecSecurity) GetXNomadToken() string {
	if o == nil {
		return ""
	}
	return o.XNomadToken
}

type GetQuotaSpecRequest struct {
	// A Nomad ACL token.
	XNomadToken *string `header:"style=simple,explode=false,name=X-Nomad-Token"`
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
	// The quota spec identifier.
	SpecName string `pathParam:"style=simple,explode=false,name=specName"`
	// If present, results will include stale reads.
	Stale *string `queryParam:"style=form,explode=true,name=stale"`
	// Provided with IndexParam to wait for change.
	Wait *string `queryParam:"style=form,explode=true,name=wait"`
}

func (o *GetQuotaSpecRequest) GetXNomadToken() *string {
	if o == nil {
		return nil
	}
	return o.XNomadToken
}

func (o *GetQuotaSpecRequest) GetIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.Index
}

func (o *GetQuotaSpecRequest) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *GetQuotaSpecRequest) GetNextToken() *string {
	if o == nil {
		return nil
	}
	return o.NextToken
}

func (o *GetQuotaSpecRequest) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *GetQuotaSpecRequest) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *GetQuotaSpecRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *GetQuotaSpecRequest) GetSpecName() string {
	if o == nil {
		return ""
	}
	return o.SpecName
}

func (o *GetQuotaSpecRequest) GetStale() *string {
	if o == nil {
		return nil
	}
	return o.Stale
}

func (o *GetQuotaSpecRequest) GetWait() *string {
	if o == nil {
		return nil
	}
	return o.Wait
}

type GetQuotaSpecResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	QuotaSpec   *shared.QuotaSpec
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetQuotaSpecResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetQuotaSpecResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *GetQuotaSpecResponse) GetQuotaSpec() *shared.QuotaSpec {
	if o == nil {
		return nil
	}
	return o.QuotaSpec
}

func (o *GetQuotaSpecResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetQuotaSpecResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
