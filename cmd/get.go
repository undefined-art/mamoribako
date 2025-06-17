package cmd

import (
	"fmt"
	"mamoribako/internal"
	"os"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [key]",
	Short: "Retrieve a password",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]

		masterKey, err := internal.ReadMasterKey()
		if err != nil {
			return err
		}
		defer internal.ZeroMasterKey(masterKey)

		salt, err := os.ReadFile("salt.bin")
		if err != nil {
			return fmt.Errorf("please run 'init' first: %w", err)
		}
		derivedKey := internal.DeriveKey(masterKey.Bytes(), salt)

		store, err := internal.LoadStore("passwords.json")
		if err != nil {
			return err
		}

		encrypted, ok := store[key]
		if !ok {
			return fmt.Errorf("no password found for key: %s", key)
		}

		decrypted, err := internal.Decrypt(derivedKey, encrypted)
		if err != nil {
			return err
		}

		fmt.Printf("Password for %s: %s\n", key, decrypted)
		return nil
	},
}
