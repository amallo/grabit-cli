package core

import (
	"grabit-cli/core/message/gateways"
)

type sendTextPlainMessageUseCase struct {
	messageGateway gateways.MessageGateway
}
type SendTextPlainMessageParams struct {
	To      string
	Content string
}
type SendTextPlainMessageResponse struct {
	Url string
}

func NewSendTextPlainMessageUseCase(messageGateway gateways.MessageGateway) sendTextPlainMessageUseCase {
	return sendTextPlainMessageUseCase{messageGateway: messageGateway}
}

func (uc *sendTextPlainMessageUseCase) Execute(params SendTextPlainMessageParams) (*SendTextPlainMessageResponse, error) {
	response, error := uc.messageGateway.SendTextPlainMessage(gateways.SendTextPlainMessageRequest{To: params.To, Content: params.Content})
	if error != nil {
		return nil, error
	}
	return &SendTextPlainMessageResponse{Url: response.Url}, nil
}
