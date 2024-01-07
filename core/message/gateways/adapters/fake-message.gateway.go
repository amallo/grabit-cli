package adapters

import "grabit-cli/core/message/gateways"

type FakeMessageGateway struct {
	GeneratedUrl string
}

func (fmg FakeMessageGateway) SendTextPlainMessage(request gateways.SendTextPlainMessageRequest) (*gateways.SendMessageResponse, error) {
	return &gateways.SendMessageResponse{Url: fmg.GeneratedUrl}, nil
}
