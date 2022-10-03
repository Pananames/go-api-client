package main

import (
	"fmt"
	"log"

	"github.com/pananames/go-api-client"
)

// Get all domains using pagination
func Pagination() {
	pnClient, err := pananames.NewClient("token")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	listOptions := &pananames.GetDomainsOptions{ListOptions: pananames.ListOptions{Limit: 30, Page: 1}}
	for {
		domains, page, err := pnClient.GetDomains(listOptions)
		if err != nil {
			fmt.Println(err)
		}
		for _, d := range domains {
			fmt.Println(d.Domain, d.Status)
		}
		if listOptions.Page = page.NextPage(); listOptions.Page == 0 {
			break
		}
	}
}
