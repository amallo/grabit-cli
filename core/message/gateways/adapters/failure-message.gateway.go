package adapters

import (
	"errors"
	"grabit-cli/core/message/gateways"
)

type FailureMessageGateway struct {
}

func (fmg *FailureMessageGateway) Send(request gateways.SendMessageRequest) (*gateways.SendMessageResponse, error) {
	return nil, errors.New("SENDING FAILS")
}
