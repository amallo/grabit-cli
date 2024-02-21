package message

import (
	id_adapters "grabit-cli/core/identities/gateways/adapters"
	"grabit-cli/core/identities/models"
	"grabit-cli/core/message/gateways/adapters"
	usecases "grabit-cli/core/message/usecases"
	"testing"

	. "github.com/franela/goblin"
)

func TestDropMessage(t *testing.T) {
	g := Goblin(t)
	var messageGateway adapters.FakeMessageGateway
	var fakeIdentityGateway id_adapters.FakeIdentityGateway
	var fakeMessageIdGenerator adapters.FakeMessageIdGenerator
	var identities map[string]string
	g.Describe("audie drops content to michael", func() {
		g.BeforeEach(func() {
			identities = make(map[string]string)
			identities["audie@foo.com"] = "madfu"
			identities["michael@foo.com"] = "mic"
			fakeIdentityGateway = id_adapters.NewFakeIdentityGateway(
				identities,
			)
			messageGateway = adapters.NewFakeMessageGateway()
			messageGateway.GeneratedUrl = "https://files.com/AZERTYUIOP"
			fakeMessageIdGenerator.WillGenerateId = "message-0"
		})

		g.It("Successfully drop message content", func() {
			useCase := usecases.NewDropTextPlainMessageUseCase(&messageGateway, &fakeIdentityGateway, &fakeMessageIdGenerator)
			args := usecases.DropTextPlainMessageArgs{Recipient: "michael@foo.com", Content: "binouze ce soir 19h", Sender: "audie@foo.com", Password: "prune"}
			result, err := useCase.Execute(args)

			g.Assert(err).IsNil()
			g.Assert(result.Url).Equal("https://files.com/AZERTYUIOP")
			g.Assert(result.Sender.Name).Equal("madfu")
			g.Assert(result.Sender.Email).Equal("audie@foo.com")

			g.Assert(result.Recipient.Name).Equal("mic")
			g.Assert(result.Recipient.Email).Equal("michael@foo.com")
			g.Assert(result.MessageId).Equal("message-0")
		})

		g.It("Fails to send text plain content to unknown recipient", func() {
			useCase := usecases.NewDropTextPlainMessageUseCase(&messageGateway, &fakeIdentityGateway, &fakeMessageIdGenerator)
			args := usecases.DropTextPlainMessageArgs{Recipient: "michael@not-found.com", Content: "binouze ce soir 19h", Sender: "audie@foo.com", Password: "prune"}
			result, err := useCase.Execute(args)

			g.Assert(result).IsNil()
			g.Assert(err.Code()).Equal(models.ErrUnknownIdentity)
		})

		g.It("Fails to drop text plain content from unknown identity", func() {
			useCase := usecases.NewDropTextPlainMessageUseCase(&messageGateway, &fakeIdentityGateway, &fakeMessageIdGenerator)
			args := usecases.DropTextPlainMessageArgs{Recipient: "michael@foo.com", Content: "binouze ce soir 19h", Sender: "audie@not-found.com", Password: "prune"}
			result, err := useCase.Execute(args)
			g.Assert(result).IsNil()
			g.Assert(err.Code()).Equal(models.ErrUnknownIdentity)

		})

		g.It("Fails to drop text plain content", func() {
			failureMessageGateway := &adapters.FailureMessageGateway{}
			useCase := usecases.NewDropTextPlainMessageUseCase(failureMessageGateway, &fakeIdentityGateway, &fakeMessageIdGenerator)
			params := usecases.DropTextPlainMessageArgs{Recipient: "michael@foo.com", Content: "binouze ce soir 19h", Sender: "audie@foo.com", Password: "prune"}
			result, err := useCase.Execute(params)

			g.Assert(result).IsNil()
			g.Assert(err.Code()).Equal(usecases.ErrDropMessageFailure)
		})

	})
}
