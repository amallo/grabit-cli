package adapters

import (
	"errors"
	"grabit-cli/core/identities/gateways"
)

type FakeIdentityGateway struct {
	WillHaveIdentityEmail       string
	WillFaildentityRegistration error
}

func (lig *FakeIdentityGateway) LoadCurrent(response chan<- *gateways.LoadIdentityResponse) error {
	go func() {
		if len(lig.WillHaveIdentityEmail) > 0 {
			response <- &gateways.LoadIdentityResponse{Email: lig.WillHaveIdentityEmail}
			return
		}
		response <- nil
	}()
	if len(lig.WillHaveIdentityEmail) == 0 {
		return errors.New("IDENTITY NOT FOUND")
	}
	return nil
}

func (lig *FakeIdentityGateway) Register(request gateways.RegisterIdentityRequest) error {
	return lig.WillFaildentityRegistration
}
