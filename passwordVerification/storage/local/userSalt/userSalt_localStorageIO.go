package userSalt 

import (
	"errors"
	"log"
	"os"
    storageUtil "github.com/Ajahks/Passkie/passwordVerification/storage/local"
)

const FILE_PATH = "userSaltDB.txt"
const LOCAL_FILE_PATH = storageUtil.LOCAL_DIR + "/" + FILE_PATH

// Stores user salts on a local file
func PutUserSalt(userhash string, salt []byte) {
    data, err := os.ReadFile(LOCAL_FILE_PATH)
    if err != nil {
        userSaltMap := make(map[string][]byte)
        userSaltMap[userhash] = salt

        storageUtil.WriteMapToFile(userSaltMap, FILE_PATH)

    } else {
        userSaltMap := storageUtil.DeserializeFileData[[]byte](data)

        userSaltMap[userhash] = salt

        storageUtil.WriteMapToFile(userSaltMap, FILE_PATH)
    }
}

// Reads salts on a local file 
func GetUserSalt(userhash string) ([]byte, error) {
    data, err := os.ReadFile(LOCAL_FILE_PATH)
    if err != nil {
        return nil, err 
    }

    userSaltMap := storageUtil.DeserializeFileData[[]byte](data)

    salt, ok := userSaltMap[userhash]
    if !ok {
        log.Printf("User %s does not exist in the DB!\n", userhash)
        return nil, errors.New("User does not exist in the DB!")
    }

    return salt, nil
}

// Removes a user salt from the storage
func RemoveUserSalt(userhash string) {
    data, err := os.ReadFile(LOCAL_FILE_PATH)
    if err != nil {
        log.Printf("Failed to read DB file: %s\n", err)
        return
    }
 
    userSaltMap := storageUtil.DeserializeFileData[[]byte](data)

    delete(userSaltMap, userhash)

    storageUtil.WriteMapToFile(userSaltMap, FILE_PATH)
}

