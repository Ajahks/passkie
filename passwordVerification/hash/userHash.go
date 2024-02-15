package hash

import "crypto/sha256"

func HashUsername(username string, masterPassword string) []byte {
    passwordBytes := []byte(masterPassword)
    usernameBytes := []byte(username)

    saltedUsernameBytes := append(usernameBytes, passwordBytes...) 

    h := sha256.New()
    h.Write(saltedUsernameBytes)
    hashedUsername := h.Sum(nil)
     
    return hashedUsername
}

