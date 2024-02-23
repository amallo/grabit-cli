package adapters

import (
	"grabit-cli/core/identities/gateways"
	"grabit-cli/core/identities/models"
)

type FakeIdentityGateway struct {
	WillRegisterIdentity models.Identity
}

func NewFakeIdentityGateway() FakeIdentityGateway {
	return FakeIdentityGateway{}
}
func (g FakeIdentityGateway) Register(request gateways.RegisterIdentityRequest) (*gateways.RegisterIdentityResponse, error) {
	return &gateways.RegisterIdentityResponse{Identity: g.WillRegisterIdentity}, nil
}
