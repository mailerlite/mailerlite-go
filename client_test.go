package mailerlite_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/mailerlite/mailerlite-go"
	"github.com/stretchr/testify/assert"
)

const (
	testKey = "valid-api-key"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

func TestPointerFunctions(t *testing.T) {
	assert.False(t, *mailerlite.Bool(false))
	assert.Equal(t, *mailerlite.Int(1), 1)
	assert.Equal(t, *mailerlite.Int64(1), int64(1))
	assert.Equal(t, *mailerlite.String("test"), string("test"))
}

func TestNewClient(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	assert.Equal(t, client.APIKey(), testKey)
	assert.Equal(t, client.Client(), http.DefaultClient)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
		}
	})

	client.SetHttpClient(testClient)
	assert.Equal(t, testClient, client.Client())

	client.SetAPIKey("valid-api-key-2")
}

func TestCanChangeAPIKey(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	assert.Equal(t, client.APIKey(), testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.Header.Get("Authorization"), "Bearer valid-api-key-2")
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
		}
	})

	client.SetHttpClient(testClient)

	client.SetAPIKey("valid-api-key-2")

	_, res, err := client.Timezone.List(context.TODO())
	if err != nil {
		return
	}
	assert.Equal(t, res.StatusCode, http.StatusOK)
}

func TestCanSetUserAgent(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.Header.Get("user-agent"), fmt.Sprintf("go-mailerlite/%v", mailerlite.Version))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
		}
	})

	client.SetHttpClient(testClient)

	languages, res, err := client.Campaign.Languages(context.TODO())
	if err != nil {
		return
	}

	assert.Equal(t, res.StatusCode, http.StatusOK)
	assert.NotEmpty(t, languages.Data)
}
func TestCanMakeApiCall(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.Header.Get("user-agent"), fmt.Sprintf("go-mailerlite/%v", mailerlite.Version))
		assert.Equal(t, req.URL.String(), "https://connect.mailerlite.com/api/timezones")
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Timezone.List(ctx)
	if err != nil {
		return
	}

	assert.Equal(t, res.StatusCode, http.StatusOK)

}

func TestWillHandleError(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		// return nil to force error from mock server
		return nil
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	listOptions := &mailerlite.ListSubscriberOptions{}

	_, _, err := client.Subscriber.List(ctx, listOptions)

	assert.Error(t, err)
}

