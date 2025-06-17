package cmd

import (
	"fmt"
	"mamoribako/internal"
	"os"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [key] [password]",
	Short: "Add a password",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		key := args[0]
		password := args[1]

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

		encrypted, err := internal.Encrypt(derivedKey, []byte(password))
		if err != nil {
			return err
		}

		store[key] = encrypted
		if err := internal.SaveStore("passwords.json", store); err != nil {
			return err
		}

		fmt.Println("Password added for key:", key)
		return nil
	},
}
