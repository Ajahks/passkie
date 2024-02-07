package hash

import "crypto/sha256"

// Hashing function for the master password.  This password will not be salted because this is not meant to be stored
func hashPassword(password string) []byte {
    h := sha256.New()
    h.Write([]byte(password)) 

    return h.Sum(nil) 
}
