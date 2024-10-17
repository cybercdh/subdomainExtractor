package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/publicsuffix"
)

// extractSubdomains returns subdomains down to the specified depth
func extractSubdomains(fullDomain string, depth int) (string, error) {
	// Get the public suffix (TLD+1)
	eTLDPlusOne, err := publicsuffix.EffectiveTLDPlusOne(fullDomain)
	if err != nil {
		return "", err
	}

	// Split the domain into parts
	parts := strings.Split(fullDomain, ".")
	tldParts := strings.Split(eTLDPlusOne, ".")

	// Find the index of the TLD+1 in the parts
	tldIndex := len(parts) - len(tldParts)

	// Calculate the subdomain parts based on the depth
	startIndex := tldIndex - depth
	if startIndex < 0 {
		startIndex = 0
	}

	// Join the subdomain parts with the TLD+1
	subdomain := strings.Join(parts[startIndex:], ".")

	return subdomain, nil
}

func main() {
	// Define command-line flag for depth (default 0)
	depthPtr := flag.Int("d", 0, "Depth of subdomains (default: 0)")
	flag.Parse()

	// Create a scanner to read from stdin
	scanner := bufio.NewScanner(os.Stdin)

	// Process each line (each domain) from stdin
	for scanner.Scan() {
		domain := scanner.Text()

		// Extract subdomains at the specified depth
		result, err := extractSubdomains(domain, *depthPtr)
		if err != nil {
			log.Printf("Error processing domain %s: %v\n", domain, err)
			continue
		}

		// Print the result
		fmt.Println(result)
	}

	// Handle any errors from scanning stdin
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
}
