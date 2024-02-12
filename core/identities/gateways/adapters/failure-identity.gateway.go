package adapters

import (
	"grabit-cli/core/identities/gateways"
)

type FailureIdentityGateway struct {
	WillFailLoadCurrent error
	WillFailRegister    error
}

func (g *FailureIdentityGateway) LoadCurrent(email string) (*gateways.LoadIdentityResponse, error) {
	return nil, g.WillFailLoadCurrent
}

func (g *FailureIdentityGateway) Register(request gateways.RegisterIdentityRequest) error {
	return g.WillFailRegister
}
