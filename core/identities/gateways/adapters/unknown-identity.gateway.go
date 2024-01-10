package adapters

import (
	"errors"
	"grabit-cli/core/identities/gateways"
)

type UnknownIdentityGateway struct {
}

func (lig *UnknownIdentityGateway) LoadCurrent() (*gateways.LoadIdentityResponse, error) {
	return nil, errors.New("UNKNOWN IDENTITY")
}

func (lig *UnknownIdentityGateway) Register(request gateways.RegisterIdentityRequest) error {
	return nil
}
