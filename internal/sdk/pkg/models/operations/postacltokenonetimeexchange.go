// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"nomad/internal/sdk/pkg/models/shared"
)

type PostACLTokenOnetimeExchangeSecurity struct {
	XNomadToken string `security:"scheme,type=apiKey,subtype=header,name=X-Nomad-Token"`
}

func (o *PostACLTokenOnetimeExchangeSecurity) GetXNomadToken() string {
	if o == nil {
		return ""
	}
	return o.XNomadToken
}

type PostACLTokenOnetimeExchangeRequest struct {
	OneTimeTokenExchangeRequest shared.OneTimeTokenExchangeRequest `request:"mediaType=application/json"`
	// A Nomad ACL token.
	XNomadToken *string `header:"style=simple,explode=false,name=X-Nomad-Token"`
	// Can be used to ensure operations are only run once.
	IdempotencyToken *string `queryParam:"style=form,explode=true,name=idempotency_token"`
	// Filters results based on the specified namespace.
	Namespace *string `queryParam:"style=form,explode=true,name=namespace"`
	// Filters results based on the specified region.
	Region *string `queryParam:"style=form,explode=true,name=region"`
}

func (o *PostACLTokenOnetimeExchangeRequest) GetOneTimeTokenExchangeRequest() shared.OneTimeTokenExchangeRequest {
	if o == nil {
		return shared.OneTimeTokenExchangeRequest{}
	}
	return o.OneTimeTokenExchangeRequest
}

func (o *PostACLTokenOnetimeExchangeRequest) GetXNomadToken() *string {
	if o == nil {
		return nil
	}
	return o.XNomadToken
}

func (o *PostACLTokenOnetimeExchangeRequest) GetIdempotencyToken() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyToken
}

func (o *PostACLTokenOnetimeExchangeRequest) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *PostACLTokenOnetimeExchangeRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

type PostACLTokenOnetimeExchangeResponse struct {
	ACLToken *shared.ACLToken
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostACLTokenOnetimeExchangeResponse) GetACLToken() *shared.ACLToken {
	if o == nil {
		return nil
	}
	return o.ACLToken
}

func (o *PostACLTokenOnetimeExchangeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostACLTokenOnetimeExchangeResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *PostACLTokenOnetimeExchangeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostACLTokenOnetimeExchangeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
