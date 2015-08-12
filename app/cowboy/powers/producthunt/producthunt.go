package producthunt

import (
	"log"
	"os"
)

const (
	APIBaseURL = "https://api.producthunt.com"
)

var (
	requestHeaders = map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
		"Host":         "api.producthunt.com",
	}
)

func init() {
	accessToken := os.Getenv("PRODUCTHUNT_ACCESS_TOKEN")
	if accessToken == "" {
		log.Fatal("[ERROR] Environment variable PRODUCTHUNT_ACCESS_TOKEN is not set")
	}

	requestHeaders["Authorization"] = "Bearer " + accessToken
}
