package adapters

import (
	"errors"
	"grabit-cli/core/identities/gateways"
)

type UnknownIdentityGateway struct {
}

func (lig *UnknownIdentityGateway) LoadCurrent(response chan<- *gateways.LoadIdentityResponse) error {
	go func() {
		response <- nil
	}()
	return errors.New("UNKNOWN IDENTITY")
}
