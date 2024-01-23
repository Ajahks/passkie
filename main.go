package main

import (
	"fmt"

	"github.com/Ajahks/PasskieCredentialEncryption/encryption"
)

func main() {
	masterPassword := "testPassword"
	var credentials map[string]string
	credentials = make(map[string]string)

	credentials["testUser"] = "testUserPassword"
    credentials["testUser2"] = "testUserPassword2"
    fmt.Printf("Credentials map: %v\n", credentials)

    encryptedCredentials := encryption.EncryptCredentials(masterPassword, credentials)
    fmt.Printf("Encrypted Credentials map: %v\n", string(encryptedCredentials))

    decryptedCredentials := encryption.DecryptCredentials(masterPassword, encryptedCredentials)
    fmt.Printf("Decrypted Credentials map: %v\n", decryptedCredentials)
}

