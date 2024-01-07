package message

import (
	"grabit-cli/core/message/gateways/adapters"
	core "grabit-cli/core/message/usecases"
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("audie sends content to michael", func() {
		g.It("Successfully send message content", func() {
			messageGateway := &adapters.FakeMessageGateway{GeneratedUrl: "https://files.com/AZERTYUIOP"}
			messageEncrypter := adapters.FakeMessageEncrypter{WillEncryptTextPlainAs: "ENCRYPTED"}
			useCase := core.NewSendTextPlainMessageUseCase(messageGateway, messageEncrypter)
			params := core.SendTextPlainMessageParams{To: "michael@foo.com", Content: "binouze ce soir 19h"}
			response, err := useCase.Execute(params)

			g.Assert(err).IsNil()
			g.Assert(response.Url).Equal("https://files.com/AZERTYUIOP")
			g.Assert(messageGateway.WillSentTextPlainContent).Equal("ENCRYPTED")

		})

	})
}
