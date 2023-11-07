// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"nomad/v2/internal/sdk/pkg/models/shared"
)

type PostSnapshotSecurity struct {
	XNomadToken string `security:"scheme,type=apiKey,subtype=header,name=X-Nomad-Token"`
}

func (o *PostSnapshotSecurity) GetXNomadToken() string {
	if o == nil {
		return ""
	}
	return o.XNomadToken
}

type PostSnapshotRequest struct {
	CSISnapshotCreateRequest shared.CSISnapshotCreateRequest `request:"mediaType=application/json"`
	// A Nomad ACL token.
	XNomadToken *string `header:"style=simple,explode=false,name=X-Nomad-Token"`
	// Can be used to ensure operations are only run once.
	IdempotencyToken *string `queryParam:"style=form,explode=true,name=idempotency_token"`
	// Filters results based on the specified namespace.
	Namespace *string `queryParam:"style=form,explode=true,name=namespace"`
	// Filters results based on the specified region.
	Region *string `queryParam:"style=form,explode=true,name=region"`
}

func (o *PostSnapshotRequest) GetCSISnapshotCreateRequest() shared.CSISnapshotCreateRequest {
	if o == nil {
		return shared.CSISnapshotCreateRequest{}
	}
	return o.CSISnapshotCreateRequest
}

func (o *PostSnapshotRequest) GetXNomadToken() *string {
	if o == nil {
		return nil
	}
	return o.XNomadToken
}

func (o *PostSnapshotRequest) GetIdempotencyToken() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyToken
}

func (o *PostSnapshotRequest) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *PostSnapshotRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

type PostSnapshotResponse struct {
	CSISnapshotCreateResponse *shared.CSISnapshotCreateResponse
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostSnapshotResponse) GetCSISnapshotCreateResponse() *shared.CSISnapshotCreateResponse {
	if o == nil {
		return nil
	}
	return o.CSISnapshotCreateResponse
}

func (o *PostSnapshotResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostSnapshotResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *PostSnapshotResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostSnapshotResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
