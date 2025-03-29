package passkie

import (
	"errors"

	"github.com/Ajahks/passkie/credentialEncryption/encryption"
	"github.com/Ajahks/passkie/credentialEncryption/hash"
	passwordverification "github.com/Ajahks/passkie/passwordVerification"
	"github.com/Ajahks/passkie/storage/localStorage/credentialsDb"
)

func StoreCredentialsForSite(
	siteBaseUrl string,
	username string,
	masterPassword string,
	credentials map[string]string,
) error {
	ok := passwordverification.VerifyPasswordForUser(username, masterPassword)
	if !ok {
		return errors.New("InvalidMasterPassword: Cannot store credentials because masterPassword for user is incorrect!")
	}

	credentialsList, err := RetrieveCredentialsForSite(siteBaseUrl, username, masterPassword)
	if err != nil {
		credentialsList = []map[string]string{credentials}
	} else {
		credentialsList = append(credentialsList, credentials)
	}

	hashedSite := hash.HashUrl(siteBaseUrl, masterPassword)
	encryptedCredentials := encryption.EncryptCredentials(masterPassword, credentialsList)
	credentialsDb.PutCredentialsForSiteHash(string(hashedSite), username, encryptedCredentials)

	return nil
}

func RetrieveCredentialsForSite(siteBaseUrl string, username string, masterPassword string) ([]map[string]string, error) {
	ok := passwordverification.VerifyPasswordForUser(username, masterPassword)
	if !ok {
		return nil, errors.New("InvalidMasterPassword: Cannot retrieve credentials because masterPassword for user is incorrect!")
	}

	hashedSite := hash.HashUrl(siteBaseUrl, masterPassword)
	encryptedCredentials, err := credentialsDb.GetCredentialsForSiteHash(string(hashedSite), username)
	if err != nil {
		return nil, err
	}

	decryptedCredentialsList, err := encryption.DecryptCredentials[[]map[string]string](masterPassword, encryptedCredentials)
	if err != nil {
		return nil, err
	}

	return decryptedCredentialsList, nil
}

func CreateNewUser(username string, masterPassword string) error {
	err := passwordverification.SetPasswordForNewUser(username, masterPassword)
	if err != nil {
		return err
	}

	return nil
}

func RemoveCredentialsForSite(siteBaseUrl string, username string, masterPassword string) error {
	ok := passwordverification.VerifyPasswordForUser(username, masterPassword)
	if !ok {
		return errors.New("InvalidMasterpassword: Cannot remove credentials because masterPassword is incorrect!")
	}

	hashedSite := hash.HashUrl(siteBaseUrl, masterPassword)
	err := credentialsDb.RemoveCredentialsForSiteHash(string(hashedSite), username)
	if err != nil {
		return err
	}

	return nil
}

func RemoveSingleCredentialsForSite(siteBaseUrl string, username string, masterPassword string, credentialsIndex int) error {
	ok := passwordverification.VerifyPasswordForUser(username, masterPassword)
	if !ok {
		return errors.New("InvalidMasterPassword: Cannot remove credentials because masterPassword for user is incorrect!")
	}

	credentialsList, err := RetrieveCredentialsForSite(siteBaseUrl, username, masterPassword)
	if err != nil {
		return errors.New("No credentials found for site: " + siteBaseUrl)
	} else {
		if credentialsIndex < 0 || credentialsIndex >= len(credentialsList) {
			return errors.New("Invalid index to remove from credentials")
		}

		if len(credentialsList) == 1 {
			return RemoveCredentialsForSite(siteBaseUrl, username, masterPassword)
		}

		credentialsList = append(credentialsList[:credentialsIndex], credentialsList[credentialsIndex+1:]...)
	}

	hashedSite := hash.HashUrl(siteBaseUrl, masterPassword)
	encryptedCredentials := encryption.EncryptCredentials(masterPassword, credentialsList)
	credentialsDb.PutCredentialsForSiteHash(string(hashedSite), username, encryptedCredentials)

	return nil
}

func RemoveUser(username string, masterPassword string) error {
	ok := passwordverification.VerifyPasswordForUser(username, masterPassword)
	if !ok {
		return errors.New("InvalidMasterpassword: Cannot remove credentials because masterPassword is incorrect!")
	}

	err := passwordverification.RemoveUser(username, masterPassword)
	if err != nil {
		return nil
	}

	err = credentialsDb.RemoveUserCredentials(username)
	if err != nil {
		return nil
	}

	return nil
}
