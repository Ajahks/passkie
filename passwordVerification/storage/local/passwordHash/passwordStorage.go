package passwordHash 

import (
	"bytes"
	"encoding/gob"
	"errors"
	"log"
	"os"
)

const LOCAL_DIR = "localDb" 
const LOCAL_FILE_PATH = LOCAL_DIR + "/passwordDB.txt"

func PutPasswordHash(userhash string, passwordHash []byte) {
    data, err := os.ReadFile(LOCAL_FILE_PATH)
    if err != nil {
        userPasswordMap := make(map[string][]byte)
        userPasswordMap[userhash] = passwordHash 

        writeMapToFile(userPasswordMap)

    } else {
        userPasswordMap := deserializeFileDataForUserPasswordMap(data)

        userPasswordMap[userhash] = passwordHash 

        writeMapToFile(userPasswordMap)
    }
}

func GetPasswordHash(userhash string) ([]byte, error) {
    data, err := os.ReadFile(LOCAL_FILE_PATH)
    if err != nil {
        return nil, err 
    }

    userPasswordMap := deserializeFileDataForUserPasswordMap(data)

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
    }
 
    userPasswordMap := deserializeFileDataForUserPasswordMap(data)

    delete(userPasswordMap, userhash)

    writeMapToFile(userPasswordMap)
}

func deserializeFileDataForUserPasswordMap(data []byte) map[string][]byte {
    b := bytes.NewBuffer(data)
    d := gob.NewDecoder(b)

    var decodedUserPasswordMap map[string][]byte
    err := d.Decode(&decodedUserPasswordMap)
    if err != nil {
        panic(err)
    }

    return decodedUserPasswordMap 
}

func writeMapToFile(userPasswordMap map[string][]byte) {
    os.Mkdir(LOCAL_DIR, os.ModePerm)
    file, err := os.Create(LOCAL_FILE_PATH)
    if err != nil {
        log.Fatalf("failed creating file: %s", err)
    }
    defer file.Close()
        
    b := new(bytes.Buffer)
    e := gob.NewEncoder(b)

    err = e.Encode(userPasswordMap)
    if err != nil {
         panic(err)
    }

    file.Write(b.Bytes())
}