func TestWillHandleAPIError(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Request:    req,
			Body: io.NopCloser(strings.NewReader(`{"message":"The given data was invalid.",
			"errors": {"filter": ["The filter must be an array."]}}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	listOptions := &mailerlite.ListSubscriberOptions{}

	_, _, err := client.Subscriber.List(ctx, listOptions)

	if err, ok := err.(*mailerlite.ErrorResponse); ok {
		assert.Equal(t, "The given data was invalid.", err.Message)
		assert.Equal(t, 1, len(err.Errors))
	}

	assert.Error(t, err)
	assert.IsType(t, err, &mailerlite.ErrorResponse{})
	assert.Equal(t, err.Error(), "GET https://connect.mailerlite.com/api/subscribers: 422 The given data was invalid. map[filter:[The filter must be an array.]]")
}

func TestWillHandleAPIError202(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusCreated,
			Request:    req,
			Body:       io.NopCloser(strings.NewReader("")),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	listOptions := &mailerlite.ListSubscriberOptions{}

	_, _, err := client.Subscriber.List(ctx, listOptions)

	if err, ok := err.(*mailerlite.ErrorResponse); ok {
		assert.Equal(t, "The given data was invalid.", err.Message)
		assert.Equal(t, 1, len(err.Errors))
	}

	assert.Error(t, err)
}

func TestWillHandleAPIFilters(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), "https://connect.mailerlite.com/api/subscribers?filter%5Bstatus%5D=active")
		return &http.Response{
			StatusCode: http.StatusAccepted,
			Request:    req,
			Body: io.NopCloser(strings.NewReader(`{
				"data": [
				  {
					"id": "123456789",
					"email": "client@example.com",
					"status": "active"
				  }
				]
			}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	listOptions := &mailerlite.ListSubscriberOptions{
		Filters: &[]mailerlite.Filter{{Name: "status", Value: "active"}},
	}

	subscribers, _, _ := client.Subscriber.List(ctx, listOptions)

	assert.Equal(t, len(subscribers.Data), 1)
	assert.Equal(t, subscribers.Data[0].Status, "active")
}

func TestWillHandleMultipleAPIFilters(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), "https://connect.mailerlite.com/api/subscribers?filter%5Bname%5D=groupName&filter%5Bstatus%5D=active")

		return &http.Response{
			StatusCode: http.StatusAccepted,
			Request:    req,
			Body: io.NopCloser(strings.NewReader(`{
				"data": [
				  {
					"id": "123456789",
					"email": "client@example.com",
					"status": "active"
				  }
				]
			}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	filters := &[]mailerlite.Filter{
		{Name: "status", Value: "active"},
		{Name: "name", Value: "groupName"},
	}

	listOptions := &mailerlite.ListSubscriberOptions{
		Filters: filters,
	}

	subscribers, _, _ := client.Subscriber.List(ctx, listOptions)

	assert.Equal(t, len(subscribers.Data), 1)
	assert.Equal(t, subscribers.Data[0].Status, "active")
}

func TestWillHandleAPIAuthError(t *testing.T) {

	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusUnauthorized,
			Request:    req,
			Body:       io.NopCloser(strings.NewReader(`{"message": "Unauthenticated."}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	client.SetAPIKey("invalid-api-key")

	listOptions := &mailerlite.ListSubscriberOptions{}

	_, _, err := client.Subscriber.List(ctx, listOptions)

	assert.Error(t, err)
	assert.IsType(t, err, &mailerlite.AuthError{})
	assert.Equal(t, err.Error(), "GET https://connect.mailerlite.com/api/subscribers: 401 Unauthenticated. map[]")
}

func TestWillHandleAPIRateError(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	header := http.Header{}
	header.Set(mailerlite.HeaderRateLimit, "60")
	header.Set(mailerlite.HeaderRateRemaining, "0")
	header.Set(mailerlite.HeaderRateRetryAfter, "59")

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		res := &http.Response{
			StatusCode: http.StatusTooManyRequests,
			Request:    req,
			Header:     header,
			Body:       io.NopCloser(strings.NewReader(`{"message": "Too Many Attempts."}`)),
		}

		return res
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	listOptions := &mailerlite.ListSubscriberOptions{}

	_, res, err := client.Subscriber.List(ctx, listOptions)

	assert.Equal(t, http.StatusTooManyRequests, res.StatusCode)
	assert.IsType(t, &mailerlite.RateLimitError{}, err)

	retryAfter := time.Duration(59) * time.Second

	if err, ok := err.(*mailerlite.RateLimitError); ok {
		assert.Equal(t, "Too Many Attempts.", err.Message)
		assert.Equal(t, 0, err.Rate.Remaining)
		assert.Equal(t, &retryAfter, err.Rate.RetryAfter)
	}

	assert.Equal(t, "GET https://connect.mailerlite.com/api/subscribers: 429 Too Many Attempts. [retry after 59s]", err.Error())

}

func TestWillHandleAPIRateErrorAndNoRemoteCall(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	header := http.Header{}
	header.Set(mailerlite.HeaderRateLimit, "120")
	header.Set(mailerlite.HeaderRateRemaining, "0")
	header.Set(mailerlite.HeaderRateRetryAfter, "59")

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		res := &http.Response{
			StatusCode: http.StatusTooManyRequests,
			Request:    req,
			Header:     header,
			Body:       io.NopCloser(strings.NewReader(`{"message": "Too Many Attempts."}`)),
		}

		return res
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	listOptions := &mailerlite.ListSubscriberOptions{}

	_, res, err := client.Subscriber.List(ctx, listOptions)

	assert.Equal(t, http.StatusTooManyRequests, res.StatusCode)
	assert.IsType(t, &mailerlite.RateLimitError{}, err)

	retryAfter := time.Duration(59) * time.Second

	if err, ok := err.(*mailerlite.RateLimitError); ok {
		assert.Equal(t, "Too Many Attempts.", err.Message)
		assert.Equal(t, 0, err.Rate.Remaining)
		assert.Equal(t, &retryAfter, err.Rate.RetryAfter)
	}

	assert.Equal(t, "GET https://connect.mailerlite.com/api/subscribers: 429 Too Many Attempts. [retry after 59s]", err.Error())

	_, res, err = client.Subscriber.List(ctx, listOptions)
	assert.Equal(t, http.StatusForbidden, res.StatusCode)
	assert.IsType(t, &mailerlite.RateLimitError{}, err)

	assert.Equal(t, "GET https://connect.mailerlite.com/api/subscribers: 403 API rate limit of 120 still exceeded until 59s, not making remote request. [retry after 59s]", err.Error())

}
