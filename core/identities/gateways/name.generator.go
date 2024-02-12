package gateways

type NameGenerator interface {
	Generate(seed string) string
}
