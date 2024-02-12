package gateways

import (
	id_models "grabit-cli/core/identities/models"
	"grabit-cli/core/message/models"
)

type DropMessageRequest struct {
	Message  models.Message
	Password string
}

type DropMessageResponse struct {
	Url       string
	MessageId string
}
type GrabMessageRequest struct {
	MessageId string
	Password  string
	Identity  id_models.Identity
}

type GrabMessageResponse struct {
	Content string
}

type MessageGateway interface {
	Drop(request DropMessageRequest) (*DropMessageResponse, error)
	Grab(request GrabMessageRequest) (*GrabMessageResponse, error)
}
