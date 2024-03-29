package userSaltDb 

import (
	"os"
	"reflect"
	"testing"
    "github.com/Ajahks/passkie/storage/localStorage"
)

func TestPutUserSaltCreatesANewFileIfNonExistent(t *testing.T) {
    localstorage.SetTestDb()
    username := "testUsername"
    salt := []byte("testUserSalt")

    PutUserSalt(username, salt)

    _, err := os.ReadFile(getFilePath())
    if err != nil {
        t.Fatalf("PutUserSalt did not create a new file: %s", getFilePath())        
    }
    localstorage.CleanDB()
}

func TestPutUserSaltAndGetUserSaltRetrievesTheOriginalSalt(t *testing.T) {
    localstorage.SetTestDb()
    username := "testUsername"
    salt := []byte("testUserSalt")

    PutUserSalt(username, salt)
    retrievedSalt, err := GetUserSalt(username)

    if err != nil {
        t.Fatalf("Failed to retrieve user salt for user %s: %s", username, err)
    }
    if !reflect.DeepEqual(salt, retrievedSalt) {
        t.Fatalf("GetUserSalt does not return original salt!\nOriginal Salt: %s\n Retrieved Salt: %s", string(salt), string(retrievedSalt))
    }
    localstorage.CleanDB()
}

func TestPutUserSaltWithMultipleUsersMaintainsCorrectMapping(t *testing.T) {
    localstorage.SetTestDb()
    username1 := "testUsername1"
    salt1 := []byte("testUserSalt1")
    username2 := "testUsername2"
    salt2 := []byte("testUserSalt2")

    PutUserSalt(username1, salt1)
    PutUserSalt(username2, salt2)
    retrievedSalt1, err := GetUserSalt(username1)
    if err != nil {
        t.Fatalf("Failed to retrieve user salt for user %s: %s", username1, err)
    }
    retrievedSalt2, err := GetUserSalt(username2)
    if err != nil {
        t.Fatalf("Failed to retrieve user salt for user %s: %s", username2, err)
    }

    if !reflect.DeepEqual(salt1, retrievedSalt1) {
        t.Fatalf("GetUserSalt does not return original salt1!\nOriginal Salt1: %s\n Retrieved Salt: %s", string(salt1), string(retrievedSalt1))
    }
    if !reflect.DeepEqual(salt2, retrievedSalt2) {
        t.Fatalf("GetUserSalt does not return original salt2!\nOriginal Salt2: %s\n Retrieved Salt: %s", string(salt2), string(retrievedSalt2))
    }
    localstorage.CleanDB()
}

func TestPutUserWithADifferentSaltOverridesTheOldSalt(t *testing.T) {
    localstorage.SetTestDb()
    username := "testUsername"
    salt1 := []byte("testUserSalt1")
    salt2 := []byte("testUserSalt2")

    PutUserSalt(username, salt1)
    PutUserSalt(username, salt2)
    retrievedSalt, err := GetUserSalt(username)
    if err != nil {
        t.Fatalf("Failed to retrieve user salt for user %s: %s", username, err)
    }

    
    if !reflect.DeepEqual(salt2, retrievedSalt) {
        t.Fatalf("GetUserSalt does not return the updated salt!\nExpected salt: %s\n Retrieved Salt: %s", string(salt2), string(retrievedSalt))
    }
    localstorage.CleanDB()
}

func TestRemoveUserSaltRemovesTheUserSaltFromTheMap(t *testing.T) {
    localstorage.SetTestDb()
    username := "testUsername"
    salt := []byte("testUserSalt")

    PutUserSalt(username, salt)
    RemoveUserSalt(username)

    retrievedSalt, err := GetUserSalt(username)

    if err == nil {
        t.Fatalf("Salt should not have been returned after retrieval. Retrieved salt: %s", string(retrievedSalt))
    }

    if retrievedSalt != nil {
        t.Fatalf("Salt should not have been returned after retrieval. Retrieved salt: %s", string(retrievedSalt))
    }
    localstorage.CleanDB()
}

