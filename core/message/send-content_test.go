package message

import (
	id_adapters "grabit-cli/core/identities/gateways/adapters"
	"grabit-cli/core/message/gateways/adapters"
	core "grabit-cli/core/message/usecases"
	"testing"

	. "github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := Goblin(t)
	var messageGateway *adapters.FakeMessageGateway
	var fakeIdentityGateway *id_adapters.FakeIdentityGateway
	var identities map[string]string
	g.Describe("audie drops content to michael", func() {
		g.BeforeEach(func() {
			identities = make(map[string]string)
			identities["audie@foo.com"] = "madfu"
			identities["michael@foo.com"] = "mic"
			fakeIdentityGateway = &id_adapters.FakeIdentityGateway{
				WillHaveIdentities: identities,
			}
			messageGateway = &adapters.FakeMessageGateway{GeneratedUrl: "https://files.com/AZERTYUIOP"}
			messageGateway.WillSentTextContent = "binouze ce soir 19h"

		})

		g.It("Successfully send message content", func() {
			useCase := core.NewSendTextPlainMessageUseCase(messageGateway, fakeIdentityGateway)
			args := core.SendTextPlainMessageArgs{To: "michael@foo.com", Content: "binouze ce soir 19h", From: "audie@foo.com", Password: "prune"}
			result, err := useCase.Execute(args)

			g.Assert(err).IsNil()
			g.Assert(result.Url).Equal("https://files.com/AZERTYUIOP")
			g.Assert(messageGateway.WillSentTextContent).Equal("binouze ce soir 19h")
			g.Assert(result.From.Name).Equal("madfu")
			g.Assert(result.From.Email).Equal("audie@foo.com")

			g.Assert(result.To.Name).Equal("mic")
			g.Assert(result.To.Email).Equal("michael@foo.com")
		})

		g.It("Fails to send text plain content to unknown recipient", func() {
			useCase := core.NewSendTextPlainMessageUseCase(messageGateway, fakeIdentityGateway)
			args := core.SendTextPlainMessageArgs{To: "michael@not-found.com", Content: "binouze ce soir 19h", From: "audie@foo.com", Password: "prune"}
			result, err := useCase.Execute(args)

			g.Assert(result).IsNil()
			g.Assert(err.Error()).Equal("UNKNOWN_RECIPIENT")
		})

		g.It("Fails to send text plain content from unknown identity", func() {
			useCase := core.NewSendTextPlainMessageUseCase(messageGateway, fakeIdentityGateway)
			args := core.SendTextPlainMessageArgs{To: "michael@foo.com", Content: "binouze ce soir 19h", From: "audie@not-found.com", Password: "prune"}
			result, err := useCase.Execute(args)

			g.Assert(result).IsNil()
			g.Assert(err.Error()).Equal("UNKNOWN_IDENTITY")
		})

		g.It("Fails to send text plain content when transmission fails", func() {
			failureMessageGateway := &adapters.FailureMessageGateway{}
			useCase := core.NewSendTextPlainMessageUseCase(failureMessageGateway, fakeIdentityGateway)
			params := core.SendTextPlainMessageArgs{To: "michael@foo.com", Content: "binouze ce soir 19h", From: "audie@foo.com", Password: "prune"}
			result, err := useCase.Execute(params)

			g.Assert(result).IsNil()
			g.Assert(err.Error()).Equal("TRANSMISSION_FAILURE")
		})

	})
}
