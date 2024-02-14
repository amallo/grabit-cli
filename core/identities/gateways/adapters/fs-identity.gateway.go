package adapters

import (
	"errors"
	common_gateways "grabit-cli/core/common/gateways"
	"grabit-cli/core/identities/gateways"
)

type FsIdentityGateway struct {
	identityPath       IdentityPath
	publicKeyFileName  string
	privateKeyFileName string
	logger             common_gateways.Logger
}

type IdentityPath struct {
	Dir      string
	FileName string
}

const (
	rsaBits = 2048
)

func NewFsIdentityGateway(identityPath IdentityPath, publicKeyFileName string, privateKeyFileName string, logger common_gateways.Logger) FsIdentityGateway {
	return FsIdentityGateway{identityPath: identityPath, publicKeyFileName: publicKeyFileName, privateKeyFileName: privateKeyFileName, logger: logger}
}

func (lig *FsIdentityGateway) LoadCurrent(email string) (*gateways.LoadIdentityResponse, error) {
	lig.logger.Trace("Checking home folder...")
	return nil, errors.New("ERROR")
	/*homeDir, err := os.UserHomeDir()
	if err != nil {
		lig.logger.Error("Cannot read home folder", err)
		return nil, errors.New("HOME NOT FOUND")
	}
	lig.logger.Info("Home folder ok")

	identityFilePath := filepath.Join(homeDir, lig.identityPath.Dir, lig.identityPath.FileName)
	nameContent, err := os.ReadFile(identityFilePath)
	if err != nil {
		lig.logger.Error("Cannot read identity", err)
		return nil, errors.New("IDENTITY FILE NOT FOUND")
	}

	lig.logger.Trace("Checking public key file")
	publicKeyFilePath := filepath.Join(homeDir, lig.identityPath.Dir, lig.publicKeyFileName)
	publicKeyContent, err := os.ReadFile(publicKeyFilePath)
	if err != nil {
		lig.logger.Error("Cannot read public key", err)
		return nil, errors.New("PUBLIC KEY NOT FOUND")
	}

	key, err := crypto.NewKeyFromArmored(string(publicKeyContent))
	if err != nil {
		lig.logger.Error("Cannot open public key", err)
		return nil, errors.New("CANNOT OPEN PUBLIC KEY")
	}
	if key == nil {
		lig.logger.Error("Cannot read public key", err)
		return nil, errors.New("CANNOT read PUBLIC KEY")
	}
	publicEntity := key.GetEntity()

	cleanName := string(nameContent[:])
	for key, value := range publicEntity.Identities {
		if key == cleanName && value.UserId != nil && value.UserId.Email == email {
			return &gateways.LoadIdentityResponse{Name: cleanName}, nil
		}
	}
	err = errors.New("PUBLIC KEY NOT MATCH")
	lig.logger.Error("Public key not match your email identity ", err)
	return nil, err*/

}

func (lig *FsIdentityGateway) Register(request gateways.RegisterIdentityRequest) error {
	/*homeDir, err := os.UserHomeDir()
	if err != nil {
		lig.logger.Error("Error while reading home golder :", err)
		return errors.New("CANNOT_READ_HOME")
	}
	lig.logger.Trace("Creating your identity file")
	identityFilePath := filepath.Join(homeDir, lig.identityPath.Dir, lig.identityPath.FileName)

	nameContent := []byte(fmt.Sprintf("%s\n", request.Name))

	err = os.MkdirAll(filepath.Join(homeDir, lig.identityPath.Dir), 0755)
	if err != nil {
		lig.logger.Error("Error while creating ~/.grabit", err)
		return errors.New("CANNOT_CREATE_GRABIT_HOME")
	}

	identityFile, err := os.Create(identityFilePath)
	if err != nil {
		lig.logger.Error("Error while creating .identity file", err)
		return errors.New("CANNOT_CREATE_IDENTITY")
	}
	defer identityFile.Close()
	_, err = identityFile.Write(nameContent)
	if err != nil {
		lig.logger.Error("Error while writing .identity file", err)
		return errors.New("CANNOT_WRITE_IDENTITY")
	}
	lig.logger.Info("Your .identity file is ready")

	privateKey, err := crypto.GenerateKey(request.Name, request.Email, request.PassPhrase, rsaBits)
	if err != nil {
		lig.logger.Error("Error while generating private key", err)
		return errors.New("CANNOT_GENERATE_PRIVATE_KEY")
	}
	armoredPrivateKey, err := privateKey.Armor()
	if err != nil {
		lig.logger.Error("Error while generating armored private key", err)
		return errors.New("CANNOT_GENERATE_ARMORED PRIVATE_KEY")
	}
	privateKeyFilePath := filepath.Join(homeDir, lig.identityPath.Dir, lig.privateKeyFileName)
	privateKeyFile, err := os.Create(privateKeyFilePath)
	if err != nil {
		lig.logger.Error("Error while creating .identity file", err)
		return errors.New("CANNOT_CREATE_PRIVATE_KEY")
	}
	defer privateKeyFile.Close()

	_, err = privateKeyFile.Write([]byte(armoredPrivateKey))
	if err != nil {
		lig.logger.Error("Error while writing private key file", err)
		return errors.New("CANNOT_WRITE_PRIVATE_KEY")
	}

	publicKeyFilePath := filepath.Join(homeDir, lig.identityPath.Dir, lig.publicKeyFileName)
	publicKeyFile, err := os.Create(publicKeyFilePath)
	if err != nil {
		lig.logger.Error("Error while creating .identity file", err)
		return errors.New("CANNOT_CREATE_IDENTITY")
	}
	defer publicKeyFile.Close()
	publicKey, err := privateKey.GetArmoredPublicKey()
	if err != nil {
		lig.logger.Error("Error while getting public armored key", err)
		return errors.New("CANNOT_WRITE_PUBLIC_KEY")
	}

	_, err = publicKeyFile.Write([]byte(publicKey))
	if err != nil {
		lig.logger.Error("Error while writing private key file", err)
		return errors.New("CANNOT_WRITE_PUBLIC_KEY")
	}*/

	return nil
}
