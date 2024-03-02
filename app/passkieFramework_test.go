package passkieApp

import (
	"reflect"
	"testing"

	localstorage "github.com/Ajahks/Passkie/storage/localStorage"
)

func TestCreateNewUserDoesntThrowErrForNewUser(t *testing.T) {
    defer localstorage.CleanDB()
    username := "testUser"
    password := "testPassword"

    err := CreateNewUser(username, password)

    if err != nil {
        t.Fatalf("Failed to create new user! %v", err) 
    }
}

func TestCreateNewUserForExistingUserThrowsException(t *testing.T) {
    defer localstorage.CleanDB()
    username := "testUser"
    password := "testPassword"

    CreateNewUser(username, password)
    err := CreateNewUser(username, password)

    if err == nil {
        t.Fatalf("Create new user should have failed! %v", err) 
    }
}

func TestStoreCredentialsForSiteForInvalidUserReturnsError(t *testing.T) {
    defer localstorage.CleanDB()
    username := "testUser"
    password := "testPassword"
    site := "testUrl.com"
    credentials := make(map[string]string)
    credentials["testCredential"] = "test"

    err := StoreCredentialsForSite(site, username, password, credentials)

    if err == nil {
        t.Fatalf("StoreCredentialsForSite did not fail for unknown user!")
    }
}

func TestStoreCredentialsForSiteForInvalidPasswordReturnsError(t *testing.T) {
    defer localstorage.CleanDB()
    username := "testUser"
    password := "testPassword"
    site := "testUrl.com"
    credentials := make(map[string]string)
    credentials["testCredential"] = "test"
    CreateNewUser(username, password)

    err := StoreCredentialsForSite(site, username, "wrongPassword", credentials)

    if err == nil {
        t.Fatalf("StoreCredentialsForSite did not fail for invalid password!")
    }
}

func TestRetrieveCredentialsForSiteForInvalidUserReturnsError(t *testing.T) {
    defer localstorage.CleanDB()
    username := "testUser"
    password := "testPassword"
    site := "testUrl.com"
    credentials := make(map[string]string)
    credentials["testCredential"] = "test"

    result, err := RetrieveCredentialsForSite(site, username, password)

    if err == nil {
        t.Errorf("RetrieveCredentialsForSite did not fail for unknown user!")
    }
    if result != nil {
        t.Errorf("RetrieveCredentialsForSite should not have returned credentials for unknown user!")
    }
}

func TestRetrieveCredentialsForSiteForInvalidPasswordReturnsError(t *testing.T) {
    defer localstorage.CleanDB()
    username := "testUser"
    password := "testPassword"
    site := "testUrl.com"
    credentials := make(map[string]string)
    credentials["testCredential"] = "test"
    CreateNewUser(username, password)

    result, err := RetrieveCredentialsForSite(site, username, "wrongPassword")

    if err == nil {
        t.Errorf("RetrieveCredentialsForSite did not fail for invalid password!")
    }
    if result != nil {
        t.Errorf("RetrieveCredentialsForSite should not have returned credentials for invalid password!")
    }
}

func TestRetrieveCrednetialsForSiteForUnknownSiteReturnsError(t *testing.T) {
    defer localstorage.CleanDB()
    username := "testUser"
    password := "testPassword"
    site := "testUrl.com"
    credentials := make(map[string]string)
    credentials["testCredential"] = "test"
    CreateNewUser(username, password)

    result, err := RetrieveCredentialsForSite(site, username, password)

    if err == nil {
        t.Errorf("RetrieveCredentialsForSite did not fail for unknown user!")
    }
    if result != nil {
        t.Errorf("RetrieveCredentialsForSite should not have returned credentials for unknown user!")
    }
}

func TestStoreCredentialsForSiteThenRetrieveCredentialsForSiteGetsOriginalCredentials(t *testing.T) {
    defer localstorage.CleanDB()
    username := "testUser"
    password := "testPassword"
    site := "testUrl.com"
    credentials := make(map[string]string)
    credentials["testField1"] = "test"
    credentials["testField2"] = "yeet"
    CreateNewUser(username, password)

    StoreCredentialsForSite(site, username, password, credentials)
    result, err := RetrieveCredentialsForSite(site, username, password)
    if err != nil {
        t.Fatalf("Failed to retrieve credentials for site: %v", err)
    }

    if !reflect.DeepEqual(credentials, result) {
        t.Fatalf("Retrieved credentials are not the same as original! Original: %v, Retrieved: %v", credentials, result) 
    }
}

func TestStoreCredentialsForSameSiteTwiceVerifyRetrieveCredentialsGetsLastStored(t *testing.T) {
    defer localstorage.CleanDB()
    username := "testUser"
    password := "testPassword"
    site := "testUrl.com"
    credentials1 := make(map[string]string)
    credentials1["testField1"] = "test"
    credentials1["testField2"] = "yeet"
    credentials2 := make(map[string]string)
    credentials2["testField1"] = "differentCred"
    credentials2["testField2"] = "yeet420"
    CreateNewUser(username, password)

    StoreCredentialsForSite(site, username, password, credentials1)
    StoreCredentialsForSite(site, username, password, credentials2)
    result, err := RetrieveCredentialsForSite(site, username, password)
    if err != nil {
        t.Fatalf("Failed to retrieve credentials for site: %v", err)
    }

    if !reflect.DeepEqual(credentials2, result) {
        t.Fatalf("Retrieved credentials are not the same as expected! Expected: %v, Retrieved: %v", credentials2, result) 
    }
}

