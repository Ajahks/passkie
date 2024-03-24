package credentialsDb

import (
	"errors"
	"os"

	localstorage "github.com/Ajahks/passkie/storage/localStorage"
)

const FILE_NAME = "credentialsDB.txt"

func PutCredentialsForSiteHash(sitehash string, username string, encryptedCredentials []byte) {
    data, err := os.ReadFile(getFilePath(username, FILE_NAME))
    if err != nil {
        siteCredentialsMap := make(map[string][]byte)
        siteCredentialsMap[sitehash] = encryptedCredentials 

        localstorage.WriteMapToFile(siteCredentialsMap, FILE_NAME, getEncodedUsername(username))

    } else {
        siteCredentialsMap := localstorage.DeserializeFileData[[]byte](data) 
        siteCredentialsMap[sitehash] = encryptedCredentials 

        localstorage.WriteMapToFile(siteCredentialsMap, FILE_NAME, getEncodedUsername(username))
    }
}

func GetCredentialsForSiteHash(sitehash string, username string) ([]byte, error) {
    data, err := os.ReadFile(getFilePath(username, FILE_NAME))
    if err != nil {
        return nil, err 
    }

    siteCredentialsMap := localstorage.DeserializeFileData[[]byte](data) 
    encryptedCredentials, ok := siteCredentialsMap[sitehash]
    if !ok {
        return nil, errors.New("Site does not exist in the DB!")
    }

    return encryptedCredentials, nil
}

func RemoveCredentialsForSiteHash(sitehash string, username string) error {
    data, err := os.ReadFile(getFilePath(username, FILE_NAME))
    if err != nil {
        return err
    }
 
    siteCredentialsMap := localstorage.DeserializeFileData[[]byte](data) 
    delete(siteCredentialsMap, sitehash)

    localstorage.WriteMapToFile(siteCredentialsMap, FILE_NAME, getEncodedUsername(username))
    return nil
}

func RemoveUserCredentials(username string) error {
    err := os.Remove(getFilePath(username, FILE_NAME))
    if err != nil { return err }

    err = os.Remove(localstorage.DB_PATH() + "/" + getEncodedUsername(username))
    if err != nil { return err } 
    return nil
}

