package salt

import (
	"crypto/rand"

	db "github.com/Ajahks/Passkie/passwordVerification/storage/local/userSalt"
)

func GetSaltForUserHash(userHash []byte) []byte {
    userHashString := string(userHash)

    salt, err := db.GetUserSalt(userHashString)
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
    db.PutUserSalt(userHashString, salt)
}

func RemoveSaltForUserHash(userHash []byte) {
    userHashString := string(userHash)
    db.RemoveUserSalt(userHashString)
}

