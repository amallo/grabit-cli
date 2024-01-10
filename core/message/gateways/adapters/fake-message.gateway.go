package adapters

import "grabit-cli/core/message/gateways"

type FakeMessageGateway struct {
	GeneratedUrl             string
	WillSentTextPlainContent string
}

func (fmg *FakeMessageGateway) Send(request gateways.SendMessageRequest) (*gateways.SendMessageResponse, error) {
	fmg.WillSentTextPlainContent = request.Message.Content
	return &gateways.SendMessageResponse{Url: fmg.GeneratedUrl}, nil
}
