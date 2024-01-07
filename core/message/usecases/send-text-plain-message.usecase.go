package core

import (
	"grabit-cli/core/message/gateways"
)

type sendTextPlainMessageUseCase struct {
	messageGateway   gateways.MessageGateway
	messageEncrypter gateways.MessageEncrypter
}
type SendTextPlainMessageParams struct {
	To      string
	Content string
}
type SendTextPlainMessageResponse struct {
	Url string
}

func NewSendTextPlainMessageUseCase(messageGateway gateways.MessageGateway, messageEncrypter gateways.MessageEncrypter) sendTextPlainMessageUseCase {
	return sendTextPlainMessageUseCase{messageGateway: messageGateway, messageEncrypter: messageEncrypter}
}

func (uc *sendTextPlainMessageUseCase) Execute(params SendTextPlainMessageParams) (*SendTextPlainMessageResponse, error) {
	responseChan := make(chan *gateways.SendMessageResponse, 1)
	defer close(responseChan)

	message, error := uc.messageEncrypter.EncryptPlainText(params.To, params.Content)
	if error != nil {
		return nil, error
	}
	request := gateways.SendMessageRequest{Message: *message}
	error = uc.messageGateway.Send(request, responseChan)
	if error != nil {
		return nil, error
	}
	response := <-responseChan
	return &SendTextPlainMessageResponse{Url: response.Url}, nil
}
