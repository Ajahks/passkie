package passkie_test

import (
    "reflect"
    "testing"

    "github.com/Ajahks/passkie"
    localstorage "github.com/Ajahks/passkie/storage/localStorage"
)

func TestCreateNewUserDoesntThrowErrForNewUser(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()
    username := "testUser"
    password := "testPassword"

    err := passkie.CreateNewUser(username, password)

    if err != nil {
        t.Fatalf("Failed to create new user! %v", err)
    }
}

func TestCreateNewUserForExistingUserThrowsException(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()
    username := "testUser"
    password := "testPassword"

    passkie.CreateNewUser(username, password)
    err := passkie.CreateNewUser(username, password)

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

    err := passkie.StoreCredentialsForSite(site, username, password, credentials)

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
    passkie.CreateNewUser(username, password)

    err := passkie.StoreCredentialsForSite(site, username, "wrongPassword", credentials)

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

    result, err := passkie.RetrieveCredentialsForSite(site, username, password)

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
    passkie.CreateNewUser(username, password)

    result, err := passkie.RetrieveCredentialsForSite(site, username, "wrongPassword")

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
    passkie.CreateNewUser(username, password)

    result, err := passkie.RetrieveCredentialsForSite(site, username, password)

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
    passkie.CreateNewUser(username, password)

    passkie.StoreCredentialsForSite(site, username, password, credentials)
    result, err := passkie.RetrieveCredentialsForSite(site, username, password)
    if err != nil {
        t.Fatalf("Failed to retrieve credentials for site: %v", err)
    }

    expectedResult := []map[string]string{credentials}
    if !reflect.DeepEqual(expectedResult, result) {
        t.Fatalf("Retrieved credentials are not the same as original! Original: %v, Retrieved: %v", expectedResult, result)
    }
}

