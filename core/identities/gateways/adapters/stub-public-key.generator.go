package adapters

type StubPublicKeyGenerator struct {
	WillGenerateKey string
}

func (k StubPublicKeyGenerator) Generate(email string, name string) string {
	return k.WillGenerateKey
}
