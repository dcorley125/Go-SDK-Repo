package configmanager

import "github.com/dcorley125/Go-SDK-Repo/pkg/testsdkconfig"

type ConfigManager struct {
	Pets testsdkconfig.Config
}

func NewConfigManager(config testsdkconfig.Config) *ConfigManager {
	return &ConfigManager{
		Pets: config,
	}
}

func (c *ConfigManager) SetBaseUrl(baseUrl string) {
	c.Pets.SetBaseUrl(baseUrl)
}

func (c *ConfigManager) GetPets() *testsdkconfig.Config {
	return &c.Pets
}
