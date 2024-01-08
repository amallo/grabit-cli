package adapters

import (
	"grabit-cli/core/identities/gateways"
)

type KnownRecipientGateway struct {
	WillLoadPublicKey string
}

func (lig *KnownRecipientGateway) FetchPublicKey(email string, response chan<- *gateways.FetchPublicKeyRequestResponse) error {
	go func() {
		response <- &gateways.FetchPublicKeyRequestResponse{PublicKey: lig.WillLoadPublicKey}
	}()
	return nil
}
