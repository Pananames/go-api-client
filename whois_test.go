package pananames

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetWhoisInfo(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	opts := &GetWhoisInfoOptions{Preview: true}
	mux.HandleFunc(apiVerPath+"domains/test.com/whois", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, `preview=true`, r.URL.RawQuery)
		writeFixture(t, w, "whois.json")
	})

	want := &WhoisInfo{
		WhoisPrivacy:      true,
		Preview:           true,
		RegistrantContact: wantContact,
		AdminContact:      wantContact,
		TechContact:       wantContact,
		BillingContact:    wantContact,
	}

	got, err := client.GetWhoisInfo("test.com", opts)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestGetWhoisPrivacy(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"domains/test.com/whois_privacy", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "whois_privacy.json")
	})

	want := &WhoisPrivacy{
		Domain:  "test.com",
		Enabled: true,
	}

	got, err := client.GetWhoisPrivacy("test.com")
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestUpdateWhoisInfo(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	opts := &UpdateWhoisInfoOptions{
		RegistrantContact: wantContact,
		AdminContact:      wantContact,
		TechContact:       wantContact,
		BillingContact:    wantContact,
	}
	mux.HandleFunc(apiVerPath+"domains/test.com/whois", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)
		want, _ := json.Marshal(opts)
		require.Equal(t, string(want), getBody(t, r))
		writeFixture(t, w, "set_whois.json")
	})

	want := &WhoisInfo{
		WhoisPrivacy:      true,
		Preview:           true,
		RegistrantContact: wantContact,
		AdminContact:      wantContact,
		TechContact:       wantContact,
		BillingContact:    wantContact,
	}
	wantNotice := "string"

	got, gotNotice, err := client.UpdateWhoisInfo("test.com", opts)
	require.NoError(t, err)
	require.Equal(t, want, got)
	require.Equal(t, wantNotice, gotNotice)
}

func TestEnableWhoisPrivacy(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"domains/test.com/whois_privacy", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)
		writeFixture(t, w, "enable_whois_privacy.json")
	})

	want := &WhoisPrivacy{
		Domain:  "test.com",
		Enabled: true,
	}

	got, err := client.EnableWhoisPrivacy("test.com")
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestDisableWhoisPrivacy(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"domains/test.com/whois_privacy", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		writeFixture(t, w, "disable_whois_privacy.json")
	})

	want := &WhoisPrivacy{
		Domain:  "test.com",
		Enabled: false,
	}

	got, err := client.DisableWhoisPrivacy("test.com")
	require.NoError(t, err)
	require.Equal(t, want, got)
}
