package pananames

import (
	"fmt"
	"net/http"
	"net/url"
)

// Represents a redirect info
type Redirect struct {
	Url            string `json:"url,omitempty"`
	MaskingEnabled bool   `json:"masking_enabled,omitempty"`
	MaskingTitle   string `json:"masking_title,omitempty"`
	MaskingDesc    string `json:"masking_desc,omitempty"`
	MaskingKwd     string `json:"masking_kwd,omitempty"`
}

// Represents a redirect bulk info
type RedirectBulk struct {
	Redirect
	DomainList []*DomainRedirect `json:"domain_list"`
}

// Represents a domain redirect info
type DomainRedirect struct {
	Domain       string `json:"domain"`
	DomainQueued bool   `json:"domain_queued"`
	Error        string `json:"error"`
}

// Available options for EnableDomainRedirect()
type EnableDomainRedirectOptions Redirect

// Available option for EnableBulkDomainRedirect()
type EnableBulkDomainRedirectOptions struct {
	EnableDomainRedirectOptions
	DomainList []string `json:"domain_list,omitempty"`
}

// Get current redirect URL and mode for the domain
func (c *Client) GetDomainRedirect(domain string, options ...RequestOptionFunc) (*Redirect, error) {
	u := fmt.Sprintf("domains/%s/redirect", url.PathEscape(domain))

	req, err := c.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, err
	}

	result := new(Redirect)
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Enable or update redirect for the domain
func (c *Client) EnableDomainRedirect(domain string, opt *EnableDomainRedirectOptions, options ...RequestOptionFunc) (*Redirect, error) {
	u := fmt.Sprintf("domains/%s/redirect", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, err
	}
	result := new(Redirect)
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Disable redirect for the domain
func (c *Client) DisableDomainRedirect(domain string, options ...RequestOptionFunc) error {
	u := fmt.Sprintf("domains/%s/redirect", url.PathEscape(domain))
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

// Puts in the queue the task of redirect for an array of domains
// The report of the operation will be sent by email
func (c *Client) EnableBulkDomainRedirect(opt *EnableBulkDomainRedirectOptions, options ...RequestOptionFunc) (*RedirectBulk, error) {
	u := "domains/bulk_redirect"
	req, err := c.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, err
	}
	result := new(RedirectBulk)
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
