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
	"github.com/stretchr/testify/require"
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

func TestWillHandleAPIErrorCode(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusConflict,
			Request:    req,
			Body: io.NopCloser(strings.NewReader(`{"message":"The email changed since you read it. Fetch it again, then retry with its new etag.",
			"code":"version_conflict","etag":"2"}`)),
		}
	})

	client.SetHttpClient(testClient)

	_, _, err := client.Timezone.List(context.TODO())

	var errResp *mailerlite.ErrorResponse
	require.ErrorAs(t, err, &errResp)
	assert.Equal(t, mailerlite.ErrorCodeVersionConflict, errResp.Code)
	assert.Equal(t, "2", errResp.Etag)
	assert.Equal(t, "GET https://connect.mailerlite.com/api/timezones: 409 version_conflict: The email changed since you read it. Fetch it again, then retry with its new etag. map[]", err.Error())
}

func TestWillKeepAPIErrorCodeOnRateLimitError(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	header := http.Header{}
	header.Set(mailerlite.HeaderRateRemaining, "0")

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusTooManyRequests,
			Request:    req,
			Header:     header,
			Body:       io.NopCloser(strings.NewReader(`{"message":"Too Many Attempts.","code":"too_many_requests"}`)),
		}
	})

	client.SetHttpClient(testClient)

	_, _, err := client.Timezone.List(context.TODO())

	var rateErr *mailerlite.RateLimitError
	require.ErrorAs(t, err, &rateErr)
	assert.Equal(t, "too_many_requests", rateErr.Code)
}

func TestWillHandleAPIFilterWithSpaces(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://connect.mailerlite.com/api/files?filter%5Bname%5D=summer+sale+%7Bbanner%7D", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Request:    req,
			Body:       io.NopCloser(strings.NewReader(`{"data":[]}`)),
		}
	})

	client.SetHttpClient(testClient)

	options := &mailerlite.ListFileOptions{
		Filters: &[]mailerlite.Filter{{Name: "name", Value: "summer sale {banner}"}},
	}

	_, _, err := client.File.List(context.TODO(), options)
	require.NoError(t, err)
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

func TestWillHandleAPIFiltersWithSpaces(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://connect.mailerlite.com/api/groups?filter%5Bname%5D=My+Value", req.URL.String())

		return &http.Response{
			StatusCode: http.StatusOK,
			Request:    req,
			Body:       io.NopCloser(strings.NewReader(`{"data": []}`)),
		}
	})

	client.SetHttpClient(testClient)

	listOptions := &mailerlite.ListGroupOptions{
		Filters: &[]mailerlite.Filter{{Name: "name", Value: "My Value"}},
	}

	_, _, err := client.Group.List(context.TODO(), listOptions)
	assert.NoError(t, err)
}

func TestWillHandleNonStringAPIFilters(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://connect.mailerlite.com/api/groups?filter%5Bcount%5D=5", req.URL.String())

		return &http.Response{
			StatusCode: http.StatusOK,
			Request:    req,
			Body:       io.NopCloser(strings.NewReader(`{"data": []}`)),
		}
	})

	client.SetHttpClient(testClient)

	listOptions := &mailerlite.ListGroupOptions{
		Filters: &[]mailerlite.Filter{{Name: "count", Value: 5}},
	}

	_, _, err := client.Group.List(context.TODO(), listOptions)
	assert.NoError(t, err)
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

func TestCanSendIdempotencyKey(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, "create-block-1", req.Header.Get(mailerlite.HeaderIdempotencyKey))
		return &http.Response{
			StatusCode: http.StatusCreated,
			Header:     http.Header{mailerlite.HeaderIdempotentReplayed: []string{"true"}},
			Body:       io.NopCloser(strings.NewReader(`{"data":{"id":"57","type":"divider","editor_content":{}},"etag":"43"}`)),
		}
	})

	client.SetHttpClient(testClient)

	ctx := mailerlite.WithIdempotencyKey(context.TODO(), "create-block-1")
	block := &mailerlite.CreateEmailBlockOptions{
		Type:          mailerlite.EmailBlockTypeDivider,
		EditorContent: map[string]interface{}{},
		Etag:          "42",
	}

	_, res, err := client.Email.CreateBlock(ctx, "123", block)
	require.NoError(t, err)

	assert.True(t, res.IsIdempotentReplay())
}

func TestDoesNotSendIdempotencyKeyByDefault(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		_, ok := req.Header[mailerlite.HeaderIdempotencyKey]
		assert.False(t, ok)
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":[]}`)),
		}
	})

	client.SetHttpClient(testClient)

	_, res, err := client.Timezone.List(context.TODO())
	require.NoError(t, err)

	assert.False(t, res.IsIdempotentReplay())
}

func TestRejectsInvalidIdempotencyKey(t *testing.T) {
	keys := map[string]string{
		"empty":         "",
		"too long":      strings.Repeat("a", 256),
		"new line":      "key\nX-Injected: 1",
		"tab":           "key\tvalue",
		"not ASCII":     "clé",
		"delete (0x7f)": "key\x7f",
	}

	for name, key := range keys {
		t.Run(name, func(t *testing.T) {
			client := mailerlite.NewClient(testKey)

			requests := 0
			client.SetHttpClient(NewTestClient(func(req *http.Request) *http.Response {
				requests++
				return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(`{}`))}
			}))

			ctx := mailerlite.WithIdempotencyKey(context.TODO(), key)
			_, _, err := client.Email.CreateVersion(ctx, "123")

			assert.ErrorIs(t, err, mailerlite.ErrInvalidIdempotencyKey)
			assert.Equal(t, 0, requests)
		})
	}
}

func TestAcceptsLongestIdempotencyKey(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	key := strings.Repeat("~", 255)
	client.SetHttpClient(NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, key, req.Header.Get(mailerlite.HeaderIdempotencyKey))
		return &http.Response{StatusCode: http.StatusCreated, Body: io.NopCloser(strings.NewReader(`{"data":{"id":"5"}}`))}
	}))

	_, _, err := client.Email.CreateVersion(mailerlite.WithIdempotencyKey(context.TODO(), key), "123")
	require.NoError(t, err)
}
