// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/hashicorp/terraform-provider-nomad/v4/internal/sdk/pkg/models/shared"
	"net/http"
)

type PostACLBootstrapSecurity struct {
	XNomadToken string `security:"scheme,type=apiKey,subtype=header,name=X-Nomad-Token"`
}

func (o *PostACLBootstrapSecurity) GetXNomadToken() string {
	if o == nil {
		return ""
	}
	return o.XNomadToken
}

type PostACLBootstrapRequest struct {
	// A Nomad ACL token.
	XNomadToken *string `header:"style=simple,explode=false,name=X-Nomad-Token"`
	// Can be used to ensure operations are only run once.
	IdempotencyToken *string `queryParam:"style=form,explode=true,name=idempotency_token"`
	// Filters results based on the specified namespace.
	Namespace *string `queryParam:"style=form,explode=true,name=namespace"`
	// Filters results based on the specified region.
	Region *string `queryParam:"style=form,explode=true,name=region"`
}

func (o *PostACLBootstrapRequest) GetXNomadToken() *string {
	if o == nil {
		return nil
	}
	return o.XNomadToken
}

func (o *PostACLBootstrapRequest) GetIdempotencyToken() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyToken
}

func (o *PostACLBootstrapRequest) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *PostACLBootstrapRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

type PostACLBootstrapResponse struct {
	ACLToken *shared.ACLToken
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostACLBootstrapResponse) GetACLToken() *shared.ACLToken {
	if o == nil {
		return nil
	}
	return o.ACLToken
}

func (o *PostACLBootstrapResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostACLBootstrapResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *PostACLBootstrapResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostACLBootstrapResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
