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
