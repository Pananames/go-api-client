package pananames

import "net/http"

// Represents a balance info
type Balance struct {
	Balance float64 `json:"balance"`
}

// Represents a payment info
type Payment struct {
	TxID       string  `json:"txid"`
	TxDate     *PnTime `json:"txdate"`
	TxType     string  `json:"txtype"`
	Domain     string  `json:"domain"`
	Period     string  `json:"period"`
	Descripton string  `json:"description"`
	Total      float64 `json:"total"`
}

// Available options for GetAccountPayments()
type GetAccountPaymentsOptions struct {
	ListOptions
	ID         int    `url:"id"`
	DomainLike string `url:"domain_like,omitempty"`
	PayType    string `url:"pay_type"`
	DateFrom   string `url:"date_from"`
	DateEnd    string `url:"date_end"`
}

// Get current balance
func (c *Client) GetAccountBalance(options ...RequestOptionFunc) (*Balance, error) {
	u := "account/balance"

	req, err := c.NewRequest(http.MethodGet, u, nil, options)
	if err != nil {
		return nil, err
	}

	result := new(Balance)
	_, err = c.Do(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Get paged list of payments from your account
func (c *Client) GetAccountPayments(opt *GetAccountPaymentsOptions, options ...RequestOptionFunc) ([]*Payment, *Pagination, error) {
	u := "account/payments"
	req, err := c.NewRequest(http.MethodGet, u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var result []*Payment
	resp, err := c.Do(req, &result)
	if err != nil {
		return nil, nil, err
	}

	return result, &resp.Meta.Pagination, nil
}
