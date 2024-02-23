package gateways

import "grabit-cli/core/identities/models"

type RegisterIdentityRequest struct {
	Email    string
	Password string
}
type RegisterIdentityResponse struct {
	Identity models.Identity
}
type IdentityGateway interface {
	Register(request RegisterIdentityRequest) (*RegisterIdentityResponse, error)
}
