// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/hashicorp/terraform-provider-nomad/v2/internal/sdk/pkg/models/shared"
	"net/http"
)

type PostDeploymentAllocationHealthSecurity struct {
	XNomadToken string `security:"scheme,type=apiKey,subtype=header,name=X-Nomad-Token"`
}

func (o *PostDeploymentAllocationHealthSecurity) GetXNomadToken() string {
	if o == nil {
		return ""
	}
	return o.XNomadToken
}

type PostDeploymentAllocationHealthRequest struct {
	DeploymentAllocHealthRequest shared.DeploymentAllocHealthRequest `request:"mediaType=application/json"`
	// A Nomad ACL token.
	XNomadToken *string `header:"style=simple,explode=false,name=X-Nomad-Token"`
	// Deployment ID.
	DeploymentID string `pathParam:"style=simple,explode=false,name=deploymentID"`
	// Can be used to ensure operations are only run once.
	IdempotencyToken *string `queryParam:"style=form,explode=true,name=idempotency_token"`
	// Filters results based on the specified namespace.
	Namespace *string `queryParam:"style=form,explode=true,name=namespace"`
	// Filters results based on the specified region.
	Region *string `queryParam:"style=form,explode=true,name=region"`
}

func (o *PostDeploymentAllocationHealthRequest) GetDeploymentAllocHealthRequest() shared.DeploymentAllocHealthRequest {
	if o == nil {
		return shared.DeploymentAllocHealthRequest{}
	}
	return o.DeploymentAllocHealthRequest
}

func (o *PostDeploymentAllocationHealthRequest) GetXNomadToken() *string {
	if o == nil {
		return nil
	}
	return o.XNomadToken
}

func (o *PostDeploymentAllocationHealthRequest) GetDeploymentID() string {
	if o == nil {
		return ""
	}
	return o.DeploymentID
}

func (o *PostDeploymentAllocationHealthRequest) GetIdempotencyToken() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyToken
}

func (o *PostDeploymentAllocationHealthRequest) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *PostDeploymentAllocationHealthRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

type PostDeploymentAllocationHealthResponse struct {
	// HTTP response content type for this operation
	ContentType              string
	DeploymentUpdateResponse *shared.DeploymentUpdateResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostDeploymentAllocationHealthResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostDeploymentAllocationHealthResponse) GetDeploymentUpdateResponse() *shared.DeploymentUpdateResponse {
	if o == nil {
		return nil
	}
	return o.DeploymentUpdateResponse
}

func (o *PostDeploymentAllocationHealthResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostDeploymentAllocationHealthResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
