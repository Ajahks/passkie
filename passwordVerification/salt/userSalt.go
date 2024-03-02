package salt

import (
	"crypto/rand"

	"github.com/Ajahks/passkie/storage/localStorage/userSaltDb"
)

func GetSaltForUserHash(userHash []byte) []byte {
    userHashString := string(userHash)

    salt, err := userSaltDb.GetUserSalt(userHashString)
    if err != nil {
        newSalt := generateRandomSalt()

        PutSaltForUserHash(userHash, newSalt)

        return newSalt
    }

    return salt
}

func generateRandomSalt() []byte {
    saltLen := 32
    salt := make([]byte, saltLen)

    _, err := rand.Read(salt)
    if err != nil {
        panic(err)
    }

    return salt
}

func PutSaltForUserHash(userHash []byte, salt []byte) {
    userHashString := string(userHash)
    userSaltDb.PutUserSalt(userHashString, salt)
}

func RemoveSaltForUserHash(userHash []byte) {
    userHashString := string(userHash)
    userSaltDb.RemoveUserSalt(userHashString)
}

