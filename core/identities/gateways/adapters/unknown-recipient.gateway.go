package adapters

import (
	"errors"
	"grabit-cli/core/identities/gateways"
)

type UnknownRecipientGateway struct {
}

func (lig *UnknownRecipientGateway) FetchPublicKey(email string, response chan<- *gateways.FetchPublicKeyRequestResponse) error {
	go func() {
		response <- nil
	}()
	return errors.New("UNKNOWN_RECIPIENT")
}
