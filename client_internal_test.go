package mailerlite

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRequestEncodesPatchBody(t *testing.T) {
	client := NewClient("valid-api-key")

	req, err := client.newRequest(http.MethodPatch, "/emails/1/settings", map[string]string{"etag": "1"})
	require.NoError(t, err)

	b, err := io.ReadAll(req.Body)
	require.NoError(t, err)
	assert.Equal(t, `{"etag":"1"}`, strings.TrimRight(string(b), "\r\n"))
}

func TestNewRequestSendsNoBodyForNilBody(t *testing.T) {
	client := NewClient("valid-api-key")

	bodies := map[string]interface{}{
		"nil":         nil,
		"nil pointer": (*CreateCampaign)(nil),
		"nil map":     map[string]string(nil),
	}

	for name, body := range bodies {
		t.Run(name, func(t *testing.T) {
			req, err := client.newRequest(http.MethodPost, "/emails/1/versions", body)
			require.NoError(t, err)

			b, err := io.ReadAll(req.Body)
			require.NoError(t, err)
			assert.Len(t, b, 0)
			assert.Equal(t, int64(0), req.ContentLength)
		})
	}
}
