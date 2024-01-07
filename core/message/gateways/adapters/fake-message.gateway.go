package adapters

import "grabit-cli/core/message/gateways"

type FakeMessageGateway struct {
	GeneratedUrl string
}

func (fmg FakeMessageGateway) SendTextPlainMessage(request gateways.SendTextPlainMessageRequest, response chan<- *gateways.SendMessageResponse) error {
	go func() {
		response <- &gateways.SendMessageResponse{Url: fmg.GeneratedUrl}
	}()
	return nil
}
