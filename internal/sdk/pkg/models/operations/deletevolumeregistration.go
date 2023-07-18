// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteVolumeRegistrationSecurity struct {
	XNomadToken string `security:"scheme,type=apiKey,subtype=header,name=X-Nomad-Token"`
}

type DeleteVolumeRegistrationRequest struct {
	// A Nomad ACL token.
	XNomadToken *string `header:"style=simple,explode=false,name=X-Nomad-Token"`
	// Used to force the de-registration of a volume.
	Force *string `queryParam:"style=form,explode=true,name=force"`
	// Can be used to ensure operations are only run once.
	IdempotencyToken *string `queryParam:"style=form,explode=true,name=idempotency_token"`
	// Filters results based on the specified namespace.
	Namespace *string `queryParam:"style=form,explode=true,name=namespace"`
	// Filters results based on the specified region.
	Region *string `queryParam:"style=form,explode=true,name=region"`
	// Volume unique identifier.
	VolumeID string `pathParam:"style=simple,explode=false,name=volumeId"`
}

type DeleteVolumeRegistrationResponse struct {
	ContentType string
	Headers     map[string][]string
	StatusCode  int
	RawResponse *http.Response
}