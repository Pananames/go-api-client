package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/FozzyHosting/go-pananames"
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
	timeoutContext := pn.WithContext(ct)
	redeemInfo, err := pnCli.RedeemDomain("test.com", timeoutContext)
	if err != nil {
		log.Fatalf("Failed to redeem domain: %v", err)
	}
	fmt.Println(redeemInfo.NewExpirationDate)
}
