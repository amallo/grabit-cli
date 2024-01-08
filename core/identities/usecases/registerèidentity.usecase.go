package usecases

type registerIdentityUseCase struct {
}

type RegisterIdentityResponse struct {
	Email     string
	PublicKey string
}

func NewRegisterIdentityUseCase() registerIdentityUseCase {
	return registerIdentityUseCase{}
}

func (uc *registerIdentityUseCase) Execute() (*RegisterIdentityResponse, error) {
	return &RegisterIdentityResponse{Email: "audie@foo.com", PublicKey: "public key"}, nil
}
