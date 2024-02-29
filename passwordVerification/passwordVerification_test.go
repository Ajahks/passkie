package passwordverification

import (
	"testing"

	"github.com/Ajahks/Passkie/storage/localStorage"
)

func TestSetPasswordForNewUserTwiceFails(t *testing.T) {
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

