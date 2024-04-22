package shared

import (
	"github.com/dcorley125/Go-SDK-Repo/internal/clients/rest/httptransport"
)

type TestSdkError struct {
	Err      error
	Body     []byte
	Metadata TestSdkErrorMetadata
}

type TestSdkErrorMetadata struct {
	Headers    map[string]string
	StatusCode int
}

func NewTestSdkError[T any](transportError *httptransport.ErrorResponse[T]) *TestSdkError {
	return &TestSdkError{
		Err:  transportError.GetError(),
		Body: transportError.GetBody(),
		Metadata: TestSdkErrorMetadata{
			StatusCode: transportError.GetStatusCode(),
			Headers:    transportError.GetHeaders(),
		},
	}
}

func (e *TestSdkError) Error() string {
	return e.Err.Error()
}
