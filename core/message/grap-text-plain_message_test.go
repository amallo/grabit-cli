package message

import (
	"errors"
	common_errors "grabit-cli/core/common/errors"
	id_adapters "grabit-cli/core/identities/gateways/adapters"
	"grabit-cli/core/message/gateways/adapters"
	"grabit-cli/core/message/models"
	core "grabit-cli/core/message/usecases"
	"testing"

	. "github.com/franela/goblin"
)

func TestGrabMessage(t *testing.T) {
	g := Goblin(t)
	var messageGateway adapters.FakeMessageGateway
	var fakeIdentityGateway *id_adapters.FakeIdentityGateway
	var fakeMessageIdGenerator adapters.FakeMessageIdGenerator
	var identities map[string]string
	g.Describe("audie grabs content from michael", func() {
		g.BeforeEach(func() {
			identities = make(map[string]string)
			identities["audie@foo.com"] = "madfu"
			identities["michael@foo.com"] = "mic"
			fakeIdentityGateway = &id_adapters.FakeIdentityGateway{
				WillHaveIdentities: identities,
			}
			messageGateway = adapters.NewFakeMessageGateway()
			messageGateway.GeneratedUrl = "https://files.com/AZERTYUIOP"
			messageGateway.WillDropMessage["msg-0"] = models.Message{Content: "binouze ce soir 19h"}
			fakeMessageIdGenerator.WillGenerateId = "message-0"
		})

		g.It("Successfully grab message content", func() {
			useCase := core.NewGrabMessageUseCase(&messageGateway, fakeIdentityGateway, &fakeMessageIdGenerator)
			args := core.GrabMessageArgs{MessageId: "msg-0", Email: "michael@foo.com", Password: "prune"}
			result, err := useCase.Execute(args)

			g.Assert(err).IsNil()
			g.Assert(result.Content).Equal("binouze ce soir 19h")
		})

		g.It("Fails to grab non existing message", func() {
			failureGateway := adapters.FailureMessageGateway{GrabMessageFailure: errors.New("CANNOT_RETRIEVE_MESSAGE")}
			useCase := core.NewGrabMessageUseCase(&failureGateway, fakeIdentityGateway, &fakeMessageIdGenerator)
			args := core.GrabMessageArgs{MessageId: "msg-not-found", Email: "michael@foo.com", Password: "prune"}
			result, err := useCase.Execute(args)
			g.Assert(errors.Is(err, common_errors.NotFoundError{Category: "Message", Id: "msg-not-found", CausedBy: "CANNOT_RETRIEVE_MESSAGE"})).IsTrue()
			g.Assert(result).IsNil()
		})

	})
}
