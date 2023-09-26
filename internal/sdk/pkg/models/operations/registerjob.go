// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"nomad/internal/sdk/pkg/models/shared"
)

type RegisterJobSecurity struct {
	XNomadToken string `security:"scheme,type=apiKey,subtype=header,name=X-Nomad-Token"`
}

type RegisterJobRequest struct {
	JobRegisterRequest shared.JobRegisterRequest `request:"mediaType=application/json"`
	// A Nomad ACL token.
	XNomadToken *string `header:"style=simple,explode=false,name=X-Nomad-Token"`
	// Can be used to ensure operations are only run once.
	IdempotencyToken *string `queryParam:"style=form,explode=true,name=idempotency_token"`
	// Filters results based on the specified namespace.
	Namespace *string `queryParam:"style=form,explode=true,name=namespace"`
	// Filters results based on the specified region.
	Region *string `queryParam:"style=form,explode=true,name=region"`
}

type RegisterJobResponse struct {
	// HTTP response content type for this operation
	ContentType         string
	Headers             map[string][]string
	JobRegisterResponse *shared.JobRegisterResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}
