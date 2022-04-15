package pananames

import (
	"fmt"
	"net/http"
	"net/url"
)

// Represents a whois info
type WhoisInfo struct {
	WhoisPrivacy      bool     `json:"whois_privacy"`
	Preview           bool     `json:"preview"`
	RegistrantContact *Contact `json:"registrant_contact"`
	AdminContact      *Contact `json:"admin_contact"`
	TechContact       *Contact `json:"tech_contact"`
	BillingContact    *Contact `json:"billing_contact"`
}

// Represents a whois privacy info
type WhoisPrivacy struct {
	Domain  string `json:"domain"`
	Enabled bool   `json:"enabled"`
}

// Available options for GetWhoisInfo()
type GetWhoisInfoOptions struct {
	Preview bool `url:"preview,omitempty"`
}

// Available options for UpdateWhoisInfo()
type UpdateWhoisInfoOptions struct {
	RegistrantContact *Contact `json:"registrant_contact,omitempty"`
	AdminContact      *Contact `json:"admin_contact,omitempty"`
	TechContact       *Contact `json:"tech_contact,omitempty"`
	BillingContact    *Contact `json:"billing_contact,omitempty"`
}

// Get WHOIS information for the domain. It works only for your domains
func (c *Client) GetWhoisInfo(domain string, opt *GetWhoisInfoOptions, options ...RequestOptionFunc) (*WhoisInfo, error) {
	u := fmt.Sprintf("domains/%s/whois", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, err
	}

	result := new(WhoisInfo)
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Update WHOIS information for the domain
// Return notice if confirmation is needed as second param
func (c *Client) UpdateWhoisInfo(domain string, opt *UpdateWhoisInfoOptions, options ...RequestOptionFunc) (*WhoisInfo, string, error) {
	u := fmt.Sprintf("domains/%s/whois", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, "", err
	}

	result := new(WhoisInfo)
	resp, err := c.Do(req, result)
	if err != nil {
		return nil, "", err
	}

	return result, resp.Meta.Notice, nil
}

// Get WHOIS privacy status for the domain
func (c *Client) GetWhoisPrivacy(domain string, options ...RequestOptionFunc) (*WhoisPrivacy, error) {
	u := fmt.Sprintf("domains/%s/whois_privacy", url.PathEscape(domain))

	req, err := c.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, err
	}

	result := new(WhoisPrivacy)
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Enable WHOIS privacy of the domain
func (c *Client) EnableWhoisPrivacy(domain string, options ...RequestOptionFunc) (*WhoisPrivacy, error) {
	u := fmt.Sprintf("domains/%s/whois_privacy", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodPut, u, nil, options)
	if err != nil {
		return nil, err
	}
	result := new(WhoisPrivacy)
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Disable WHOIS privacy of the domain
func (c *Client) DisableWhoisPrivacy(domain string, options ...RequestOptionFunc) (*WhoisPrivacy, error) {
	u := fmt.Sprintf("domains/%s/whois_privacy", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}
	result := new(WhoisPrivacy)
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
