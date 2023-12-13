package apollo

type Config struct{}

func (c *Config) Get(key string) (interface{}, error) {
	return nil, nil
}

func (c *Config) String() string {
	return ""
}
