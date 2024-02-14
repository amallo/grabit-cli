package usecases

import (
	core_errors "grabit-cli/core/common/errors"
	"grabit-cli/core/identities/gateways"
	"grabit-cli/core/identities/models"
)

type registerIdentityUseCase struct {
	identityGateway gateways.IdentityGateway
	idNameGenerator gateways.NameGenerator
}

type RegisterIdentityParams struct {
	Email      string
	PassPhrase string
}
type registerIdentityResult struct {
	Name string
}

func NewRegisterIdentityUseCase(identityGateway gateways.IdentityGateway, nameGenerator gateways.NameGenerator) registerIdentityUseCase {
	return registerIdentityUseCase{identityGateway: identityGateway, idNameGenerator: nameGenerator}
}

func (uc *registerIdentityUseCase) Execute(request RegisterIdentityParams) (*registerIdentityResult, core_errors.Error) {
	_, err := uc.identityGateway.LoadCurrent(request.Email)
	if err == nil {
		return nil, core_errors.Err(models.ErrIdentityAlreadyRegistered, err)
	}
	name := uc.idNameGenerator.Generate(request.Email)

	err = uc.identityGateway.Register(gateways.RegisterIdentityRequest{Email: request.Email, Name: name, PassPhrase: request.PassPhrase})
	if err != nil {
		return nil, core_errors.Err(models.ErrCannotRegisterIdentity, err)
	}
	result := registerIdentityResult{Name: name}
	return &result, nil
}
