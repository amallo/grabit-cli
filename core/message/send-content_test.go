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
			messageGateway := adapters.FakeMessageGateway{GeneratedUrl: "https://files.com/AZERTYUIOP"}
			useCase := core.NewSendTextPlainMessageUseCase(messageGateway)
			params := core.SendTextPlainMessageParams{To: "michael@foo.com", Content: "binouze ce soir 19h"}
			response, err := useCase.Execute(params)
			g.Assert(err == nil)
			g.Assert(response.Url == "https://files.com/AZERTYUIOP")
		})

	})
}
