package adapters

import (
	"grabit-cli/core/identities/gateways"
)

type KnownRecipientGateway struct {
	WillLoadPublicKey string
}

func (lig *KnownRecipientGateway) FetchPublicKey(email string) (*gateways.FetchPublicKeyResponse, error) {
	return &gateways.FetchPublicKeyResponse{PublicKey: lig.WillLoadPublicKey}, nil
}
