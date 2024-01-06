// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/hashicorp/terraform-provider-nomad/v3/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetSnapshotsSecurity struct {
	XNomadToken string `security:"scheme,type=apiKey,subtype=header,name=X-Nomad-Token"`
}

func (o *GetSnapshotsSecurity) GetXNomadToken() string {
	if o == nil {
		return ""
	}
	return o.XNomadToken
}

type GetSnapshotsRequest struct {
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
	// Filters volume lists by plugin ID.
	PluginID *string `queryParam:"style=form,explode=true,name=plugin_id"`
	// Constrains results to jobs that start with the defined prefix
	Prefix *string `queryParam:"style=form,explode=true,name=prefix"`
	// Filters results based on the specified region.
	Region *string `queryParam:"style=form,explode=true,name=region"`
	// If present, results will include stale reads.
	Stale *string `queryParam:"style=form,explode=true,name=stale"`
	// Provided with IndexParam to wait for change.
	Wait *string `queryParam:"style=form,explode=true,name=wait"`
}

func (o *GetSnapshotsRequest) GetXNomadToken() *string {
	if o == nil {
		return nil
	}
	return o.XNomadToken
}

func (o *GetSnapshotsRequest) GetIndex() *int64 {
	if o == nil {
		return nil
	}
	return o.Index
}

func (o *GetSnapshotsRequest) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *GetSnapshotsRequest) GetNextToken() *string {
	if o == nil {
		return nil
	}
	return o.NextToken
}

func (o *GetSnapshotsRequest) GetPerPage() *int64 {
	if o == nil {
		return nil
	}
	return o.PerPage
}

func (o *GetSnapshotsRequest) GetPluginID() *string {
	if o == nil {
		return nil
	}
	return o.PluginID
}

func (o *GetSnapshotsRequest) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *GetSnapshotsRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *GetSnapshotsRequest) GetStale() *string {
	if o == nil {
		return nil
	}
	return o.Stale
}

func (o *GetSnapshotsRequest) GetWait() *string {
	if o == nil {
		return nil
	}
	return o.Wait
}

type GetSnapshotsResponse struct {
	CSISnapshotListResponse *shared.CSISnapshotListResponse
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetSnapshotsResponse) GetCSISnapshotListResponse() *shared.CSISnapshotListResponse {
	if o == nil {
		return nil
	}
	return o.CSISnapshotListResponse
}

func (o *GetSnapshotsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSnapshotsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *GetSnapshotsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSnapshotsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
