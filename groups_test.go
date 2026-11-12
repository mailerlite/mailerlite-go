package mailerlite_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/mailerlite/mailerlite-go"
	"github.com/stretchr/testify/assert"
)

func TestCanListGroupSubscribersWithCursor(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/groups/1/subscribers?cursor=abc&filter%5Bstatus%5D=active&include=groups&limit=10", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":[]}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Group.Subscribers(ctx, &mailerlite.ListGroupSubscriberOptions{GroupID: "1", Cursor: "abc", Limit: 10, Include: "groups", Filters: &[]mailerlite.Filter{{Name: "status", Value: "active"}}})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanImportSubscribersToGroup(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/groups/1/import-subscribers", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"subscribers":[{"email":"test@example.com","name":"Test"}],"resubscribe":true,"trigger_automations":true}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"import_progress_url":"https://connect.mailerlite.com/api/subscribers/import/5"}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	imp, res, err := client.Group.ImportSubscribers(ctx, "1", &mailerlite.ImportSubscribersOptions{Subscribers: []mailerlite.ImportSubscriber{{Email: "test@example.com", Name: "Test"}}, Resubscribe: true, TriggerAutomations: mailerlite.Bool(true)})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "https://connect.mailerlite.com/api/subscribers/import/5", imp.ImportProgressURL)
}

func TestCanImportSubscribersToGroupWithoutTriggeringAutomations(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"subscribers":[{"email":"test@example.com"}],"trigger_automations":false}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"import_progress_url":"https://connect.mailerlite.com/api/subscribers/import/5"}`)),
		}
	})

	client.SetHttpClient(testClient)

	_, _, err := client.Group.ImportSubscribers(context.TODO(), "1", &mailerlite.ImportSubscribersOptions{Subscribers: []mailerlite.ImportSubscriber{{Email: "test@example.com"}}, TriggerAutomations: mailerlite.Bool(false)})
	assert.NoError(t, err)
}

func TestCanImportSubscribersToGroupOmitsTriggerAutomationsWhenNil(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"subscribers":[{"email":"test@example.com"}]}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"import_progress_url":"https://connect.mailerlite.com/api/subscribers/import/5"}`)),
		}
	})

	client.SetHttpClient(testClient)

	_, _, err := client.Group.ImportSubscribers(context.TODO(), "1", &mailerlite.ImportSubscribersOptions{Subscribers: []mailerlite.ImportSubscriber{{Email: "test@example.com"}}})
	assert.NoError(t, err)
}
