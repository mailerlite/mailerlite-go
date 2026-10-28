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

const testEmailDocument = `{
	"id": "123",
	"type": "builder_inline_email",
	"emailable_type": "campaigns",
	"emailable_id": "456",
	"settings": {
		"brand": {"mode": "account", "font_family": "Inter, sans-serif", "background_color": "#ffffff", "text_color": "#000000", "border_color": "#e5e7eb"},
		"layout": {"width_mode": "default", "alignment": "center"},
		"typography": {
			"link": {"color": "#2cb191"},
			"heading": {"font_family": "Inter, sans-serif", "font_weight": "700", "line_height": "150%", "letter_spacing": "0", "h1": {"color": "#000000", "size": "32"}},
			"paragraph": {"font_family": "Inter, sans-serif", "font_weight": "400", "line_height": "150%", "letter_spacing": "0", "p1": {"color": "#000000", "size": "16"}}
		},
		"button": {"font_family": "Inter, sans-serif", "font_weight": "600", "letter_spacing": "0", "background_color": "#2cb191", "text_color": "#ffffff", "border_radius": "4"},
		"footer": {
			"show_preferences": false,
			"company_name": {"editor_content": {"content": [{"type": "text", "text": "Acme"}]}},
			"unsubscribe_text": "Unsubscribe",
			"social_links": [{"url": "https://example.com", "type": "website", "label": "Website", "link_id": 7}]
		}
	},
	"blocks": [
		{"id": "1", "type": "paragraph", "editor_content": {"attrs": {"textAlign": "left"}, "content": [{"type": "text", "text": "Hello"}]}},
		{"id": "2", "type": "bulleted_list", "editor_content": {}, "children": [{"id": "3", "type": "paragraph", "editor_content": {}}]}
	]
}`

func TestCanGetEmail(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/emails/123", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`{"data":` + testEmailDocument + `,"etag":"42"}`)),
		}
	})

	client.SetHttpClient(testClient)

	root, res, err := client.Email.Get(context.TODO(), "123")
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "42", root.Etag)
	assert.Equal(t, "123", root.Data.ID)
	assert.Equal(t, "456", root.Data.EmailableID)
	assert.Equal(t, "account", *root.Data.Settings.Brand.Mode)
	assert.Equal(t, "32", *root.Data.Settings.Typography.Heading.H1.Size)
	assert.Nil(t, root.Data.Settings.Typography.Heading.H2)
	assert.Equal(t, "4", *root.Data.Settings.Button.BorderRadius)
	assert.False(t, *root.Data.Settings.Footer.ShowPreferences)
	assert.Nil(t, root.Data.Settings.Footer.Address)
	assert.Equal(t, 7, *(*root.Data.Settings.Footer.SocialLinks)[0].LinkID)
	require.Len(t, root.Data.Blocks, 2)
	assert.Equal(t, "paragraph", root.Data.Blocks[0].Type)
	assert.Equal(t, map[string]interface{}{"textAlign": "left"}, root.Data.Blocks[0].EditorContent["attrs"])
	assert.Equal(t, "3", root.Data.Blocks[1].Children[0].ID)
	assert.Nil(t, root.Data.Blocks[0].Placement)
}
