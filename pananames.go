package pananames

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	baseURL    = "https://api.pananames.com"
	apiVerPath = "/merchant/v2/"
	userAgent  = "go-pananames"
)

// Custom type for DateTime data, inherits time.Time
// Needs to avoid JSON unmarshall error when DateTime in JSON returns as empty string ""
type PnTime struct {
	time.Time
}

type PnDate struct {
	time.Time
}

// Represents api client
type Client struct {
	httpClient *http.Client
	baseURL    *url.URL
	token      string
	userAgent  string
}

// Represents api response
type Response struct {
	Data json.RawMessage `json:"data"`
	Meta Meta            `json:"meta"`
}

// Represents meta field from response
type Meta struct {
	Pagination
	Notice string `json:"notice"`
}

// Represents pagination info
type Pagination struct {
	Total int `json:"total_entries"`
	Limit int `json:"per_page"`
	Page  int `json:"current_page"`
	Pages int `json:"total_pages"`
}

// Represents error api response with Response struct
type ErrorResponse struct {
	Response *http.Response
	Errors   []struct {
		Code        int    `json:"code"`
		Message     string `json:"message"`
		Description string `json:"description"`
	} `json:"errors"`
}

// Represents pagination options
type ListOptions struct {
	Limit int `url:"per_page,omitempty"`
	Page  int `url:"current_page,omitempty"`
}

// Represents option func to customize api client
type Option func(*Client) error

// Represents option func to customize API request
type RequestOptionFunc func(*http.Request) error

// Custom Unmarshall for PnTime
func (t *PnTime) UnmarshalJSON(data []byte) error {
	if string(data) == `""` {
		return nil
	}
	return t.Time.UnmarshalJSON(data)
}

func (d *PnDate) UnmarshalJSON(data []byte) error {
	val := strings.Trim(string(data), `"`)
	if val == "" {
		return nil
	}

	var err error
	if d.Time, err = time.Parse(time.DateOnly, val); err != nil {
		return err
	}

	return nil
}

func (e *ErrorResponse) Error() string {
	path, _ := url.QueryUnescape(e.Response.Request.URL.Path)
	u := fmt.Sprintf("%s://%s%s", e.Response.Request.URL.Scheme, e.Response.Request.URL.Host, path)
	var errs []string
	for _, err := range e.Errors {
		errs = append(errs, fmt.Sprintf("Error code: %d, Message: '%s', Description: '%s'", err.Code, err.Message, err.Description))
	}
	return fmt.Sprintf("%s %s: %d:\n%s", e.Response.Request.Method, u, e.Response.StatusCode, strings.Join(errs, "\n"))
}

// Method returns next page if available
func (p *Pagination) NextPage() int {
	if p.Pages-p.Page > 0 {
		p.Page++
		return p.Page
	}
	return 0
}

// Method returns previous page if available
func (p *Pagination) PreviousPage() int {
	if p.Page > 1 {
		p.Page--
		return p.Page
	}
	return 0
}

// Run API request with provided context
func WithContext(ctx context.Context) RequestOptionFunc {
	return func(req *http.Request) error {
		*req = *req.WithContext(ctx)
		return nil
	}
}

// WithBaseURL Set BaseURL for api client
func WithBaseURL(baseURL string) Option {
	return func(c *Client) error {
		return c.setBaseURL(baseURL)
	}
}

// WithHTTPClient Set configured http.Client for api client
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) error {
		c.httpClient = httpClient

		return nil
	}
}

// WithUserAgent Set user agent for api client
func WithUserAgent(userAgent string) Option {
	return func(c *Client) error {
		c.userAgent = userAgent

		return nil
	}
}

// Set BaseURL, validate it and add api path to it
func (c *Client) setBaseURL(urlStr string) error {
	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}
	baseURL.Path = apiVerPath
	c.baseURL = baseURL
	return nil
}

// Exec all passed option func to the client
func (c *Client) parseOptions(opts ...Option) error {
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return err
		}
	}
	return nil
}

