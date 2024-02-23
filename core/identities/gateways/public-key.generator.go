package gateways

type PublicKeyGenerator interface {
	Generate(email string, name string) string
}
