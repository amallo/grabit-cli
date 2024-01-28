package identities

import (
	"errors"
	"grabit-cli/core/identities/gateways/adapters"
	"grabit-cli/core/identities/usecases"
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)
	var fakeIdentityGateway *adapters.FakeIdentityGateway
	g.Describe("audie registers a new identity", func() {
		g.BeforeEach(func() {
			fakeIdentityGateway = &adapters.FakeIdentityGateway{}
		})
		g.It("Successfully registers a new identity", func() {

			useCase := usecases.NewRegisterIdentityUseCase(fakeIdentityGateway)
			request := usecases.RegisterIdentityParams{Email: "audie@foo.com", Name: "audie", PassPhrase: "lets-go-deeper"}
			err := useCase.Execute(request)

			g.Assert(err).IsNil()
		})

		g.It("Fails to register identity twice", func() {
			fakeIdentityGateway.WillHaveIdentityEmail = "audie@foo.com"
			fakeIdentityGateway.WillHaveIdentityName = "audie"
			useCase := usecases.NewRegisterIdentityUseCase(fakeIdentityGateway)
			request := usecases.RegisterIdentityParams{Email: "audie@foo.com", Name: "audie", PassPhrase: "lets-go-deeper"}
			err := useCase.Execute(request)

			g.Assert(err.Error()).Equal("ALREADY_HAVE_IDENTITY")
		})

		g.It("Fails to register identity for unknown reason", func() {
			fakeIdentityGateway.WillFaildentityRegistration = errors.New("IDENTITY_REGISTRATION_FAILED")
			useCase := usecases.NewRegisterIdentityUseCase(fakeIdentityGateway)
			request := usecases.RegisterIdentityParams{Email: "audie@foo.com", Name: "audie", PassPhrase: "lets-go-deeper"}
			err := useCase.Execute(request)

			g.Assert(err.Error()).Equal("IDENTITY_REGISTRATION_FAILED")
		})
	})

	g.Describe("I want to check my own identity", func() {
		g.BeforeEach(func() {
			fakeIdentityGateway = &adapters.FakeIdentityGateway{WillHaveIdentityEmail: "audie@app2b.io", WillHaveIdentityName: "audie"}
		})
		g.It("Successfully retrieves my identity", func() {
			useCase := usecases.NewGetCurrentIdentityUseCase(fakeIdentityGateway)
			response := useCase.Execute(usecases.GetCurrentIdentityRequest{Email: "audie@app2b.io"})
			g.Assert(response.Identity.Name).Equal("audie")
			g.Assert(response.Identity.Name).Equal("audie@app2b.io")
		})

		g.It("Fails to retrieve unknown identity", func() {
			unknownIdentityGateway := &adapters.FakeIdentityGateway{}
			useCase := usecases.NewGetCurrentIdentityUseCase(unknownIdentityGateway)
			response := useCase.Execute(usecases.GetCurrentIdentityRequest{Email: "baba-not-registered@gmail.com"})
			g.Assert(response).IsNil()
		})
	})
}
