package pananames

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

var wantTransferInInfo = &TransferIn{
	Domain:            "test.com",
	TransferStatus:    "waiting registrant confirmation",
	InitDate:          &PnTime{wantDate},
	PremiumPrice:      123,
	WhoisPrivacy:      true,
	RegistrantContact: wantContact,
	AdminContact:      wantContact,
	TechContact:       wantContact,
	BillingContact:    wantContact,
	NameServers:       &NameServers{"string"},
	NameServerRecords: []*NameServerRecord{{
		ID:       "string",
		Name:     "string",
		Type:     "A",
		Value:    "string",
		Priority: 0,
		TTL:      0}},
}

func TestGetTransfersIn(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"transfers_in", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "transfers_in.json")
	})
	want := []*TransferIn{wantTransferInInfo}
	wantPage := &Pagination{Total: 1, Limit: 30, Page: 1, Pages: 1}
	got, gotPage, err := client.GetTransfersIn(nil)
	require.NoError(t, err)
	require.Equal(t, want, got)
	require.Equal(t, wantPage, gotPage)
}

func TestInitTransfersIn(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	opts := &InitTransferInOptions{
		Domain:            "test.com",
		AuthCode:          "123",
		PremiumPrice:      123,
		WhoisPrivacy:      true,
		RegistrantContact: wantContact,
		AdminContact:      wantContact,
		TechContact:       wantContact,
		BillingContact:    wantContact,
		NameServers:       &NameServers{"string"},
		NameServerRecords: []*NameServerRecord{{
			ID:       "string",
			Name:     "string",
			Type:     "A",
			Value:    "string",
			Priority: 0,
			TTL:      0,
		}},
	}

	mux.HandleFunc(apiVerPath+"transfers_in", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		want, _ := json.Marshal(opts)
		require.Equal(t, string(want), getBody(t, r))
		writeFixture(t, w, "transfers_in_init.json")
	})

	want := wantTransferInInfo
	got, err := client.InitTransferIn(opts)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestCancelTransfersIn(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	opts := &CancelTransferInOptions{Domain: "test.com"}
	mux.HandleFunc(apiVerPath+"transfers_in", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		want, _ := json.Marshal(opts)
		require.Equal(t, string(want), getBody(t, r))
	})

	err := client.CancelTransferIn(opts)
	require.NoError(t, err)
}

func TestCancelTransfersOut(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"domains/test.com/transfer_out", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
	})

	err := client.CancelTransferOut("test.com")
	require.NoError(t, err)
}

func TestInitTransfersOut(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"domains/test.com/transfer_out", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)
	})

	err := client.InitTransferOut("test.com")
	require.NoError(t, err)
}
