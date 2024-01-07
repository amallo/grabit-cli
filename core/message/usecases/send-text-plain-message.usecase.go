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
	responseChan := make(chan *gateways.SendMessageResponse, 1)
	defer close(responseChan)

	error := uc.messageGateway.SendTextPlainMessage(gateways.SendTextPlainMessageRequest{To: params.To, Content: params.Content}, responseChan)
	if error != nil {
		return nil, error
	}
	response := <-responseChan
	return &SendTextPlainMessageResponse{Url: response.Url}, nil
}
