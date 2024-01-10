package gateways

import "grabit-cli/core/message/models"

type SendMessageRequest struct {
	Message models.Message
	To      string
}

type SendMessageResponse struct {
	Url string
}
type MessageGateway interface {
	Send(request SendMessageRequest) (*SendMessageResponse, error)
}
type MessageEncrypter interface {
	EncryptPlainText(publicKey string, text string) (*models.Message, error)
}
