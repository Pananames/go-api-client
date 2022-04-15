package pananames

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAccountBalance(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"account/balance", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "balance.json")
	})
	want := &Balance{12.34}
	got, err := client.GetAccountBalance()
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestGetAccountPayments(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"account/payments", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "payments.json")
	})
	want := []*Payment{
		{
			TxID:   "12345",
			TxDate: &PnTime{wantDate},
			TxType: "create",
			Domain: "test.com",
			Period: "1",
			Total:  -1.23,
		},
	}
	wantPage := &Pagination{Total: 1, Limit: 30, Page: 1, Pages: 1}
	got, gotPage, err := client.GetAccountPayments(nil)
	require.NoError(t, err)
	require.Equal(t, want, got)
	require.Equal(t, wantPage, gotPage)
}
