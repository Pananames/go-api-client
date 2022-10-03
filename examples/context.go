package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pananames/go-api-client"
)

// Redeem domain with context
// Context can be passed to any func
func RedeemWithContext() {
	pnClient, err := pananames.NewClient("token")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ct, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	timeoutContext := pananames.WithContext(ct)
	redeemInfo, err := pnClient.RedeemDomain("test.com", timeoutContext)
	if err != nil {
		log.Fatalf("Failed to redeem domain: %v", err)
	}
	fmt.Println(redeemInfo.NewExpirationDate)
}
