package gateways

import "grabit-cli/core/message/models"

type SendMessageRequest struct {
	Message models.Message
}

type SendMessageResponse struct {
	Url string
}
type MessageGateway interface {
	Send(request SendMessageRequest, smr chan<- *SendMessageResponse) error
}
type MessageEncrypter interface {
	EncryptPlainText(to string, text string) (*models.Message, error)
}
