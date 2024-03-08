package activeuserdb

import (
	"os"

	"github.com/Ajahks/passkie/storage/localStorage"
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
        return false 
    }

    return activeUser
}

func RemoveActiveUser(userhash string) error {
    data, err := os.ReadFile(LOCAL_FILE_PATH)
    if err != nil {
        return err
    }
 
    activeUserMap := localstorage.DeserializeFileData[bool](data) 
    delete(activeUserMap, userhash)

    localstorage.WriteMapToFile(activeUserMap, FILE_NAME)
    return nil
}

