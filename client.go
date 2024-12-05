package mailerlite

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	Version    = "0.0.2"
	APIVersion = "2023-18-04"

	defaultBaseURL   = "https://connect.mailerlite.com/api"
	defaultUserAgent = "go-mailerlite" + "/" + Version

	HeaderAPIVersion     = "X-Version"
	HeaderRateLimit      = "X-RateLimit-Limit"
	HeaderRateRemaining  = "X-RateLimit-Remaining"
	HeaderRateRetryAfter = "Retry-After"
)

// Client - base api client
type Client struct {
	client *http.Client // HTTP client used to communicate with the API.

	apiBase    *url.URL // apiBase the base used when communicating with the API.
	apiVersion string   // apiVersion the version used when communicating with the API.
	apiKey     string   // apiKey used when communicating with the API.

	userAgent string // userAgent User agent used when communicating with the API.

	rateMu     sync.Mutex // rateMu protects the rate during getting rate limits from client
	rateLimits Rate       // Rate limits for the client as determined by the most recent API calls.

	common service // common service

	Subscriber SubscriberService // Subscriber service
	Group      GroupService      // Group service
	Field      FieldService      // Field service
	Form       FormService       // Form service
	Segment    SegmentService    // Segment service
	Webhook    WebhookService    // Webhook service
	Campaign   CampaignService   // Campaign service
	Automation AutomationService // Automation service
	Timezone   TimezoneService   // Timezone service

}

type service struct {
	client *Client
}

// Response is a MailerLite API response. This wraps the standard http.Response
type Response struct {
	*http.Response

	// Explicitly specify the Rate type so Rate's String() receiver doesn't
	// propagate to Response.
	Rate Rate
}

// ErrorResponse is a MailerLite API error response. This wraps the standard http.Response
type ErrorResponse struct {
	Response *http.Response      // HTTP response that caused this error
	Message  string              `json:"message"` // error message
	Errors   map[string][]string `json:"errors"`
}

