package mailerlite_test

import (
	"bytes"
	"context"
	"github.com/mailerlite/mailerlite-go"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
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
