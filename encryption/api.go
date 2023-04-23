package encryption

import (
	asymmetric "encdec/encryption/asymetric"
	"fmt"
	"os"
)

func GenerateKeys() error {
	keys, err := asymmetric.CreateKeyPair()
	if err != nil {
		return fmt.Errorf("failed to create keys: %v", err)
	}

	err = keys.Save()
	if err != nil {
		return fmt.Errorf("failed to save keys: %v", err)
	}

	return nil
}

func EncryptMessage(clearText string) error {
	publicKey, err := asymmetric.LoadPublicKey()
	if err != nil {
		return fmt.Errorf("failed to load private key: %v", err)
	}

	encrypted, err := asymmetric.EncryptString(publicKey, clearText)
	if err != nil {
		return fmt.Errorf("failed to enctypt message: %v", err)
	}

	err = os.WriteFile("./enc.txt", []byte(*encrypted), 0644)
	if err != nil {
		return fmt.Errorf("failed to save encrypted text to file: %v", err)
	}

	return nil
}

func DecryptMessage() error {
	privateKey, err := asymmetric.LoadPrivateKey()
	if err != nil {
		return fmt.Errorf("failed to load private key: %v", err)
	}

	encryptedText, err := os.ReadFile("./enc.txt")
	if err != nil {
		return fmt.Errorf("faild to load encrypted text: %v", err)
	}

	clearText, err := asymmetric.DecryptString(privateKey, string(encryptedText))
	if err != nil {
		return fmt.Errorf("failed to enctypt message: %v", err)
	}

	err = os.WriteFile("./clear.txt", []byte(*clearText), 0644)
	if err != nil {
		return fmt.Errorf("failed to write clear text: %v", err)
	}
	return nil
}
