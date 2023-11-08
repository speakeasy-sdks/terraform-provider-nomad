// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/hashicorp/terraform-provider-nomad/v2/internal/sdk/pkg/models/shared"
	"net/http"
)

type PostDeploymentPromoteSecurity struct {
	XNomadToken string `security:"scheme,type=apiKey,subtype=header,name=X-Nomad-Token"`
}

func (o *PostDeploymentPromoteSecurity) GetXNomadToken() string {
	if o == nil {
		return ""
	}
	return o.XNomadToken
}

type PostDeploymentPromoteRequest struct {
	DeploymentPromoteRequest shared.DeploymentPromoteRequest `request:"mediaType=application/json"`
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

func (o *PostDeploymentPromoteRequest) GetDeploymentPromoteRequest() shared.DeploymentPromoteRequest {
	if o == nil {
		return shared.DeploymentPromoteRequest{}
	}
	return o.DeploymentPromoteRequest
}

func (o *PostDeploymentPromoteRequest) GetXNomadToken() *string {
	if o == nil {
		return nil
	}
	return o.XNomadToken
}

func (o *PostDeploymentPromoteRequest) GetDeploymentID() string {
	if o == nil {
		return ""
	}
	return o.DeploymentID
}

func (o *PostDeploymentPromoteRequest) GetIdempotencyToken() *string {
	if o == nil {
		return nil
	}
	return o.IdempotencyToken
}

func (o *PostDeploymentPromoteRequest) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *PostDeploymentPromoteRequest) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

type PostDeploymentPromoteResponse struct {
	// HTTP response content type for this operation
	ContentType              string
	DeploymentUpdateResponse *shared.DeploymentUpdateResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostDeploymentPromoteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostDeploymentPromoteResponse) GetDeploymentUpdateResponse() *shared.DeploymentUpdateResponse {
	if o == nil {
		return nil
	}
	return o.DeploymentUpdateResponse
}

func (o *PostDeploymentPromoteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostDeploymentPromoteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
