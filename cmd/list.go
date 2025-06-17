package cmd

import (
	"fmt"
	"mamoribako/internal"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all stored keys",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := internal.LoadStore("passwords.json")
		if err != nil {
			return err
		}
		fmt.Println("Stored keys:")
		for k := range store {
			fmt.Println("-", k)
		}
		return nil
	},
}
