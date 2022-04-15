package pananames

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

var wantNameServerRecord = &NameServerRecord{
	ID:       "string",
	Name:     "string",
	Type:     "string",
	Value:    "string",
	Priority: 1,
	TTL:      1,
}

var wantDNSSec = &DNSSec{
	Domain:  "test.com",
	DSData:  "string",
	Enabled: true,
}

func TestGetNameServers(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"domains/test.com/name_servers", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "nameservers.json")
	})

	want := &NameServers{"ns1.test.com", "ns2.test.com"}
	got, err := client.GetNameServers("test.com")
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestSetNameServers(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	opts := &SetNameServersOptions{
		NameServers: []string{"ns1.test.com", "ns2.test.com"},
	}
	mux.HandleFunc(apiVerPath+"domains/test.com/name_servers", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)
		want, _ := json.Marshal(opts)
		require.Equal(t, string(want), getBody(t, r))
		writeFixture(t, w, "nameservers.json")
	})

	want := &NameServers{"ns1.test.com", "ns2.test.com"}
	got, err := client.SetNameServers("test.com", opts)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestDeleteNameServers(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"domains/test.com/name_servers", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
	})

	err := client.DeleteNameServers("test.com")
	require.NoError(t, err)
}

func TestGetChildNameServers(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"domains/test.com/child_name_servers", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "child_nameservers.json")
	})

	want := []*ChildNameServer{
		{Hostname: "ns1.test.com", IPv4: "string", IPv6: "string"},
		{Hostname: "ns2.test.com", IPv4: "string", IPv6: "string"},
	}
	got, err := client.GetChildNameServers("test.com")
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestUpdateChildNameServer(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	opts := &ChildNameServerOptions{
		Hostname: "ns1.test.com",
		IPv4:     "string",
		IPv6:     "string",
	}
	mux.HandleFunc(apiVerPath+"domains/test.com/child_name_servers", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)
		want, _ := json.Marshal(opts)
		require.Equal(t, string(want), getBody(t, r))
		writeFixture(t, w, "child_nameserver.json")
	})

	want := &ChildNameServer{
		Hostname: "ns1.test.com",
		IPv4:     "string",
		IPv6:     "string",
	}
	got, err := client.UpdateChildNameServer("test.com", opts)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestAddChildNameServer(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	opts := &ChildNameServerOptions{
		Hostname: "ns1.test.com",
		IPv4:     "string",
		IPv6:     "string",
	}
	mux.HandleFunc(apiVerPath+"domains/test.com/child_name_servers", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		want, _ := json.Marshal(opts)
		require.Equal(t, string(want), getBody(t, r))
		writeFixture(t, w, "child_nameserver.json")
	})

	want := &ChildNameServer{
		Hostname: "ns1.test.com",
		IPv4:     "string",
		IPv6:     "string",
	}
	got, err := client.AddChildNameServer("test.com", opts)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestDeleteChildNameServer(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	opts := &DeleteChildNameServerOptions{Hostname: "ns1.test.com"}
	mux.HandleFunc(apiVerPath+"domains/test.com/child_name_servers", func(w http.ResponseWriter, r *http.Request) {
		want, _ := json.Marshal(opts)
		require.Equal(t, string(want), getBody(t, r))
		require.Equal(t, http.MethodDelete, r.Method)
	})

	err := client.DeleteChildNameServer("test.com", opts)
	require.NoError(t, err)
}

func TestGetNameServerRecords(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"domains/test.com/name_server_records", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "name_server_records.json")
	})

	want := []*NameServerRecord{wantNameServerRecord, wantNameServerRecord}
	got, err := client.GetNameServerRecords("test.com")
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestAddNameServerRecord(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	opts := wantNameServerRecord
	mux.HandleFunc(apiVerPath+"domains/test.com/name_server_records", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		want, _ := json.Marshal(opts)
		require.Equal(t, string(want), getBody(t, r))
		writeFixture(t, w, "name_server_record.json")
	})

	want := wantNameServerRecord
	got, err := client.AddNameServerRecord("test.com", opts)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestUpdateNameServerRecord(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	opts := wantNameServerRecord
	mux.HandleFunc(apiVerPath+"domains/test.com/name_server_records", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)
		want, _ := json.Marshal(opts)
		require.Equal(t, string(want), getBody(t, r))
		writeFixture(t, w, "name_server_record.json")
	})

	want := wantNameServerRecord

	got, err := client.UpdateNameServerRecord("test.com", opts)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestDeleteNameServerRecord(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	opts := &DeleteNameServerRecordsOptions{ID: "1"}
	mux.HandleFunc(apiVerPath+"domains/test.com/name_server_records", func(w http.ResponseWriter, r *http.Request) {
		want, _ := json.Marshal(opts)
		require.Equal(t, string(want), getBody(t, r))
		require.Equal(t, http.MethodDelete, r.Method)
	})

	err := client.DeleteNameServerRecord("test.com", opts)
	require.NoError(t, err)
}

func TestSetBulkNameServerRecords(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	opts := []*NameServerRecord{wantNameServerRecord, wantNameServerRecord}
	mux.HandleFunc(apiVerPath+"domains/test.com/bulk_name_server_records", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)
		want, _ := json.Marshal(&opts)
		require.Equal(t, string(want), getBody(t, r))
		writeFixture(t, w, "name_server_records.json")
	})

	want := []*NameServerRecord{wantNameServerRecord, wantNameServerRecord}

	got, err := client.SetBulkNameServerRecords("test.com", opts)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestUpdateBulkNameServerRecords(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	opts := []*NameServerRecord{wantNameServerRecord, wantNameServerRecord}
	mux.HandleFunc(apiVerPath+"domains/test.com/bulk_name_server_records", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)
		want, _ := json.Marshal(&opts)
		require.Equal(t, string(want), getBody(t, r))
		writeFixture(t, w, "name_server_records.json")
	})

	want := []*NameServerRecord{wantNameServerRecord, wantNameServerRecord}

	got, err := client.UpdateBulkNameServerRecords("test.com", opts)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestGetDNSSec(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"domains/test.com/dnssec", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		writeFixture(t, w, "dnssec.json")
	})

	want := wantDNSSec
	got, err := client.GetDNSSec("test.com")
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestEnableDNSSec(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	opts := &EnableDNSSecOptions{
		DS: "string",
	}
	mux.HandleFunc(apiVerPath+"domains/test.com/dnssec", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPut, r.Method)
		want, _ := json.Marshal(opts)
		require.Equal(t, string(want), getBody(t, r))
		writeFixture(t, w, "dnssec.json")
	})

	want := wantDNSSec
	got, err := client.EnableDNSSec("test.com", opts)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestDisableDNSSec(t *testing.T) {
	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc(apiVerPath+"domains/test.com/dnssec", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodDelete, r.Method)
		writeFixture(t, w, "dnssec.json")
	})

	want := wantDNSSec
	got, err := client.DisableDNSSec("test.com")
	require.NoError(t, err)
	require.Equal(t, want, got)
}
