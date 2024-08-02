package improvement

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/charmbracelet/glow"
)

func FindSubdomains(baseURL string) ([]string, error) {
	subdomains := []string{"www.", "api.", "admin.", "blog.", "forum."}
	foundSubdomains := []string{}

	for _, subdomain := range subdomains {
		fmt.Printf("Checking subdomain %s...\n", subdomain)
		for i := 1; i <= 10; i++ {
			url := fmt.Sprintf("%s%s%d", baseURL, subdomain, i)
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("Error checking %s: %v\n", url, err)
				continue
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("Error reading response body for %s: %v\n", url, err)
				continue
			}

			if resp.StatusCode == 200 {
				fmt.Printf("Found subdomain: %s\n", url)
				foundSubdomains = append(foundSubdomains, url)
				printResponseBodyWithGlow(string(body))
			}
		}
	}

	return foundSubdomains, nil
}

func printResponseBodyWithGlow(body string) {
	// Clear the screen
	glow.Clear()

	// Print the response body
	glow.Println(body)

	// Refresh the screen
	glow.Flush()
}