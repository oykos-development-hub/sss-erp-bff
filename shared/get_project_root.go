package shared

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetProjectRoot() (string, error) {
	if len(os.Args) > 1 {
		environment := os.Args[1]

		if environment == "staging" {
			fmt.Println("Staging environment!")
			return "", nil
		} else if environment == "development" {
			fmt.Println("Development environment!")
			return "/var/www/vhosts/oykos.me/sss-erp-bff.oykos.me/github/sss-erp-bff", nil
		}
		return "", fmt.Errorf("wrong environment flag passed %s", environment)
	}

	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	// Traverse up the directory tree until we find a directory containing a main.go file
	for {
		// Check if we're at the root directory
		if filepath.Dir(cwd) == cwd {
			return "", fmt.Errorf("inable to find project root")
		}
		// Check if the current directory contains a main.go file
		if _, err := os.Stat(filepath.Join(cwd, "main.go")); err == nil {
			return cwd, nil
		}
		// Move up one directory
		cwd = filepath.Dir(cwd)
	}
}
