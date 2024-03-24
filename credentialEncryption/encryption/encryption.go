package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/gob"
	"fmt"
	"log"

	passkieHash "github.com/Ajahks/passkie/credentialEncryption/hash"
)

type CredentialDecodingError struct {
    cause error
}

func (e *CredentialDecodingError) Error() string {
    return fmt.Sprintf("Failed to decode credentials: %v", e.cause)
}

func EncryptCredentials[T any](masterPassword string, credentials T) []byte {
    byteBuffer := new(bytes.Buffer)
    e := gob.NewEncoder(byteBuffer)
    err := e.Encode(credentials)
    if err != nil {
       log.Fatal("Failed to encode credentials")
    }
    credentialBytes := byteBuffer.Bytes()

    passwordHash := passkieHash.HashPassword(masterPassword)

    aesCipher, err := aes.NewCipher(passwordHash)
    if err != nil {
        log.Fatal("Error found when initializing cipher with private key", err)
    }
    gcm, err := cipher.NewGCM(aesCipher)
    if err != nil {
        log.Fatal("Error while initializing GCM cipher", err)
    }
    nonce := make([]byte, gcm.NonceSize())
    _, err = rand.Read(nonce)
    if err != nil {
        log.Fatal("Error while creating random nonce value", err)
    }

    ciphertext := gcm.Seal(nonce, nonce, credentialBytes, nil)

    return ciphertext 
}

func DecryptCredentials[T any](masterPassword string, encryptedCredentials []byte) (T, error) {
    passwordHash := passkieHash.HashPassword(masterPassword) 
    var decodedCredentials T 

    aesCipher, err := aes.NewCipher(passwordHash)
    if err != nil {
        return decodedCredentials, fmt.Errorf("Error when initializing cipher with private key: %w", err)
    }
    gcm, err := cipher.NewGCM(aesCipher)
    if err != nil {
        return decodedCredentials, fmt.Errorf("Error while initializing GCM cipher: %w", err) 
    }

    nonceSize := gcm.NonceSize()
    nonce, encryptedCredentials := encryptedCredentials[:nonceSize], encryptedCredentials[nonceSize:]

    decryptedCredentials, err := gcm.Open(nil, nonce, encryptedCredentials, nil)
    if err != nil {
        return decodedCredentials, fmt.Errorf("Error while decrypting credentials: %w", err)
    }

    byteBuffer := new(bytes.Buffer)
    byteBuffer.Write(decryptedCredentials)
    decoder := gob.NewDecoder(byteBuffer)

    err = decoder.Decode(&decodedCredentials)
    if err != nil {
        return decodedCredentials, &CredentialDecodingError{ cause: err } 
    }

    return decodedCredentials, nil
}