// Rate represents the rate limit for the current client.
type Rate struct {
	// The number of requests per minute the client is currently limited to.
	Limit int `json:"limit"`

	// The number of remaining requests the client can make this minute.
	Remaining int `json:"remaining"`

	// Retry After
	RetryAfter *time.Duration `json:"retry"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %+v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message, r.Errors)
}

// AuthError occurs when using HTTP Authentication fails
type AuthError ErrorResponse

func (r *AuthError) Error() string { return (*ErrorResponse)(r).Error() }

// NewClient - creates a new client instance.
func NewClient(apiKey string) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)

	client := &Client{
		apiBase:   baseURL,
		apiKey:    apiKey,
		userAgent: defaultUserAgent,
		client:    http.DefaultClient,
	}

	client.common.client = client
	client.Subscriber = &subscriberService{&client.common}
	client.Group = &groupService{&client.common}
	client.Field = &fieldService{&client.common}
	client.Form = &formService{&client.common}
	client.Segment = &segmentService{&client.common}
	client.Webhook = &webhookService{&client.common}
	client.Campaign = &campaignService{&client.common}
	client.Automation = &automationService{&client.common}
	client.Timezone = &timezoneService{&client.common}

	return client
}

// APIKey - Get api key after it has been created
func (c *Client) APIKey() string {
	return c.apiKey
}

// Client - Get the current client
func (c *Client) Client() *http.Client {
	return c.client
}

// SetHttpClient - Set the client if you want more control over the client implementation
func (c *Client) SetHttpClient(client *http.Client) {
	c.client = client
}

// SetAPIKey - Set the client api key
func (c *Client) SetAPIKey(apikey string) {
	c.apiKey = apikey
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	reqURL := fmt.Sprintf("%s%s", c.apiBase, path)
	reqBodyBytes := new(bytes.Buffer)

	if method == http.MethodPost ||
		method == http.MethodPut ||
		method == http.MethodDelete {
		err := json.NewEncoder(reqBodyBytes).Encode(body)
		if err != nil {
			return nil, err
		}
	} else if method == http.MethodGet {
		reqURL, _ = addOptions(reqURL, body)
	}

	req, err := http.NewRequest(method, reqURL, reqBodyBytes)
	if err != nil {
		return nil, err
	}

	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}

	req.Header.Set(HeaderAPIVersion, APIVersion)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Accept", "application/json")

	return req, nil
}

func (c *Client) do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	req = req.WithContext(ctx)

	// If we've hit rate limit, don't make further requests before Reset time.
	if err := c.checkRateLimitBeforeDo(req); err != nil {
		return &Response{
			Response: err.Response,
			Rate:     err.Rate,
		}, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}

	response := newResponse(resp)

	c.rateMu.Lock()
	c.rateLimits = response.Rate
	c.rateMu.Unlock()

	err = checkResponse(resp)
	if err != nil {
		defer resp.Body.Close()
		return response, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
		if err != nil {
			return nil, err
		}
	}

	return response, err
}

// checkRateLimitBeforeDo does not make any network calls, but uses existing knowledge from
// current client state in order to quickly check if *RateLimitError can be immediately returned
// from Client.do, and if so, returns it so that Client.do can skip making a network API call unnecessarily.
// Otherwise it returns nil, and Client.do should proceed normally.
func (c *Client) checkRateLimitBeforeDo(req *http.Request) *RateLimitError {
	c.rateMu.Lock()
	rate := c.rateLimits
	c.rateMu.Unlock()
	if rate.Remaining == 0 && rate.RetryAfter != nil && time.Now().Before(time.Now().Add(*rate.RetryAfter)) {
		// Create a fake response.
		resp := &http.Response{
			Status:     http.StatusText(http.StatusForbidden),
			StatusCode: http.StatusForbidden,
			Request:    req,
			Header:     make(http.Header),
			Body:       io.NopCloser(strings.NewReader("")),
		}
		return &RateLimitError{
			Rate:     rate,
			Response: resp,
			Message:  fmt.Sprintf("API rate limit of %v still exceeded until %v, not making remote request.", rate.Limit, rate.RetryAfter),
		}
	}

	return nil
}

// newResponse creates a new Response for the provided http.Response.
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	response.Rate = parseRate(r)
	return response
}

func checkResponse(r *http.Response) error {
	if r.StatusCode == http.StatusAccepted {
		return nil
	}

	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := io.ReadAll(r.Body)

	if err == nil && len(data) > 0 {
		err := json.Unmarshal(data, errorResponse)
		if err != nil {
			errorResponse.Message = string(data)
		}
	}

	switch {
	case r.StatusCode == http.StatusUnauthorized:
		return (*AuthError)(errorResponse)
	case r.StatusCode == http.StatusTooManyRequests && r.Header.Get(HeaderRateRemaining) == "0":
		return &RateLimitError{
			Rate:     parseRate(r),
			Response: errorResponse.Response,
			Message:  errorResponse.Message,
		}
	default:
		return errorResponse
	}
}

func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)

	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	origURL, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	origValues := origURL.Query()

	newValues, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	for k, v := range newValues {
		if k == "Filters" {
			for _, fv := range v {
				if fv == "" {
					continue
				}
				split := strings.Fields(strings.Trim(fv, "{}"))
				filterKey := fmt.Sprintf("filter[%s]", split[0])
				origValues.Add(filterKey, split[1])
			}
			continue
		} else {
			for _, fv := range v {
				if fv == "" {
					continue
				}
				origValues.Add(k, fv)
			}
		}
	}

	origURL.RawQuery = origValues.Encode()

	return origURL.String(), nil
}

func parseRate(r *http.Response) Rate {
	var rate Rate
	if limit := r.Header.Get(HeaderRateLimit); limit != "" {
		rate.Limit, _ = strconv.Atoi(limit)
	}
	if remaining := r.Header.Get(HeaderRateRemaining); remaining != "" {
		rate.Remaining, _ = strconv.Atoi(remaining)
	}

	if retry := r.Header.Get(HeaderRateRetryAfter); retry != "" {
		// The "Retry-After" header value will be
		// an integer which represents the number of seconds that one should
		// wait before resuming making requests.
		retryAfterSeconds, _ := strconv.ParseInt(retry, 10, 64) // Error handling is noop.
		retryAfter := time.Duration(retryAfterSeconds) * time.Second
		rate.RetryAfter = &retryAfter
	}
	return rate
}

// RateLimitError occurs when MailerLite returns 403 Forbidden response with a rate limit
// remaining value of 0.
type RateLimitError struct {
	Rate     Rate           // Rate specifies last known rate limit for the client
	Response *http.Response // HTTP response that caused this error
	Message  string         `json:"message"` // error message
}

func (r *RateLimitError) Error() string {
	return fmt.Sprintf("%v %v: %d %v [retry after %v]",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message, r.Rate.RetryAfter)
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }
