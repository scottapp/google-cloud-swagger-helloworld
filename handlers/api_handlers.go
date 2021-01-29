package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/scottapp/google-cloud-swagger-helloworld/rest/restapi/operations"
)

type ApiHandler struct {
}

func NewApiHandler() (*ApiHandler, error) {
	h := ApiHandler{}
	return &h, nil
}

func (h *ApiHandler) GetGreetingHandler(params operations.GetGreetingParams) middleware.Responder {
	var name string
	if params.Name == nil {
		name = "World"
	} else {
		name = *params.Name
	}
	return operations.NewGetGreetingOK().WithPayload("Hello " + name)
}
