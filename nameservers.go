package pananames

import (
	"fmt"
	"net/http"
	"net/url"
)

// Represents a list of name servers
type NameServers []string

// Represents a name server record info
type NameServerRecord struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Value    string `json:"value"`
	Priority int    `json:"priority"`
	TTL      int    `json:"ttl"`
}

// Represents a DNSSec info
type DNSSec struct {
	Domain  string `json:"domain"`
	DSData  string `json:"ds_data"`
	Enabled bool   `json:"enabled"`
}

// Represents a child name server info
type ChildNameServer struct {
	Hostname string `json:"hostname"`
	IPv4     string `json:"ipv4,omitempty"`
	IPv6     string `json:"ipv6,omitempty"`
}

// Available options for SetNameServers()
type SetNameServersOptions struct {
	NameServers NameServers `json:"name_servers"`
}

// Available options for AddChildNameServer() and UpdateChildNameServer()
type ChildNameServerOptions ChildNameServer

// Available options for DeleteChildNameServer()
type DeleteChildNameServerOptions struct {
	Hostname string `json:"hostname,omitempty"`
}

// Available options for DeleteNameServerRecords()
type DeleteNameServerRecordsOptions struct {
	ID string `json:"id,omitempty"`
}

// Available options for UpdateDNSSec()
type EnableDNSSecOptions struct {
	DS string `json:"ds,omitempty"`
}

// Get name servers list for the domain
func (c *Client) GetNameServers(domain string, options ...RequestOptionFunc) (*NameServers, error) {
	u := fmt.Sprintf("domains/%s/name_servers", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, err
	}

	result := new(NameServers)
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Set name servers for the domain
func (c *Client) SetNameServers(domain string, opt *SetNameServersOptions, options ...RequestOptionFunc) (*NameServers, error) {
	u := fmt.Sprintf("domains/%s/name_servers", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, err
	}

	result := new(NameServers)
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Validate SetNameServersOptions for required options
func (opt *SetNameServersOptions) Validate() error {
	if opt == nil {
		return fmt.Errorf("%T can't be nil", opt)
	}
	return nil
}

// Delete name servers for the domain
func (c *Client) DeleteNameServers(domain string, options ...RequestOptionFunc) error {
	u := fmt.Sprintf("domains/%s/name_servers", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return err
}

// Get a list of child name servers for the domain
func (c *Client) GetChildNameServers(domain string, options ...RequestOptionFunc) ([]*ChildNameServer, error) {
	u := fmt.Sprintf("domains/%s/child_name_servers", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, err
	}

	var result []*ChildNameServer
	_, err = c.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Create a new child name server for the domain
func (c *Client) AddChildNameServer(domain string, opt *ChildNameServerOptions, options ...RequestOptionFunc) (*ChildNameServer, error) {
	u := fmt.Sprintf("domains/%s/child_name_servers", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, err
	}

	result := new(ChildNameServer)
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Update an existing child name server for the domain
func (c *Client) UpdateChildNameServer(domain string, opt *ChildNameServerOptions, options ...RequestOptionFunc) (*ChildNameServer, error) {
	u := fmt.Sprintf("domains/%s/child_name_servers", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, err
	}

	result := new(ChildNameServer)
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Delete a child name server for the domain by name
func (c *Client) DeleteChildNameServer(domain string, opt *DeleteChildNameServerOptions, options ...RequestOptionFunc) error {
	u := fmt.Sprintf("domains/%s/child_name_servers", url.PathEscape(domain))

	req, err := c.NewRequest(http.MethodDelete, u, opt, options)
	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return err
}

// Get name server records list for the domain
func (c *Client) GetNameServerRecords(domain string, options ...RequestOptionFunc) ([]*NameServerRecord, error) {
	u := fmt.Sprintf("domains/%s/name_server_records", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, err
	}

	var result []*NameServerRecord
	_, err = c.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Create a new name server record for the domain
func (c *Client) AddNameServerRecord(domain string, opt *NameServerRecord, options ...RequestOptionFunc) (*NameServerRecord, error) {
	u := fmt.Sprintf("domains/%s/name_server_records", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, err
	}

	result := new(NameServerRecord)
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Update an existing name server record for the domain
func (c *Client) UpdateNameServerRecord(domain string, opt *NameServerRecord, options ...RequestOptionFunc) (*NameServerRecord, error) {
	u := fmt.Sprintf("domains/%s/name_server_records", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, err
	}

	result := new(NameServerRecord)
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Delete a specific name server record for the domain
func (c *Client) DeleteNameServerRecord(domain string, opt *DeleteNameServerRecordsOptions, options ...RequestOptionFunc) error {
	u := fmt.Sprintf("domains/%s/name_server_records", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodDelete, u, opt, options)
	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return err
}

// Create a list of new name server records for the domain
func (c *Client) SetBulkNameServerRecords(domain string, opt []*NameServerRecord, options ...RequestOptionFunc) ([]*NameServerRecord, error) {
	u := fmt.Sprintf("domains/%s/bulk_name_server_records", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, err
	}

	var result []*NameServerRecord
	_, err = c.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Update list of existing name server records for the domain
func (c *Client) UpdateBulkNameServerRecords(domain string, opt []*NameServerRecord, options ...RequestOptionFunc) ([]*NameServerRecord, error) {
	u := fmt.Sprintf("domains/%s/bulk_name_server_records", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, err
	}

	var result []*NameServerRecord
	_, err = c.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Get DNSSec status for the domain
func (c *Client) GetDNSSec(domain string, options ...RequestOptionFunc) (*DNSSec, error) {
	u := fmt.Sprintf("domains/%s/dnssec", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, err
	}

	result := new(DNSSec)
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Enable DNSSec for the domain
func (c *Client) EnableDNSSec(domain string, opt *EnableDNSSecOptions, options ...RequestOptionFunc) (*DNSSec, error) {
	u := fmt.Sprintf("domains/%s/dnssec", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, err
	}

	result := new(DNSSec)
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Disable DNSSec for the domain
func (c *Client) DisableDNSSec(domain string, options ...RequestOptionFunc) (*DNSSec, error) {
	u := fmt.Sprintf("domains/%s/dnssec", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}

	result := new(DNSSec)
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