// Creates a new instance of api client
func NewClient(token string, opts ...Option) (*Client, error) {
	c := &Client{
		userAgent:  userAgent,
		token:      token,
		httpClient: &http.Client{Timeout: time.Second * 30},
	}
	_ = c.setBaseURL(baseURL)
	if err := c.parseOptions(opts...); err != nil {
		return nil, err
	}
	return c, nil
}

// Make an http request, check and parse response
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	// Parse data field from response
	if v != nil {
		result := &Response{}
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return nil, fmt.Errorf("status: %d, unable to decode response, unknown format: %v", resp.StatusCode, err)
		}
		// Decode the data field
		if result.Data == nil {
			return nil, fmt.Errorf("status: %d, missing data from response", resp.StatusCode)
		}
		if err := json.Unmarshal(result.Data, v); err != nil {
			return result, fmt.Errorf("status: %d, unable to parse response data: %s", resp.StatusCode, err)
		}
		fixZeroDate(v)
		return result, nil
	}

	return nil, err
}

// Creates and validates a new request
// sets required headers
func (c *Client) NewRequest(method, path string, opt interface{}, options []RequestOptionFunc) (*http.Request, error) {
	u := *c.baseURL
	u.Path = c.baseURL.Path + path

	// Prepare headers
	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")
	reqHeaders.Set("SIGNATURE", c.token)
	reqHeaders.Set("User-Agent", c.userAgent)

	// Validate and marshall request body if any
	var body []byte
	if method == http.MethodPost || method == http.MethodPut || method == http.MethodDelete {
		reqHeaders.Set("Content-Type", "application/json")
		if opt != nil {
			var err error
			if body, err = json.Marshal(opt); err != nil {
				return nil, err
			}
		}
	} else {
		if opt != nil {
			q, err := query.Values(opt)
			if err != nil {
				return nil, err
			}
			u.RawQuery = q.Encode()
		}
	}
	// Create a new request
	req, err := http.NewRequest(method, u.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	// Set request option funcs
	for _, fn := range options {
		if fn == nil {
			continue
		}
		if err := fn(req); err != nil {
			return nil, err
		}
	}

	// Set the request specific headers
	for k, v := range reqHeaders {
		req.Header[k] = v
	}

	return req, nil
}

// Checks the API response for errors
func CheckResponse(r *http.Response) error {
	switch r.StatusCode {
	case 200, 201, 202, 204, 304:
		return nil
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if data != nil {
		errorResponse := &ErrorResponse{Response: r}
		if err := json.Unmarshal(data, errorResponse); err != nil {
			return fmt.Errorf("status: %d, can't parse error, unknown format, raw data: %s", r.StatusCode, data)
		}
		return errorResponse
	}
	return fmt.Errorf("status: %d, empty response", r.StatusCode)
}

// Parse value, find zero *PnTime and set it to nil to avoid misleading
func fixZeroDate(value interface{}) {
	v := reflect.ValueOf(value)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			zeroPnTime(v.Index(i))
		}
	case reflect.Map:
		for _, k := range v.MapKeys() {
			zeroPnTime(v.MapIndex(k))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			zeroPnTime(v.Field(i))
		}
	}
}

// Set &PnTime or &PnDate to nil if it zero PnTime or PnDate
// If not, continue parse with fixZeroDate()
func zeroPnTime(v reflect.Value) {
	if v.IsZero() {
		return
	}
	if v.Type() == reflect.TypeOf(&PnTime{}) {
		if t, ok := v.Elem().Interface().(PnTime); ok {
			if t.IsZero() {
				v.Set(reflect.Zero(v.Type()))
			}
		}
	} else if v.Type() == reflect.TypeOf(&PnDate{}) {
		if t, ok := v.Elem().Interface().(PnDate); ok {
			if t.IsZero() {
				v.Set(reflect.Zero(v.Type()))
			}
		}
	} else {
		fixZeroDate(v.Interface())
	}
}
