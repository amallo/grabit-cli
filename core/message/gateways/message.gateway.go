package gateways

import "grabit-cli/core/message/models"

type DropMessageRequest struct {
	Message  models.Message
	Password string
}

type DropMessageResponse struct {
	Url       string
	MessageId string
}
type GrabMessageRequest struct {
	Id       string
	Email    string
	Password string
}

type GrabMessageResponse struct {
	Content string
}
type MessageGateway interface {
	Drop(request DropMessageRequest) (*DropMessageResponse, error)
	Grab(request GrabMessageRequest) (*GrabMessageResponse, error)
}
