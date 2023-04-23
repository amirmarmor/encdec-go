package asymmetric

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encdec/log"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
)

type PublicKey struct {
	publicKey *rsa.PublicKey
}

type Keys struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func CreateKeyPair() (*Keys, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, err
	}

	k := Keys{
		privateKey: privateKey,
		publicKey:  &privateKey.PublicKey,
	}

	return &k, nil
}

func LoadPrivateKey() (*rsa.PrivateKey, error) {
	pemFile, err := os.ReadFile("./priv.pem")
	if err != nil {
		return nil, err
	}
	log.Info(string(pemFile))

	pemBlock, _ := pem.Decode(pemFile)
	privateKey, err := x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
	if err != nil {
		return nil, err
	}

	log.Info("private key loaded")
	return privateKey, nil
}

func LoadPublicKey() (*rsa.PublicKey, error) {
	pemFile, err := os.ReadFile("./pub.pem")
	if err != nil {
		return nil, err
	}
	log.Info(string(pemFile))

	pemBlock, _ := pem.Decode(pemFile)

	publicKey, err := x509.ParsePKIXPublicKey(pemBlock.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to generate public key")
	}

	return rsaPublicKey, nil
}

func (k *Keys) Save() error {
	err := k.savePrivateKey()
	if err != nil {
		return fmt.Errorf("failed to save private key: %v", err)
	}

	err = k.savePublicKey()
	if err != nil {
		return fmt.Errorf("failed to save public key: %v", err)
	}

	return nil
}

func (k *Keys) savePrivateKey() error {
	privateKeyBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(k.privateKey),
		},
	)

	err := os.WriteFile("./priv.pem", privateKeyBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (k *Keys) savePublicKey() error {
	pubASN1, err := x509.MarshalPKIXPublicKey(k.publicKey)
	if err != nil {
		return err
	}

	publicKeyBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubASN1,
		},
	)

	err = os.WriteFile("./pub.pem", publicKeyBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func EncryptString(publicKey *rsa.PublicKey, clearText string) (*string, error) {
	sha := sha256.New()
	random := rand.Reader

	encryptedBytes, err := rsa.EncryptOAEP(
		sha,
		random,
		publicKey,
		[]byte(clearText),
		nil,
	)

	if err != nil {
		return nil, err
	}

	encryptedBase64 := base64.StdEncoding.EncodeToString(encryptedBytes)
	return &encryptedBase64, nil
}

func DecryptString(privateKey *rsa.PrivateKey, base64Cipher string) (*string, error) {
	encryptedBytes, err := base64.StdEncoding.DecodeString(base64Cipher)
	if err != nil {
		return nil, err
	}

	sha := sha256.New()
	random := rand.Reader

	clearTextBytes, err := rsa.DecryptOAEP(
		sha,
		random,
		privateKey,
		encryptedBytes,
		nil,
	)

	if err != nil {
		return nil, err
	}

	clearText := string(clearTextBytes)
	return &clearText, nil
}
