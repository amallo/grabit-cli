package gateways

type FetchPublicKeyResponse struct {
	PublicKey string
}

type RecipientGateway interface {
	FetchPublicKey(email string, rir chan<- *FetchPublicKeyResponse) error
}
