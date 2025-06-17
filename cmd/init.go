package cmd

import (
	"crypto/rand"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the password manager (generate salt)",
	RunE: func(cmd *cobra.Command, args []string) error {
		salt := make([]byte, 16)
		if _, err := rand.Read(salt); err != nil {
			return err
		}
		// Save salt to file
		if err := os.WriteFile("salt.bin", salt, 0600); err != nil {
			return err
		}
		fmt.Println("Initialization complete. Salt generated and saved.")
		return nil
	},
}
