package adapters

import "grabit-cli/core/message/gateways"

type FakeMessageGateway struct {
	GeneratedUrl        string
	WillSentTextContent string
}

func (fmg *FakeMessageGateway) Drop(request gateways.DropMessageRequest) (*gateways.DropMessageResponse, error) {
	fmg.WillSentTextContent = request.Message.Content
	return &gateways.DropMessageResponse{Url: fmg.GeneratedUrl}, nil
}
