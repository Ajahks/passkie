package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/gob"
	passkieHash "github.com/Ajahks/passkie/credentialEncryption/hash"
	"log"
)

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

func DecryptCredentials[T any](masterPassword string, encryptedCredentials []byte) T {
    passwordHash := passkieHash.HashPassword(masterPassword) 

    aesCipher, err := aes.NewCipher(passwordHash)
    if err != nil {
        log.Fatal("Error found when initializing cipher with private key", err)
    }
    gcm, err := cipher.NewGCM(aesCipher)
    if err != nil {
        log.Fatal("Error while initializing GCM cipher", err)
    }

    nonceSize := gcm.NonceSize()
    nonce, encryptedCredentials := encryptedCredentials[:nonceSize], encryptedCredentials[nonceSize:]

    decryptedCredentials, err := gcm.Open(nil, nonce, encryptedCredentials, nil)
    if err != nil {
        log.Fatal("Error while decoding credentials", err)
    }

    byteBuffer := new(bytes.Buffer)
    byteBuffer.Write(decryptedCredentials)
    var decodedMap T 
    decoder := gob.NewDecoder(byteBuffer)

    err = decoder.Decode(&decodedMap)
    if err != nil {
        log.Fatal("Error while decoding decryptedCrednetials into map", err)
    }

    return decodedMap
}
