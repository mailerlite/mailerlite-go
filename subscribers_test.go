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

func TestCanListSubscribersWithAllParams(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/subscribers?cursor=abc&end_date=2026-01-31+00%3A00%3A00&filter%5Bstatus%5D=active&include=groups&limit=10&page=2&sort=-created_at&start_date=2026-01-01+00%3A00%3A00&with_fields=1&with_groups=1", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":[]}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Subscriber.List(ctx, &mailerlite.ListSubscriberOptions{Cursor: "abc", Limit: 10, Page: 2, Sort: "-created_at", Include: "groups", StartDate: "2026-01-01 00:00:00", EndDate: "2026-01-31 00:00:00", WithFields: true, WithGroups: true, Filters: &[]mailerlite.Filter{{Name: "status", Value: "active"}}})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanUpsertSubscriberWithResubscribeAndNotes(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/subscribers", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"email":"test@example.com","resubscribe":true,"notes":["VIP"]}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{}}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Subscriber.Upsert(ctx, &mailerlite.UpsertSubscriber{Email: "test@example.com", Resubscribe: true, Notes: []string{"VIP"}})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanListSubscriberActivityLog(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/subscribers/1/activity-log?filter%5Blog_name%5D=email_open&limit=10&page=2", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":[{"id":"1","created_at":"2026-01-01 00:00:00","updated_at":"2026-01-02 00:00:00"}]}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Subscriber.ActivityLog(ctx, &mailerlite.ListActivityOptions{SubscriberID: "1", Page: 2, Limit: 10, Filters: &[]mailerlite.Filter{{Name: "log_name", Value: "email_open"}}})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanGetSubscriberImport(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/subscribers/import/1", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{"id":"1","account_id":"2","total":3,"processed":1,"percent":33.3,"done":false,"duplicate_count":3,"is_stopped_due_to_free_account_limit_exceeded":true}}`)),
		}
	})

	client.SetHttpClient(testClient)

	imp, res, err := client.Subscriber.GetImport(context.TODO(), "1")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "2", imp.Data.AccountID)
	assert.Equal(t, 3, imp.Data.DuplicateCount)
	assert.True(t, imp.Data.IsStoppedDueToFreeAccountLimitExceeded)
	assert.Equal(t, 33.3, imp.Data.Percent)
}

func TestCanDecodeActivityLogWithEmptyProperties(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(strings.NewReader(`{"data":[
				{"id":"1","log_name":"activated","properties":[],"created_at":"2026-01-01 00:00:00"},
				{"id":"2","log_name":"added_to_group","properties":{"group_id":"5"}}
			]}`)),
		}
	})

	client.SetHttpClient(testClient)

	log, _, err := client.Subscriber.ActivityLog(context.TODO(), &mailerlite.ListActivityOptions{SubscriberID: "1"})
	assert.NoError(t, err)
	assert.Len(t, log.Data, 2)
	assert.Equal(t, "activated", log.Data[0].LogName)
	assert.Equal(t, "2026-01-01 00:00:00", log.Data[0].CreatedAt)
	assert.Empty(t, log.Data[0].Properties)
	assert.Equal(t, "5", log.Data[1].Properties["group_id"])
}
