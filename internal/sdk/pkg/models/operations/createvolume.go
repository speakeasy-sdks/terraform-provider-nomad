// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/hashicorp/terraform-provider-nomad/v3/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateVolumeSecurity struct {
	XNomadToken string `security:"scheme,type=apiKey,subtype=header,name=X-Nomad-Token"`
}

func (o *CreateVolumeSecurity) GetXNomadToken() string {
	if o == nil {
		return ""
	}
	return o.XNomadToken
}

type CreateVolumeRequest struct {
	CSIVolumeCreateRequest shared.CSIVolumeCreateRequest `request:"mediaType=application/json"`
	// A Nomad ACL token.
	XNomadToken *string `header:"style=simple,explode=false,name=X-Nomad-Token"`
	// The action to perform on the Volume (create, detach, delete).
	Action string `pathParam:"style=simple,explode=false,name=action"`
	// Can be used to ensure operations are only run once.
	IdempotencyToken *string `queryParam:"style=form,explode=true,name=idempotency_token"`
	// Filters results based on the specified namespace.
	Namespace *string `queryParam:"style=form,explode=true,name=namespace"`
	// Filters results based on the specified region.
	Region *string `queryParam:"style=form,explode=true,name=region"`
	// Volume unique identifier.
	VolumeID string `pathParam:"style=simple,explode=false,name=volumeId"`
}

func (o *CreateVolumeRequest) GetCSIVolumeCreateRequest() shared.CSIVolumeCreateRequest {
	if o == nil {
		return shared.CSIVolumeCreateRequest{}
	}
	return o.CSIVolumeCreateRequest
}

func (o *CreateVolumeRequest) GetXNomadToken() *string {
	if o == nil {
		return nil
	}
	return o.XNomadToken
}

func (o *CreateVolumeRequest) GetAction() string {
	if o == nil {
		return ""
	}
	return o.Action
}

func (o *CreateVolumeRequest) GetIdempotencyToken() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyToken
}

func (o *CreateVolumeRequest) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *CreateVolumeRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *CreateVolumeRequest) GetVolumeID() string {
	if o == nil {
		return ""
	}
	return o.VolumeID
}

type CreateVolumeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateVolumeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateVolumeResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *CreateVolumeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateVolumeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
