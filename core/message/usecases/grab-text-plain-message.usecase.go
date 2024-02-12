package core

import (
	"errors"
	common_errors "grabit-cli/core/common/errors"
	identities_gateway "grabit-cli/core/identities/gateways"
	"grabit-cli/core/identities/models"
	"grabit-cli/core/message/gateways"
)

type grabMessageUseCase struct {
	messageGateway     gateways.MessageGateway
	identityGateway    identities_gateway.IdentityGateway
	messageIdGenerator gateways.MessageIdGenerator
}
type GrabMessageArgs struct {
	Email     string
	Password  string
	MessageId string
}
type GrabMessageResult struct {
	Content string
}

func NewGrabMessageUseCase(messageGateway gateways.MessageGateway,
	identityGateway identities_gateway.IdentityGateway,
	messageIdGenerator gateways.MessageIdGenerator,
) grabMessageUseCase {
	return grabMessageUseCase{messageGateway: messageGateway, identityGateway: identityGateway, messageIdGenerator: messageIdGenerator}
}

func (uc grabMessageUseCase) Execute(params GrabMessageArgs) (*GrabMessageResult, error) {
	identityResponse, err := uc.identityGateway.LoadCurrent(params.Email)
	if err != nil {
		return nil, errors.New("UNKNOWN_IDENTITY")
	}
	identity := models.Identity{Email: params.Email, Name: identityResponse.Name}
	grabRequest := gateways.GrabMessageRequest{MessageId: params.MessageId, Password: params.Password, Identity: identity}
	grabResponse, err := uc.messageGateway.Grab(grabRequest)
	if err != nil {
		return nil, common_errors.NotFoundError{Category: "Message", CausedBy: err.Error(), Id: params.MessageId}
	}
	return &GrabMessageResult{Content: grabResponse.Content}, nil
}
