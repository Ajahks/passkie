package encryption

import (
    "testing"
)

func TestEncryptionDecryption(t *testing.T) {
    masterPassword := "testMasterPassword"
	var originalCredentials map[string]string
	originalCredentials = make(map[string]string)
    originalCredentials["testUser1"] = "testPassword1"

    encryptedCredentials := EncryptCredentials(masterPassword, originalCredentials)
    decryptedCredentials := DecryptCredentials(masterPassword, encryptedCredentials)

    if len(decryptedCredentials) != 1 {
        t.Fatalf(
            `Length of decrypted map does not map length of original map.  Original: %v, Decrypted: %v`,
            len(originalCredentials),
            len(decryptedCredentials),
        )
    }
    decryptedPassword, ok := decryptedCredentials["testUser1"]
    if !ok {
        t.Fatal("decryptedCredentials does not contain testUser1")
    }
    if originalCredentials["testUser1"] != decryptedPassword {
        t.Fatalf(
            `Decrypted password does not match original for testUser1: Expected: %v, Received: %v`, 
            originalCredentials["testUser1"], 
            decryptedPassword,
        )
    }   
}

func TestEncryptionDecryptionWithMultipleUsers(t *testing.T) {
    masterPassword := "testMasterPassword"
	var originalCredentials map[string]string
	originalCredentials = make(map[string]string)
    originalCredentials["testUser1"] = "testPassword1"
    originalCredentials["testUser2"] = "testPassword2"
    originalCredentials["testUser3"] = "testPassword3"

    encryptedCredentials := EncryptCredentials(masterPassword, originalCredentials)
    decryptedCredentials := DecryptCredentials(masterPassword, encryptedCredentials)

    if len(decryptedCredentials) != 3 {
        t.Fatalf(
            `Length of decrypted map does not map length of original map.  Original: %v, Decrypted: %v`,
            len(originalCredentials),
            len(decryptedCredentials),
        )
    }
    decryptedPassword1, ok := decryptedCredentials["testUser1"]
    if !ok {
        t.Fatal("decryptedCredentials does not contain testUser1")
    }
    decryptedPassword2, ok := decryptedCredentials["testUser2"]
    if !ok {
        t.Fatal("decryptedCredentials does not contain testUser2")
    }
    decryptedPassword3, ok := decryptedCredentials["testUser3"]
    if !ok {
        t.Fatal("decryptedCredentials does not contain testUser3")
    }
    if originalCredentials["testUser1"] != decryptedPassword1 {
        t.Fatalf(
            `Decrypted password does not match original for testUser1: Expected: %v, Received: %v`, 
            originalCredentials["testUser1"], 
            decryptedPassword1,
        )
    }   
    if originalCredentials["testUser2"] != decryptedPassword2 {
        t.Fatalf(
            `Decrypted password does not match original for testUser2: Expected: %v, Received: %v`, 
            originalCredentials["testUser2"], 
            decryptedPassword2,
        )
    }   
    if originalCredentials["testUser3"] != decryptedPassword3 {
        t.Fatalf(
            `Decrypted password does not match original for testUser3: Expected: %v, Received: %v`, 
            originalCredentials["testUser3"], 
            decryptedPassword3,
        )
    }   
}
