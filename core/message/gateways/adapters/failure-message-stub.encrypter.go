package adapters

import (
	"errors"
	"grabit-cli/core/message/models"
)

type FailureMessageStubEncrypter struct {
	WillFailWith string
}

func (fme FailureMessageStubEncrypter) EncryptPlainText(publicKey string, text string) (*models.Message, error) {
	return nil, errors.New("ENCRYPT TEXT FAILED")
}
