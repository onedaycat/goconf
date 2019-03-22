package goconf

type Option func(o *opts)

type opts struct {
	yaml      bool
	yamlPaths []string
	yamlBytes []byte
	env       bool
	envPrefix string
}

func WithYaml(paths ...string) Option {
	return func(o *opts) {
		o.yaml = true
		o.yamlPaths = paths
	}
}

func WithYamlFromBytes(yaml []byte) Option {
	return func(o *opts) {
		o.yaml = true
		o.yamlBytes = yaml
	}
}

func WithEnv(prefix string) Option {
	return func(o *opts) {
		o.env = true
		o.envPrefix = prefix
	}
}
