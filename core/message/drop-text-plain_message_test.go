package message

import (
	"grabit-cli/core/message/gateways/adapters"
	usecases "grabit-cli/core/message/usecases"
	"testing"

	. "github.com/franela/goblin"
)

func TestDropMessage(t *testing.T) {
	g := Goblin(t)
	var messageGateway adapters.FakeMessageGateway
	var fakeMessageIdGenerator adapters.FakeMessageIdGenerator
	g.Describe("audie drops content to michael", func() {
		g.BeforeEach(func() {

			messageGateway = adapters.NewFakeMessageGateway()
			messageGateway.GeneratedUrl = "https://files.com/AZERTYUIOP"
			fakeMessageIdGenerator.WillGenerateId = "message-0"
		})

		g.It("Successfully drop message content", func() {
			useCase := usecases.NewDropTextPlainMessageUseCase(&messageGateway, &fakeMessageIdGenerator)
			args := usecases.DropTextPlainMessageArgs{Recipient: "michael@foo.com", Content: "binouze ce soir 19h", Sender: "audie@foo.com", Password: "prune"}
			result, err := useCase.Execute(args)

			g.Assert(err).IsNil()
			g.Assert(result.Url).Equal("https://files.com/AZERTYUIOP")
			g.Assert(result.Recipient.Email).Equal("michael@foo.com")
			g.Assert(result.MessageId).Equal("message-0")
		})

		g.It("Fails to drop text plain content", func() {
			failureMessageGateway := &adapters.FailureMessageGateway{}
			useCase := usecases.NewDropTextPlainMessageUseCase(failureMessageGateway, &fakeMessageIdGenerator)
			params := usecases.DropTextPlainMessageArgs{Recipient: "michael@foo.com", Content: "binouze ce soir 19h", Sender: "audie@foo.com", Password: "prune"}
			result, err := useCase.Execute(params)

			g.Assert(result).IsNil()
			g.Assert(err.Code()).Equal(usecases.ErrDropMessageFailure)
		})

	})
}
