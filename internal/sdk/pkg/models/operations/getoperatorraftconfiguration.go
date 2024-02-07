// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/hashicorp/terraform-provider-nomad/v5/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetOperatorRaftConfigurationSecurity struct {
	XNomadToken string `security:"scheme,type=apiKey,subtype=header,name=X-Nomad-Token"`
}

func (o *GetOperatorRaftConfigurationSecurity) GetXNomadToken() string {
	if o == nil {
		return ""
	}
	return o.XNomadToken
}

type GetOperatorRaftConfigurationRequest struct {
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

func (o *GetOperatorRaftConfigurationRequest) GetXNomadToken() *string {
	if o == nil {
		return nil
	}
	return o.XNomadToken
}

func (o *GetOperatorRaftConfigurationRequest) GetIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.Index
}

func (o *GetOperatorRaftConfigurationRequest) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *GetOperatorRaftConfigurationRequest) GetNextToken() *string {
	if o == nil {
		return nil
	}
	return o.NextToken
}

func (o *GetOperatorRaftConfigurationRequest) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *GetOperatorRaftConfigurationRequest) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *GetOperatorRaftConfigurationRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *GetOperatorRaftConfigurationRequest) GetStale() *string {
	if o == nil {
		return nil
	}
	return o.Stale
}

func (o *GetOperatorRaftConfigurationRequest) GetWait() *string {
	if o == nil {
		return nil
	}
	return o.Wait
}

type GetOperatorRaftConfigurationResponse struct {
	// HTTP response content type for this operation
	ContentType       string
	RaftConfiguration *shared.RaftConfiguration
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetOperatorRaftConfigurationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetOperatorRaftConfigurationResponse) GetRaftConfiguration() *shared.RaftConfiguration {
	if o == nil {
		return nil
	}
	return o.RaftConfiguration
}

func (o *GetOperatorRaftConfigurationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetOperatorRaftConfigurationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
