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

func TestCanGetEmailHTML(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/emails/123/html", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`{"data":{"html":"<!doctype html><p>Hello</p>"},"etag":"42"}`)),
		}
	})

	client.SetHttpClient(testClient)

	root, _, err := client.Email.GetHTML(context.TODO(), "123")
	require.NoError(t, err)

	assert.Equal(t, "<!doctype html><p>Hello</p>", root.Data.HTML)
	assert.Equal(t, "42", root.Etag)
}

const testEmailBlock = `{"data":{"id":"57","type":"paragraph","editor_content":{"content":[{"type":"text","text":"Hello"}]},"placement":{"after":"56","before":null}},"etag":"43"}`

func TestCanGetEmailBlock(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/emails/123/blocks/57", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(testEmailBlock)),
		}
	})

	client.SetHttpClient(testClient)

	root, _, err := client.Email.GetBlock(context.TODO(), "123", "57")
	require.NoError(t, err)

	assert.Equal(t, "43", root.Etag)
	assert.Equal(t, "57", root.Data.ID)
	assert.Equal(t, "paragraph", root.Data.Type)
	require.NotNil(t, root.Data.Placement)
	assert.Equal(t, "56", *root.Data.Placement.After)
	assert.Nil(t, root.Data.Placement.Before)
}

func TestCanCreateEmailBlock(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/emails/123/blocks", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"type":"paragraph","editor_content":{"content":[{"text":"Hello","type":"text"}]},"placement":{"position":"after","block_id":"56"},"etag":"42"}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusCreated,
			Body:       io.NopCloser(bytes.NewBufferString(testEmailBlock)),
		}
	})

	client.SetHttpClient(testClient)

	block := &mailerlite.CreateEmailBlockOptions{
		Type: mailerlite.EmailBlockTypeParagraph,
		EditorContent: map[string]interface{}{
			"content": []interface{}{map[string]interface{}{"type": "text", "text": "Hello"}},
		},
		Placement: &mailerlite.EmailBlockPosition{
			Position: mailerlite.EmailBlockPositionAfter,
			BlockID:  "56",
		},
		Etag: "42",
	}

	root, res, err := client.Email.CreateBlock(context.TODO(), "123", block)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.Equal(t, "57", root.Data.ID)
	assert.Equal(t, "43", root.Etag)
}

func TestCanUpdateEmailBlock(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPatch, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/emails/123/blocks/57", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"editor_content":{"attrs":{"textAlign":"center"}},"etag":"42"}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(testEmailBlock)),
		}
	})

	client.SetHttpClient(testClient)

	block := &mailerlite.UpdateEmailBlockOptions{
		EditorContent: map[string]interface{}{"attrs": map[string]interface{}{"textAlign": "center"}},
		Etag:          "42",
	}

	root, _, err := client.Email.UpdateBlock(context.TODO(), "123", "57", block)
	require.NoError(t, err)

	assert.Equal(t, "57", root.Data.ID)
	assert.Equal(t, "43", root.Etag)
}

func TestCanMoveEmailBlock(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/emails/123/blocks/57/move", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"placement":{"position":"start"},"etag":"42"}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(testEmailBlock)),
		}
	})

	client.SetHttpClient(testClient)

	move := &mailerlite.MoveEmailBlockOptions{
		Placement: mailerlite.EmailBlockPosition{Position: mailerlite.EmailBlockPositionStart},
		Etag:      "42",
	}

	root, _, err := client.Email.MoveBlock(context.TODO(), "123", "57", move)
	require.NoError(t, err)

	assert.Equal(t, "57", root.Data.ID)
	assert.Equal(t, "43", root.Etag)
}
