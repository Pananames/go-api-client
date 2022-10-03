package main

import (
	"fmt"
	"log"

	"github.com/pananames/go-api-client"
)

// Get, Set, Delete NameServers
func NameServers() {
	pnClient, err := pananames.NewClient("token")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	domain := "test.com"

	// get NS
	nameServers, err := pnClient.GetNameServers(domain)
	for _, ns := range *nameServers {
		fmt.Println(ns)
	}
	// set NS
	setNSOptions := &pananames.SetNameServersOptions{NameServers: pananames.NameServers{"ns1.test.com", "ns2.test.com"}}
	_, err = pnClient.SetNameServers(domain, setNSOptions)
	if err != nil {
		log.Fatalf("Failed to set domain name servers: %v", err)
	}

	// delete NS
	err = pnClient.DeleteNameServers(domain)
	if err != nil {
		log.Fatalf("Failed to delete domain name servers: %v", err)
	}
}
