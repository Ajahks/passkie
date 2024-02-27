package passwordverification

import (
	"bytes"
	"errors"

	"github.com/Ajahks/Passkie/passwordVerification/hash"
	"github.com/Ajahks/Passkie/passwordVerification/salt"
	"github.com/Ajahks/Passkie/passwordVerification/storage/local/activeUserDb"
	passwordDB "github.com/Ajahks/Passkie/passwordVerification/storage/local/passwordHash"
)

func SetPasswordForNewUser(masterPassword string, username string) error {
    unsaltedUserHash := hash.HashUsername(username, "")
    if activeuserdb.IsUserHashActive(string(unsaltedUserHash)) {
        return errors.New("Failed to set password for new user: User already exists")
    }
    activeuserdb.AddActiveUser(string(unsaltedUserHash))

    userhash := hash.HashUsername(username, masterPassword)
    usersalt := salt.GetSaltForUserHash(userhash)
    salt.PutSaltForUserHash(userhash, usersalt)

    passwordhash := hash.HashPassword(masterPassword, usersalt)
    passwordDB.PutPasswordHash(string(userhash), passwordhash)

    return nil
}

func VerifyPasswordForUser(masterPassword string, username string) bool {
    hashedUser := hash.HashUsername(username, masterPassword)

    userHashSalt := salt.GetSaltForUserHash(hashedUser)
    passwordHash := hash.HashPassword(masterPassword, userHashSalt)

    savedPasswordHash, err := passwordDB.GetPasswordHash(string(hashedUser)) 
    if err != nil {
        return false 
    }

    if bytes.Equal(savedPasswordHash, passwordHash) {
        return true
    }
    return false
}

