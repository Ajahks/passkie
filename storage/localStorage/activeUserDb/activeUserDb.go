package activeuserdb

import (
	"log"
	"os"

	"github.com/Ajahks/Passkie/storage/localStorage"
)

const FILE_NAME = "activeUsers.txt"
const LOCAL_FILE_PATH = localstorage.LOCAL_DIR + "/" + FILE_NAME

func AddActiveUser(userhash string) {
    data, err := os.ReadFile(LOCAL_FILE_PATH)
    if err != nil {
        activeUserMap := make(map[string]bool)
        activeUserMap[userhash] = true 

        localstorage.WriteMapToFile(activeUserMap, FILE_NAME)

    } else {
        activeUserMap := localstorage.DeserializeFileData[bool](data) 
        activeUserMap[userhash] = true 

        localstorage.WriteMapToFile(activeUserMap, FILE_NAME)
    }
}

func IsUserHashActive(userhash string) bool {
    data, err := os.ReadFile(LOCAL_FILE_PATH)
    if err != nil {
        return false 
    }

    activeUserMap := localstorage.DeserializeFileData[bool](data) 
    activeUser, ok := activeUserMap[userhash]
    if !ok {
        log.Printf("User %s does not exist in the DB!\n", userhash)
        return false 
    }

    return activeUser
}

func RemoveActiveUser(userhash string) {
    data, err := os.ReadFile(LOCAL_FILE_PATH)
    if err != nil {
        log.Printf("Failed to read DB file: %s\n", err)
        return
    }
 
    activeUserMap := localstorage.DeserializeFileData[bool](data) 
    delete(activeUserMap, userhash)

    localstorage.WriteMapToFile(activeUserMap, FILE_NAME)
}

