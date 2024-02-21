package usecases

import (
	core_errors "grabit-cli/core/common/errors"
	identities_gateway "grabit-cli/core/identities/gateways"
	"grabit-cli/core/identities/models"
	"grabit-cli/core/message/gateways"
	models2 "grabit-cli/core/message/models"
)

type dropTextPlainMessageUseCase struct {
	messageGateway     gateways.MessageGateway
	identityGateway    identities_gateway.IdentityGateway
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
	identityGateway identities_gateway.IdentityGateway,
	messageIdGenerator gateways.MessageIdGenerator,
) dropTextPlainMessageUseCase {
	return dropTextPlainMessageUseCase{messageGateway: messageGateway, identityGateway: identityGateway, messageIdGenerator: messageIdGenerator}
}

func (uc *dropTextPlainMessageUseCase) Execute(params DropTextPlainMessageArgs) (*DropTextPlainMessageResult, core_errors.Error) {
	senderIdentityResponse, err := uc.identityGateway.LoadCurrent(params.Sender)
	if err != nil {
		return nil, core_errors.Err(models.ErrUnknownIdentity, err)
	}

	recipientIdentityResponse, err := uc.identityGateway.LoadCurrent(params.Recipient)
	if err != nil {
		return nil, core_errors.Err(models.ErrUnknownIdentity, err)
	}

	senderIdentity := models.Identity{Email: params.Sender, Name: senderIdentityResponse.Name}
	recipientIdentity := models.Identity{Email: params.Recipient, Name: recipientIdentityResponse.Name}
	newMessageId := uc.messageIdGenerator.Generate()
	message := models2.Message{Content: params.Content, From: senderIdentity, To: recipientIdentity, Id: newMessageId}

	dropRequest := gateways.DropMessageRequest{Message: message, Password: params.Password}
	dropResponse, err := uc.messageGateway.Drop(dropRequest)
	if err != nil {
		return nil, core_errors.Err(ErrDropMessageFailure, err)
	}
	return &DropTextPlainMessageResult{Url: dropResponse.Url, Sender: senderIdentity, Recipient: recipientIdentity, MessageId: message.Id}, nil
}
