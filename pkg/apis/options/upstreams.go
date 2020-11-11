package options

import "time"

// Upstreams is a collection of definitions for upstream servers.
type Upstreams []Upstream

// Upstream represents the configuration for an upstream server.
// Requests will be proxied to this upstream if the path matches the request path.
type Upstream struct {
	// ID should be a unique identifier for the upstream.
	// This value is required for all upstreams.
	ID string `yaml:"id,omitempty"`

	// Path is used to map requests to the upstream server.
	// The closest match will take precedence and all Paths must be unique.
	Path string `yaml:"path,omitempty"`

	// The URI of the upstream server. This may be an HTTP(S) server of a File
	// based URL. It may include a path, in which case all requests will be served
	// under that path.
	// Eg:
	// - http://localhost:8080
	// - https://service.localhost
	// - https://service.localhost/path
	// - file://host/path
	// If the URI's path is "/base" and the incoming request was for "/dir",
	// the upstream request will be for "/base/dir".
	URI string `yaml:"uri,omitempty"`

	// InsecureSkipTLSVerify will skip TLS verification of upstream HTTPS hosts.
	// This option is insecure and will allow potential Man-In-The-Middle attacks
	// betweem OAuth2 Proxy and the usptream server.
	// Defaults to false.
	InsecureSkipTLSVerify bool `yaml:"insecureSkipTLSVerify,omitempty"`

	// Static will make all requests to this upstream have a static response.
	// The response will have a body of "Authenticated" and a response code
	// matching StaticCode.
	// If StaticCode is not set, the response will return a 200 response.
	Static bool `yaml:"static,omitempty"`

	// StaticCode determines the response code for the Static response.
	// This option can only be used with Static enabled.
	StaticCode *int `yaml:"staticCode,omitempty"`

	// FlushInterval is the period between flushing the response buffer when
	// streaming response from the upstream.
	// Defaults to 1 second.
	FlushInterval *time.Duration `yaml:"flushInterval,omitempty"`

	// PassHostHeader determines whether the request host header should be proxied
	// to the upstream server.
	// Defaults to true.
	PassHostHeader *bool `yaml:"passHostHeader,omitempty"`

	// ProxyWebSockets enables proxying of websockets to upstream servers
	// Defaults to true.
	ProxyWebSockets *bool `yaml:"proxyWebSockets,omitempty"`
}
