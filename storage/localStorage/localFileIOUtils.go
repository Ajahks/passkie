package localstorage 

import (
	"bytes"
	"encoding/gob"
	"os"
)

var APPCONFIG_DIR = "passkie"
var DATABASE_DIR = "localDb"

func initDb() {
    homeConfigDir, err := os.UserConfigDir()
    if err != nil { homeConfigDir = "." }
    
    os.Mkdir(homeConfigDir + "/" + APPCONFIG_DIR, os.ModePerm)
    os.Mkdir(homeConfigDir + "/" + APPCONFIG_DIR + "/" + DATABASE_DIR, os.ModePerm) 
}

func SetTestDb() {
    DATABASE_DIR = "localDb-test"
}

func DB_PATH() string {
    homeConfigDir, err := os.UserConfigDir()
    // if the home config directory cannot be found, just return local path
    if err != nil { return APPCONFIG_DIR + "/" + DATABASE_DIR }

    return homeConfigDir + "/" + APPCONFIG_DIR + "/" + DATABASE_DIR
}

func CleanDB() error {
    err := os.RemoveAll(DB_PATH())
    if err != nil { return err }
    return nil
}

func WriteMapToFile[V any](dataMap map[string]V, filename string, subdirectories ...string) error {
    initDb()

    localFilePath := DB_PATH() 
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

