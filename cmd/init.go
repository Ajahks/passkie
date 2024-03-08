package cmd

import (
	"bytes"
	"fmt"
	"strings"

	passkieApp "github.com/Ajahks/passkie/app"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes passkie with a user and master password",
	Long: `Sets the username and master password for passkie.

User must enter a master password twice.  User must also not be already created.
    `,
	Run: func(cmd *cobra.Command, args []string) {
        user = strings.ToLower(user)
        fmt.Printf("Initializing with username: %s\n", user)

        fmt.Println("Enter a master password:")
        password, err := term.ReadPassword(0)
        if err != nil {
            fmt.Println("Failed to read password!")
            return
        }
        if !validatePassword(password) { return }

        fmt.Println("Re-enter master password:")
        passwordVerification, err := term.ReadPassword(0)
        if err != nil {
            fmt.Println("Failed to read password!")
            return
        }

        if !bytes.Equal(password, passwordVerification) {
            fmt.Println("Passwords do not match!")
            return 
        }

        err = passkieApp.CreateNewUser(user, string(password))
        if err != nil {
            fmt.Printf("Error found creating user: %v\n", err)
            return
        }
        fmt.Printf("User %s successfuly created!", user)
	},
}

func validatePassword(password []byte) bool {
    if len(password) == 0 {
        fmt.Println("Password must not be empty!")
        return false 
    }

    return true
}

func init() {
	rootCmd.AddCommand(initCmd)
    initCmd.Flags().StringVarP(&user, "user", "u", "default", "passkie username. Non case sensitive. Default:'default'")
}
