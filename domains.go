package pananames

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Represents a domain info
type Domain struct {
	Domain           string             `json:"domain"`
	DomainIDN        string             `json:"domain_idn"`
	Premium          bool               `json:"premium"`
	AutoRenew        bool               `json:"auto_renew"`
	WhoisPrivacy     bool               `json:"whois_privacy"`
	LockStatus       string             `json:"lock_status"`
	RegistrationDate *PnTime            `json:"registration_date"`
	ExpirationDate   *PnTime            `json:"expiration_date"`
	DeletionDate     *PnDate            `json:"deletion_date"`
	Status           string             `json:"status"`
	NameServers      *NameServers       `json:"name_servers"`
	ChildNameServers []*ChildNameServer `json:"child_name_servers"`
}

// Represents a domain check info
type DomainCheck struct {
	Domain                string             `json:"domain"`
	DomainIDN             string             `json:"domain_idn"`
	Available             bool               `json:"available"`
	Premium               bool               `json:"premium"`
	Prices                *Prices            `json:"prices"`
	PromoPrices           *Prices            `json:"promo_prices"`
	PromoTwoYearsPrices   *Prices            `json:"promo_two_years_prices"`
	PromoMultiYearsPrices map[string]*Prices `json:"promo_multi_years_prices"`
	Claim                 bool               `json:"claim"`
	AddReq                bool               `json:"add_req"`
}

// Represents prices info
type Prices struct {
	Currency string  `json:"currency"`
	Register float64 `json:"register"`
	Renew    float64 `json:"renew"`
	Transfer float64 `json:"transfer"`
	Redeem   float64 `json:"redeem"`
}

// Represents a contact info
type Contact struct {
	Org     string   `json:"org,omitempty"`
	Name    string   `json:"name,omitempty"`
	Email   string   `json:"email,omitempty"`
	Address string   `json:"address,omitempty"`
	City    string   `json:"city,omitempty"`
	State   string   `json:"state,omitempty"`
	Zip     string   `json:"zip,omitempty"`
	Country string   `json:"country,omitempty"`
	Phone   string   `json:"phone,omitempty"`
	Extras  []string `json:"extras,omitempty"`
}

// Represents a claim contact info
type ClaimContact struct {
	Name         string `json:"name,omitempty"`
	Email        string `json:"email,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
	Zip          string `json:"zip,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Organization string `json:"organization,omitempty"`
	Street       string `json:"street,omitempty"`
	CountryCode  string `json:"country_code,omitempty"`
}

// Represents a claim info
type Claim struct {
	TradeMark               string        `json:"trade_mark"`
	Jurisdiction            string        `json:"jurisdiction"`
	JurisdictionCountryCode string        `json:"jurisdiction_country_code"`
	Goods                   string        `json:"goods"`
	RegistrantContact       *ClaimContact `json:"registrant_contact"`
	AgentContact            *ClaimContact `json:"agent_contact"`
	Description             []string      `json:"description"`
}

// Represents an auto renew info
type AutoRenew struct {
	Domain    string `json:"domain"`
	AutoRenew bool   `json:"auto_renew"`
}

// Represents a domain renew info
type Renew struct {
	Domain            string  `json:"domain"`
	NewExpirationDate *PnTime `json:"new_expiration_date"`
}

// Represents a domain redeem info
type Redeem Renew

// Available options for GetDomains()
type GetDomainsOptions struct {
	ListOptions
	DomainLike string `url:"domain_like,omitempty"`
	Status     string `url:"status,omitempty"`
	LockStatus string `url:"lock_status,omitempty"`
}

// Available options for CheckDomainsBulk()
type CheckDomainsBulkOptions struct {
	Domains []string `url:"domains,omitempty"`
}

// helper struct for CheckDomainsBulk() to convert slice into domains string
type checkDomainsBulkOptions struct {
	Domains string `url:"domains,omitempty"`
}

// Available options for RegisterDomain()
type RegisterDomainOptions struct {
	Domain            string   `json:"domain,omitempty"`
	Period            int      `json:"period,omitempty"`
	WhoisPrivacy      bool     `json:"whois_privacy"`
	PremiumPrice      float64  `json:"premium_price,omitempty"`
	ClaimsAccepted    bool     `json:"claims_accepted,omitempty"`
	AddReqAccepted    bool     `json:"add_req_accepted,omitempty"`
	RegistrantContact *Contact `json:"registrant_contact,omitempty"`
	AdminContact      *Contact `json:"admin_contact,omitempty"`
	TechContact       *Contact `json:"tech_contact,omitempty"`
	BillingContact    *Contact `json:"billing_contact,omitempty"`
}

