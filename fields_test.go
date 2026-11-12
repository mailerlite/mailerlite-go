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

func TestCanListFields(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/fields?filter%5Bkeyword%5D=city&filter%5Btype%5D=text&limit=10&page=2&sort=name", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":[{"id":"1","is_default":true}]}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Field.List(ctx, &mailerlite.ListFieldOptions{Page: 2, Limit: 10, Sort: "name", Filters: &[]mailerlite.Filter{{Name: "keyword", Value: "city"}, {Name: "type", Value: "text"}}})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}
