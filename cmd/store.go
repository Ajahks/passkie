/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	passkieApp "github.com/Ajahks/passkie/app"
	"github.com/spf13/cobra"
)

// storeCmd represents the store command
var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Starts workflow for storing a new set of credentials for the user",
	Long: `Initiates a cli workflow to store a new set of credentials for the user

Must pass in base url which will be what the credentials map to.  (Example: https://example.com/, http://test.net/)
Can pass in a username, or will default to a 'default' user if not passed. 
When the workflow starts, it will be a series of questions for the user.
- MasterPassword: the password used to initialize the user.  Will try 3 times before failing if password is incorrect.
- Then it will go through the credential adding workflow:
  1. Ask for credential field name (i.e. username, email, password)
  2. Ask for the credential itself
  3. Ask if there are more fields and rerun the first two steps.
- Credentials will be stored
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("storeCredentials called")

        fmt.Printf("Storing credentials for user: %s\n", user)
        fmt.Printf("Storing credentials for url: %s\n", url)
        password, err := verifyMasterPasswordWorkflow() 
        if err != nil {
            return
        }

        credentialsMap := inputCredentialsWorkflow()

        fmt.Println("Storing credentials for url!")
        passkieApp.StoreCredentialsForSite(url, user, password, credentialsMap)
	},
}

func inputCredentialsWorkflow() map[string]string {
    credentialsMap := make(map[string]string)
    
    for {
        var fieldName string
        fmt.Print("Enter the name of the credentials field (ex: 'Username' 'Password'): ")
        scanner := bufio.NewScanner(os.Stdin)
        if scanner.Scan() {
            fieldName = scanner.Text()
        }

        var credential string
        fmt.Printf("Enter the credentials to store for the field, %s: ", fieldName)
        if scanner.Scan() {
            credential = scanner.Text()
        }

        credentialsMap[fieldName] = credential

        var continuePrompt string
        fmt.Print("Add more credentials? [y/N]: ")
        _, err := fmt.Scanln(&continuePrompt)
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
	rootCmd.AddCommand(storeCmd)

    storeCmd.Flags().StringVarP(&user, "user", "u", "default", "passkie username. default:'default'")
    storeCmd.Flags().StringVarP(&url, "site", "s", "", "Base url to retrieve credentials for (Ex: http://example.com/, https://test.com/) REQUIRED")
    storeCmd.MarkFlagRequired("site")
}
