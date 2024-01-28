package core

import (
	"errors"
	identities_gateway "grabit-cli/core/identities/gateways"
	"grabit-cli/core/identities/models"
	"grabit-cli/core/message/gateways"
	models2 "grabit-cli/core/message/models"
)

type sendTextPlainMessageUseCase struct {
	messageGateway  gateways.MessageGateway
	identityGateway identities_gateway.IdentityGateway
}
type SendTextPlainMessageArgs struct {
	To       string
	From     string
	Password string
	Content  string
}
type SendTextPlainMessageResult struct {
	Url  string
	From models.Identity
	To   models.Identity
}

func NewSendTextPlainMessageUseCase(messageGateway gateways.MessageGateway,
	identityGateway identities_gateway.IdentityGateway,
) sendTextPlainMessageUseCase {
	return sendTextPlainMessageUseCase{messageGateway: messageGateway, identityGateway: identityGateway}
}

func (uc *sendTextPlainMessageUseCase) Execute(params SendTextPlainMessageArgs) (*SendTextPlainMessageResult, error) {
	senderIdentityResponse, error := uc.identityGateway.LoadCurrent(params.From)
	if error != nil {
		return nil, errors.New("UNKNOWN_IDENTITY")
	}

	recipientIdentityResponse, error := uc.identityGateway.LoadCurrent(params.To)
	if error != nil {
		return nil, errors.New("UNKNOWN_RECIPIENT")
	}

	fromIdentity := models.Identity{Email: params.From, Name: senderIdentityResponse.Name}
	toIdentity := models.Identity{Email: params.To, Name: recipientIdentityResponse.Name}
	message := models2.Message{Content: params.Content, From: fromIdentity, To: toIdentity}

	dropRequest := gateways.DropMessageRequest{Message: message, Password: params.Password}
	result, error := uc.messageGateway.Drop(dropRequest)
	if error != nil {
		return nil, errors.New("TRANSMISSION_FAILURE")
	}
	return &SendTextPlainMessageResult{Url: result.Url, From: fromIdentity, To: toIdentity}, nil
}
