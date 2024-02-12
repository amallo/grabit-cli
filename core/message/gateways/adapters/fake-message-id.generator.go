package adapters

type FakeMessageIdGenerator struct {
	WillGenerateId string
}

func (g *FakeMessageIdGenerator) Generate() string {
	return g.WillGenerateId
}
