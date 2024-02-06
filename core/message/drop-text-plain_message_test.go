package message

import (
	"errors"
	"fmt"
	common_errors "grabit-cli/core/common/errors"
	id_adapters "grabit-cli/core/identities/gateways/adapters"
	"grabit-cli/core/message/gateways/adapters"
	core "grabit-cli/core/message/usecases"
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
			useCase := core.NewDropTextPlainMessageUseCase(&messageGateway, &fakeIdentityGateway, &fakeMessageIdGenerator)
			args := core.DropTextPlainMessageArgs{To: "michael@foo.com", Content: "binouze ce soir 19h", From: "audie@foo.com", Password: "prune"}
			result, err := useCase.Execute(args)

			g.Assert(err).IsNil()
			g.Assert(result.Url).Equal("https://files.com/AZERTYUIOP")
			g.Assert(result.From.Name).Equal("madfu")
			g.Assert(result.From.Email).Equal("audie@foo.com")

			g.Assert(result.To.Name).Equal("mic")
			g.Assert(result.To.Email).Equal("michael@foo.com")
			g.Assert(result.MessageId).Equal("message-0")
		})

		g.It("Fails to send text plain content to unknown recipient", func() {
			useCase := core.NewDropTextPlainMessageUseCase(&messageGateway, &fakeIdentityGateway, &fakeMessageIdGenerator)
			args := core.DropTextPlainMessageArgs{To: "michael@not-found.com", Content: "binouze ce soir 19h", From: "audie@foo.com", Password: "prune"}
			result, err := useCase.Execute(args)

			g.Assert(result).IsNil()
			g.Assert(err.Error()).Equal("UNKNOWN_RECIPIENT")
		})

		g.It("Fails to drop text plain content from unknown identity", func() {
			failureIdentityGateway := id_adapters.FailureIdentityGateway{WillFailLoadCurrent: errors.New("UNKNOWN_IDENTITY")}
			useCase := core.NewDropTextPlainMessageUseCase(&messageGateway, &failureIdentityGateway, &fakeMessageIdGenerator)
			args := core.DropTextPlainMessageArgs{To: "michael@foo.com", Content: "binouze ce soir 19h", From: "audie@not-found.com", Password: "prune"}
			result, err := useCase.Execute(args)
			fmt.Println("err", err)
			g.Assert(result).IsNil()
			g.Assert(errors.Is(err, common_errors.UnknownIdentityError{Email: "audie@not-found.com", CausedBy: "UNKNOWN_IDENTITY"})).IsTrue()

		})

		g.It("Fails to drop text plain content", func() {
			failureMessageGateway := &adapters.FailureMessageGateway{}
			useCase := core.NewDropTextPlainMessageUseCase(failureMessageGateway, &fakeIdentityGateway, &fakeMessageIdGenerator)
			params := core.DropTextPlainMessageArgs{To: "michael@foo.com", Content: "binouze ce soir 19h", From: "audie@foo.com", Password: "prune"}
			result, err := useCase.Execute(params)

			g.Assert(result).IsNil()
			g.Assert(err.Error()).Equal("DROP_MESSAGE_FAILURE")
		})

	})
}
