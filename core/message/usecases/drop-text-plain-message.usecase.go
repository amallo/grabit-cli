package core

import (
	"errors"
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
	To       string
	From     string
	Password string
	Content  string
}
type DropTextPlainMessageResult struct {
	Url       string
	From      models.Identity
	To        models.Identity
	MessageId string
}

func NewDropTextPlainMessageUseCase(messageGateway gateways.MessageGateway,
	identityGateway identities_gateway.IdentityGateway,
	messageIdGenerator gateways.MessageIdGenerator,
) dropTextPlainMessageUseCase {
	return dropTextPlainMessageUseCase{messageGateway: messageGateway, identityGateway: identityGateway, messageIdGenerator: messageIdGenerator}
}

func (uc *dropTextPlainMessageUseCase) Execute(params DropTextPlainMessageArgs) (*DropTextPlainMessageResult, error) {
	senderIdentityResponse, err := uc.identityGateway.LoadCurrent(params.From)
	if err != nil {
		return nil, errors.New("UNKNOWN_IDENTITY")
	}

	recipientIdentityResponse, err := uc.identityGateway.LoadCurrent(params.To)
	if err != nil {
		return nil, errors.New("UNKNOWN_RECIPIENT")
	}

	fromIdentity := models.Identity{Email: params.From, Name: senderIdentityResponse.Name}
	toIdentity := models.Identity{Email: params.To, Name: recipientIdentityResponse.Name}
	newMessageId := uc.messageIdGenerator.Generate()
	message := models2.Message{Content: params.Content, From: fromIdentity, To: toIdentity, Id: newMessageId}

	dropRequest := gateways.DropMessageRequest{Message: message, Password: params.Password}
	dropResponse, err := uc.messageGateway.Drop(dropRequest)
	if err != nil {
		return nil, errors.New("DROP_MESSAGE_FAILURE")
	}
	return &DropTextPlainMessageResult{Url: dropResponse.Url, From: fromIdentity, To: toIdentity, MessageId: message.Id}, nil
}
