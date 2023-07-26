package shared

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
)

func GetDataRoot() string {
	var root = "http://localhost:8080/mocked-data"
	var environment string

	if len(os.Args) > 1 {
		environment = os.Args[1]
	} else if os.Getenv("ENVIRONMENT") != "" {
		environment = os.Getenv("ENVIRONMENT")
	}

	if environment == "staging" {
		fmt.Println("Staging environment!")
	} else if environment == "development" {
		// Disable certificate verification
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		fmt.Println("Development environment!")
		root = "https://sss-erp-bff.oykos.me/mocked-data"
	}

	return root
}
