package pananames

import (
	"fmt"
	"net/http"
	"net/url"
)

// Represents a TLD info
type TLD struct {
	TLD                   string                      `json:"tld"`
	IDN                   bool                        `json:"idn"`
	DNSSec                bool                        `json:"dnssec"`
	Prices                *Prices                     `json:"prices"`
	PromoPrices           *Prices                     `json:"promo_prices"`
	PromoUntil            *PnTime                     `json:"promo_untill"`
	PromoTwoYearsUntil    *PnTime                     `json:"promo_two_years_untill"`
	PromoTwoYearsPrices   *Prices                     `json:"promo_two_years_prices"`
	PromoMultiYearsPrices map[string]*PromoMultiYears `json:"promo_multi_years_prices"`
}

// Represents a promo for multi years prices
type PromoMultiYears struct {
	PromoMultiYearsPrices *Prices `json:"promo_multi_years_prices"`
	PromoMultiYearsUntil  *PnTime `json:"promo_multi_years_untill"`
}

// Represents a TLD notice info
type TLDNotice struct {
	TLD     string   `json:"tld"`
	Notices []string `json:"notices"`
}

// Represents a domain status info
type DomainStatus struct {
	Domain string `json:"domain"`
	Status string `json:"status"`
}

// Represents an Email info
type Email struct {
	Email          string          `json:"email"`
	FirstEmailDate *PnTime         `json:"first_email_date"`
	VerifyDate     *PnTime         `json:"verify_date"`
	SuspendDate    *PnTime         `json:"suspend_date"`
	Status         string          `json:"status"`
	Domains        []*DomainStatus `json:"domains"`
}

// Available GetEmails() options
type GetEmailsOptions struct {
	ListOptions
	EmailLike   string `url:"email_like,omitempty"`
	Status      string `url:"status,omitempty"`
	EmailStatus string `url:"email_status,omitempty"`
}

// Get Registration Notices for all TLDs
func (c *Client) GetTLDAddReqList(options ...RequestOptionFunc) ([]*TLDNotice, error) {
	u := "add_req_list"
	req, err := c.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, err
	}

	var result []*TLDNotice
	_, err = c.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Get full list of available TLDs
func (c *Client) GetTLDs(options ...RequestOptionFunc) ([]*TLD, error) {
	u := "tlds"

	req, err := c.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, err
	}

	var result []*TLD
	_, err = c.Do(req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Get account related emails
func (c *Client) GetEmails(opt *GetEmailsOptions, options ...RequestOptionFunc) ([]*Email, *Pagination, error) {
	u := "emails"

	req, err := c.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}
	var result []*Email
	resp, err := c.Do(req, &result)
	if err != nil {
		return nil, nil, err
	}

	return result, &resp.Meta.Pagination, nil
}

// Get Registration Notices for TLD
func (c *Client) GetTLDAddReq(tld string, options ...RequestOptionFunc) (*TLDNotice, error) {
	u := fmt.Sprintf("tlds/%s/add_req", url.PathEscape(tld))
	req, err := c.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, err
	}

	result := new(TLDNotice)
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
