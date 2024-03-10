package passwordverification

import (
	"testing"

	"github.com/Ajahks/passkie/storage/localStorage"
)

func TestSetPasswordForNewUserTwiceFails(t *testing.T) {
    localstorage.SetTestDb()
    username := "testUsername"
    masterPassword := "testPassword"

    err := SetPasswordForNewUser(username, masterPassword)
    if err != nil {
        t.Errorf("Failed to set password for new user the first time: %v", err)
    }
    err = SetPasswordForNewUser(username, masterPassword)
    if err == nil {
        t.Error("Setting password for the same user a second time should have returned an error")
    }

    localstorage.CleanDB()
}

func TestSetPasswordForNewUserThenVerifyPasswordWithCorrectPasswordReturnsTrue(t *testing.T) {
    localstorage.SetTestDb()
    username := "testUsername"
    masterPassword := "testPassword"

    err := SetPasswordForNewUser(username, masterPassword)
    if err != nil {
        t.Errorf("Failed to set password for new user the first time: %v", err)
    }
    ok := VerifyPasswordForUser(username, masterPassword)

    if !ok {
        t.Error("Password is deemed incorrect for the user")
    }

    localstorage.CleanDB()
}

func TestSetPasswordForNewUserThenVerifyPasswordWithWrongPasswordReturnsFalse(t *testing.T) {
    localstorage.SetTestDb()
    username := "testUsername"
    masterPassword := "testPassword"

    err := SetPasswordForNewUser(username, masterPassword)
    if err != nil {
        t.Errorf("Failed to set password for new user the first time: %v", err)
    }
    ok := VerifyPasswordForUser(username, "fakePassword")

    if ok {
        t.Error("Password is deemed correct for the user when it isn't")
    }

    localstorage.CleanDB()
}

func TestUpdatePasswordForUserThenVerifyNewPasswordSucceeds(t *testing.T) {
    localstorage.SetTestDb()
    username := "testUsername"
    firstPassword := "testPassword"
    secondPassword := "testPassword2"
    SetPasswordForNewUser(username, firstPassword)

    err := UpdatePasswordForUser(username, firstPassword, secondPassword)
    if err != nil {
        t.Errorf("Failed to update password for user: %v", err)
    }
    ok := VerifyPasswordForUser(username, secondPassword)
    if !ok {
        t.Error("New password is failing to verify after password update")
    }

    localstorage.CleanDB()
}

func TestUpdatePasswordForUserThenVerifyOldPasswordFails(t *testing.T) {
    localstorage.SetTestDb()
    username := "testUsername"
    firstPassword := "testPassword"
    secondPassword := "testPassword2"
    SetPasswordForNewUser(username, firstPassword)

    err := UpdatePasswordForUser(username, firstPassword, secondPassword)
    if err != nil {
        t.Errorf("Failed to update password for user: %v", err)
    }
    ok := VerifyPasswordForUser(username, firstPassword)
    if ok {
        t.Error("First password should no longer work after update but it is still working!")
    }

    localstorage.CleanDB()
}

func TestUpdatePasswordForUserWithNewUserThrowsError(t *testing.T) {
    localstorage.SetTestDb()
    username := "testUsername"
    firstPassword := "testPassword"
    secondPassword := "testPassword2"

    err := UpdatePasswordForUser(username, firstPassword, secondPassword)
    if err == nil {
        t.Errorf("Updating password for new user should have thrown an error")
    }

    localstorage.CleanDB()
}

func TestUpdatePasswordForUserWithWrongCurrentPasswordFails(t *testing.T) {
    localstorage.SetTestDb()
    username := "testUsername"
    firstPassword := "testPassword"
    secondPassword := "testPassword2"
    SetPasswordForNewUser(username, firstPassword)

    err := UpdatePasswordForUser(username, secondPassword, secondPassword)
    if err == nil {
        t.Errorf("Updating password for wrong current password should have thrown an error")
    }

    localstorage.CleanDB()
}

func TestRemoveUserStopsPasswordVerificationFromWorkingEvenWithCorrectPassword(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()
    username := "testUsername"
    password := "testPassword"
    SetPasswordForNewUser(username, password)
    
    err := RemoveUser(username, password)
    if err != nil {
        t.Fatalf("Failed to DeactivateUser: %v", err)
    }
    ok := VerifyPasswordForUser(username, password)
    
    if ok {
        t.Errorf("VerifyPasswordForUser still succeeds after deactivation!")
    }
}

func TestRemoveUserAllowsSetPasswordForNewUserAgain(t *testing.T) {
    localstorage.SetTestDb()
    defer localstorage.CleanDB()
    
    username := "testUsername"
    password1 := "testPassword1"
    password2 := "testPassword2"
    SetPasswordForNewUser(username, password1)

    err := RemoveUser(username, password1)
    if err != nil {
        t.Fatalf("Failed to RemoveUser: %v", err)
    }
    err = SetPasswordForNewUser(username, password2)
    if err != nil {
        t.Fatalf("Failed to recreate new user after removal: %v", err)
    }

    ok := VerifyPasswordForUser(username, password2)
    if !ok {
        t.Fatalf("New password is returning false for VerifyPassword")
    }
    ok = VerifyPasswordForUser(username, password1)
    if ok {
        t.Fatalf("Old password should not be working after RemoveUser")
    }
}
