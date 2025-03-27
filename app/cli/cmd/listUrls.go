package cmd

import (
	"fmt"
	"github.com/Ajahks/passkie/app/cli/storage/urlDb"
	"strings"

	"github.com/spf13/cobra"
)

// listUrlsCmd represents the listUrls command
var listUrlsCmd = &cobra.Command{
	Use:   "listUrls",
	Short: "Retrieves a list of urls that the user has credentials for",
	Long: `Retrieves a list of all urls that the user has credentials for

The user will be the username to query for.

The master password will be asked by the cli and verified.
`,
	Run: func(cmd *cobra.Command, args []string) {
		user = strings.ToLower(user)
		fmt.Printf("Retrieving url list for user: %s\n", user)

		_, err := verifyMasterPasswordWorkflow()
		if err != nil {
			return
		}

		urlList, err := urldb.ListUrlsForUser(user)
		if err != nil {
			fmt.Printf("Failed to find urls for user %s. Error: %v\n", user, err)
			return
		}

		fmt.Println("Urls:")
		fmt.Println("[")
		for _, url := range urlList {
			fmt.Printf("  %s\n", url)
		}
		fmt.Println("]")
	},
}

func init() {
	rootCmd.AddCommand(listUrlsCmd)

	listUrlsCmd.Flags().StringVarP(&user, "user", "u", "default", "passkie username. default:'default'")
}
