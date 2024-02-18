package salt

import (
	"crypto/rand"

	"github.com/Ajahks/Passkie/passwordVerification/storage"
)

func GetSaltForUserHash(userHash []byte) []byte {
    userHashString := string(userHash)

    salt, err := storage.GetUserSalt(userHashString)
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
    storage.PutUserSalt(userHashString, salt)
}

