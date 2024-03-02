package passwordverification

import (
	"bytes"
	"errors"

	"github.com/Ajahks/passkie/passwordVerification/hash"
	"github.com/Ajahks/passkie/passwordVerification/salt"
	"github.com/Ajahks/passkie/storage/localStorage/activeUserDb"
	"github.com/Ajahks/passkie/storage/localStorage/passwordHashDb"
)

func SetPasswordForNewUser(username string, masterPassword string) error {
    unsaltedUserHash := hash.HashUsername(username, "")
    if activeuserdb.IsUserHashActive(string(unsaltedUserHash)) {
        return errors.New("Failed to set password for new user: User already exists")
    }
    activeuserdb.AddActiveUser(string(unsaltedUserHash))

    userhash := hash.HashUsername(username, masterPassword)
    usersalt := salt.GetSaltForUserHash(userhash)
    salt.PutSaltForUserHash(userhash, usersalt)

    passwordhash := hash.HashPassword(masterPassword, usersalt)
    passwordHashDb.PutPasswordHash(string(userhash), passwordhash)

    return nil
}

func VerifyPasswordForUser(username string, masterPassword string) bool {
    hashedUser := hash.HashUsername(username, masterPassword)

    userHashSalt := salt.GetSaltForUserHash(hashedUser)
    passwordHash := hash.HashPassword(masterPassword, userHashSalt)

    savedPasswordHash, err := passwordHashDb.GetPasswordHash(string(hashedUser)) 
    if err != nil {
        return false 
    }

    if bytes.Equal(savedPasswordHash, passwordHash) {
        return true
    }
    return false
}

func UpdatePasswordForUser(username string, currentPassword string, newPassword string) error {
    unsaltedUserHash := hash.HashUsername(username, "")
    if !activeuserdb.IsUserHashActive(string(unsaltedUserHash)) {
        return errors.New("Cannot update user password: User does not exist")
    }

    if !VerifyPasswordForUser(username, currentPassword) {
        return errors.New("Cannot update user password: current password is incorrect")
    }

    newUserHash := hash.HashUsername(username, newPassword)
    newUserSalt := salt.GetSaltForUserHash(newUserHash)
    salt.PutSaltForUserHash(newUserHash, newUserSalt)

    newPasswordhash := hash.HashPassword(newPassword, newUserSalt)
    passwordHashDb.PutPasswordHash(string(newUserHash), newPasswordhash)

    oldUserHash := hash.HashUsername(username, currentPassword)
    salt.RemoveSaltForUserHash(oldUserHash)
    passwordHashDb.RemovePasswordHash(string(oldUserHash))

    return nil
}

