package message

import (
	usecases "grabit-cli/core/message/usecases"
	"testing"

	. "github.com/franela/goblin"
)

func TestDropMessage(t *testing.T) {
	g := Goblin(t)

	var messageFixture MessageFixture
	g.Describe("audie drops content to michael", func() {
		g.BeforeEach(func() {
			messageFixture = NewMessageFixture(g)
		})
		g.It("Successfully drop message content", func() {
			messageFixture.GivenGeneratedMessageId("message-0")
			messageFixture.GivenGeneratedUrl("https://files.com/AZERTYUIOP")

			args := usecases.DropTextPlainMessageArgs{Recipient: "michael@foo.com", Content: "binouze ce soir 19h", Sender: "audie@foo.com", Password: "prune"}
			messageFixture.WhenDroppingTextPlainMessage(args)

			messageFixture.ThenDroppedEmailShouldBe("michael@foo.com")
			messageFixture.ThenDroppedMessageIdShouldBe("message-0")
			messageFixture.ThenDroppedUrlShouldBe("https://files.com/AZERTYUIOP")
		})

		g.It("Fails to drop text plain content", func() {
			args := usecases.DropTextPlainMessageArgs{Recipient: "michael@foo.com", Content: "binouze ce soir 19h", Sender: "audie@foo.com", Password: "prune"}
			messageFixture.WhenDroppingTextPlainMessageWithFailure(args)
			messageFixture.ThenDroppingFailureShouldBe(usecases.ErrDropMessageFailure)
		})

	})
}
