package options

type AlphaOptions struct {
	Upstreams Upstreams `yaml:"upstreams,omitempty"`

	InjectRequestHeaders  []Header `yaml:"injectRequestHeaders,omitempty"`
	InjectResponseHeaders []Header `yaml:"injectResponseHeaders,omitempty"`
}

func (a *AlphaOptions) MergeInto(opts *Options) {
	opts.UpstreamServers = a.Upstreams
	opts.InjectRequestHeaders = a.InjectRequestHeaders
	opts.InjectResponseHeaders = a.InjectResponseHeaders
}

func (a *AlphaOptions) ExtractFrom(opts *Options) {
	a.Upstreams = opts.UpstreamServers
	a.InjectRequestHeaders = opts.InjectRequestHeaders
	a.InjectResponseHeaders = opts.InjectResponseHeaders
}
