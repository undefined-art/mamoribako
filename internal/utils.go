package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadLine(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(line), nil
}

func Confirm(prompt string) (bool, error) {
	for {
		input, err := ReadLine(prompt + " (y/n): ")
		if err != nil {
			return false, err
		}
		switch strings.ToLower(input) {
		case "y", "yes":
			return true, nil
		case "n", "no":
			return false, nil
		default:
			fmt.Println("Please enter y or n.")
		}
	}
}
