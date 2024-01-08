package identities

import (
	core "grabit-cli/core/identities/usecases"
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("audie registers a new identity", func() {
		/*g.BeforeEach(func() {
			messageEncrypter = &adapters.FakeMessageEncrypter{WillEncryptTextPlainAs: "ENCRYPTED"}
			messageGateway = &adapters.FakeMessageGateway{GeneratedUrl: "https://files.com/AZERTYUIOP"}
			knownRecipientGateway = &id_adapters.KnownRecipientGateway{WillLoadPublicKey: "public key"}
		})*/

		g.It("Successfully registers a new identity", func() {
			useCase := core.NewRegisterIdentityUseCase()
			response, err := useCase.Execute()

			g.Assert(err).IsNil()
			g.Assert(response.Email).Equal("audie@foo.com")
			g.Assert(response.PublicKey).Equal("public key")
		})
	})
}
