// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/hashicorp/terraform-provider-nomad/v5/internal/sdk/models/shared"
	"net/http"
)

type GetMetricsSummaryRequest struct {
	// The format the user requested for the metrics summary (e.g. prometheus)
	Format *string `queryParam:"style=form,explode=true,name=format"`
}

func (o *GetMetricsSummaryRequest) GetFormat() *string {
	if o == nil {
		return nil
	}
	return o.Format
}

type GetMetricsSummaryResponse struct {
	// HTTP response content type for this operation
	ContentType    string
	MetricsSummary *shared.MetricsSummary
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetMetricsSummaryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMetricsSummaryResponse) GetMetricsSummary() *shared.MetricsSummary {
	if o == nil {
		return nil
	}
	return o.MetricsSummary
}

func (o *GetMetricsSummaryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMetricsSummaryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
