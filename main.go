package main

import (
	"encdec/encryption"
	"encdec/log"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Error("Must give a command")
		return
	}

	cmd := os.Args[1]
	switch cmd {
	case "g":
		err := encryption.GenerateKeys()
		if err != nil {
			panic(err)
		}
		log.V1("Keys generated and saved as pem file")
	case "s":
		if len(os.Args) < 3 || os.Args[2] == "" {
			log.Fatal("no message given for encryption")
		}
		err := encryption.EncryptMessage(os.Args[2])
		if err != nil {
			log.Fatal("failed to encrypt", err)
		}
	case "l":
		err := encryption.DecryptMessage()
		if err != nil {
			log.Fatal("failed to decrypt", err)
		}
	case "":
	default:
		fmt.Println("Invalid command")
	}
}
