package shared

import (
	"github.com/dcorley125/Go-SDK-Repo/internal/clients/rest/httptransport"
)

type TestSdkResponse[T any] struct {
	Data     T
	Metadata TestSdkResponseMetadata
}

type TestSdkResponseMetadata struct {
	Headers    map[string]string
	StatusCode int
}

func NewTestSdkResponse[T any](resp *httptransport.Response[T]) *TestSdkResponse[T] {
	return &TestSdkResponse[T]{
		Data: resp.Data,
		Metadata: TestSdkResponseMetadata{
			StatusCode: resp.StatusCode,
			Headers:    resp.Headers,
		},
	}
}
