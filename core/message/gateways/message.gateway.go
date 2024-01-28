package gateways

import "grabit-cli/core/message/models"

type DropMessageRequest struct {
	Message  models.Message
	Password string
}

type DropMessageResponse struct {
	Url string
}
type MessageGateway interface {
	Drop(request DropMessageRequest) (*DropMessageResponse, error)
}
