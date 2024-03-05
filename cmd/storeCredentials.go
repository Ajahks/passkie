/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	passkieApp "github.com/Ajahks/passkie/app"
	passwordverification "github.com/Ajahks/passkie/passwordVerification"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var storeUser string

// storeCredentialsCmd represents the storeCredentials command
var storeCredentialsCmd = &cobra.Command{
	Use:   "storeCredentials",
	Short: "Starts workflow for storing a new set of credentials for the user",
	Long: `Initiates a cli workflow to store a new set of credentials for the user

Can pass in a username, or will default to a 'default' user if not passed. 
When the workflow starts, it will be a series of questions for the user.
- MasterPassword: the password used to initialize the user.  Will try 3 times before failing if password is incorrect.
- Base url of site: the base url of the site in the format http://mysite.com or https://mysite.com/
- Then it will go through the credential adding workflow:
  1. Ask for credential field name (i.e. username, email, password)
  2. Ask for the credential itself
  3. Ask if there are more fields and rerun the first two steps.
- Credentials will be stored
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("storeCredentials called")

        fmt.Printf("Storing credentials for user: %s\n", storeUser)
        password, err := verifyMasterPasswordWorkflow() 
        if err != nil {
            return
        }

        fmt.Print("Input base url (Example: http://example.com/, https://example.com/): ")
        var url string
        _, err = fmt.Scanln(&url)
        if err != nil {
            fmt.Printf("Failed to read input: %v\n", err)
        }
        fmt.Printf("Adding credentials for url: %s\n", url)

        credentialsMap := inputCredentialsWorkflow()

        fmt.Println("Storing credentials for url!")
        passkieApp.StoreCredentialsForSite(url, storeUser, password, credentialsMap)
	},
}

func verifyMasterPasswordWorkflow() (string, error) {
    var retryCount = 3
    for i := 0; i < retryCount; i++ {
        fmt.Println("Enter a master password:")
        password, err := term.ReadPassword(0)
        if err != nil {
            fmt.Println("Failed to read password!")
            continue
        }

        if passwordverification.VerifyPasswordForUser(storeUser, string(password)) {
            return string(password), nil 
        }
        fmt.Println("Incorrect password try again!")
    }
    
    fmt.Println("Master password was incorrect 3 times! Ending session")
    return "", errors.New("Incorrect master password")
}

func inputCredentialsWorkflow() map[string]string {
    credentialsMap := make(map[string]string)
    
    for {
        var fieldName string
        fmt.Print("Enter the name of the credentials field (ex: 'Username' 'Password'): ")
        _, err := fmt.Scanln(&fieldName)
        if err != nil {
            fmt.Printf("Failed to read input: %v\n", err)
            break
        }

        var credential string
        fmt.Printf("Enter the credentials to store for the field, %s: ", fieldName)
        _, err = fmt.Scanln(&credential)
        if err != nil {
            fmt.Printf("Failed to read input: %v\n", err)
            break
        }

        credentialsMap[fieldName] = credential

        var continuePrompt string
        fmt.Print("Add more credentials? [y/N]: ")
        fmt.Scanln(&continuePrompt)
        if err != nil {
            fmt.Printf("Failed to read input: %v\n", err)
            break
        }

        if continuePrompt != "y" && continuePrompt != "Y" {
            break 
        }
    }

    return credentialsMap
}

func init() {
	rootCmd.AddCommand(storeCredentialsCmd)

    storeCredentialsCmd.Flags().StringVarP(&storeUser, "user", "u", "default", "passkie username. default:'default'")
}
