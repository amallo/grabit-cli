package adapters

import (
	"errors"
	"fmt"
	"grabit-cli/core/identities/gateways"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FsIdentityGateway struct {
}

func (lig *FsIdentityGateway) LoadCurrent() (*gateways.LoadIdentityResponse, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error while reading home golder:", err)
		return nil, errors.New("IDENTITY NOT FOUND")
	}
	identityFilePath := filepath.Join(homeDir, ".grabit", ".identity")

	content, err := ioutil.ReadFile(identityFilePath)
	if err != nil {
		fmt.Println("Error while reading identity file:", err)
		return nil, errors.New("IDENTITY NOT FOUND")
	}

	return &gateways.LoadIdentityResponse{Email: string(content)}, nil
}

func (lig *FsIdentityGateway) Register(request gateways.RegisterIdentityRequest) error {
	// home folder
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error while reading home golder :", err)
		return errors.New("CANNOT_READ_HOME")
	}
	identityFilePath := filepath.Join(homeDir, ".grabit", ".identity")

	emailContent := []byte(fmt.Sprintf("%s\n", request.Email))

	err = os.MkdirAll(filepath.Join(homeDir, ".grabit"), 0755)
	if err != nil {
		fmt.Println("Error while creating ~/.grabit :", err)
		return errors.New("CANNOT_CREATE_GRABIT_HOME")
	}

	file, err := os.Create(identityFilePath)
	if err != nil {
		fmt.Println("Error while creating .identity file :", err)
		return errors.New("CANNOT_CREATE_IDENTITY")
	}
	defer file.Close()

	_, err = file.Write(emailContent)
	if err != nil {
		fmt.Println("Error while writing.identity :", err)
		return errors.New("CANNOT_WRITE_IDENTITY")
	}
	return nil
}
