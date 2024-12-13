package localstorage

import (
	"os"
	"testing"
)

func TestCreateTheFolderAndCleanDbRemovesTheFolder(t *testing.T) {
	SetTestDb()
	initDb()
	localFilePath := DB_PATH() + "/RandomFile"
	file, _ := os.Create(localFilePath)
	file.Close()
	_, err := os.ReadFile(localFilePath)
	if err != nil {
		t.Errorf("Failed to create a file, this test is not going to work properly! %v", err)
	}

	err = CleanDB()
	if err != nil {
		t.Errorf("CleanDB failed to delete the db file, error; %v.", err)
	}

	_, err = os.ReadFile(localFilePath)
	if err == nil {
		t.Errorf("CleanDB failed to delete the db file created at %v", localFilePath)
	}
}

func TestWriteMapToFileCreatesAFile(t *testing.T) {
	SetTestDb()
	defer CleanDB()
	testMap := map[string]string{"test": "value"}
	filename := "testFile.txt"
	localFilePath := DB_PATH() + "/" + filename

	WriteMapToFile[string](testMap, filename)

	_, err := os.ReadFile(localFilePath)
	if err != nil {
		t.Errorf("WriteMapToFile failed to generate file: %s, error: %v", localFilePath, err)

	}
}

func TestWriteToFileWithOneSubdirectoryCreatesNestedFile(t *testing.T) {
	SetTestDb()
	defer CleanDB()
	testMap := map[string]string{"test": "value"}
	filename := "testFile.txt"
	subdirectory := "test"
	localFilePath := DB_PATH() + "/" + subdirectory + "/" + filename

	WriteMapToFile[string](testMap, filename, subdirectory)

	_, err := os.ReadFile(localFilePath)
	if err != nil {
		t.Errorf("WriteMapToFile failed to generate file in correct path: %s, error: %v", localFilePath, err)

	}
}

func TestWriteToFileWithMultipleSubdirectoriesCreatesNestedFile(t *testing.T) {
	SetTestDb()
	defer CleanDB()
	testMap := map[string]string{"test": "value"}
	filename := "testFile.txt"
	subdirectory1 := "test"
	subdirectory2 := "test2"
	localFilePath := DB_PATH() + "/" + subdirectory1 + "/" + subdirectory2 + "/" + filename

	WriteMapToFile[string](testMap, filename, subdirectory1, subdirectory2)

	_, err := os.ReadFile(localFilePath)
	if err != nil {
		t.Errorf("WriteMapToFile failed to generate file in correct path: %s, error %v", localFilePath, err)

	}
}

func TestWriteMapToFileAndDeserializeFileDataOfFileGetsOriginalData(t *testing.T) {
	SetTestDb()
	defer CleanDB()
	testMap := map[string]string{"test": "value"}
	filename := "testFile.txt"
	localFilePath := DB_PATH() + "/" + filename

	WriteMapToFile[string](testMap, filename)
	data, err := os.ReadFile(localFilePath)
	if err != nil {
		t.Errorf("WriteMapToFile failed to generate file: %s: %v", localFilePath, err)
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
}
