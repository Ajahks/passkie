package encryption_test

import (
	"bytes"
	"testing"

	"github.com/Ajahks/passkie/credentialEncryption/encryption"
)

func TestEncryptUrlDecryptUrlReturnsOriginalResponse(t *testing.T) {
    password := "testPassword"
    url := "testUrl.com"

    encryptedUrl := encryption.EncryptUrl(url, password) 
    decryptedUrl, err := encryption.DecryptUrl(encryptedUrl, password)
    
    if err != nil {
        t.Fatalf("Failed to decrypt url: %v", err)
    }
    if decryptedUrl != url {
        t.Fatalf("Decrypted url does not match the original url! Original: %v, Decrytped: %v", url, decryptedUrl)
    }
}

func TestEncryptUrlReturnsADifferentStringFromTheOriginalUrl(t *testing.T) {
    password := "testPassword"
    url := "testUrl.com"

    encryptedUrl := encryption.EncryptUrl(url, password) 

    if url == string(encryptedUrl) {
        t.Fatal("Encrypted url should not match the original url") 
    }
}

func TestEncryptUrlAndDecryptUrlWithDifferentPasswordsReturnsDifferentResults(t *testing.T) {
    password1 := "testPass1"
    password2 := "testPass2"
    url := "testUrl.com"

    encryptedUrl := encryption.EncryptUrl(url, password1)
    decryptedUrl, _ := encryption.DecryptUrl(encryptedUrl, password2)

    if decryptedUrl == url {
        t.Fatal("Decrypted url with a different password should not match the original url")
    }
}

func TestEncryptUrlTwiceWithDifferentPasswordReturnsADifferentResult(t *testing.T) {
    password1 := "testPass1"
    password2 := "testPass2"
    url := "testUrl.com"

    encryptedUrl1 := encryption.EncryptUrl(url, password1)
    encryptedUrl2 := encryption.EncryptUrl(url, password2)

    if bytes.Equal(encryptedUrl1, encryptedUrl2) {
        t.Fatal("Encrypted urls with different passwords should not lead to matching encryptions!")
    }
}

