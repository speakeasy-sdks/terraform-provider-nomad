// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"nomad/internal/sdk/pkg/models/shared"
)

type PostQuotaSpecSecurity struct {
	XNomadToken string `security:"scheme,type=apiKey,subtype=header,name=X-Nomad-Token"`
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

type PostQuotaSpecResponse struct {
	ContentType string
	Headers     map[string][]string
	StatusCode  int
	RawResponse *http.Response
}
