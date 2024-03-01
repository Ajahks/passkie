package localstorage 

import (
	"os"
	"testing"
)

func TestCreateTheFolderAndCleanDbRemovesTheFolder(t *testing.T) {
    os.Mkdir(LOCAL_DIR, os.ModePerm)
    localFilePath := LOCAL_DIR + "/RandomFile" 
    os.Create(localFilePath)
    _, err := os.ReadFile(localFilePath)
    if err != nil {
        t.Errorf("Failed to create a file, this test is not going to work properly!")    
    }

    CleanDB()

    _, err = os.ReadFile(localFilePath)
    if err == nil {
        t.Errorf("CleanDB failed to delete the db file created")
    }
}

func TestWriteMapToFileCreatesAFile(t *testing.T) {
    testMap := map[string]string{"test":"value"} 
    filename := "testFile.txt"
    localFilePath := LOCAL_DIR + "/" + filename 

    WriteMapToFile[string](testMap, filename)

    _, err := os.ReadFile(localFilePath)
    if err != nil {
        t.Errorf("WriteMapToFile failed to generate file: %s", localFilePath)
 
    }

    CleanDB()
}

func TestWriteToFileWithOneSubdirectoryCreatesNestedFile(t *testing.T) {
    testMap := map[string]string{"test":"value"} 
    filename := "testFile.txt"
    subdirectory := "test"
    localFilePath := LOCAL_DIR + "/" + subdirectory + "/" + filename 

    WriteMapToFile[string](testMap, filename, subdirectory)

    _, err := os.ReadFile(localFilePath)
    if err != nil {
        t.Errorf("WriteMapToFile failed to generate file in correct path: %s", localFilePath)
 
    }

    CleanDB()
}

func TestWriteToFileWithMultipleSubdirectoriesCreatesNestedFile(t *testing.T) {
    testMap := map[string]string{"test":"value"} 
    filename := "testFile.txt"
    subdirectory1 := "test"
    subdirectory2 := "test2"
    localFilePath := LOCAL_DIR + "/" + subdirectory1 + "/" + subdirectory2 + "/" + filename 

    WriteMapToFile[string](testMap, filename, subdirectory1, subdirectory2)

    _, err := os.ReadFile(localFilePath)
    if err != nil {
        t.Errorf("WriteMapToFile failed to generate file in correct path: %s", localFilePath)
 
    }

    CleanDB()
}

func TestWriteMapToFileAndDeserializeFileDataOfFileGetsOriginalData(t *testing.T) {
    testMap := map[string]string{"test":"value"} 
    filename := "testFile.txt"
    localFilePath := LOCAL_DIR + "/" + filename 

    WriteMapToFile[string](testMap, filename)
    data, err := os.ReadFile(localFilePath)
    if err != nil {
        t.Errorf("WriteMapToFile failed to generate file: %s", localFilePath)
    }
    deserializedMap := DeserializeFileData[string](data)

    if len(deserializedMap) != 1 {
        t.Errorf("deserialized map is not the same len as the original! original: %v, got: %v", testMap, deserializedMap)
    }
    resultValue, ok := deserializedMap["test"]
    if !ok {
        t.Errorf("Failed to find key 'test' in deserialized map! Deserialized map: %v", deserializedMap)
    }
    if resultValue != "value" {
        t.Errorf("Deserialized map value is not expected: Exptected: 'value', Got: '%s'", resultValue)
    }

    CleanDB()
}

