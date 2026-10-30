package mailerlite_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/mailerlite/mailerlite-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCanListFiles(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/files?filter%5Btype%5D=image&limit=10&page=1&sort=-size", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(`{"data":[
			{"id":"1","type":"image","name":"logo.png","url":"https://example.com/logo.png","size":2048,"width":120,"height":40,"created_at":"2026-10-28 12:00:00"},
			{"id":"2","type":"pdf","name":"guide.pdf","url":"https://example.com/guide.pdf","size":4096,"width":null,"height":null,"created_at":"2026-10-27 12:00:00"}],
			"links":{"first":"https://connect.mailerlite.com/api/files?page=1","last":null,"prev":null,"next":null},
			"meta":{"current_page":1,"from":1,"last_page":1,"links":[],"path":"https://connect.mailerlite.com/api/files","per_page":10,"to":2,"total":2}}`)),
		}
	})

	client.SetHttpClient(testClient)

	options := &mailerlite.ListFileOptions{
		Filters: &[]mailerlite.Filter{{
			Name:  "type",
			Value: mailerlite.FileTypeImage,
		}},
		Sort:  mailerlite.SortBySizeDescending,
		Page:  1,
		Limit: 10,
	}

	root, _, err := client.File.List(context.TODO(), options)
	require.NoError(t, err)

	require.Len(t, root.Data, 2)
	assert.Equal(t, "logo.png", root.Data[0].Name)
	assert.Equal(t, 2048, root.Data[0].Size)
	assert.Equal(t, 120, *root.Data[0].Width)
	assert.Nil(t, root.Data[1].Width)
	assert.Nil(t, root.Data[1].Height)
	assert.Equal(t, 2, root.Meta.Total)
}
