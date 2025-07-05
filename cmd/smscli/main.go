// cmd/smscli/main.go
package main

import (
	"fmt"
	"os"

	"sms-go/client"
	"sms-go/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Config error: %v\n", err)
		os.Exit(1)
	}

	// Initialize the appropriate SMS provider
	provider, err := client.NewProvider(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Provider error: %v\n", err)
		os.Exit(1)
	}

	// Example usage; replace with CLI flags or other input as needed
	ref := "ref123"
	recipients := []string{"+85366501580"}
	content := "Hello from Go!"
	lang := "E"

	results, err := provider.SendMessage(ref, recipients, content, lang)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Send error: %v\n", err)
		os.Exit(1)
	}

	for _, r := range results {
		fmt.Printf("To %s: JobID %s\n", r.Recipient, r.JobID)
	}
}
