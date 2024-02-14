package identities

import (
	"errors"
	"grabit-cli/core/identities/gateways/adapters"
	"grabit-cli/core/identities/models"
	"grabit-cli/core/identities/usecases"
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)
	var fakeIdentityGateway adapters.FakeIdentityGateway
	var fakeNameGenerator adapters.FakeNameGenerator
	var identities map[string]string
	g.Describe("audie registers a new identity", func() {
		g.BeforeEach(func() {
			identities = make(map[string]string)
			fakeIdentityGateway = adapters.NewFakeIdentityGateway(identities)
			fakeNameGenerator = adapters.FakeNameGenerator{WillGenerateName: "madfu"}
		})
		g.It("Successfully registers a new identity", func() {
			useCase := usecases.NewRegisterIdentityUseCase(&fakeIdentityGateway, &fakeNameGenerator)
			params := usecases.RegisterIdentityParams{Email: "audie@foo.com", PassPhrase: "lets-go-deeper"}
			result, err := useCase.Execute(params)

			g.Assert(err).IsNil()
			g.Assert(result.Name).Equal("madfu")
		})

		g.It("Fails to register identity twice", func() {
			identities["audie@foo.com"] = "audie"
			useCase := usecases.NewRegisterIdentityUseCase(&fakeIdentityGateway, &fakeNameGenerator)
			params := usecases.RegisterIdentityParams{Email: "audie@foo.com", PassPhrase: "lets-go-deeper"}
			_, err := useCase.Execute(params)

			g.Assert(err.Code()).Equal(models.ErrIdentityAlreadyRegistered)
		})
		g.It("Fails to register identity for unknown reason", func() {
			fakeIdentityGateway.WillFaildentityRegistration = errors.New("FAIL")
			useCase := usecases.NewRegisterIdentityUseCase(&fakeIdentityGateway, &fakeNameGenerator)
			params := usecases.RegisterIdentityParams{Email: "audie@foo.com", PassPhrase: "lets-go-deeper"}
			_, err := useCase.Execute(params)
			g.Assert(err.Code()).Equal(models.ErrCannotRegisterIdentity)
		})

	})
}
