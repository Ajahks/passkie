package passwordverification

import (
	"bytes"

	"github.com/Ajahks/Passkie/passwordVerification/hash"
	"github.com/Ajahks/Passkie/passwordVerification/salt"
	passwordDB "github.com/Ajahks/Passkie/passwordVerification/storage/local/passwordHash"
)

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

