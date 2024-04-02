// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/hashicorp/terraform-provider-nomad/v5/internal/sdk/models/shared"
	"net/http"
)

type PostACLTokenRequest struct {
	ACLToken shared.ACLToken `request:"mediaType=application/json"`
	// A Nomad ACL token.
	XNomadToken *string `header:"style=simple,explode=false,name=X-Nomad-Token"`
	// Can be used to ensure operations are only run once.
	IdempotencyToken *string `queryParam:"style=form,explode=true,name=idempotency_token"`
	// Filters results based on the specified namespace.
	Namespace *string `queryParam:"style=form,explode=true,name=namespace"`
	// Filters results based on the specified region.
	Region *string `queryParam:"style=form,explode=true,name=region"`
	// The token accessor ID.
	TokenAccessor string `pathParam:"style=simple,explode=false,name=tokenAccessor"`
}

func (o *PostACLTokenRequest) GetACLToken() shared.ACLToken {
	if o == nil {
		return shared.ACLToken{}
	}
	return o.ACLToken
}

func (o *PostACLTokenRequest) GetXNomadToken() *string {
	if o == nil {
		return nil
	}
	return o.XNomadToken
}

func (o *PostACLTokenRequest) GetIdempotencyToken() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyToken
}

func (o *PostACLTokenRequest) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *PostACLTokenRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *PostACLTokenRequest) GetTokenAccessor() string {
	if o == nil {
		return ""
	}
	return o.TokenAccessor
}

type PostACLTokenResponse struct {
	ACLToken *shared.ACLToken
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostACLTokenResponse) GetACLToken() *shared.ACLToken {
	if o == nil {
		return nil
	}
	return o.ACLToken
}

func (o *PostACLTokenResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostACLTokenResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *PostACLTokenResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostACLTokenResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}