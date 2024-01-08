package gateways

import "grabit-cli/core/identities/models"

type LoadIdentityResponse struct {
	Identity models.Identity
}

type IdentityGateway interface {
	LoadCurrent(lir chan<- *LoadIdentityResponse) error
}
