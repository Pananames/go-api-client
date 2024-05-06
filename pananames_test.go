package pananames

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var wantDate = time.Date(2020, time.January, 2, 3, 4, 5, 0, time.UTC)
var wantDateOnly = time.Date(2020, time.January, 2, 0, 0, 0, 0, time.UTC)

var wantContact = &Contact{
	Org:     "string",
	Name:    "string",
	Email:   "string",
	Address: "string",
	City:    "string",
	State:   "string",
	Zip:     "string",
	Country: "string",
	Phone:   "string",
	Extras:  []string{},
}

var wantClaimContact = &ClaimContact{
	Name:         "string",
	Email:        "string",
	City:         "string",
	State:        "string",
	Zip:          "string",
	Phone:        "string",
	Organization: "string",
	Street:       "string",
	CountryCode:  "string",
}

// setup a test http server
func setup(t *testing.T) (*http.ServeMux, *httptest.Server, *Client) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	client, err := NewClient("secret", BaseURL(server.URL))
	if err != nil {
		server.Close()
		t.Fatalf("Failed to create a new client: %v", err)
	}
	return mux, server, client
}

// close the httptest server
func teardown(s *httptest.Server) {
	s.Close()
}

// returns body from http.Request as string
func getBody(t *testing.T, r *http.Request) string {
	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(r.Body)
	if err != nil {
		t.Fatalf("Failed to Read Body: %v", err)
	}
	return buffer.String()
}

// write fixture to specified io.Writer
func writeFixture(t *testing.T, w io.Writer, fixturePath string) {
	f, err := os.Open("testdata/" + fixturePath)
	if err != nil {
		t.Fatalf("failed to open fixture file: %v", err)
	}

	if _, err = io.Copy(w, f); err != nil {
		t.Fatalf("failed to copy fixture to writer: %v", err)
	}
}

func TestRequestHeaders(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"domains", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "secret", r.Header.Get("SIGNATURE"))
		require.Equal(t, "go-pananames", r.Header.Get("User-Agent"))
		require.Equal(t, "application/json", r.Header.Get("Accept"))
		writeFixture(t, w, "domains.json")
	})

	_, _, err := client.GetDomains(nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRequestOptions(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"domains", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, "current_page=1&per_page=10&status=suspended", r.URL.RawQuery)
		writeFixture(t, w, "domains.json")
	})

	_, _, err := client.GetDomains(&GetDomainsOptions{ListOptions: ListOptions{Limit: 10, Page: 1}, Status: "suspended"})
	if err != nil {
		t.Fatal(err)
	}
}
