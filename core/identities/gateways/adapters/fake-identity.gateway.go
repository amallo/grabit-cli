package adapters

import (
	"errors"
	"grabit-cli/core/identities/gateways"
	"grabit-cli/core/identities/models"
)

type FakeIdentityGateway struct {
	WillHaveIdentityEmail       string
	WillHaveIdentityName        string
	WillFaildentityRegistration error
	WillHaveReceiverIdentity    models.Identity
	WillHaveSenderIdentity      models.Identity

	WillHaveIdentities map[string]string
}

func NewFakeIdentityGateway(identities map[string]string) FakeIdentityGateway {
	return FakeIdentityGateway{
		WillHaveIdentities: identities,
	}
}

func (lig *FakeIdentityGateway) LoadCurrent(email string) (*gateways.LoadIdentityResponse, error) {
	name, ok := lig.WillHaveIdentities[email]
	if ok {
		return &gateways.LoadIdentityResponse{Name: name}, nil
	}
	return nil, errors.New("IDENTITY NOT FOUND")
}

func (lig *FakeIdentityGateway) Register(request gateways.RegisterIdentityRequest) error {
	return lig.WillFaildentityRegistration
}
