package pananames

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetTLDAddReqList(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"add_req_list", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "add_req_list.json")
	})
	want := []*TLDNotice{
		{
			TLD:     "CLUB",
			Notices: []string{"THIS NOTICE IS FOR CLUB"},
		},
		{
			TLD:     "INFO",
			Notices: []string{"This tld notice is for INFO"},
		},
	}
	got, err := client.GetTLDAddReqList()
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestGetTLDs(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"tlds", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "tlds.json")
	})
	want := []*TLD{
		{
			TLD:                   "XYZ",
			IDN:                   true,
			DNSSec:                true,
			Prices:                &Prices{Currency: "usd", Register: 9.79, Renew: 9.79, Transfer: 9.79, Redeem: 65.79},
			PromoTwoYearsPrices:   &Prices{Currency: "usd", Register: 2.64},
			PromoTwoYearsUntil:    &PnTime{wantDate},
			PromoMultiYearsPrices: map[string]*PromoMultiYears{"2": {PromoMultiYearsPrices: &Prices{Currency: "usd", Register: 2.64}, PromoMultiYearsUntil: &PnTime{wantDate}}},
		},
	}
	got, err := client.GetTLDs()
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestGetEmails(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"emails", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "emails.json")
	})
	want := []*Email{
		{
			Email:          "test@test1.com",
			FirstEmailDate: &PnTime{wantDate},
			SuspendDate:    &PnTime{wantDate},
			Status:         "unverified",
			Domains:        []*DomainStatus{{Domain: "test1.com", Status: "ok"}},
		},
		{
			Email:          "test@test2.com",
			FirstEmailDate: &PnTime{wantDate},
			SuspendDate:    &PnTime{wantDate},
			Status:         "unverified",
			Domains:        []*DomainStatus{{Domain: "test2.com", Status: "ok"}},
		},
	}
	wantPage := &Pagination{Total: 2, Limit: 30, Page: 1, Pages: 1}
	got, gotPage, err := client.GetEmails(nil)
	require.NoError(t, err)
	require.Equal(t, want, got)
	require.Equal(t, wantPage, gotPage)
}
func TestGetTLDAddReq(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"tlds/app/add_req", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "add_req.json")
	})
	want := &TLDNotice{TLD: "APP", Notices: []string{"Google notice"}}
	got, err := client.GetTLDAddReq("app")
	require.NoError(t, err)
	require.Equal(t, want, got)
}
