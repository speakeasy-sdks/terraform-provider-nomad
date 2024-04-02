// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetQuotasRequest struct {
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
	// If present, results will include stale reads.
	Stale *string `queryParam:"style=form,explode=true,name=stale"`
	// Provided with IndexParam to wait for change.
	Wait *string `queryParam:"style=form,explode=true,name=wait"`
}

func (o *GetQuotasRequest) GetXNomadToken() *string {
	if o == nil {
		return nil
	}
	return o.XNomadToken
}

func (o *GetQuotasRequest) GetIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.Index
}

func (o *GetQuotasRequest) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *GetQuotasRequest) GetNextToken() *string {
	if o == nil {
		return nil
	}
	return o.NextToken
}

func (o *GetQuotasRequest) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *GetQuotasRequest) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *GetQuotasRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *GetQuotasRequest) GetStale() *string {
	if o == nil {
		return nil
	}
	return o.Stale
}

func (o *GetQuotasRequest) GetWait() *string {
	if o == nil {
		return nil
	}
	return o.Wait
}

type GetQuotasResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	Anies       []interface{}
}

func (o *GetQuotasResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetQuotasResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *GetQuotasResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetQuotasResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetQuotasResponse) GetAnies() []interface{} {
	if o == nil {
		return nil
	}
	return o.Anies
}