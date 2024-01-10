package adapters

import (
	"errors"
	"grabit-cli/core/identities/gateways"
)

type UnknownRecipientGateway struct {
}

func (lig *UnknownRecipientGateway) FetchPublicKey(email string) (*gateways.FetchPublicKeyResponse, error) {
	return nil, errors.New("UNKNOWN_RECIPIENT")
}
