// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/hashicorp/terraform-provider-nomad/v2/internal/sdk/pkg/models/shared"
	"net/http"
)

type PostQuotaSpecSecurity struct {
	XNomadToken string `security:"scheme,type=apiKey,subtype=header,name=X-Nomad-Token"`
}

func (o *PostQuotaSpecSecurity) GetXNomadToken() string {
	if o == nil {
		return ""
	}
	return o.XNomadToken
}

type PostQuotaSpecRequest struct {
	QuotaSpec shared.QuotaSpec `request:"mediaType=application/json"`
	// A Nomad ACL token.
	XNomadToken *string `header:"style=simple,explode=false,name=X-Nomad-Token"`
	// Can be used to ensure operations are only run once.
	IdempotencyToken *string `queryParam:"style=form,explode=true,name=idempotency_token"`
	// Filters results based on the specified namespace.
	Namespace *string `queryParam:"style=form,explode=true,name=namespace"`
	// Filters results based on the specified region.
	Region *string `queryParam:"style=form,explode=true,name=region"`
	// The quota spec identifier.
	SpecName string `pathParam:"style=simple,explode=false,name=specName"`
}

func (o *PostQuotaSpecRequest) GetQuotaSpec() shared.QuotaSpec {
	if o == nil {
		return shared.QuotaSpec{}
	}
	return o.QuotaSpec
}

func (o *PostQuotaSpecRequest) GetXNomadToken() *string {
	if o == nil {
		return nil
	}
	return o.XNomadToken
}

func (o *PostQuotaSpecRequest) GetIdempotencyToken() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyToken
}

func (o *PostQuotaSpecRequest) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *PostQuotaSpecRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *PostQuotaSpecRequest) GetSpecName() string {
	if o == nil {
		return ""
	}
	return o.SpecName
}

type PostQuotaSpecResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostQuotaSpecResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostQuotaSpecResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *PostQuotaSpecResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostQuotaSpecResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
