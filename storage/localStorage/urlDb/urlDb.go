// URL DB for unencrypted URLs.
// This DB should stay LOCAL, and should never be migrated to an online DB since the urls are non encrypted.
// The purpose of this DB is just for quality of life for front end clients to access a list of sites.  
package urldb

import (
	"os"

	"github.com/Ajahks/passkie/storage/localStorage"
)

const FILE_NAME = "urls.txt"

func getFilePath() string {
    return localstorage.DB_PATH() + "/" + FILE_NAME 
}

func AddActiveUrl(url string) {
    data, err := os.ReadFile(getFilePath())
    if err != nil {
        activeUrlMap := make(map[string]bool)
        activeUrlMap[url] = true 

        localstorage.WriteMapToFile(activeUrlMap, FILE_NAME)
    } else {
        activeUrlMap := localstorage.DeserializeFileData[bool](data) 
        activeUrlMap[url] = true 

        localstorage.WriteMapToFile(activeUrlMap, FILE_NAME)
    }
}

func IsUrlActive(url string) bool {
    data, err := os.ReadFile(getFilePath())
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

func RemoveActiveUrl(url string) error {
    data, err := os.ReadFile(getFilePath())
    if err != nil {
        return err
    }
 
    activeUrlMap := localstorage.DeserializeFileData[bool](data) 
    delete(activeUrlMap, url)

    localstorage.WriteMapToFile(activeUrlMap, FILE_NAME)
    return nil
}

