/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
    "fmt"
    urldb "github.com/Ajahks/passkie/app/cli/storage/urlDb"
    "strconv"
    "strings"

    passkieApp "github.com/Ajahks/passkie"
    "github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
    Use:   "remove",
    Short: "Removes either a user or site's credentials from the users DB",
    Long: `Completely removes the user and all stored credentials if no site provided to the command line.

If a user and a site is provided, only removes the site's credentials from the user's credential database`,
    Run: func(cmd *cobra.Command, args []string) {
        user = strings.ToLower(user)

        masterPassword, err := verifyMasterPasswordWorkflow()
        if err != nil {
            return
        }

        if len(url) == 0 {
            runUserDeletionWorkflow(masterPassword)
            return
        }

        runSiteDeletionWorkflow(masterPassword)
    },
}

func runSiteDeletionWorkflow(masterPassword string) {
    credentialIndex := -1
    if len(index) != 0 {
        convertedIndex, err := strconv.Atoi(index)
        if err == nil {
            if convertedIndex < 0 {
                fmt.Printf("Invalid index provided: %v\n", index)
                return
            }
            credentialIndex = convertedIndex
        } else {
            fmt.Printf("Malformed index provided: %v. %v\n", index, err)
            return
        }
    }

    if credentialIndex != -1 {
        fmt.Printf("Deleting credentials at index %v for site: %s .  This cannot be undone!\n", credentialIndex, url)
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

        err = passkieApp.RemoveSingleCredentialsForSite(url, user, masterPassword, credentialIndex)
        if err != nil {
            fmt.Printf("Failed to remove credentials for site: %v\n", err)
        }

        // Check if there are still credentials in the site or if it got deleted
        _, err = passkieApp.RetrieveCredentialsForSite(url, user, masterPassword)
        if err != nil {
            fmt.Println("All credentials were removed from the site.  Removing site from local db")
            urldb.RemoveActiveUrlForUser(url, user)
            return
        }

    } else {
        fmt.Printf("Deleting all credentials for site: %s .  This cannot be undone!\n", url)
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

        err = passkieApp.RemoveCredentialsForSite(url, user, masterPassword)
        if err != nil {
            fmt.Printf("Failed to remove credentials for site: %v\n", err)
        }

        fmt.Println("Removing url from local db")
        urldb.RemoveActiveUrlForUser(url, user)
    }
}

func runUserDeletionWorkflow(masterPassword string) {
    fmt.Printf("No site specified this will delete the user '%s' and all their credentials! This cannot be undone!\n", user)
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

    err = passkieApp.RemoveUser(user, masterPassword)
    if err != nil {
        fmt.Printf("Failed to remove the user: %v\n", err)
    }
}

func init() {
    rootCmd.AddCommand(removeCmd)
    removeCmd.Flags().StringVarP(&user, "user", "u", "default", "passkie username. default:'default'")
    removeCmd.Flags().StringVarP(&url, "site", "s", "", "Base url to remove credentials (Ex: http://example.com/, https://test.com/).  Exclude this if you want to delete the entire user")
    removeCmd.Flags().StringVarP(&index, "index", "i", "", "Index of a specific credential to remove.  Exclude this if you want to delete the entire site.  Only usable with -s flag")
}
