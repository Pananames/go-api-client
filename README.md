# pananames-go-client

Pananames API v2 Golang Client

[![Test Status](https://github.com/Pananames/go-api-client/actions/workflows/test.yml/badge.svg)](https://github.com/Pananames/go-api-client/actions/workflows/test.yml)

## Developer Documentation

The actual API Documentation available on this [link](https://docs.pananames.com/).

## Usage

```go
import "github.com/pananames/go-api-client"

pnClient, err := pananames.NewClient("token")
if err != nil {
  log.Fatalf("Failed to create client: %v", err)
}

// get all domains; default limit per_page is 30
domainsInfo, _, err := pnClient.GetDomains(nil)
if err != nil {
	log.Fatalf("Failed to get domains info: %v", err)
}

// get only domains with 'suspended' status
listOptions := &pananames.GetDomainsOptions{Status: "suspended"}
domainsInfo, _, err = pnClient.GetDomains(listOptions)
if err != nil {
	log.Fatalf("Failed to get domains info: %v", err)
}

for _,d := range domainsInfo {
	fmt.Println(d.Domain)
}
```

### Examples

The [examples](examples) directory contains serveral examples of using this library.
