package mailerlite_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
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

func TestCanCreateSubscriber(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.Method, http.MethodPost)
		assert.Equal(t, req.URL.String(), "https://connect.mailerlite.com/api/subscribers")
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`{"data":{}}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	cases := []struct {
		name    string
		options *mailerlite.CreateSubscriberRequest
	}{
		{
			name: "with mutable fields",
			options: &mailerlite.CreateSubscriberRequest{
				Email: "test@test.com",
				MutableSubscriberFields: mailerlite.MutableSubscriberFields{
					Groups: []string{"123", "456"},
				},
			},
		},
		{
			name: "without mutable fields",
			options: &mailerlite.CreateSubscriberRequest{
				Email: "test@test.com",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, res, err := client.Subscriber.Create(ctx, tc.options)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, res.StatusCode, http.StatusOK)
		})
	}
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
