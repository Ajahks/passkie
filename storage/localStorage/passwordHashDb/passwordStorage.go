package passwordHashDb 

import (
	"errors"
	"os"
    "github.com/Ajahks/passkie/storage/localStorage"
)

const FILE_NAME = "passwordDB.txt"

func getFilePath() string {
    return localstorage.DB_PATH() + "/" + FILE_NAME
}

func PutPasswordHash(userhash string, passwordHash []byte) {
    data, err := os.ReadFile(getFilePath())
    if err != nil {
        userPasswordMap := make(map[string][]byte)
        userPasswordMap[userhash] = passwordHash 

        localstorage.WriteMapToFile(userPasswordMap, FILE_NAME)

    } else {
        userPasswordMap := localstorage.DeserializeFileData[[]byte](data) 
        userPasswordMap[userhash] = passwordHash 

        localstorage.WriteMapToFile(userPasswordMap, FILE_NAME)
    }
}

func GetPasswordHash(userhash string) ([]byte, error) {
    data, err := os.ReadFile(getFilePath())
    if err != nil {
        return nil, err 
    }

    userPasswordMap := localstorage.DeserializeFileData[[]byte](data) 
    passwordHash, ok := userPasswordMap[userhash]
    if !ok {
        return nil, errors.New("User does not exist in the DB!")
    }

    return passwordHash, nil
}

func RemovePasswordHash(userhash string) error {
    data, err := os.ReadFile(getFilePath())
    if err != nil {
        return err 
    }
 
    userPasswordMap := localstorage.DeserializeFileData[[]byte](data) 
    delete(userPasswordMap, userhash)

    localstorage.WriteMapToFile(userPasswordMap, FILE_NAME)
    return nil
}

