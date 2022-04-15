package pananames

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

var wantRedirect = &Redirect{
	Url:            "https://test.com",
	MaskingEnabled: true,
	MaskingTitle:   "string",
	MaskingDesc:    "string",
	MaskingKwd:     "string",
}

func TestGetDomainRedirect(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"domains/test.com/redirect", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "redirect.json")
	})

	want := wantRedirect

	got, err := client.GetDomainRedirect("test.com")
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestEnableDomainRedirect(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	opts := EnableDomainRedirectOptions(*wantRedirect)
	mux.HandleFunc(apiVerPath+"domains/test.com/redirect", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)
		want, _ := json.Marshal(&opts)
		require.Equal(t, string(want), getBody(t, r))
		writeFixture(t, w, "redirect.json")
	})

	want := wantRedirect

	got, err := client.EnableDomainRedirect("test.com", &opts)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestDisableDomainRedirect(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"domains/test.com/redirect", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
	})

	err := client.DisableDomainRedirect("test.com")
	require.NoError(t, err)
}

func TestEnableBulkDomainRedirect(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	opts := &EnableBulkDomainRedirectOptions{
		EnableDomainRedirectOptions: EnableDomainRedirectOptions(*wantRedirect),
		DomainList:                  []string{"test.com", "test.app"},
	}
	mux.HandleFunc(apiVerPath+"domains/bulk_redirect", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)
		want, _ := json.Marshal(opts)
		require.Equal(t, string(want), getBody(t, r))
		writeFixture(t, w, "redirect_bulk.json")
	})

	want := &RedirectBulk{
		Redirect: *wantRedirect,
		DomainList: []*DomainRedirect{
			{Domain: "test.com", DomainQueued: true, Error: "string"},
			{Domain: "test.app", DomainQueued: true, Error: "string"},
		},
	}

	got, err := client.EnableBulkDomainRedirect(opts)
	require.NoError(t, err)
	require.Equal(t, want, got)
}
