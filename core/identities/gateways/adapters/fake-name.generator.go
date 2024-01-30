package adapters

type FakeNameGenerator struct {
	WillGenerateName string
}

func (gen *FakeNameGenerator) Generate(seed string) string {
	return gen.WillGenerateName
}
