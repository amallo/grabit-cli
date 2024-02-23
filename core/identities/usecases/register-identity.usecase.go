package usecases

import (
	core_errors "grabit-cli/core/common/errors"
	"grabit-cli/core/identities/gateways"
	"grabit-cli/core/identities/models"
)

type useCase struct {
	identityGateway gateways.IdentityGateway
}

type RegisterIdentityArgs struct {
	Email    string
	Password string
}
type RegisterIdentityResult struct {
	Key   string
	Name  string
	Email string
}

func NewRegisterIdentityUseCase(identityGateway gateways.IdentityGateway) useCase {
	return useCase{identityGateway: identityGateway}
}

func (u useCase) Execute(args RegisterIdentityArgs) (*RegisterIdentityResult, core_errors.Error) {
	registerResponse, err := u.identityGateway.Register(gateways.RegisterIdentityRequest{Email: args.Email, Password: args.Password})
	if err != nil {
		return nil, core_errors.Err(models.ErrCannotRegisterIdentity, err)
	}
	return &RegisterIdentityResult{Key: registerResponse.Identity.Key, Name: registerResponse.Identity.Name, Email: registerResponse.Identity.Email}, nil
}
