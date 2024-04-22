package testsdkconfig

type Config struct {
	BaseUrl *string
}

func NewConfig() Config {
	baseUrl := DEFAULT_ENVIRONMENT
	newConfig := Config{
		BaseUrl: &baseUrl,
	}

	return newConfig
}

func (c *Config) SetBaseUrl(baseUrl string) {
	c.BaseUrl = &baseUrl
}
