package pets

import (
	"context"
	restClient "github.com/dcorley125/Go-SDK-Repo/internal/clients/rest"
	"github.com/dcorley125/Go-SDK-Repo/internal/clients/rest/httptransport"
	"github.com/dcorley125/Go-SDK-Repo/internal/configmanager"
	"github.com/dcorley125/Go-SDK-Repo/pkg/shared"
	"github.com/dcorley125/Go-SDK-Repo/pkg/testsdkconfig"
)

type PetsService struct {
	manager *configmanager.ConfigManager
}

func NewPetsService(manager *configmanager.ConfigManager) *PetsService {
	return &PetsService{
		manager: manager,
	}
}

func (api *PetsService) getConfig() *testsdkconfig.Config {
	return api.manager.GetPets()
}

func (api *PetsService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *PetsService) ListPets(ctx context.Context, params ListPetsRequestParams) (*shared.TestSdkResponse[[]Pet], *shared.TestSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[[]Pet](config)

	request := httptransport.NewRequest(ctx, "GET", "/pets", config)

	request.Options = params

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewTestSdkError[[]Pet](err)
	}

	return shared.NewTestSdkResponse[[]Pet](resp), nil
}

func (api *PetsService) CreatePets(ctx context.Context, pet Pet) (*shared.TestSdkResponse[any], *shared.TestSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[any](config)

	request := httptransport.NewRequest(ctx, "POST", "/pets", config)

	request.Body = pet

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewTestSdkError[any](err)
	}

	return shared.NewTestSdkResponse[any](resp), nil
}

func (api *PetsService) ShowPetById(ctx context.Context, petId string) (*shared.TestSdkResponse[Pet], *shared.TestSdkError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Pet](config)

	request := httptransport.NewRequest(ctx, "GET", "/pets/{petId}", config)

	request.SetPathParam("petId", petId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewTestSdkError[Pet](err)
	}

	return shared.NewTestSdkResponse[Pet](resp), nil
}
