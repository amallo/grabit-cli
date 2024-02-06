package adapters

import (
	"grabit-cli/core/message/gateways"
	"grabit-cli/core/message/models"
)

type FakeMessageGateway struct {
	GeneratedUrl    string
	WillDropMessage map[string]models.Message
}

func NewFakeMessageGateway() FakeMessageGateway {
	return FakeMessageGateway{WillDropMessage: make(map[string]models.Message)}
}

func (fmg *FakeMessageGateway) Drop(request gateways.DropMessageRequest) (*gateways.DropMessageResponse, error) {
	fmg.WillDropMessage[request.Message.Id] = request.Message
	return &gateways.DropMessageResponse{Url: fmg.GeneratedUrl, MessageId: request.Message.Id}, nil
}

func (fmg *FakeMessageGateway) Grab(request gateways.GrabMessageRequest) (*gateways.GrabMessageResponse, error) {
	droppedMessage := fmg.WillDropMessage[request.MessageId]
	delete(fmg.WillDropMessage, request.MessageId)
	return &gateways.GrabMessageResponse{Content: droppedMessage.Content}, nil
}
