package main

import (
	"fmt"
	"log"

	"github.com/pananames/go-api-client"
)

// Check bulk domains
func CheckDomainsBulk() {
	pnClient, err := pananames.NewClient("token")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	domains := []string{"test.com", "test.org", "test.info"}
	checkBulkOptions := pananames.CheckDomainsBulkOptions{Domains: domains}
	checkBulkDomainInfo, err := pnClient.CheckDomainsBulk(&checkBulkOptions)
	if err != nil {
		log.Fatalf("Failed to get bulk domain info: %v", err)
	}
	for _, d := range checkBulkDomainInfo {
		fmt.Printf("Domain: %s, Available: %v, Premium: %v\n", d.Domain, d.Available, d.Premium)
	}
}

// Get info for single domain or for all domains
func GetDomain() {
	pnClient, err := pananames.NewClient("token")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// single domain
	domain := "test.com"
	domainInfo, err := pnClient.GetDomain(domain)
	if err != nil {
		log.Fatalf("Failed to get domain info: %v", err)
	}
	fmt.Println(domainInfo.NameServers)

	// all domains
	domainsInfo, _, err := pnClient.GetDomains(nil)
	for _, d := range domainsInfo {
		fmt.Println(d.Domain)
	}

	// with filter Status
	listOptions := &pananames.GetDomainsOptions{Status: "suspended"}
	domainsInfo, _, err = pnClient.GetDomains(listOptions)
	for _, d := range domainsInfo {
		fmt.Println(d.Domain)
	}
}

// Check domain availability and register it
func RegisterDomain() {
	pnClient, err := pananames.NewClient("token")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	domain := "test.com"
	checkDomainInfo, err := pnClient.CheckDomain(domain)
	if checkDomainInfo.Available {

		contact := pananames.Contact{
			Org:     "Org",
			Name:    "Name",
			Email:   "test@email.com",
			Address: "Address",
			City:    "City",
			State:   "State",
			Zip:     "Zip",
			Country: "US",
			Phone:   "+1.234567890",
		}
		regOptions := pananames.RegisterDomainOptions{
			Domain:            domain,
			Period:            1,
			WhoisPrivacy:      true,
			RegistrantContact: &contact,
			AdminContact:      &contact,
			BillingContact:    &contact,
			TechContact:       &contact,
		}
		regResp, err := pnClient.RegisterDomain(&regOptions)
		if err != nil {
			log.Fatalf("Failed to register domain: %v", err)
		}
		fmt.Println(regResp.ExpirationDate)
	}
}

func DeleteDomain() {
	pnClient, err := pananames.NewClient("token")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	domain := "test.com"
	err := pnClient.DeleteDomain(domain)
	if err != nil {
		log.Fatalf("Failed to delete domain: %v", err)
	}
}