func TestStoreCredentialsForSameSiteTwiceVerifyRetrieveCredentialsGetsBoth(t *testing.T) {
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
    passkie.CreateNewUser(username, password)

    passkie.StoreCredentialsForSite(site, username, password, credentials1)
    passkie.StoreCredentialsForSite(site, username, password, credentials2)
    result, err := passkie.RetrieveCredentialsForSite(site, username, password)
    if err != nil {
        t.Fatalf("Failed to retrieve credentials for site: %v", err)
    }

    expectedResult := []map[string]string{credentials1, credentials2}
    if !reflect.DeepEqual(expectedResult, result) {
        t.Fatalf("Retrieved credentials are not the same as expected! Expected: %v, Retrieved: %v", expectedResult, result)
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
    passkie.CreateNewUser(username, password)
    passkie.StoreCredentialsForSite(site, username, password, credentials)

    err := passkie.RemoveCredentialsForSite(site, username, "wrongPassword")

    if err == nil {
        t.Fatalf("RemoveCredentialsForSite should have returned exception with wrong password!")
    }
    retrievedCredentials, err := passkie.RetrieveCredentialsForSite(site, username, password)
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
    passkie.CreateNewUser(username, password)
    passkie.StoreCredentialsForSite(site, username, password, credentials)

    err := passkie.RemoveCredentialsForSite(site, username, password)
    if err != nil {
        t.Fatalf("Failed to remove credentials for site: %v", err)
    }

    retrievedCredentials, err := passkie.RetrieveCredentialsForSite(site, username, password)
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
    passkie.CreateNewUser(username, password)

    err := passkie.RemoveCredentialsForSite(site, username, password)

    if err == nil {
        t.Fatal("RemoveCredentialsForSite should have returned an error for a non existent site!")
    }
}

func TestRemoveSingleCredentialsFromSiteRemovesCredentialsProperly(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()

    username := "testUser"
    password := "testPassword"
    site := "testUrl.com"
    credentials1 := make(map[string]string)
    credentials1["testField1"] = "test"
    credentials1["testField2"] = "yeet"
    passkie.CreateNewUser(username, password)
    passkie.StoreCredentialsForSite(site, username, password, credentials1)
    credentials2 := make(map[string]string)
    credentials2["testField3"] = "hello"
    credentials2["testField4"] = "yuh"
    passkie.StoreCredentialsForSite(site, username, password, credentials2)

    err := passkie.RemoveSingleCredentialsForSite(site, username, password, 0)
    if err != nil {
        t.Fatalf("Failed to remove credentials for site: %v", err)
    }

    retrievedCredentials, err := passkie.RetrieveCredentialsForSite(site, username, password)
    if err != nil {
        t.Errorf("Retrieved credential after valid single credential remove should not have thrown an error! %v", err)
    }
    if retrievedCredentials == nil {
        t.Errorf("Retrieved credentials were not returned after valid remove single credential with multiple credentials stored")
    }

    expectedResult := []map[string]string{credentials2}
    if !reflect.DeepEqual(expectedResult, retrievedCredentials) {
        t.Fatalf("Retrieved credentials are not the same as expected! Expected: %v, Retrieved: %v", expectedResult, retrievedCredentials)
    }
}

func TestRemoveSingleCredentialWithASingleCredentialJustRemovesTheSite(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()

    username := "testUser"
    password := "testPassword"
    site := "testUrl.com"
    credentials1 := make(map[string]string)
    credentials1["testField1"] = "test"
    credentials1["testField2"] = "yeet"
    passkie.CreateNewUser(username, password)
    passkie.StoreCredentialsForSite(site, username, password, credentials1)

    err := passkie.RemoveSingleCredentialsForSite(site, username, password, 0)
    if err != nil {
        t.Fatalf("Failed to remove credentials for site: %v", err)
    }

    retrievedCredentials, err := passkie.RetrieveCredentialsForSite(site, username, password)
    if err == nil {
        t.Error("Retrieved credential after valid remove should have returned an error!")
    }
    if retrievedCredentials != nil {
        t.Errorf("Retrieved credentials were still returned after valid removal call: %v", retrievedCredentials)
    }
}

func TestRemoveSingleCredentialWithInvalidIndexShouldReturnError(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()

    username := "testUser"
    password := "testPassword"
    site := "testUrl.com"
    passkie.CreateNewUser(username, password)

    err := passkie.RemoveSingleCredentialsForSite(site, username, password, 1)

    if err == nil {
        t.Fatal("RemoveCredentialsForSite should have returned an error for a non existent site!")
    }
}

func TestRemoveSingleCredentialWithNegativeIndexShouldReturnError(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()

    username := "testUser"
    password := "testPassword"
    site := "testUrl.com"
    passkie.CreateNewUser(username, password)

    err := passkie.RemoveSingleCredentialsForSite(site, username, password, -1)

    if err == nil {
        t.Fatal("RemoveCredentialsForSite should have returned an error for a non existent site!")
    }
}

func TestRemoveUserOnNonExistentUserThrowsError(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()

    err := passkie.RemoveUser("fakeUsername", "fakePassword")

    if err == nil {
        t.Fatal("RemoveUser should have returned an error for a non existent user!")
    }
}

func TestRemoveUserWithWrongPasswordThrowsErrorAndKeepsUser(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()

    username := "testUser"
    password := "testPassword"
    passkie.CreateNewUser(username, password)

    err := passkie.RemoveUser(username, "wrongPassword")

    if err == nil {
        t.Error("RemoveUser with the wrong password should have thrown an error")
    }
    err = passkie.CreateNewUser(username, "wrongPassword")
    if err == nil {
        t.Error("CreateNewUser should not work for a user that was not deleted")
    }
}

func TestRemoveUserOnExistingUserAllowsNewUserCreation(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()

    username := "testUser"
    password := "testPassword"
    passkie.CreateNewUser(username, password)

    err := passkie.RemoveUser(username, password)

    if err != nil {
        t.Errorf("Failed to RemoveUser: %v", err)
    }
    err = passkie.CreateNewUser(username, password)
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
    passkie.CreateNewUser(username, originalPassword)
    site := "testUrl.com"
    credentials := make(map[string]string)
    credentials["testField1"] = "test"
    credentials["testField2"] = "yeet"

    passkie.RemoveUser(username, originalPassword)
    passkie.CreateNewUser(username, newPassword)
    err := passkie.StoreCredentialsForSite(site, username, originalPassword, credentials)

    if err == nil {
        t.Fatal("StoreCredentials with the old credentials of a previously deleted user should not work!")
    }
}

func TestRemoveUserOnExistingUserDoesNotSaveOldCredentials(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()

    username := "testUser"
    password := "testPassword"
    passkie.CreateNewUser(username, password)
    site := "testUrl.com"
    credentials := make(map[string]string)
    credentials["testField1"] = "test"
    credentials["testField2"] = "yeet"
    passkie.StoreCredentialsForSite(site, username, password, credentials)

    passkie.RemoveUser(username, password)
    passkie.CreateNewUser(username, password)
    res, err := passkie.RetrieveCredentialsForSite(site, username, password)
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
    passkie.CreateNewUser(username1, password1)
    passkie.CreateNewUser(username2, password2)

    passkie.RemoveUser(username2, password2)

    err := passkie.CreateNewUser(username1, password1)
    if err == nil {
        t.Errorf("Should not be able to recreate username1 when only username2 was deleted")
    }
    err = passkie.CreateNewUser(username2, password2)
    if err != nil {
        t.Errorf("Should have been able to recreate username 2 after it was deleted. Error: %v", err)
    }
}
