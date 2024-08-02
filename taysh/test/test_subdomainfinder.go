package test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestFindSubdomains(t *testing.T) {
	baseURL := "http://example.com"

	subdomains := []string{"www.", "api.", "admin.", "blog.", "forum."}

	for _, subdomain := range subdomains {
		for i := 1; i <= 10; i++ {
			url := fmt.Sprintf("%s%s%d", baseURL, subdomain, i)
			resp, err := http.Get(url)
			if err != nil {
				t.Fatalf("Error checking %s: %v", url, err)
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Error reading response body for %s: %v", url, err)
			}

			if resp.StatusCode == 200 {
				t.Logf("Found subdomain: %s", url)
			}
		}
	}
}