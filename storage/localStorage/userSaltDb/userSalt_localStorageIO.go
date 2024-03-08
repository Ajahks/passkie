package userSaltDb 

import (
	"errors"
	"os"
    "github.com/Ajahks/passkie/storage/localStorage"
)

const FILE_PATH = "userSaltDB.txt"
const LOCAL_FILE_PATH = localstorage.LOCAL_DIR + "/" + FILE_PATH

// Stores user salts on a local file
func PutUserSalt(userhash string, salt []byte) {
    data, err := os.ReadFile(LOCAL_FILE_PATH)
    if err != nil {
        userSaltMap := make(map[string][]byte)
        userSaltMap[userhash] = salt

        localstorage.WriteMapToFile(userSaltMap, FILE_PATH)

    } else {
        userSaltMap := localstorage.DeserializeFileData[[]byte](data)

        userSaltMap[userhash] = salt

        localstorage.WriteMapToFile(userSaltMap, FILE_PATH)
    }
}

// Reads salts on a local file 
func GetUserSalt(userhash string) ([]byte, error) {
    data, err := os.ReadFile(LOCAL_FILE_PATH)
    if err != nil {
        return nil, err 
    }

    userSaltMap := localstorage.DeserializeFileData[[]byte](data)

    salt, ok := userSaltMap[userhash]
    if !ok {
        return nil, errors.New("User does not exist in the DB!")
    }

    return salt, nil
}

// Removes a user salt from the storage
func RemoveUserSalt(userhash string) error {
    data, err := os.ReadFile(LOCAL_FILE_PATH)
    if err != nil {
        return err
    }
 
    userSaltMap := localstorage.DeserializeFileData[[]byte](data)

    delete(userSaltMap, userhash)

    localstorage.WriteMapToFile(userSaltMap, FILE_PATH)
    return nil
}

