package internal

import (
	"fmt"
	"os"

	"github.com/awnumar/memguard"
	"golang.org/x/term"
)

func ReadMasterKey() (*memguard.LockedBuffer, error) {
	fmt.Print("Enter master key: ")
	bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()

	if err != nil {
		return nil, err
	}

	buf := memguard.NewBufferFromBytes(bytePassword)

	for i := range bytePassword {
		bytePassword[i] = 0
	}
	return buf, nil
}

func ZeroMasterKey(buf *memguard.LockedBuffer) {
	if buf != nil {
		buf.Destroy()
	}
}

func InitMemguard() {
	memguard.CatchInterrupt()
	memguard.Purge()
}

func SetupSecureMemoryHandling() {
	InitMemguard()
	// Additional setup if needed
}

func SaveLockedBufferToBytes(buf *memguard.LockedBuffer) []byte {
	return buf.Bytes()
}
