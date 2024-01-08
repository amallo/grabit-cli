package gateways

type FetchPublicKeyRequestResponse struct {
	PublicKey string
}

type RecipientGateway interface {
	FetchPublicKey(email string, rir chan<- *FetchPublicKeyRequestResponse) error
}
