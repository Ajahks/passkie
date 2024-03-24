package credentialsDb

import (
	"encoding/base64"

	localstorage "github.com/Ajahks/passkie/storage/localStorage"
)

const LIST_FILE_NAME = "credentialsListDB.txt"

func getEncodedUsername(username string) string {
    return base64.StdEncoding.EncodeToString([]byte(username))
}

func getFilePath(username string, filename string) string {
    // Not the most secure storing the usernames as base64 encoding, but slightly better than plaintext
    return localstorage.DB_PATH() + "/" + getEncodedUsername(username) + "/" + filename 
}

