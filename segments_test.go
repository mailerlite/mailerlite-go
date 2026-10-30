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

func TestCanListSegmentsWithSort(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/segments?limit=10&page=2&sort=-name", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":[]}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Segment.List(ctx, &mailerlite.ListSegmentOptions{Page: 2, Limit: 10, Sort: "-name"})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanUpdateSegmentWithFilter(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPut, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/segments/1", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"name":"VIP","filter":[[{"args":["groups","$in",["2"]],"operator":"in_any"}]]}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{}}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Segment.UpdateWithOptions(ctx, "1", &mailerlite.UpdateSegmentOptions{Name: "VIP", Filter: []interface{}{[]interface{}{map[string]interface{}{"operator": "in_any", "args": []interface{}{"groups", "$in", []string{"2"}}}}}})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanListSegmentSubscribersWithCursor(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/segments/1/subscribers?cursor=abc&filter%5Bstatus%5D=active&limit=10", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":[]}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Segment.Subscribers(ctx, &mailerlite.ListSegmentSubscriberOptions{SegmentID: "1", Cursor: "abc", Limit: 10, Filters: &[]mailerlite.Filter{{Name: "status", Value: "active"}}})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanListSegmentSubscribersWithAfter(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/segments/1/subscribers?after=123&limit=10", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":[]}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Segment.Subscribers(ctx, &mailerlite.ListSegmentSubscriberOptions{SegmentID: "1", After: 123, Limit: 10})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanUpdateSegmentName(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPut, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/segments/1", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"name":"VIP"}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{"id":"1","name":"VIP"}}`)),
		}
	})

	client.SetHttpClient(testClient)

	seg, _, err := client.Segment.Update(context.TODO(), "1", "VIP")
	assert.NoError(t, err)
	assert.Equal(t, "VIP", seg.Data.Name)
}
