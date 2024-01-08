package adapters

import (
	"grabit-cli/core/message/models"
)

type FakeMessageEncrypter struct {
	WillEncryptTextPlainAs string
}

func (fme FakeMessageEncrypter) EncryptPlainText(publicKey string, text string) (*models.Message, error) {
	return &models.Message{Content: fme.WillEncryptTextPlainAs}, nil
}
