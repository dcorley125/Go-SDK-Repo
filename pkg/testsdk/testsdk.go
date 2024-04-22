package testsdk

import (
	"github.com/dcorley125/Go-SDK-Repo/internal/configmanager"
	"github.com/dcorley125/Go-SDK-Repo/pkg/pets"
	"github.com/dcorley125/Go-SDK-Repo/pkg/testsdkconfig"
)

type TestSdk struct {
	Pets    *pets.PetsService
	manager *configmanager.ConfigManager
}

func NewTestSdk(config testsdkconfig.Config) *TestSdk {
	manager := configmanager.NewConfigManager(config)
	return &TestSdk{
		Pets:    pets.NewPetsService(manager),
		manager: manager,
	}
}

func (t *TestSdk) SetBaseUrl(baseUrl string) {
	t.manager.SetBaseUrl(baseUrl)
}

// c029837e0e474b76bc487506e8799df5e3335891efe4fb02bda7a1441840310c
