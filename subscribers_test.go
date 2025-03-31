package mailerlite_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/mailerlite/mailerlite-go"
	"github.com/stretchr/testify/assert"
)

func TestCanListSubscrbers(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), "https://connect.mailerlite.com/api/subscribers")
		return &http.Response{
			StatusCode: http.StatusAccepted,
			Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	options := &mailerlite.ListSubscriberOptions{}

	_, res, err := client.Subscriber.List(ctx, options)
	if err != nil {
		return
	}

	assert.Equal(t, res.StatusCode, http.StatusAccepted)

}

func TestCanGetSingleSubscrber(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), "https://connect.mailerlite.com/api/subscribers/123456")
		return &http.Response{
			StatusCode: http.StatusAccepted,
			Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	options := &mailerlite.GetSubscriberOptions{
		SubscriberID: "123456",
	}

	_, res, err := client.Subscriber.Get(ctx, options)
	if err != nil {
		return
	}

	assert.Equal(t, res.StatusCode, http.StatusAccepted)
}

func TestCanGetSingleSubscrberByEmail(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), "https://connect.mailerlite.com/api/subscribers/test@test.com")
		return &http.Response{
			StatusCode: http.StatusAccepted,
			Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	options := &mailerlite.GetSubscriberOptions{
		Email: "test@test.com",
	}

	_, res, err := client.Subscriber.Get(ctx, options)
	if err != nil {
		return
	}

	assert.Equal(t, res.StatusCode, http.StatusAccepted)
}

func TestCanCreateSubscrber(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.Method, http.MethodPost)
		assert.Equal(t, req.URL.String(), "https://connect.mailerlite.com/api/subscribers")
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	options := &mailerlite.Subscriber{
		Email: "test@test.com",
	}

	_, res, err := client.Subscriber.Create(ctx, options)
	if err != nil {
		return
	}

	assert.Equal(t, res.StatusCode, http.StatusOK)
}

func TestCanDeleteSubscrber(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.Method, http.MethodDelete)
		assert.Equal(t, req.URL.String(), "https://connect.mailerlite.com/api/subscribers/1234")
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	res, err := client.Subscriber.Delete(ctx, "1234")
	if err != nil {
		assert.Fail(t, "Delete request threw an error")
	}

	assert.Equal(t, res.StatusCode, http.StatusOK)
}

func TestCanForgetSubscrber(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.Method, http.MethodPost)
		assert.Equal(t, req.URL.String(), "https://connect.mailerlite.com/api/subscribers/1234/forget")
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Subscriber.Forget(ctx, "1234")
	if err != nil {
		return
	}

	assert.Equal(t, res.StatusCode, http.StatusOK)
}

func TestCanCreateSubscrberWithGroupDeprecated(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.Method, http.MethodPost)
		assert.Equal(t, req.URL.String(), "https://connect.mailerlite.com/api/subscribers")
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, strings.TrimRight(string(b), "\r\n"), `{"email":"test@test.com","groups":[{"id":"1234","name":"","active_count":0,"sent_count":0,"opens_count":0,"open_rate":{"float":0,"string":""},"clicks_count":0,"click_rate":{"float":0,"string":""},"unsubscribed_count":0,"unconfirmed_count":0,"bounced_count":0,"junk_count":0,"created_at":""}]}`)
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	options := &mailerlite.Subscriber{
		Email:  "test@test.com",
		Groups: []mailerlite.Group{{ID: "1234"}},
	}

	_, res, err := client.Subscriber.Create(ctx, options)
	if err != nil {
		return
	}

	assert.Equal(t, res.StatusCode, http.StatusOK)
}

func TestCanUpsertSubscrberWithGroup(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.Method, http.MethodPost)
		assert.Equal(t, req.URL.String(), "https://connect.mailerlite.com/api/subscribers")
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, strings.TrimRight(string(b), "\r\n"), `{"email":"test@test.com","groups":["1234"]}`)
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	options := &mailerlite.UpsertSubscriber{
		Email:  "test@test.com",
		Groups: []string{"1234"},
	}

	_, res, err := client.Subscriber.Upsert(ctx, options)
	if err != nil {
		return
	}

	assert.Equal(t, res.StatusCode, http.StatusOK)
}

func TestCanUpdateSubscrberWithGroup(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.Method, http.MethodPut)
		assert.Equal(t, req.URL.String(), "https://connect.mailerlite.com/api/subscribers/1")
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, strings.TrimRight(string(b), "\r\n"), `{"id":"1","email":"test@test.com","groups":["1234"]}`)
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`OK`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	options := &mailerlite.UpdateSubscriber{
		ID:     "1",
		Email:  "test@test.com",
		Groups: []string{"1234"},
	}

	_, res, err := client.Subscriber.Update(ctx, options)
	if err != nil {
		return
	}

	assert.Equal(t, res.StatusCode, http.StatusOK)
}
