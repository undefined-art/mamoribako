package main

import (
	"mamoribako/cmd"
	"mamoribako/internal"
)

func main() {
	internal.SetupSecureMemoryHandling()
	cmd.Execute()
}
