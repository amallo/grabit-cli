package gateways

type NameIdentityGenerator interface {
	Generate(seed string) string
}
