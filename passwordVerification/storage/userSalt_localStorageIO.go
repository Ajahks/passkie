package storage

import (
	"bytes"
	"encoding/gob"
	"log"
	"os"
)

// Stores user salts on a local file
func PutUserSalt(username string, salt []byte) {
    data, err := os.ReadFile("userSaltDB.txt")
    if err != nil {
        // file doesn't exist create a new one

        userSaltMap := make(map[string][]byte)
        userSaltMap[username] = salt

        writeMapToFile(userSaltMap)

    } else {
        // Deserialize data
        b := bytes.NewBuffer(data)
        d := gob.NewDecoder(b)
        var decodedUserSaltMap map[string][]byte
        err = d.Decode(&decodedUserSaltMap)
        if err != nil {
            panic(err)
        }

        decodedUserSaltMap[username] = salt

        writeMapToFile(decodedUserSaltMap)
    }
}

func GetUserSalt(username string) []byte {
    data, err := os.ReadFile("userSaltDB.txt")
    if err != nil {
        panic(err)
    }


    b := bytes.NewBuffer(data)
    d := gob.NewDecoder(b)
    var decodedUserSaltMap map[string][]byte
    err = d.Decode(&decodedUserSaltMap)
    if err != nil {
        panic(err)
    }

    salt, ok := decodedUserSaltMap[username]
    if !ok {
        log.Panicf("User %s does not exist in the DB!", username)
    }

    return salt
}

func writeMapToFile(userSaltMap map[string][]byte) {
    file, err := os.Create("userSaltDB.txt")
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

