package adapters

import (
	"errors"
	"grabit-cli/core/message/gateways"
)

type FailureMessageGateway struct {
}

func (fmg *FailureMessageGateway) Send(request gateways.SendMessageRequest, response chan<- *gateways.SendMessageResponse) error {
	go func() {
		response <- nil
	}()
	return errors.New("SENDING FAILS")
}
