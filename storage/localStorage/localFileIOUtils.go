package localstorage 

import (
	"bytes"
	"encoding/gob"
	"log"
	"os"
)

const LOCAL_DIR = "localDb"

func CleanDB() {
    os.RemoveAll(LOCAL_DIR)
}

func WriteMapToFile[V any](dataMap map[string]V, filename string) {
    os.Mkdir(LOCAL_DIR, os.ModePerm)
    localFilePath := LOCAL_DIR + "/" + filename
    file, err := os.Create(localFilePath)
    if err != nil {
        log.Fatalf("failed creating file: %s", err)
    }
    defer file.Close()
        
    b := new(bytes.Buffer)
    e := gob.NewEncoder(b)

    err = e.Encode(dataMap)
    if err != nil {
         panic(err)
    }

    file.Write(b.Bytes())
}

func DeserializeFileData[V any](data []byte) map[string]V {
    b := bytes.NewBuffer(data)
    d := gob.NewDecoder(b)

    var decodedMap map[string]V
    err := d.Decode(&decodedMap)
    if err != nil {
        panic(err)
    }

    return decodedMap 
}