// Available options for RenewDomain()
type RenewDomainOptions struct {
	Period       string  `json:"period,omitempty"`
	PremiumPrice float64 `json:"premium_price,omitempty"`
}

// Get paged list of domains available in your account
func (c *Client) GetDomains(opt *GetDomainsOptions, options ...RequestOptionFunc) ([]*Domain, *Pagination, error) {
	u := "domains"
	req, err := c.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var result []*Domain
	resp, err := c.Do(req, &result)
	if err != nil {
		return nil, nil, err
	}

	return result, &resp.Meta.Pagination, nil
}

// Register a domain name
// The premium price can be fetched via the CheckDomain() method
func (c *Client) RegisterDomain(opt *RegisterDomainOptions, options ...RequestOptionFunc) (*Domain, error) {
	u := "domains"
	req, err := c.NewRequest(http.MethodPost, u, opt, options)
	if err != nil {
		return nil, err
	}
	result := new(Domain)
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Get information about the domain
func (c *Client) GetDomain(domain string, options ...RequestOptionFunc) (*Domain, error) {
	u := fmt.Sprintf("domains/%s", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, err
	}

	result := new(Domain)
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Delete domain
func (c *Client) DeleteDomain(domain string, options ...RequestOptionFunc) error {
	u := fmt.Sprintf("domains/%s", domain)
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

// Check domain availability and pricing
// This method provides crucial information needed for registration of a domain, as well as domain renewal, transfer and redemption costs
func (c *Client) CheckDomain(domain string, options ...RequestOptionFunc) (*DomainCheck, error) {
	u := fmt.Sprintf("domains/%s/check", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, err
	}

	result := new(DomainCheck)
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Bulk check the domains availability
// Get information about the domains availability, prices and claim
func (c *Client) CheckDomainsBulk(opt *CheckDomainsBulkOptions, options ...RequestOptionFunc) ([]*DomainCheck, error) {
	u := "domains/bulk_check"
	opts := &checkDomainsBulkOptions{}
	if opt != nil && len(opt.Domains) > 0 {
		opts.Domains = strings.Join(opt.Domains, ",")
	} else {
		return nil, fmt.Errorf("%T can't be nil", opt)
	}

	req, err := c.NewRequest(http.MethodGet, u, opts, options)
	if err != nil {
		return nil, err
	}

	var result []*DomainCheck
	_, err = c.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Get claim information for the domain
func (c *Client) GetDomainClaim(domain string, options ...RequestOptionFunc) ([]*Claim, error) {
	u := fmt.Sprintf("domains/%s/claim", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, err
	}

	var result []*Claim
	_, err = c.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Get list of status codes set for the domain
// EPP Status Code with description: https://www.icann.org/resources/pages/epp-status-codes-2014-06-16-en
func (c *Client) GetDomainStatusCodes(domain string, options ...RequestOptionFunc) ([]string, error) {
	u := fmt.Sprintf("domains/%s/status_codes", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, err
	}

	result := []string{}
	_, err = c.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Enable auto renew of the domain
func (c *Client) EnableDomainAutoRenew(domain string, options ...RequestOptionFunc) (*AutoRenew, error) {
	u := fmt.Sprintf("domains/%s/auto_renew", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodPut, u, nil, options)
	if err != nil {
		return nil, err
	}
	result := new(AutoRenew)
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Disable auto renew of the domain
func (c *Client) DisableDomainAutoRenew(domain string, options ...RequestOptionFunc) (*AutoRenew, error) {
	u := fmt.Sprintf("domains/%s/auto_renew", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodDelete, u, nil, options)
	if err != nil {
		return nil, err
	}
	result := new(AutoRenew)
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Renew tne domain. The domain may be renewed only for a period 1 to 10 years
// An information about minimum and maximum registration period is available via GetTLDs() method
func (c *Client) RenewDomain(domain string, opt *RenewDomainOptions, options ...RequestOptionFunc) (*Renew, error) {
	u := fmt.Sprintf("domains/%s/renew", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodPut, u, opt, options)
	if err != nil {
		return nil, err
	}
	result := new(Renew)
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Restore domain name during Redemption Grace Period
func (c *Client) RedeemDomain(domain string, options ...RequestOptionFunc) (*Redeem, error) {
	u := fmt.Sprintf("domains/%s/redeem", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodPut, u, nil, options)
	if err != nil {
		return nil, err
	}
	result := new(Redeem)
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Resend verification email
func (c *Client) ResendDomainEmail(domain string, options ...RequestOptionFunc) error {
	u := fmt.Sprintf("domains/%s/resend", url.PathEscape(domain))
	req, err := c.NewRequest(http.MethodPut, u, nil, options)
	if err != nil {
		return err
	}
	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
