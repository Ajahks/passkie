package hash

import "crypto/sha256"

func HashUrl(url string, masterPassword string) []byte {
    h := sha256.New()
    urlBytes := []byte(url)
    hashedPassword := HashPassword(masterPassword)
    saltedUrlBytes := append(urlBytes, hashedPassword...)
    h.Write(saltedUrlBytes) 

    hashedUrl := h.Sum(nil)

    return hashedUrl 
}

