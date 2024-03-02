package credentialsDb

import (
	"encoding/base64"
	"errors"
	"log"
	"os"

	localstorage "github.com/Ajahks/passkie/storage/localStorage"
)

const FILE_NAME = "credentialsDB.txt"

func getEncodedUsername(username string) string {
    return base64.StdEncoding.EncodeToString([]byte(username))
}

func getLocalFilePath(username string) string {
    // Not the most secure storing the usernames as base64 encoding, but slightly better than plaintext
    return localstorage.LOCAL_DIR + "/" + getEncodedUsername(username) + "/" + FILE_NAME
}

func PutCredentialsForSiteHash(sitehash string, username string, encryptedCredentials []byte) {
    data, err := os.ReadFile(getLocalFilePath(username))
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
    data, err := os.ReadFile(getLocalFilePath(username))
    if err != nil {
        return nil, err 
    }

    siteCredentialsMap := localstorage.DeserializeFileData[[]byte](data) 
    encryptedCredentials, ok := siteCredentialsMap[sitehash]
    if !ok {
        log.Printf("Site %s does not exist in the DB!\n", sitehash)
        return nil, errors.New("Site does not exist in the DB!")
    }

    return encryptedCredentials, nil
}

func RemoveCredentialsForSiteHash(sitehash string, username string) {
    data, err := os.ReadFile(getLocalFilePath(username))
    if err != nil {
        log.Printf("Failed to read DB file: %s\n", err)
        return
    }
 
    siteCredentialsMap := localstorage.DeserializeFileData[[]byte](data) 
    delete(siteCredentialsMap, sitehash)

    localstorage.WriteMapToFile(siteCredentialsMap, FILE_NAME, getEncodedUsername(username))
}

