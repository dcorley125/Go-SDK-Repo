package rest

import (
	"github.com/dcorley125/Go-SDK-Repo/internal/clients/rest/handlers"
	"github.com/dcorley125/Go-SDK-Repo/internal/clients/rest/hooks"
	"github.com/dcorley125/Go-SDK-Repo/internal/clients/rest/httptransport"
	"github.com/dcorley125/Go-SDK-Repo/pkg/testsdkconfig"
)

type RestClient[T any] struct {
	handlers *handlers.HandlerChain[T]
}

func NewRestClient[T any](config testsdkconfig.Config) *RestClient[T] {
	defaultHeadersHandler := handlers.NewDefaultHeadersHandler[T]()
	retryHandler := handlers.NewRetryHandler[T]()
	responseValidationHandler := handlers.NewResponseValidationHandler[T]()
	unmarshalHandler := handlers.NewUnmarshalHandler[T]()
	requestValidationHandler := handlers.NewRequestValidationHandler[T]()
	hookHandler := handlers.NewHookHandler[T](hooks.NewDefaultHook())
	terminatingHandler := handlers.NewTerminatingHandler[T]()

	handlers := handlers.BuildHandlerChain[T]().
		AddHandler(defaultHeadersHandler).
		AddHandler(retryHandler).
		AddHandler(responseValidationHandler).
		AddHandler(unmarshalHandler).
		AddHandler(requestValidationHandler).
		AddHandler(hookHandler).
		AddHandler(terminatingHandler)

	return &RestClient[T]{
		handlers: handlers,
	}
}

func (client *RestClient[T]) Call(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[T]) {
	return client.handlers.CallApi(request)
}
