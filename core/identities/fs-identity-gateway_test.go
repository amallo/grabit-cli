package identities

import (
	"grabit-cli/core/identities/gateways"
	"grabit-cli/core/identities/gateways/adapters"
	"testing"

	. "github.com/franela/goblin"
)

func TestFsIdentityGateway(t *testing.T) {
	g := Goblin(t)

	g.Describe("audie manage its identity", func() {
		g.BeforeEach(func() {

		})
		g.It("Audie registers its identity", func() {
			publicKeyGenerator := adapters.StubPublicKeyGenerator{WillGenerateKey: "key"}
			fsIdentityGateway := adapters.NewFsIdentityGenerator(publicKeyGenerator)
			r, err := fsIdentityGateway.Register(gateways.RegisterIdentityRequest{Email: "audie@app2.io", Password: "prune"})
			g.Assert(err).IsNil()
			g.Assert(r.Identity.Key).Equal("key")
			g.Assert(r.Identity.Name).Equal("audie-le-hunt")
		})
	})
}
