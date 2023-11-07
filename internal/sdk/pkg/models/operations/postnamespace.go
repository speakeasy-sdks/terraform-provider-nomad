// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"nomad/v2/internal/sdk/pkg/models/shared"
)

type PostNamespaceSecurity struct {
	XNomadToken string `security:"scheme,type=apiKey,subtype=header,name=X-Nomad-Token"`
}

func (o *PostNamespaceSecurity) GetXNomadToken() string {
	if o == nil {
		return ""
	}
	return o.XNomadToken
}

type PostNamespaceRequest struct {
	Namespace1 shared.Namespace `request:"mediaType=application/json"`
	// A Nomad ACL token.
	XNomadToken *string `header:"style=simple,explode=false,name=X-Nomad-Token"`
	// Can be used to ensure operations are only run once.
	IdempotencyToken *string `queryParam:"style=form,explode=true,name=idempotency_token"`
	// Filters results based on the specified namespace.
	NamespaceQueryParameter *string `queryParam:"style=form,explode=true,name=namespace"`
	// The namespace identifier.
	NamespaceName string `pathParam:"style=simple,explode=false,name=namespaceName"`
	// Filters results based on the specified region.
	Region *string `queryParam:"style=form,explode=true,name=region"`
}

func (o *PostNamespaceRequest) GetNamespace1() shared.Namespace {
	if o == nil {
		return shared.Namespace{}
	}
	return o.Namespace1
}

func (o *PostNamespaceRequest) GetXNomadToken() *string {
	if o == nil {
		return nil
	}
	return o.XNomadToken
}

func (o *PostNamespaceRequest) GetIdempotencyToken() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyToken
}

func (o *PostNamespaceRequest) GetNamespaceQueryParameter() *string {
	if o == nil {
		return nil
	}
	return o.NamespaceQueryParameter
}

func (o *PostNamespaceRequest) GetNamespaceName() string {
	if o == nil {
		return ""
	}
	return o.NamespaceName
}

func (o *PostNamespaceRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

type PostNamespaceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostNamespaceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostNamespaceResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *PostNamespaceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostNamespaceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
