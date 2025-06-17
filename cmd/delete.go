package cmd

import (
	"fmt"
	"mamoribako/internal"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [key]",
	Short: "Delete a password",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]

		store, err := internal.LoadStore("passwords.json")
		if err != nil {
			return err
		}

		if _, exists := store[key]; !exists {
			return fmt.Errorf("no password found for key: %s", key)
		}

		delete(store, key)
		if err := internal.SaveStore("passwords.json", store); err != nil {
			return err
		}

		fmt.Println("Deleted password for key:", key)
		return nil
	},
}
