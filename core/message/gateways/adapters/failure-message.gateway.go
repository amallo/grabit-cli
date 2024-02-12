package adapters

import (
	"errors"
	"grabit-cli/core/message/gateways"
)

type FailureMessageGateway struct {
	GrabMessageFailure error
}

func (fmg *FailureMessageGateway) Drop(request gateways.DropMessageRequest) (*gateways.DropMessageResponse, error) {
	return nil, errors.New("DROP MESSAGE FAILS")
}

func (fmg *FailureMessageGateway) Grab(request gateways.GrabMessageRequest) (*gateways.GrabMessageResponse, error) {
	return nil, fmg.GrabMessageFailure
}
