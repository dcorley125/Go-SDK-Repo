package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/dcorley125/Go-SDK-Repo/pkg/pets"
	"github.com/dcorley125/Go-SDK-Repo/pkg/testsdk"
	"github.com/dcorley125/Go-SDK-Repo/pkg/testsdkconfig"
)

func main() {
	loadEnv()

	config := testsdkconfig.NewConfig()
	client := testsdk.NewTestSdk(config)

	params := pets.ListPetsRequestParams{}

	response, err := client.Pets.ListPets(context.Background(), params)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", response)
}

func loadEnv() error {
	file, err := os.Open(".env")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
