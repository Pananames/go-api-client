package pananames

import (
	"fmt"
	"net/http"
	"net/url"
)

// Represents a transfer in info
type TransferIn struct {
	Domain            string              `json:"domain"`
	TransferStatus    string              `json:"transfer_status"`
	InitDate          *PnTime             `json:"init_date"`
	PremiumPrice      float64             `json:"premium_price"`
	WhoisPrivacy      bool                `json:"whois_privacy"`
	RegistrantContact *Contact            `json:"registrant_contact"`
	AdminContact      *Contact            `json:"admin_contact"`
	TechContact       *Contact            `json:"tech_contact"`
	BillingContact    *Contact            `json:"billing_contact"`
	NameServers       *NameServers        `json:"name_servers"`
	NameServerRecords []*NameServerRecord `json:"name_server_records"`
}

// Available options for InitTransferIn()
type InitTransferInOptions struct {
	Domain            string              `json:"domain,omitempty"`
	AuthCode          string              `json:"auth_code,omitempty"`
	PremiumPrice      float64             `json:"premium_price,omitempty"`
	WhoisPrivacy      bool                `json:"whois_privacy,omitempty"`
	RegistrantContact *Contact            `json:"registrant_contact,omitempty"`
	AdminContact      *Contact            `json:"admin_contact,omitempty"`
	TechContact       *Contact            `json:"tech_contact,omitempty"`
	BillingContact    *Contact            `json:"billing_contact,omitempty"`
	NameServers       *NameServers        `json:"name_servers,omitempty"`
	NameServerRecords []*NameServerRecord `json:"name_servers_records,omitempty"`
}

// Available options for GetTransfersIn()
type GetTransfersInOptions struct {
	ListOptions
	DomainLike string `url:"domain_like,omitempty"`
	Status     string `url:"status,omitempty"`
}

// Available options for CancelTransferIn()
type CancelTransferInOptions struct {
	Domain string `json:"domain,omitempty"`
}

// Get paged list of active transfers in
func (c *Client) GetTransfersIn(opt *GetTransfersInOptions, options ...RequestOptionFunc) ([]*TransferIn, *Pagination, error) {
	u := "transfers_in"
	req, err := c.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var result []*TransferIn
	resp, err := c.Do(req, &result)
	if err != nil {
		return nil, nil, err
	}

	return result, &resp.Meta.Pagination, nil
}

// Initiate transfer in process for domain
// You should provide correct WHOIS information
func (c *Client) InitTransferIn(opt *InitTransferInOptions, options ...RequestOptionFunc) (*TransferIn, error) {
	u := "transfers_in"
	req, err := c.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, err
	}

	result := new(TransferIn)
	_, err = c.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Cancel transfer in process for domain
func (c *Client) CancelTransferIn(opt *CancelTransferInOptions, options ...RequestOptionFunc) error {
	u := "transfers_in"
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

// Prepare a domain for transferring out
// This will unlock a domain and send the authorization code to the domain’s registrant email
func (c *Client) InitTransferOut(domain string, options ...RequestOptionFunc) error {
	u := fmt.Sprintf("domains/%s/transfer_out", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodPut, u, nil, options)
	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return err
}

// Cancel transfer out process for domain
// Domain will be locked again
func (c *Client) CancelTransferOut(domain string, options ...RequestOptionFunc) error {
	u := fmt.Sprintf("domains/%s/transfer_out", url.PathEscape(domain))
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
