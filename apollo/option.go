package apollo

import "time"

type Options struct {
	File      string
	EndPoints string
	TTL       time.Duration
}

type Option func(o *Options)

// WithFile sets the local file path
func WithFile(file string) Option {
	return func(o *Options) { o.File = file }
}

// WithEndpoints sets the remote endpoints, pull remote config to local with ttl
// endpoints schema support etcd://localhost, http://localhost
func WithEndpoints(endpoints string) Option {
	return func(o *Options) { o.EndPoints = endpoints }
}

// WithTTL sets the TTL of the remote endpoints synchronization
func WithTTL(ttl int64) Option {
	return func(o *Options) { o.TTL = time.Duration(ttl) * time.Second }
}
