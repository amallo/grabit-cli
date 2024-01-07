package adapters

import "grabit-cli/core/message/gateways"

type FakeMessageGateway struct {
	GeneratedUrl             string
	WillSentTextPlainContent string
}

func (fmg *FakeMessageGateway) Send(request gateways.SendMessageRequest, response chan<- *gateways.SendMessageResponse) error {
	fmg.WillSentTextPlainContent = request.Message.Content
	go func() {
		response <- &gateways.SendMessageResponse{Url: fmg.GeneratedUrl}
	}()
	return nil
}
