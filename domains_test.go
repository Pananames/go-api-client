package pananames

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

var wantPrices = &Prices{
	Currency: "usd",
	Register: 1.00,
	Renew:    2.00,
	Transfer: 3.00,
	Redeem:   4.00,
}
var wantDomainInfo = &Domain{
	Domain:           "test.com",
	DomainIDN:        "string",
	Premium:          true,
	AutoRenew:        true,
	WhoisPrivacy:     true,
	LockStatus:       "unlocked",
	RegistrationDate: &PnTime{wantDate},
	ExpirationDate:   &PnTime{wantDate},
	DeletionDate:     &PnTime{wantDate},
	Status:           "ok",
	NameServers:      &NameServers{"string"},
	ChildNameServers: []*ChildNameServer{{
		Hostname: "string",
		IPv4:     "string",
		IPv6:     "string"},
	}}

func TestGetDomains(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"domains", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "domains.json")
	})

	want := []*Domain{wantDomainInfo}
	wantPage := &Pagination{Total: 1, Limit: 30, Page: 1, Pages: 1}

	got, gotPage, err := client.GetDomains(nil)
	require.NoError(t, err)
	require.Equal(t, want, got)
	require.Equal(t, wantPage, gotPage)
}

func TestGetDomain(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"domains/test.com", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "domain.json")
	})

	want := wantDomainInfo

	got, err := client.GetDomain("test.com")
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestRegisterDomain(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	opts := &RegisterDomainOptions{
		Domain:            "test.com",
		Period:            1,
		WhoisPrivacy:      true,
		ClaimsAccepted:    false,
		AddReqAccepted:    false,
		RegistrantContact: wantContact,
		TechContact:       wantContact,
		BillingContact:    wantContact,
		AdminContact:      wantContact,
	}

	mux.HandleFunc(apiVerPath+"domains", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		want, _ := json.Marshal(opts)
		require.Equal(t, string(want), getBody(t, r))
		writeFixture(t, w, "domain.json")
	})

	want := wantDomainInfo
	got, err := client.RegisterDomain(opts)
	require.NoError(t, err)
	require.Equal(t, want, got)
}
func TestDeleteDomain(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"domains/test.com", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
	})

	err := client.DeleteDomain("test.com")
	require.NoError(t, err)
}

func TestCheckDomain(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"domains/test.com/check", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "domain_check.json")
	})

	want := &DomainCheck{
		Domain:                "test.xyz",
		Available:             false,
		Premium:               false,
		Prices:                wantPrices,
		PromoTwoYearsPrices:   wantPrices,
		PromoMultiYearsPrices: map[string]*Prices{"2": wantPrices},
		Claim:                 false,
		AddReq:                false,
	}

	got, err := client.CheckDomain("test.com")
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestCheckDomainsBulk(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"domains/bulk_check", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		require.Equal(t, `domains=test.xyz%2Ctest.com`, r.URL.RawQuery)
		writeFixture(t, w, "domain_bulk_check.json")
	})

	want := []*DomainCheck{
		{
			Domain:                "test.xyz",
			Available:             false,
			Premium:               false,
			Prices:                wantPrices,
			PromoTwoYearsPrices:   wantPrices,
			PromoMultiYearsPrices: map[string]*Prices{"2": wantPrices},
			Claim:                 false,
			AddReq:                false,
		},
		{
			Domain:    "test.com",
			Available: false,
			Premium:   false,
			Prices:    wantPrices,
			Claim:     false,
			AddReq:    false,
		},
	}
	opts := &CheckDomainsBulkOptions{Domains: []string{"test.xyz", "test.com"}}
	got, err := client.CheckDomainsBulk(opts)
	require.NoError(t, err)
	require.Equal(t, want, got)

	got, err = client.CheckDomainsBulk(nil)
	require.Error(t, err)
	require.Nil(t, got)
}

func TestGetDomainClaim(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"domains/test.com/claim", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "claim.json")
	})
	// want := &TLDNotices{TLD: "APP", Notices: []string{"Google notice"}}
	want := []*Claim{{
		TradeMark:               "string",
		Jurisdiction:            "string",
		JurisdictionCountryCode: "string",
		Goods:                   "string",
		RegistrantContact:       wantClaimContact,
		AgentContact:            wantClaimContact,
		Description:             []string{"string"},
	}}

	got, err := client.GetDomainClaim("test.com")
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestGetDomainStatusCodes(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"domains/test.com/status_codes", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "status_codes.json")
	})
	want := []string{"clientTransferProhibited", "clientUpdateProhibited"}

	got, err := client.GetDomainStatusCodes("test.com")
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestEnableAutoRenew(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"domains/test.com/auto_renew", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)
		writeFixture(t, w, "enable_auto_renew.json")
	})

	want := &AutoRenew{
		Domain:    "test.com",
		AutoRenew: true,
	}

	got, err := client.EnableDomainAutoRenew("test.com")
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestDisableAutoRenew(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"domains/test.com/auto_renew", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		writeFixture(t, w, "disable_auto_renew.json")
	})

	want := &AutoRenew{
		Domain:    "test.com",
		AutoRenew: false,
	}

	got, err := client.DisableDomainAutoRenew("test.com")
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestRenewDomain(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	opts := &RenewDomainOptions{
		Period: "1",
	}
	mux.HandleFunc(apiVerPath+"domains/test.com/renew", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)
		want, _ := json.Marshal(opts)
		require.Equal(t, string(want), getBody(t, r))
		writeFixture(t, w, "renew.json")
	})

	want := &Renew{
		Domain:            "test.com",
		NewExpirationDate: &PnTime{wantDate},
	}

	got, err := client.RenewDomain("test.com", opts)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestRedeemDomain(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"domains/test.com/redeem", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)
		writeFixture(t, w, "renew.json")
	})

	want := &Redeem{
		Domain:            "test.com",
		NewExpirationDate: &PnTime{wantDate},
	}

	got, err := client.RedeemDomain("test.com")
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestResendDomainEmail(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"domains/test.com/resend", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)
	})

	err := client.ResendDomainEmail("test.com")
	require.NoError(t, err)
}
