package utils

import (
	"os"
	"strings"
)

func IsDifferentDomain() bool {
	domain := os.Getenv("DOMAIN")

	if url == domain {
		return false
	}

	cleanURL := strings.TrimPrefix(url, "http://")
	cleanURL = strings.TrimPrefix(cleanURL, "https://")
	cleanURL = strings.TrimPrefix(cleanURL, "www.")
	cleanURL = strings.Split(cleanURL, "/")[0]

	return cleanURL != domain
}
