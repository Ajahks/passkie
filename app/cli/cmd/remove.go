/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes either a user or site's credentials from the users DB",
	Long: `Completely removes the user and all stored credentials if no site provided to the command line.

If a user and a site is provided, only removes the site's credentials from the user's credential database`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
        user = strings.ToLower(user)

        if len(url) == 0 {
            runUserDeletionWorkflow()
            return
        }

        runSiteDeletionWorkflow()
	},
}

func runSiteDeletionWorkflow() {
    fmt.Printf("Deleting credentials for site: %s .  This cannot be undone!", url)
    var continuePrompt string
    fmt.Print("Continue? [y/N]: ")
    _, err := fmt.Scanln(&continuePrompt)
    if err != nil {
        fmt.Printf("Failed to read input: %v\n", err)
        return
    }
    if continuePrompt != "y" && continuePrompt != "Y" {
        return 
    }

    // TODO: Remove credentials
}

func runUserDeletionWorkflow() {
    fmt.Printf("No site specified this will delete the user '%s' and all their credentials! This cannot be undone!", user)
    var continuePrompt string
    fmt.Print("Continue? [y/N]: ")
    _, err := fmt.Scanln(&continuePrompt)
    if err != nil {
        fmt.Printf("Failed to read input: %v\n", err)
        return
    }
    if continuePrompt != "y" && continuePrompt != "Y" {
        return 
    }

    // TODO: Remove user
}

func init() {
	rootCmd.AddCommand(removeCmd)
    removeCmd.Flags().StringVarP(&user, "user", "u", "default", "passkie username. default:'default'")
    removeCmd.Flags().StringVarP(&url, "site", "s", "", "Base url to retrieve credentials for (Ex: http://example.com/, https://test.com/) REQUIRED")
}

