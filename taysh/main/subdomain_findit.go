package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/exec"
	"strings"

	"github.com/charmbracelet/glow"
)

func main() {
	baseURL := flag.String("url", "http://example.com", "Base URL to check")
	flag.Parse()

	subdomainFinder := NewSubdomainFinder(*baseURL)
	err := subdomainFinder.FindSubdomains()
	if err != nil {
		fmt.Println(err)
		return
	}
}

type SubdomainFinder struct {
	baseURL string
}

func NewSubdomainFinder(baseURL string) *SubdomainFinder {
	return &SubdomainFinder{baseURL: baseURL}
}

func (f *SubdomainFinder) FindSubdomains() error {
	subdomains := []string{"www.", "api.", "admin.", "blog.", "forum."}

	for _, subdomain := range subdomains {
		fmt.Printf("Checking subdomain %s...\n", subdomain)
		for i := 1; i <= 10; i++ {
			url := fmt.Sprintf("%s%s%d", f.baseURL, subdomain, i)
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
				cmd := exec.Command("warn_script", url)
				err = cmd.Run()
				if err != nil {
					fmt.Printf("Error running warn script for %s: %v\n", url, err)
				}
				printResponseBodyWithGlow(string(body))
			}
		}
	}

	return nil
}

func printResponseBodyWithGlow(body string) {
	// Clear the screen
	glow.Clear()

	// Print the response body
	glow.Println(body)

	// Refresh the screen
	glow.Flush()
}