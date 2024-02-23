package adapters

import (
	"grabit-cli/core/identities/gateways"
	"grabit-cli/core/identities/models"
)

type FsIdentityGateway struct {
	publicKeyGenerator gateways.PublicKeyGenerator
}

func NewFsIdentityGenerator(publicKeyGenerator gateways.PublicKeyGenerator) FsIdentityGateway {
	return FsIdentityGateway{publicKeyGenerator: publicKeyGenerator}
}

func (g FsIdentityGateway) Register(request gateways.RegisterIdentityRequest) (*gateways.RegisterIdentityResponse, error) {
	return &gateways.RegisterIdentityResponse{Identity: models.Identity{Email: "audie@app2.io", Name: "audie-le-hunt", Key: "key"}}, nil
}
