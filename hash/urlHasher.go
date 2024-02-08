package hash

import "crypto/sha256"

func HashUrl(url string, salt []byte) []byte {
    h := sha256.New()
    urlBytes := []byte(url)
    saltedUrlBytes := append(urlBytes, salt...)
    h.Write(saltedUrlBytes) 

    hashedUrl := h.Sum(nil)

    return hashedUrl 
}

