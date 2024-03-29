package cmd

import (
	"errors"
	"fmt"
	"strings"

	passwordverification "github.com/Ajahks/passkie/passwordVerification"
	"golang.org/x/term"
)

var user string
var url string

func verifyMasterPasswordWorkflow() (string, error) {
    var retryCount = 3
    for i := 0; i < retryCount; i++ {
        fmt.Println("Enter a master password:")
        password, err := term.ReadPassword(0)
        if err != nil {
            fmt.Println("Failed to read password!")
            continue
        }

        user = strings.ToLower(user)
        if passwordverification.VerifyPasswordForUser(user, string(password)) {
            return string(password), nil 
        }
        fmt.Println("Incorrect password try again!")
    }
    
    fmt.Println("Master password was incorrect 3 times! User may also be incorrect. Ending session, try again")
    return "", errors.New("Incorrect master password")
}

func outputCredentials(credentialsList []map[string]string) {
    fmt.Println("Credentials:")
    fmt.Println("[")
    for index, credentials := range credentialsList {
        fmt.Printf("  %d: {\n", index)
        for field, credential := range credentials {
            fmt.Printf("    %s: %s\n", field, credential)
        }
        fmt.Println("  }")
    }
    fmt.Println("]")
}

