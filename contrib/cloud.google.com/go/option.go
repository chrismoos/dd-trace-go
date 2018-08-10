package cloud

import "context"

type config struct {
	serviceName string
	ctx         context.Context
}

func newConfig(options ...Option) *config {
	cfg := &config{
		ctx: context.Background(),
	}
	for _, opt := range options {
		opt(cfg)
	}
	return cfg
}

// An Option customizes the config.
type Option func(*config)

// WithContext sets the context in the config. This can be used to set span
// parents or pass a context through to the underlying client constructor.
func WithContext(ctx context.Context) Option {
	return func(cfg *config) {
		cfg.ctx = ctx
	}
}

// WithServiceName sets the service name in the config. The default service
// name is inferred from the API definitions based on the http request route.
func WithServiceName(serviceName string) Option {
	return func(cfg *config) {
		cfg.serviceName = serviceName
	}
}