package identities

import (
	"errors"
	"grabit-cli/core/identities/gateways/adapters"
	"grabit-cli/core/identities/usecases"
	message_adapters "grabit-cli/core/message/gateways/adapters"
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)
	var fakeIdentityGateway adapters.FakeIdentityGateway
	var fakeNameGenerator adapters.FakeNameGenerator
	var fakeMessageGateway message_adapters.FakeMessageGateway
	var identities map[string]string
	g.Describe("audie registers a new identity", func() {
		g.BeforeEach(func() {
			identities = make(map[string]string)
			fakeIdentityGateway = adapters.NewFakeIdentityGateway(identities)
			fakeNameGenerator = adapters.FakeNameGenerator{WillGenerateName: "madfu"}
			fakeMessageGateway = message_adapters.NewFakeMessageGateway()
		})
		g.It("Successfully registers a new identity", func() {
			useCase := usecases.NewRegisterIdentityUseCase(&fakeIdentityGateway, &fakeNameGenerator, &fakeMessageGateway)
			params := usecases.RegisterIdentityParams{Email: "audie@foo.com", PassPhrase: "lets-go-deeper"}
			result, err := useCase.Execute(params)

			g.Assert(err).IsNil()
			g.Assert(result.Name).Equal("madfu")
		})

		g.It("Fails to register email twice", func() {
			identities["audie@foo.com"] = "audie"
			useCase := usecases.NewRegisterIdentityUseCase(&fakeIdentityGateway, &fakeNameGenerator, &fakeMessageGateway)
			params := usecases.RegisterIdentityParams{Email: "audie@foo.com", PassPhrase: "lets-go-deeper"}
			_, err := useCase.Execute(params)

			g.Assert(err.Error()).Equal("ALREADY_HAVE_IDENTITY")
		})
		g.It("Fails to register identity for unknown reason", func() {
			failureIdentityGateway := adapters.FailureIdentityGateway{WillFailLoadCurrent: errors.New("No id found"), WillFailRegister: errors.New("Registration failure")}
			useCase := usecases.NewRegisterIdentityUseCase(&failureIdentityGateway, &fakeNameGenerator, &fakeMessageGateway)
			params := usecases.RegisterIdentityParams{Email: "audie@foo.com", PassPhrase: "lets-go-deeper"}
			_, err := useCase.Execute(params)

			g.Assert(err.Error()).Equal("IDENTITY_REGISTRATION_FAILED")
		})

	})
}
