package core

import (
	"errors"
	identities_gateway "grabit-cli/core/identities/gateways"
	"grabit-cli/core/message/gateways"
)

type sendTextPlainMessageUseCase struct {
	messageGateway   gateways.MessageGateway
	messageEncrypter gateways.MessageEncrypter
	recipientGateway identities_gateway.RecipientGateway
	identityGateway  identities_gateway.IdentityGateway
}
type SendTextPlainMessageParams struct {
	To      string
	Content string
}
type SendTextPlainMessageResponse struct {
	Url string
}

func NewSendTextPlainMessageUseCase(messageGateway gateways.MessageGateway,
	messageEncrypter gateways.MessageEncrypter,
	recipientGateway identities_gateway.RecipientGateway,
	identityGateway identities_gateway.IdentityGateway,
) sendTextPlainMessageUseCase {
	return sendTextPlainMessageUseCase{messageGateway: messageGateway, messageEncrypter: messageEncrypter, recipientGateway: recipientGateway, identityGateway: identityGateway}
}

func (uc *sendTextPlainMessageUseCase) Execute(params SendTextPlainMessageParams) (*SendTextPlainMessageResponse, error) {
	_, error := uc.identityGateway.LoadCurrent()
	if error != nil {
		return nil, errors.New("UNKNOWN_IDENTITY")
	}

	recipientIdentityResponse, error := uc.recipientGateway.FetchPublicKey(params.To)

	if error != nil {
		return nil, errors.New("UNKNOWN_RECIPIENT")
	}

	message, error := uc.messageEncrypter.EncryptPlainText(recipientIdentityResponse.PublicKey, params.Content)
	if error != nil {
		return nil, errors.New("ENCRYPTION_FAILURE")
	}

	request := gateways.SendMessageRequest{Message: *message, To: params.To}
	response, error := uc.messageGateway.Send(request)
	if error != nil {
		return nil, errors.New("TRANSMISSION_FAILURE")
	}
	return &SendTextPlainMessageResponse{Url: response.Url}, nil
}
