package hash

import "golang.org/x/crypto/argon2"

// Hashing function for the master password that will result in a salted hash that is safe to store
func HashPassword(masterPassword string, salt []byte) []byte {
 
    // Using recommended OWASP configuration settings here:
    // https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html#argon2id
    hash := argon2.IDKey([]byte(masterPassword), salt, 3, 12*1024, 1, 32)

    return hash
}
