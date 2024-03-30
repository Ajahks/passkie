package cmd

import (
	"fmt"
	"strings"

	passkieApp "github.com/Ajahks/passkie"
	"github.com/spf13/cobra"
)

// retrieveCmd represents the retrieve command
var retrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves credentials stored for site",
	Long: `Retrieves credentials stored for a given site.

The base url will be the url to query for.
The user will be the username to query for.

The master password will be asked by the cli and verified.
If credentials exist, will return those credentials.
`,
	Run: func(cmd *cobra.Command, args []string) {
        user = strings.ToLower(user)
        fmt.Printf("Retrieving credentials for user: %s\n", user)
        fmt.Printf("Retrieving credentials for site: %s\n", url)

        password, err := verifyMasterPasswordWorkflow() 
        if err != nil {
            return
        }

        credentials, err := passkieApp.RetrieveCredentialsForSite(url, user, password)
        if err != nil {
            fmt.Printf("Failed to read credentials for site %s: %v\n", url, err)
            return
        }
        
        outputCredentials(credentials)
	},
}

func init() {
	rootCmd.AddCommand(retrieveCmd)

    retrieveCmd.Flags().StringVarP(&user, "user", "u", "default", "passkie username. default:'default'")
    retrieveCmd.Flags().StringVarP(&url, "site", "s", "", "Base url to retrieve credentials for (Ex: http://example.com/, https://test.com/) REQUIRED")
    retrieveCmd.MarkFlagRequired("site")
}

