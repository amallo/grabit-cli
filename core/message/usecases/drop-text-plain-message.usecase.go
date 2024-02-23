package usecases

import (
	core_errors "grabit-cli/core/common/errors"
	"grabit-cli/core/identities/models"
	"grabit-cli/core/message/gateways"
	message_models "grabit-cli/core/message/models"
)

type DropTextPlainMessageUseCase struct {
	messageGateway     gateways.MessageGateway
	messageIdGenerator gateways.MessageIdGenerator
}
type DropTextPlainMessageArgs struct {
	Recipient string
	Sender    string
	Password  string
	Content   string
}
type DropTextPlainMessageResult struct {
	Url       string
	Sender    models.Identity
	Recipient models.Identity
	MessageId string
}

func NewDropTextPlainMessageUseCase(messageGateway gateways.MessageGateway,
	messageIdGenerator gateways.MessageIdGenerator,
) DropTextPlainMessageUseCase {
	return DropTextPlainMessageUseCase{messageGateway: messageGateway, messageIdGenerator: messageIdGenerator}
}

func (uc *DropTextPlainMessageUseCase) Execute(args DropTextPlainMessageArgs) (*DropTextPlainMessageResult, core_errors.Error) {
	sender := models.Identity{Email: args.Sender}
	recipient := models.Identity{Email: args.Recipient}
	newMessageId := uc.messageIdGenerator.Generate()
	message := message_models.Message{Content: args.Content, From: sender, To: recipient, Id: newMessageId}

	dropRequest := gateways.DropMessageRequest{Message: message, Password: args.Password}
	dropResponse, err := uc.messageGateway.Drop(dropRequest)
	if err != nil {
		return nil, core_errors.Err(ErrDropMessageFailure, err)
	}
	return &DropTextPlainMessageResult{Url: dropResponse.Url, Recipient: recipient, MessageId: message.Id}, nil
}
