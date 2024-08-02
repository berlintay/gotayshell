package main

import (
	"flag"
	"fmt"
	"os"

	"improvement"
	"test"
)

func main() {
	baseURL := flag.String("url", "http://example.com", "Base URL to check")
	flag.Parse()

	subdomains, err := improvement.FindSubdomains(*baseURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Found subdomains:")
	for _, subdomain := range subdomains {
		fmt.Println(subdomain)
	}

	test.TestFindSubdomains()
}