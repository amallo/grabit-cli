package adapters

import (
	"errors"
	"grabit-cli/core/message/gateways"
)

type FailureMessageGateway struct {
}

func (fmg *FailureMessageGateway) Drop(request gateways.DropMessageRequest) (*gateways.DropMessageResponse, error) {
	return nil, errors.New("SENDING FAILS")
}
