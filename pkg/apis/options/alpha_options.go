package options

type AlphaOptions struct {
	Upstreams Upstreams `json:"upstreams"`

	InjectRequestHeaders  []Header `json:"injectRequestHeaders"`
	InjectResponseHeaders []Header `json:"injectResponseHeaders"`
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
