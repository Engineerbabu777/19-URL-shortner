package helpers

import (
	"os"
	"strings"
)

func EnforcedHTTP(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}

func RemoveDomain(url string) bool {

	if url == os.Getenv("DOMAIN") {
		return false
	}

	// REPLACING EXPECTED URLS!
	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)

	// SPLITTING THE URL BY (/)
	newURL = strings.Split(newURL, "/")[0]

	if newURL == os.Getenv("DOMAIN") {
		return false
	}

	return false

}
