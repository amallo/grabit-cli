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
	identityChan := make(chan *identities_gateway.LoadIdentityResponse, 1)
	error := uc.identityGateway.LoadCurrent(identityChan)
	defer close(identityChan)
	if error != nil {
		return nil, errors.New("UNKNOWN_IDENTITY")
	}

	recipientChan := make(chan *identities_gateway.FetchPublicKeyRequestResponse, 1)
	error = uc.recipientGateway.FetchPublicKey(params.To, recipientChan)
	defer close(recipientChan)
	if error != nil {
		return nil, errors.New("UNKNOWN_RECIPIENT")
	}
	recipientIdentityResponse := <-recipientChan

	message, error := uc.messageEncrypter.EncryptPlainText(recipientIdentityResponse.PublicKey, params.Content)
	if error != nil {
		return nil, errors.New("ENCRYPTION_FAILURE")
	}

	sendMessageChan := make(chan *gateways.SendMessageResponse, 1)
	defer close(sendMessageChan)

	request := gateways.SendMessageRequest{Message: *message, To: params.To}
	error = uc.messageGateway.Send(request, sendMessageChan)
	if error != nil {
		return nil, errors.New("TRANSMISSION_FAILURE")
	}
	response := <-sendMessageChan
	return &SendTextPlainMessageResponse{Url: response.Url}, nil
}
