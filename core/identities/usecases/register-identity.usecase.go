package usecases

import (
	"errors"
	"grabit-cli/core/identities/gateways"
)

type registerIdentityUseCase struct {
	identityGateway gateways.IdentityGateway
}

type RegisterIdentityParams struct {
	Email      string
	Name       string
	PassPhrase string
}

func NewRegisterIdentityUseCase(identityGateway gateways.IdentityGateway) registerIdentityUseCase {
	return registerIdentityUseCase{identityGateway: identityGateway}
}

func (uc *registerIdentityUseCase) Execute(request RegisterIdentityParams) error {
	_, err := uc.identityGateway.LoadCurrent(request.Email)
	if err == nil {
		return errors.New("ALREADY_HAVE_IDENTITY")
	}
	err = uc.identityGateway.Register(gateways.RegisterIdentityRequest{Email: request.Email, Name: request.Name, PassPhrase: request.PassPhrase})
	return err
}
