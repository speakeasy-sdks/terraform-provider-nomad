// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"nomad/internal/sdk/pkg/models/shared"
	"nomad/internal/sdk/pkg/utils"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// dev-agent
	"{scheme}://{address}:{port}/v1",
	// agent
	"{scheme}://{address}:{port}/v1",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          *shared.Security
	ServerURL         string
	ServerIndex       int
	ServerDefaults    []map[string]string
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], c.ServerDefaults[c.ServerIndex]
}

// Nomad
type Nomad struct {
	ACL         *acl
	Allocations *allocations
	Deployments *deployments
	Enterprise  *enterprise
	Evaluations *evaluations
	Jobs        *jobs
	Metrics     *metrics
	Namespaces  *namespaces
	Nodes       *nodes
	Operator    *operator
	Plugins     *plugins
	Regions     *regions
	Scaling     *scaling
	Search      *search
	Status      *status
	System      *system
	Variables   *variables
	Volumes     *volumes

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*Nomad)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Nomad) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Nomad) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *Nomad) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithAddress allows setting the $name variable for url substitution
func WithAddress(address string) SDKOption {
	return func(sdk *Nomad) {
		for idx := range sdk.sdkConfiguration.ServerDefaults {
			if _, ok := sdk.sdkConfiguration.ServerDefaults[idx]["address"]; !ok {
				continue
			}

			sdk.sdkConfiguration.ServerDefaults[idx]["address"] = fmt.Sprintf("%v", address)
		}
	}
}

// WithPort allows setting the $name variable for url substitution
func WithPort(port string) SDKOption {
	return func(sdk *Nomad) {
		for idx := range sdk.sdkConfiguration.ServerDefaults {
			if _, ok := sdk.sdkConfiguration.ServerDefaults[idx]["port"]; !ok {
				continue
			}

			sdk.sdkConfiguration.ServerDefaults[idx]["port"] = fmt.Sprintf("%v", port)
		}
	}
}

type ServerScheme string

const (
	ServerSchemeHTTPS ServerScheme = "https"
	ServerSchemeHTTP  ServerScheme = "http"
)

func (e ServerScheme) ToPointer() *ServerScheme {
	return &e
}

func (e *ServerScheme) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "https":
		fallthrough
	case "http":
		*e = ServerScheme(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ServerScheme: %v", v)
	}
}

// WithScheme allows setting the $name variable for url substitution
func WithScheme(scheme ServerScheme) SDKOption {
	return func(sdk *Nomad) {
		for idx := range sdk.sdkConfiguration.ServerDefaults {
			if _, ok := sdk.sdkConfiguration.ServerDefaults[idx]["scheme"]; !ok {
				continue
			}

			sdk.sdkConfiguration.ServerDefaults[idx]["scheme"] = fmt.Sprintf("%v", scheme)
		}
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Nomad) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *Nomad) {
		sdk.sdkConfiguration.Security = &security
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Nomad {
	sdk := &Nomad{
		sdkConfiguration: sdkConfiguration{
			Language:          "terraform",
			OpenAPIDocVersion: "1.1.4",
			SDKVersion:        "1.13.0",
			GenVersion:        "2.91.2",
			ServerDefaults: []map[string]string{
				{
					"address": "127.0.0.1",
					"port":    "4646",
					"scheme":  "http",
				},
				{
					"address": "127.0.0.1",
					"port":    "4646",
					"scheme":  "https",
				},
			},
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.ACL = newACL(sdk.sdkConfiguration)

	sdk.Allocations = newAllocations(sdk.sdkConfiguration)

	sdk.Deployments = newDeployments(sdk.sdkConfiguration)

	sdk.Enterprise = newEnterprise(sdk.sdkConfiguration)

	sdk.Evaluations = newEvaluations(sdk.sdkConfiguration)

	sdk.Jobs = newJobs(sdk.sdkConfiguration)

	sdk.Metrics = newMetrics(sdk.sdkConfiguration)

	sdk.Namespaces = newNamespaces(sdk.sdkConfiguration)

	sdk.Nodes = newNodes(sdk.sdkConfiguration)

	sdk.Operator = newOperator(sdk.sdkConfiguration)

	sdk.Plugins = newPlugins(sdk.sdkConfiguration)

	sdk.Regions = newRegions(sdk.sdkConfiguration)

	sdk.Scaling = newScaling(sdk.sdkConfiguration)

	sdk.Search = newSearch(sdk.sdkConfiguration)

	sdk.Status = newStatus(sdk.sdkConfiguration)

	sdk.System = newSystem(sdk.sdkConfiguration)

	sdk.Variables = newVariables(sdk.sdkConfiguration)

	sdk.Volumes = newVolumes(sdk.sdkConfiguration)

	return sdk
}
