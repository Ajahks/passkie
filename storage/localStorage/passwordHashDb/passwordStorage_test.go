package passwordHashDb 

import (
	"os"
	"reflect"
	"testing"
    "github.com/Ajahks/passkie/storage/localStorage"
)

func TestPutPasswordHashCreatesANewFileIfNonExistent(t *testing.T) {
    localstorage.SetTestDb()
    username := "testUsername"
    passwordHash := []byte("testPasswordHash")

    PutPasswordHash(username, passwordHash)

    _, err := os.ReadFile(getFilePath())
    if err != nil {
        t.Fatalf("PutPasswordHash did not create a new file: %s", getFilePath())        
    }
    localstorage.CleanDB()
}

func TestPutPasswordHashAndGetPasswordHashRetrievesTheOriginalPasswordHash(t *testing.T) {
    localstorage.SetTestDb()
    username := "testUsername"
    passwordHash := []byte("testPasswordHash")

    PutPasswordHash(username, passwordHash)
    retrievedPasswordHash, err := GetPasswordHash(username)

    if err != nil {
        t.Fatalf("Failed to retrieve user passwordHash for user %s: %s", username, err)
    }
    if !reflect.DeepEqual(passwordHash, retrievedPasswordHash) {
        t.Fatalf("GetPasswordHash does not return original passwordHash!\nOriginal PasswordHash: %s\n Retrieved PasswordHash: %s", string(passwordHash), string(retrievedPasswordHash))
    }
    localstorage.CleanDB()
}

func TestPutPasswordHashWithMultipleUsersMaintainsCorrectMapping(t *testing.T) {
    localstorage.SetTestDb()
    username1 := "testUsername1"
    passwordHash1 := []byte("testPasswordHash1")
    username2 := "testUsername2"
    passwordHash2 := []byte("testPasswordHash2")

    PutPasswordHash(username1, passwordHash1)
    PutPasswordHash(username2, passwordHash2)
    retrievedPasswordHash1, err := GetPasswordHash(username1)
    if err != nil {
        t.Fatalf("Failed to retrieve user passwordHash for user %s: %s", username1, err)
    }
    retrievedPasswordHash2, err := GetPasswordHash(username2)
    if err != nil {
        t.Fatalf("Failed to retrieve user passwordHash for user %s: %s", username2, err)
    }

    if !reflect.DeepEqual(passwordHash1, retrievedPasswordHash1) {
        t.Fatalf("GetPasswordHash does not return original passwordHash1!\nOriginal PasswordHash1: %s\n Retrieved PasswordHash: %s", string(passwordHash1), string(retrievedPasswordHash1))
    }
    if !reflect.DeepEqual(passwordHash2, retrievedPasswordHash2) {
        t.Fatalf("GetPasswordHash does not return original passwordHash2!\nOriginal PasswordHash2: %s\n Retrieved PasswordHash: %s", string(passwordHash2), string(retrievedPasswordHash2))
    }
    localstorage.CleanDB()
}

func TestPutUserWithADifferentPasswordHashOverridesTheOldPasswordHash(t *testing.T) {
    localstorage.SetTestDb()
    username := "testUsername"
    passwordHash1 := []byte("testPasswordHash1")
    passwordHash2 := []byte("testPasswordHash2")

    PutPasswordHash(username, passwordHash1)
    PutPasswordHash(username, passwordHash2)
    retrievedPasswordHash, err := GetPasswordHash(username)
    if err != nil {
        t.Fatalf("Failed to retrieve user passwordHash for user %s: %s", username, err)
    }

    
    if !reflect.DeepEqual(passwordHash2, retrievedPasswordHash) {
        t.Fatalf("GetPasswordHash does not return the updated passwordHash!\nExpected passwordHash: %s\n Retrieved PasswordHash: %s", string(passwordHash2), string(retrievedPasswordHash))
    }
    localstorage.CleanDB()
}

func TestRemovePasswordHashRemovesTheUserPasswordHashFromTheMap(t *testing.T) {
    localstorage.SetTestDb()
    username := "testUsername"
    passwordHash := []byte("testPasswordHash")

    PutPasswordHash(username, passwordHash)
    RemovePasswordHash(username)

    retrievedPasswordHash, err := GetPasswordHash(username)

    if err == nil {
        t.Fatalf("PasswordHash should not have been returned after retrieval. Retrieved passwordHash: %s", string(retrievedPasswordHash))
    }

    if retrievedPasswordHash != nil {
        t.Fatalf("PasswordHash should not have been returned after retrieval. Retrieved passwordHash: %s", string(retrievedPasswordHash))
    }
    localstorage.CleanDB()
}

