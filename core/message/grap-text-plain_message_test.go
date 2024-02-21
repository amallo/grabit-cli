package message

import (
	"errors"
	usecases "grabit-cli/core/message/usecases"
	"testing"

	. "github.com/franela/goblin"
)

func TestGrabMessage(t *testing.T) {
	g := Goblin(t)
	var messageFixture MessageFixture
	g.Describe("audie grabs content from michael", func() {
		g.BeforeEach(func() {
			messageFixture = NewMessageFixture(g)
		})

		g.It("Successfully grab message content", func() {
			messageFixture.GivenDroppedMessage("msg-0", "binouze ce soir 19h")
			args := usecases.GrabMessageArgs{MessageId: "msg-0", Email: "michael@foo.com", Password: "prune"}
			messageFixture.WhenGrabbingMessage(args)
			messageFixture.ThenGrabbedMessageContentShouldBe("binouze ce soir 19h")
		})

		g.It("Fails to grab message", func() {
			args := usecases.GrabMessageArgs{MessageId: "msg-not-found", Email: "michael@foo.com", Password: "prune"}
			messageFixture.WhenGrabbingMessageWithFailure(args, errors.New("CANNOT_RETRIEVE_MESSAGE"))
			messageFixture.ThenGrabbingFailureShouldBe(usecases.ErrGrapMessageFailure)
		})

	})
}
