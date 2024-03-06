package localstorage 

import (
	"bytes"
	"encoding/gob"
	"os"
)

const LOCAL_DIR = "localDb"

func CleanDB() {
    os.RemoveAll(LOCAL_DIR)
}

func WriteMapToFile[V any](dataMap map[string]V, filename string, subdirectories ...string) error {
    os.Mkdir(LOCAL_DIR, os.ModePerm)
    localFilePath := LOCAL_DIR
    for _, subdirectory := range subdirectories {
        localFilePath = localFilePath + "/" + subdirectory
        os.Mkdir(localFilePath, os.ModePerm)
    }
    localFilePath = localFilePath + "/" + filename

    file, err := os.Create(localFilePath)
    if err != nil {
        return err
    }
    defer file.Close()
        
    b := new(bytes.Buffer)
    e := gob.NewEncoder(b)

    err = e.Encode(dataMap)
    if err != nil {
         panic(err)
    }

    file.Write(b.Bytes())
    return nil
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

