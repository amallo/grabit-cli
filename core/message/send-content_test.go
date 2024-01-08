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
	var messageEncrypter *adapters.FakeMessageEncrypter
	var knownRecipientGateway *id_adapters.KnownRecipientGateway
	g.Describe("audie sends content to michael", func() {
		g.BeforeEach(func() {
			messageEncrypter = &adapters.FakeMessageEncrypter{WillEncryptTextPlainAs: "ENCRYPTED"}
			messageGateway = &adapters.FakeMessageGateway{GeneratedUrl: "https://files.com/AZERTYUIOP"}
			knownRecipientGateway = &id_adapters.KnownRecipientGateway{WillLoadPublicKey: "public key"}
		})

		g.It("Successfully send message content", func() {
			useCase := core.NewSendTextPlainMessageUseCase(messageGateway, messageEncrypter, knownRecipientGateway)
			params := core.SendTextPlainMessageParams{To: "michael@foo.com", Content: "binouze ce soir 19h"}
			response, err := useCase.Execute(params)

			g.Assert(err).IsNil()
			g.Assert(response.Url).Equal("https://files.com/AZERTYUIOP")
			g.Assert(messageGateway.WillSentTextPlainContent).Equal("ENCRYPTED")
		})

		g.It("Fails to send text plain content to unknown recipient", func() {
			unknownRecipientGateway := &id_adapters.UnknownRecipientGateway{}
			useCase := core.NewSendTextPlainMessageUseCase(messageGateway, messageEncrypter, unknownRecipientGateway)
			params := core.SendTextPlainMessageParams{To: "michael@foo.com", Content: "binouze ce soir 19h"}
			response, err := useCase.Execute(params)

			g.Assert(response).IsNil()
			g.Assert(err.Error()).Equal("UNKNOWN_RECIPIENT")
		})

		g.It("Fails to send text plain content when encryption failure", func() {
			failureMessageEncrypter := adapters.FailureMessageStubEncrypter{}
			useCase := core.NewSendTextPlainMessageUseCase(messageGateway, failureMessageEncrypter, knownRecipientGateway)
			params := core.SendTextPlainMessageParams{To: "michael@foo.com", Content: "binouze ce soir 19h"}
			response, err := useCase.Execute(params)

			g.Assert(response).IsNil()
			g.Assert(err.Error()).Equal("ENCRYPTION_FAILURE")
		})

		g.It("Fails to send text plain content when transmission fails", func() {
			failureMessageGateway := &adapters.FailureMessageGateway{}
			useCase := core.NewSendTextPlainMessageUseCase(failureMessageGateway, messageEncrypter, knownRecipientGateway)
			params := core.SendTextPlainMessageParams{To: "michael@foo.com", Content: "binouze ce soir 19h"}
			response, err := useCase.Execute(params)

			g.Assert(response).IsNil()
			g.Assert(err.Error()).Equal("TRANSMISSION_FAILURE")
		})

	})
}
