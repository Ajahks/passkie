// URL DB for unencrypted URLs.
// This DB should stay LOCAL, and should never be migrated to an online DB since the urls are non encrypted.
// The purpose of this DB is just for quality of life for front end clients to access a list of sites.  
package urldb

import (
	"encoding/base64"
	"os"

	"github.com/Ajahks/passkie/storage/localStorage"
)

const FILE_NAME = "urls.txt"

func getEncodedUsername(username string) string {
    return base64.StdEncoding.EncodeToString([]byte(username))
}

func getFilePath(username string) string {
    // Not the most secure storing the usernames as base64 encoding, but slightly better than plaintext
    return localstorage.DB_PATH() + "/" + getEncodedUsername(username) + "/" + FILE_NAME
}

func AddActiveUrlForUser(url string, username string) {
    data, err := os.ReadFile(getFilePath(username))
    if err != nil {
        activeUrlMap := make(map[string]bool)
        activeUrlMap[url] = true 

        localstorage.WriteMapToFile(activeUrlMap, FILE_NAME, getEncodedUsername(username))
    } else {
        activeUrlMap := localstorage.DeserializeFileData[bool](data) 
        activeUrlMap[url] = true 

        localstorage.WriteMapToFile(activeUrlMap, FILE_NAME, getEncodedUsername(username))
    }
}

func IsUrlActiveForUser(url string, username string) bool {
    data, err := os.ReadFile(getFilePath(username))
    if err != nil {
        return false 
    }

    activeUrlMap := localstorage.DeserializeFileData[bool](data) 
    activeUrl, ok := activeUrlMap[url]
    if !ok {
        return false 
    }

    return activeUrl
}

func ListUrlsForUser(username string) ([]string, error) {
	keys := make([]string, 0)
    data, err := os.ReadFile(getFilePath(username))
    if err != nil {
        return keys, err 
    }

	urlMap := localstorage.DeserializeFileData[bool](data)
	for key := range urlMap {
		if (urlMap[key] == true) {
			keys = append(keys, key)	
		}
	}

	return keys, nil
}

func RemoveActiveUrlForUser(url string, username string) error {
    data, err := os.ReadFile(getFilePath(username))
    if err != nil {
        return err
    }
 
    activeUrlMap := localstorage.DeserializeFileData[bool](data) 
    delete(activeUrlMap, url)

    localstorage.WriteMapToFile(activeUrlMap, FILE_NAME, getEncodedUsername(username))
    return nil
}

