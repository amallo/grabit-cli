package adapters

import (
	"grabit-cli/core/identities/gateways"
)

type FailRegisterIdentityGateway struct {
	WillFailRegisteringIdentity error
}

func (g FailRegisterIdentityGateway) Register(request gateways.RegisterIdentityRequest) (*gateways.RegisterIdentityResponse, error) {
	return nil, g.WillFailRegisteringIdentity
}
