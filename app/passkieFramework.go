package passkieApp

import (
	"errors"
	"log"

	"github.com/Ajahks/Passkie/credentialEncryption/encryption"
	"github.com/Ajahks/Passkie/credentialEncryption/hash"
	passwordverification "github.com/Ajahks/Passkie/passwordVerification"
	"github.com/Ajahks/Passkie/storage/localStorage/credentialsDb"
)

func StoreCredentialsForSite(
    siteBaseUrl string,
    username string,
    masterPassword string,
    credentials map[string]string,
) error {
    ok := passwordverification.VerifyPasswordForUser(username, masterPassword)
    if !ok {
        log.Printf("Cannot store credentials because masterPassword for user is incorrect!\n")
        return errors.New("InvalidMasterPassword")
    }

    hashedSite := hash.HashUrl(siteBaseUrl, masterPassword)
    encryptedCredentials := encryption.EncryptCredentials(masterPassword, credentials)
    credentialsDb.PutCredentialsForSiteHash(string(hashedSite), username, encryptedCredentials)
    
    return nil
}

func RetrieveCredentialsForSite(siteBaseUrl string, username string, masterPassword string) (map[string]string, error) {
    ok := passwordverification.VerifyPasswordForUser(username, masterPassword)
    if !ok {
        log.Printf("Cannot store credentials because masterPassword for user is incorrect!\n")
        return nil, errors.New("InvalidMasterPassword")
    }

    hashedSite := hash.HashUrl(siteBaseUrl, masterPassword)
    encryptedCredentials, err := credentialsDb.GetCredentialsForSiteHash(string(hashedSite), username)
    if err != nil {
        log.Println("Failed to retrieve credentials for site!")
        return nil, err
    }

    decryptedCredentials := encryption.DecryptCredentials(masterPassword, encryptedCredentials) 
    return decryptedCredentials, nil
}

func CreateNewUser(username string, masterPassword string) error {
    err := passwordverification.SetPasswordForNewUser(username, masterPassword) 
    if err != nil {
        log.Printf("Failed to create new user: %s\n", username)
        return err 
    }

    return nil
}

