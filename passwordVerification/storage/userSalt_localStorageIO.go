package storage

import (
	"bytes"
	"encoding/gob"
	"errors"
	"log"
	"os"
)

const LOCAL_DIR = "localDb"
const LOCAL_FILE_PATH = LOCAL_DIR + "/userSaltDB.txt"

// Stores user salts on a local file
func PutUserSalt(username string, salt []byte) {
    data, err := os.ReadFile(LOCAL_FILE_PATH)
    if err != nil {
        userSaltMap := make(map[string][]byte)
        userSaltMap[username] = salt

        writeMapToFile(userSaltMap)

    } else {
        userSaltMap := deserializeFileData(data)

        userSaltMap[username] = salt

        writeMapToFile(userSaltMap)
    }
}

// Reads salts on a local file 
func GetUserSalt(username string) ([]byte, error) {
    data, err := os.ReadFile(LOCAL_FILE_PATH)
    if err != nil {
        return nil, err 
    }

    userSaltMap := deserializeFileData(data)

    salt, ok := userSaltMap[username]
    if !ok {
        log.Printf("User %s does not exist in the DB!\n", username)
        return nil, errors.New("User does not exist in the DB!")
    }

    return salt, nil
}

func writeMapToFile(userSaltMap map[string][]byte) {
    os.Mkdir(LOCAL_DIR, os.ModePerm)
    file, err := os.Create(LOCAL_FILE_PATH)
    if err != nil {
        log.Fatalf("failed creating file: %s", err)
    }
    defer file.Close()
        
    b := new(bytes.Buffer)
    e := gob.NewEncoder(b)

    err = e.Encode(userSaltMap)
    if err != nil {
         panic(err)
    }

    file.Write(b.Bytes())
}

func deserializeFileData(data []byte) map[string][]byte {
    b := bytes.NewBuffer(data)
    d := gob.NewDecoder(b)

    var decodedUserSaltMap map[string][]byte
    err := d.Decode(&decodedUserSaltMap)
    if err != nil {
        panic(err)
    }

    return decodedUserSaltMap
}
