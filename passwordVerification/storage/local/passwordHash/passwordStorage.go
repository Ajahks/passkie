package passwordHash 

import (
	"errors"
	"log"
	"os"
    storageUtil "github.com/Ajahks/Passkie/passwordVerification/storage/local"
)

const FILE_NAME = "passwordDB.txt"
const LOCAL_FILE_PATH = storageUtil.LOCAL_DIR + "/" + FILE_NAME

func PutPasswordHash(userhash string, passwordHash []byte) {
    data, err := os.ReadFile(LOCAL_FILE_PATH)
    if err != nil {
        userPasswordMap := make(map[string][]byte)
        userPasswordMap[userhash] = passwordHash 

        storageUtil.WriteMapToFile(userPasswordMap, FILE_NAME)

    } else {
        userPasswordMap := storageUtil.DeserializeFileData[[]byte](data) 
        userPasswordMap[userhash] = passwordHash 

        storageUtil.WriteMapToFile(userPasswordMap, FILE_NAME)
    }
}

func GetPasswordHash(userhash string) ([]byte, error) {
    data, err := os.ReadFile(LOCAL_FILE_PATH)
    if err != nil {
        return nil, err 
    }

    userPasswordMap := storageUtil.DeserializeFileData[[]byte](data) 
    passwordHash, ok := userPasswordMap[userhash]
    if !ok {
        log.Printf("User %s does not exist in the DB!\n", userhash)
        return nil, errors.New("User does not exist in the DB!")
    }

    return passwordHash, nil
}

func RemovePasswordHash(userhash string) {
    data, err := os.ReadFile(LOCAL_FILE_PATH)
    if err != nil {
        log.Printf("Failed to read DB file: %s\n", err)
        return
    }
 
    userPasswordMap := storageUtil.DeserializeFileData[[]byte](data) 
    delete(userPasswordMap, userhash)

    storageUtil.WriteMapToFile(userPasswordMap, FILE_NAME)
}

