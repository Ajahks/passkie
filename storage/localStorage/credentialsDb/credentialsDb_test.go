package credentialsDb

import (
	"bytes"
	"os"
	"testing"

	localstorage "github.com/Ajahks/passkie/storage/localStorage"
)

func TestPutCredentialsForSiteHashCreatesCorrectFileInCorrectPath(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()
    sitehash := "testSiteHash"
    username := "testUsername"
    credentials := []byte("test")
    
    PutCredentialsForSiteHash(sitehash, username, credentials)


    _, err := os.ReadFile(getFilePath(username))
    if err != nil {
        t.Errorf("Failed to find the credential DB file in localPath: %s", getFilePath(username))
    }
}

func TestPutCredentialsForSiteHashAndGetCredentialsForSiteHashGetTheCorrectCredentials(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()
    sitehash := "testSiteHash"
    username := "testUsername"
    credentials := []byte("test")

    PutCredentialsForSiteHash(sitehash, username, credentials)
    result, err := GetCredentialsForSiteHash(sitehash, username)
    if err != nil {
        t.Fatalf("Error while reading credentials for site: %v", err)
    }

    if !bytes.Equal(credentials, result) {
        t.Fatalf("Returned credentials do not match original! Original: %s, Received: %s", string(credentials), string(result)) 
    }
}

func TestGetCredentialsForNonExistentUserReturnsError(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()
    sitehash := "testSiteHash"
    username := "testUsername"

    result, err := GetCredentialsForSiteHash(sitehash, username)
    if err == nil {
        t.Fatalf("GetCredentialsForSiteHash for missing sitehash did not return an error!")
    }

    if result != nil {
        t.Fatalf("GetCredentialsForSiteHash for missing sitehash returned an unexpected result! %s", string(result))
    }
}

func TestPutCredentialsForSiteHashTwiceReturnsMostRecentUpdateWhenGet(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()
    sitehash := "testSiteHash"
    username := "testUsername"
    credentials1 := []byte("test1")
    credentials2 := []byte("test2")

    PutCredentialsForSiteHash(sitehash, username, credentials1)
    PutCredentialsForSiteHash(sitehash, username, credentials2)
    result, err := GetCredentialsForSiteHash(sitehash, username)
    if err != nil {
        t.Fatalf("Error while reading credentials for site: %v", err)
    }

    if !bytes.Equal(credentials2, result) {
        t.Fatalf("Returned credentials do not match updated credentials! Expected: %s, Received: %s", string(credentials2), string(result)) 
    }
}

func TestPutCredentialsForSiteHashWithDifferentUsersMaintainTheMappings(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()
    sitehash := "testSiteHash"
    username1 := "testUsername1"
    username2 := "testUsername2"
    credentials1 := []byte("test1")
    credentials2 := []byte("test2")

    PutCredentialsForSiteHash(sitehash, username1, credentials1)
    PutCredentialsForSiteHash(sitehash, username2, credentials2)
    result1, err := GetCredentialsForSiteHash(sitehash, username1)
    if err != nil {
        t.Fatalf("Error while reading credentials for user %s for site.  Error: %v", username1, err)
    }
    result2, err := GetCredentialsForSiteHash(sitehash, username2)
    if err != nil {
        t.Fatalf("Error while reading credentials for user %s for site.  Error: %v", username2, err)
    }

    if !bytes.Equal(credentials1, result1) {
        t.Fatalf("Returned credentials for user %s do not match credentials! Expected: %s, Received: %s", username1, string(credentials1), string(result1)) 
    }
    if !bytes.Equal(credentials2, result2) {
        t.Fatalf("Returned credentials for user %s do not match credentials! Expected: %s, Received: %s", username2, string(credentials2), string(result2)) 
    }
}

func TestRemoveCredentialsForSiteHash(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()
    sitehash := "testSiteHash"
    username := "testUsername"
    credentials := []byte("test")

    PutCredentialsForSiteHash(sitehash, username, credentials)
    RemoveCredentialsForSiteHash(sitehash, username)
    result, err := GetCredentialsForSiteHash(sitehash, username)
    if err == nil {
        t.Fatalf("GetCredentialsForSiteHash for deleted sitehash did not return an error!")
    }

    if result != nil {
        t.Fatalf("GetCredentialsForSiteHash for missing user returned an unexpected result! %s", string(result))
    }
}

func TestRemoveCredentialsForSiteHashForNonExistentSiteHashDoesNotModifyResult(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()
    sitehash := "testSiteHash"
    username := "testUsername"
    credentials := []byte("test")

    PutCredentialsForSiteHash(sitehash, username, credentials)
    data1, err := os.ReadFile(getFilePath(username))
    if err != nil {
        t.Fatalf("Failed to read DB file: %s\n", err)
    }
    RemoveCredentialsForSiteHash("otherSite", username)
    data2, err := os.ReadFile(getFilePath(username))
    if err != nil {
        t.Fatalf("Failed to read DB file: %s\n", err)
    }

    if !bytes.Equal(data1, data2) {
        t.Fatalf("RemoveCredentialsForSiteHash modified the file! Original: %s, New: %s", string(data1), string(data2))
    }
}

func TestRemoveUserCredentialsForNonExistentUserReturnsError(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()

    username := "testUsername" 

    err := RemoveUserCredentials(username)
    
    if err == nil {
        t.Fatalf("RemoveUserCredentials for non existent user should be returning an error")
    }
}

func TestRemoveUserCredentialsForExistingUserRemovesTheFolder(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()

    sitehash := "testSiteHash"
    username := "testUsername"
    credentials := []byte("test")
    PutCredentialsForSiteHash(sitehash, username, credentials)

    err := RemoveUserCredentials(username)

    if err != nil {
        t.Errorf("RemoveUserCredentials failed with an error: %v", err)
    }
    res, err := os.ReadDir(localstorage.DB_PATH() + "/" + getEncodedUsername(username))
    if err == nil {
        t.Errorf("After remove credentials, the user directory path still exsits")
    }
    if res != nil {
        t.Errorf("After remove credentials, the user directory still returns some content: %v", res)
    }
}

func TestRemoveUserCredentialsForExistingUserActuallyRemovesCredentials(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()

    sitehash := "testSiteHash"
    username := "testUsername"
    credentials := []byte("test")
    PutCredentialsForSiteHash(sitehash, username, credentials)

    RemoveUserCredentials(username)

    res, err := GetCredentialsForSiteHash(sitehash, username)
    if err == nil {
        t.Errorf("After RemoveUserCredentials, GetCredentialsForSiteHash does not return an error")
    }
    if res != nil {
        t.Errorf("After RemoveUserCredentials, GetCredentialsForSiteHash should not have returned a result. Returned result: %v", res)
    }
}

func TestRemoveUserCredentialsForExistingUserDoesNotRemoveOtherUser(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()

    sitehash := "testSiteHash"
    username1 := "testUsername1"
    credentials := []byte("test")
    PutCredentialsForSiteHash(sitehash, username1, credentials)
    username2 := "testUsername2"
    PutCredentialsForSiteHash(sitehash, username2, credentials)
    
    RemoveUserCredentials(username1)

    res, err := GetCredentialsForSiteHash(sitehash, username2)
    if err != nil {
        t.Errorf("Non removed user should not have its credentials removed! Error thrown: %v", err)
    }
    if res == nil {
        t.Errorf("Non removed user should still return back credentials!")
    }
}

