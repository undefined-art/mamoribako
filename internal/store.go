package internal

import (
	"encoding/json"
	"os"
)

type PasswordStore map[string][]byte

func LoadStore(filename string) (PasswordStore, error) {
	store := make(PasswordStore)
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return store, nil // empty store
		}
		return nil, err
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&store)
	return store, err
}

func SaveStore(filename string, store PasswordStore) error {
	tmpFile := filename + ".tmp"
	file, err := os.Create(tmpFile)
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(store); err != nil {
		file.Close()
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	return os.Rename(tmpFile, filename)
}
