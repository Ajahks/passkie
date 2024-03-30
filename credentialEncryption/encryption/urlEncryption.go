package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"log"

	passkieHash "github.com/Ajahks/passkie/credentialEncryption/hash"
)

func EncryptUrl(masterPassword string, url string) []byte {
    urlBytes := []byte(url) 

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

    ciphertext := gcm.Seal(nonce, nonce, urlBytes, nil)

    return ciphertext 
}

func DecryptUrl(masterPassword string, encryptedUrl []byte) (string, error) {
    passwordHash := passkieHash.HashPassword(masterPassword) 

    aesCipher, err := aes.NewCipher(passwordHash)
    if err != nil {
        return "", fmt.Errorf("Error when initializing cipher with private key: %w", err)
    }
    gcm, err := cipher.NewGCM(aesCipher)
    if err != nil {
        return "", fmt.Errorf("Error while initializing GCM cipher: %w", err) 
    }

    nonceSize := gcm.NonceSize()
    nonce, encryptedUrl := encryptedUrl[:nonceSize], encryptedUrl[nonceSize:]

    decryptedUrl, err := gcm.Open(nil, nonce, encryptedUrl, nil)
    if err != nil {
        return "", fmt.Errorf("Error while decrypting credentials: %w", err)
    }

    decodedUrl := string(decryptedUrl[:])
    return decodedUrl, nil
}

