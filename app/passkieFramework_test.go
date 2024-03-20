package passkieApp

import (
	"reflect"
	"testing"

	localstorage "github.com/Ajahks/passkie/storage/localStorage"
)

func TestCreateNewUserDoesntThrowErrForNewUser(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()
    username := "testUser"
    password := "testPassword"

    err := CreateNewUser(username, password)

    if err != nil {
        t.Fatalf("Failed to create new user! %v", err) 
    }
}

func TestCreateNewUserForExistingUserThrowsException(t *testing.T) {
    localstorage.SetTestDb()
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
    localstorage.SetTestDb()
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
    localstorage.SetTestDb()
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
    localstorage.SetTestDb()
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
    localstorage.SetTestDb()
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
    localstorage.SetTestDb()
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
    localstorage.SetTestDb()
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
    localstorage.SetTestDb()
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

func TestRemoveCredentialsForSiteWithInvalidMasterPasswordDoesNotRemoveSite(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()

    username := "testUser"
    password := "testPassword"
    site := "testUrl.com"
    credentials := make(map[string]string)
    credentials["testField1"] = "test"
    credentials["testField2"] = "yeet"
    CreateNewUser(username, password)
    StoreCredentialsForSite(site, username, password, credentials)

    err := RemoveCredentialsForSite(site, username, "wrongPassword") 

    if err == nil {
        t.Fatalf("RemoveCredentialsForSite should have returned exception with wrong password!")
    }
    retrievedCredentials, err := RetrieveCredentialsForSite(site, username, password)
    if err != nil {
        t.Errorf("Error returned for retrieve credentials: %v, Expected credentials should not be touched after failed removal", err)
    }
    if retrievedCredentials == nil {
        t.Error("Credentials missing after supposedly failed RemoveCredentialsForSite!")
    }
}

func TestRemoveCredentialsFromSiteRemovesCredentialsProperly(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()

    username := "testUser"
    password := "testPassword"
    site := "testUrl.com"
    credentials := make(map[string]string)
    credentials["testField1"] = "test"
    credentials["testField2"] = "yeet"
    CreateNewUser(username, password)
    StoreCredentialsForSite(site, username, password, credentials)

    err := RemoveCredentialsForSite(site, username, password)
    if err != nil {
        t.Fatalf("Failed to remove credentials for site: %v", err)
    }

    retrievedCredentials, err := RetrieveCredentialsForSite(site, username, password)
    if err == nil {
        t.Error("Retrieved credential after valid remove should have returned an error!")
    }
    if retrievedCredentials != nil {
        t.Errorf("Retrieved credentials were still returned after valid removal call: %v", retrievedCredentials) 
    }
}

func TestRemoveCredentialsFromSiteOnNonExistentSiteThrowsError(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()

    username := "testUser"
    password := "testPassword"
    site := "testUrl.com"
    CreateNewUser(username, password)

    err := RemoveCredentialsForSite(site, username, password)

    if err == nil {
        t.Fatal("RemoveCredentialsForSite should have returned an error for a non existent site!")
    }
}

func TestRemoveUserOnNonExistentUserThrowsError(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()

    err := RemoveUser("fakeUsername", "fakePassword")
    
    if err == nil {
        t.Fatal("RemoveUser should have returned an error for a non existent user!")
    }
}

func TestRemoveUserWithWrongPasswordThrowsErrorAndKeepsUser(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()

    username := "testUser"
    password := "testPassword"
    CreateNewUser(username, password)

    err := RemoveUser(username, "wrongPassword") 

    if err == nil {
        t.Error("RemoveUser with the wrong password should have thrown an error") 
    }
    err = CreateNewUser(username, "wrongPassword")
    if err == nil {
        t.Error("CreateNewUser should not work for a user that was not deleted")
    }
}

func TestRemoveUserOnExistingUserAllowsNewUserCreation(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()

    username := "testUser"
    password := "testPassword"
    CreateNewUser(username, password)

    err := RemoveUser(username, password) 

    if err != nil {
        t.Errorf("Failed to RemoveUser: %v", err)
    }
    err = CreateNewUser(username, password)
    if err != nil {
        t.Errorf("Failed to create new user after user was deleted: %v", err)
    }
}

func TestRemoveUserOnExistingUserDoesNotSaveOldMasterPassword(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()

    username := "testUser"
    originalPassword := "testPassword"
    newPassword := "testPassword2"
    CreateNewUser(username, originalPassword)
    site := "testUrl.com"
    credentials := make(map[string]string)
    credentials["testField1"] = "test"
    credentials["testField2"] = "yeet"

    RemoveUser(username, originalPassword)
    CreateNewUser(username, newPassword)
    err := StoreCredentialsForSite(site, username, originalPassword, credentials)

    if err == nil {
        t.Fatal("StoreCredentials with the old credentials of a previously deleted user should not work!")
    }
}

func TestRemoveUserOnExistingUserDoesNotSaveOldCredentials(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()

    username := "testUser"
    password := "testPassword"
    CreateNewUser(username, password)
    site := "testUrl.com"
    credentials := make(map[string]string)
    credentials["testField1"] = "test"
    credentials["testField2"] = "yeet"
    StoreCredentialsForSite(site, username, password, credentials)


    RemoveUser(username, password)
    CreateNewUser(username, password)
    res, err := RetrieveCredentialsForSite(site, username, password)
    if err == nil {
        t.Error("Retrieve credentials for previously removed user should not exist")
    }
    if res != nil {
        t.Errorf("Retrieve credentials for previously removed user should not have returned a result. Result: %v", res) 
    }
}

func TestRemoveUserOnExistingUserDoesNotAffectOtherUsers(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()

    username1 := "testUser1"
    password1 := "testPassword1"
    username2 := "testUser2"
    password2 := "testPassword2"
    CreateNewUser(username1, password1)
    CreateNewUser(username2, password2)

    RemoveUser(username2, password2) 

    err := CreateNewUser(username1, password1)
    if err == nil {
        t.Errorf("Should not be able to recreate username1 when only username2 was deleted")
    }
    err = CreateNewUser(username2, password2)
    if err != nil {
        t.Errorf("Should have been able to recreate username 2 after it was deleted. Error: %v", err)
    }
}

