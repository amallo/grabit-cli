package adapters

import (
	"grabit-cli/core/identities/gateways"
	"grabit-cli/core/identities/models"
)

type KnownIdentityGateway struct {
	WillLoadIdentity *models.Identity
}

func (lig *KnownIdentityGateway) LoadCurrent(response chan<- *gateways.LoadIdentityResponse) error {
	go func() {

		response <- &gateways.LoadIdentityResponse{Identity: *lig.WillLoadIdentity}
	}()
	return nil
}
