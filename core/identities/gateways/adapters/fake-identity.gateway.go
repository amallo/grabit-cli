package adapters

import (
	"errors"
	"grabit-cli/core/identities/gateways"
)

type FakeIdentityGateway struct {
	WillHaveIdentityEmail       string
	WillFaildentityRegistration error
}

func (lig *FakeIdentityGateway) LoadCurrent() (*gateways.LoadIdentityResponse, error) {
	if len(lig.WillHaveIdentityEmail) == 0 {
		return nil, errors.New("IDENTITY NOT FOUND")
	}
	return &gateways.LoadIdentityResponse{Email: lig.WillHaveIdentityEmail}, nil
}

func (lig *FakeIdentityGateway) Register(request gateways.RegisterIdentityRequest) error {
	return lig.WillFaildentityRegistration
}
