package usecases

import (
	"grabit-cli/core/identities/gateways"
	"grabit-cli/core/identities/models"
)

type getCurrentIdentityUseCase struct {
	identityGateway gateways.IdentityGateway
}

type getCurrentIdentityResponse struct {
	Identity models.Identity
}

type GetCurrentIdentityRequest struct {
	Email string
}

func NewGetCurrentIdentityUseCase(identityGateway gateways.IdentityGateway) getCurrentIdentityUseCase {
	return getCurrentIdentityUseCase{identityGateway: identityGateway}
}

func (uc *getCurrentIdentityUseCase) Execute(request GetCurrentIdentityRequest) *getCurrentIdentityResponse {
	response, err := uc.identityGateway.LoadCurrent(request.Email)
	if err == nil {
		return &getCurrentIdentityResponse{models.Identity{Email: request.Email, Name: response.Name}}
	}
	return nil
}
