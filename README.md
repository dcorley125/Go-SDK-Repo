# TestSdk Go SDK 1.2.0

A Go SDK for TestSdk.

- API version: 1.0.0
- SDK version: 1.2.0

## Table of Contents

- [Services](#services)

## Services

### PetsService

#### ListPets

List all pets

```GO
import (
  "fmt"
  "encoding/json"
  "github.com/dcorley125/Go-SDK-Repo/pkg/testsdkconfig"
  "github.com/dcorley125/Go-SDK-Repo/pkg/testsdk"
  "github.com/dcorley125/Go-SDK-Repo/pkg/pets"
)

config := testsdkconfig.NewConfig()
client := testsdk.NewTestSdk(config)


params := pets.ListPetsRequestParams{}


response, err := client.Pets.ListPets(context.Background(), params)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

#### CreatePets

Create a pet

```GO
import (
  "fmt"
  "encoding/json"
  "github.com/dcorley125/Go-SDK-Repo/pkg/testsdkconfig"
  "github.com/dcorley125/Go-SDK-Repo/pkg/testsdk"
  "github.com/dcorley125/Go-SDK-Repo/pkg/pets"
)

config := testsdkconfig.NewConfig()
client := testsdk.NewTestSdk(config)


request := pets.Pet{}
request.SetId(int64(123))
request.SetName("Name")

response, err := client.Pets.CreatePets(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

#### ShowPetById

Info for a specific pet

```GO
import (
  "fmt"
  "encoding/json"
  "github.com/dcorley125/Go-SDK-Repo/pkg/testsdkconfig"
  "github.com/dcorley125/Go-SDK-Repo/pkg/testsdk"
)

config := testsdkconfig.NewConfig()
client := testsdk.NewTestSdk(config)

response, err := client.Pets.ShowPetById(context.Background(), "petId")
if err != nil {
  panic(err)
}

fmt.Print(response)
```
