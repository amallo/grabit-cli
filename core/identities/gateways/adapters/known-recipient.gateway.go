package adapters

import (
	"grabit-cli/core/identities/gateways"
)

type KnownRecipientGateway struct {
	WillLoadPublicKey string
}

func (lig *KnownRecipientGateway) FetchPublicKey(email string, response chan<- *gateways.FetchPublicKeyResponse) error {
	go func() {
		response <- &gateways.FetchPublicKeyResponse{PublicKey: lig.WillLoadPublicKey}
	}()
	return nil
}
