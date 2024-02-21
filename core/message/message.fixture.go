package message

import (
	core_errors "grabit-cli/core/common/errors"
	"grabit-cli/core/message/gateways/adapters"
	"grabit-cli/core/message/models"
	usecases "grabit-cli/core/message/usecases"

	. "github.com/franela/goblin"
)

type MessageFixture struct {
	g                  *G
	messageGateway     adapters.FakeMessageGateway
	messageIdGenerator adapters.FakeMessageIdGenerator

	dropTextPlainMessageResult *usecases.DropTextPlainMessageResult
	grabMessageResult          *usecases.GrabMessageResult
	err                        core_errors.Error
}

func NewMessageFixture(g *G) MessageFixture {
	messageGateway := adapters.NewFakeMessageGateway()
	messageGateway.GeneratedUrl = "https://files.com/AZERTYUIOP"

	fakeMessageIdGenerator := adapters.FakeMessageIdGenerator{}
	fakeMessageIdGenerator.WillGenerateId = "message-0"
	return MessageFixture{messageGateway: messageGateway, messageIdGenerator: fakeMessageIdGenerator, g: g}
}

func (f *MessageFixture) GivenGeneratedUrl(url string) {
	f.messageGateway.GeneratedUrl = url
}
func (f *MessageFixture) GivenGeneratedMessageId(messageId string) {
	f.messageIdGenerator.WillGenerateId = messageId
}

func (f *MessageFixture) GivenDroppedMessage(messageId string, messageContent string) {
	f.messageGateway.WillDropMessage[messageId] = models.Message{Content: messageContent}
}
func (f *MessageFixture) WhenDroppingTextPlainMessage(args usecases.DropTextPlainMessageArgs) {
	useCase := usecases.NewDropTextPlainMessageUseCase(&f.messageGateway, &f.messageIdGenerator)
	f.dropTextPlainMessageResult, f.err = useCase.Execute(args)
}
func (f *MessageFixture) WhenDroppingTextPlainMessageWithFailure(args usecases.DropTextPlainMessageArgs) {
	failureMessageGateway := adapters.FailureMessageGateway{}
	useCase := usecases.NewDropTextPlainMessageUseCase(&failureMessageGateway, &f.messageIdGenerator)
	f.dropTextPlainMessageResult, f.err = useCase.Execute(args)
}
func (f *MessageFixture) WhenGrabbingMessage(args usecases.GrabMessageArgs) {
	useCase := usecases.NewGrabMessageUseCase(&f.messageGateway)
	f.grabMessageResult, f.err = useCase.Execute(args)
}

func (f *MessageFixture) WhenGrabbingMessageWithFailure(args usecases.GrabMessageArgs, failure error) {
	failureGateway := adapters.FailureMessageGateway{GrabMessageFailure: failure}
	useCase := usecases.NewGrabMessageUseCase(&failureGateway)
	f.grabMessageResult, f.err = useCase.Execute(args)
}

func (f *MessageFixture) ThenDroppedUrlShouldBe(expectedUrl string) {
	f.g.Assert(f.err).IsNil()
	f.g.Assert(f.dropTextPlainMessageResult.Url).Equal(expectedUrl)
}
func (f *MessageFixture) ThenDroppedEmailShouldBe(expectedEmail string) {
	f.g.Assert(f.err).IsNil()
	f.g.Assert(f.dropTextPlainMessageResult.Recipient.Email).Equal(expectedEmail)
}
func (f *MessageFixture) ThenDroppedMessageIdShouldBe(expectedMessageId string) {
	f.g.Assert(f.err).IsNil()
	f.g.Assert(f.dropTextPlainMessageResult.MessageId).Equal(expectedMessageId)
}
func (f *MessageFixture) ThenDroppingFailureShouldBe(expectedErrorCode string) {
	f.g.Assert(f.dropTextPlainMessageResult).IsNil()
	f.g.Assert(f.err.Code()).Equal(expectedErrorCode)
}
func (f *MessageFixture) ThenGrabbingFailureShouldBe(expectedErrorCode string) {
	f.g.Assert(f.grabMessageResult).IsNil()
	f.g.Assert(f.err.Code()).Equal(expectedErrorCode)
}

func (f *MessageFixture) ThenGrabbedMessageContentShouldBe(expectedMessageContent string) {
	f.g.Assert(f.err).IsNil()
	f.g.Assert(f.grabMessageResult.Content).Equal(expectedMessageContent)
}
