package identities

import (
	"errors"
	"grabit-cli/core/identities/gateways/adapters"
	"grabit-cli/core/identities/models"
	"grabit-cli/core/identities/usecases"
	"testing"

	. "github.com/franela/goblin"
)

func TestIdentities(t *testing.T) {
	g := Goblin(t)

	g.Describe("audie manage its identity", func() {
		g.BeforeEach(func() {

		})
		g.It("Audie registers its identity first time", func() {
			identityGateway := adapters.FakeIdentityGateway{}
			identityGateway.WillRegisterIdentity = models.Identity{Email: "audie@mail.com", Key: "audie-pub-key", Name: "audie-baba-le-hunt"}
			useCase := usecases.NewRegisterIdentityUseCase(identityGateway)
			result, err := useCase.Execute(usecases.RegisterIdentityArgs{Email: "audie@mail.com", Password: "prune"})
			g.Assert(err).IsNil()
			g.Assert(result.Key).Equal("audie-pub-key")
			g.Assert(result.Name).Equal("audie-baba-le-hunt")
			g.Assert(result.Email).Equal("audie@mail.com")
		})

		g.It("Audie fails to register identity", func() {
			identityGateway := adapters.FailRegisterIdentityGateway{}
			identityGateway.WillFailRegisteringIdentity = errors.New("Failure")
			useCase := usecases.NewRegisterIdentityUseCase(identityGateway)
			result, err := useCase.Execute(usecases.RegisterIdentityArgs{Email: "audie@mail.com", Password: "prune"})
			g.Assert(result).IsNil()
			g.Assert(err.Code()).Equal(models.ErrCannotRegisterIdentity)
		})

	})
}
