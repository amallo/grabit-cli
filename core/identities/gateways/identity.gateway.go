package gateways

import "grabit-cli/core/identities/models"

type LoadIdentityResponse struct {
	Name string
}

type RegisterIdentityRequest struct {
	Email      string
	Name       string
	PassPhrase string
}
type RegisteredIdentityResponse struct {
	Identity models.Identity
}

type IdentityGateway interface {
	LoadCurrent(name string) (*LoadIdentityResponse, error)
	Register(request RegisterIdentityRequest) error
}
