package usecases

import (
	core_errors "grabit-cli/core/common/errors"
	"grabit-cli/core/identities/models"
	"grabit-cli/core/message/gateways"
	models2 "grabit-cli/core/message/models"
)

type dropTextPlainMessageUseCase struct {
	messageGateway     gateways.MessageGateway
	messageIdGenerator gateways.MessageIdGenerator
}
type DropTextPlainMessageArgs struct {
	Recipient string
	Sender    string
	Password  string
	Content   string
}
type dropTextPlainMessageResult struct {
	Url       string
	Sender    models.Identity
	Recipient models.Identity
	MessageId string
}

func NewDropTextPlainMessageUseCase(messageGateway gateways.MessageGateway,
	messageIdGenerator gateways.MessageIdGenerator,
) dropTextPlainMessageUseCase {
	return dropTextPlainMessageUseCase{messageGateway: messageGateway, messageIdGenerator: messageIdGenerator}
}

func (uc *dropTextPlainMessageUseCase) Execute(params DropTextPlainMessageArgs) (*dropTextPlainMessageResult, core_errors.Error) {
	sender := models.Identity{Email: params.Sender}
	recipient := models.Identity{Email: params.Recipient}
	newMessageId := uc.messageIdGenerator.Generate()
	message := models2.Message{Content: params.Content, From: sender, To: recipient, Id: newMessageId}

	dropRequest := gateways.DropMessageRequest{Message: message, Password: params.Password}
	dropResponse, err := uc.messageGateway.Drop(dropRequest)
	if err != nil {
		return nil, core_errors.Err(ErrDropMessageFailure, err)
	}
	return &dropTextPlainMessageResult{Url: dropResponse.Url, Recipient: recipient, MessageId: message.Id}, nil
}
