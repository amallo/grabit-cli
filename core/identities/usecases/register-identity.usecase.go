package usecases

import (
	"errors"
	"grabit-cli/core/identities/gateways"
	message_gateways "grabit-cli/core/message/gateways"
)

type registerIdentityUseCase struct {
	identityGateway gateways.IdentityGateway
	idNameGenerator gateways.NameGenerator
	messageGateway  message_gateways.MessageGateway
}

type RegisterIdentityParams struct {
	Email      string
	PassPhrase string
}
type registerIdentityResult struct {
	Name string
}

func NewRegisterIdentityUseCase(identityGateway gateways.IdentityGateway, nameGenerator gateways.NameGenerator, messageGateway message_gateways.MessageGateway) registerIdentityUseCase {
	return registerIdentityUseCase{identityGateway: identityGateway, idNameGenerator: nameGenerator, messageGateway: messageGateway}
}

func (uc *registerIdentityUseCase) Execute(request RegisterIdentityParams) (*registerIdentityResult, error) {
	_, err := uc.identityGateway.LoadCurrent(request.Email)
	if err == nil {
		return nil, errors.New("ALREADY_HAVE_IDENTITY")
	}
	name := uc.idNameGenerator.Generate(request.Email)

	err = uc.identityGateway.Register(gateways.RegisterIdentityRequest{Email: request.Email, Name: name, PassPhrase: request.PassPhrase})
	if err != nil {
		return nil, errors.New("IDENTITY_REGISTRATION_FAILED")
	}
	result := registerIdentityResult{Name: name}
	return &result, nil
}
