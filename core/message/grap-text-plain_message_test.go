package message

import (
	"errors"
	"grabit-cli/core/message/gateways/adapters"
	"grabit-cli/core/message/models"
	message_usecases "grabit-cli/core/message/usecases"
	"testing"

	. "github.com/franela/goblin"
)

func TestGrabMessage(t *testing.T) {
	g := Goblin(t)
	var messageGateway adapters.FakeMessageGateway
	var fakeMessageIdGenerator adapters.FakeMessageIdGenerator
	g.Describe("audie grabs content from michael", func() {
		g.BeforeEach(func() {

			messageGateway = adapters.NewFakeMessageGateway()
			messageGateway.GeneratedUrl = "https://files.com/AZERTYUIOP"
			messageGateway.WillDropMessage["msg-0"] = models.Message{Content: "binouze ce soir 19h"}
			fakeMessageIdGenerator.WillGenerateId = "message-0"
		})

		g.It("Successfully grab message content", func() {
			useCase := message_usecases.NewGrabMessageUseCase(&messageGateway, &fakeMessageIdGenerator)
			args := message_usecases.GrabMessageArgs{MessageId: "msg-0", Email: "michael@foo.com", Password: "prune"}
			result, err := useCase.Execute(args)

			g.Assert(err).IsNil()
			g.Assert(result.Content).Equal("binouze ce soir 19h")
		})

		g.It("Fails to grab message", func() {
			failureGateway := adapters.FailureMessageGateway{GrabMessageFailure: errors.New("CANNOT_RETRIEVE_MESSAGE")}
			useCase := message_usecases.NewGrabMessageUseCase(&failureGateway, &fakeMessageIdGenerator)
			args := message_usecases.GrabMessageArgs{MessageId: "msg-not-found", Email: "michael@foo.com", Password: "prune"}
			result, err := useCase.Execute(args)
			g.Assert(result).IsNil()
			g.Assert(err.Code()).Equal(message_usecases.ErrGrapMessageFailure)
		})

	})
}
