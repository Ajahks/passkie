package passwordHash 

import (
	"os"
	"reflect"
	"testing"
    localUtil "github.com/Ajahks/Passkie/passwordVerification/storage/local"
)

func TestPutPasswordHashCreatesANewFileIfNonExistent(t *testing.T) {
    username := "testUsername"
    passwordHash := []byte("testPasswordHash")

    PutPasswordHash(username, passwordHash)

    _, err := os.ReadFile(LOCAL_FILE_PATH)
    if err != nil {
        t.Fatalf("PutPasswordHash did not create a new file: %s", LOCAL_FILE_PATH)        
    }
    localUtil.CleanDB()
}

func TestPutPasswordHashAndGetPasswordHashRetrievesTheOriginalPasswordHash(t *testing.T) {
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
    localUtil.CleanDB()
}

func TestPutPasswordHashWithMultipleUsersMaintainsCorrectMapping(t *testing.T) {
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
    localUtil.CleanDB()
}

func TestPutUserWithADifferentPasswordHashOverridesTheOldPasswordHash(t *testing.T) {
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
    localUtil.CleanDB()
}

func TestRemovePasswordHashRemovesTheUserPasswordHashFromTheMap(t *testing.T) {
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
    localUtil.CleanDB()
}

