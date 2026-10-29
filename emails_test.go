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

func TestCanDeleteEmailBlock(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodDelete, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/emails/123/blocks/57", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"etag":"42"}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`{"etag":"43"}`)),
		}
	})

	client.SetHttpClient(testClient)

	root, _, err := client.Email.DeleteBlock(context.TODO(), "123", "57", "42")
	require.NoError(t, err)

	assert.Equal(t, "43", root.Etag)
}

func TestCanValidateEmailBlock(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/emails/blocks/validate", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"type":"paragraph","editor_content":{"attrs":{"textAlign":"sideways"}}}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(`{"valid":false,"errors":{"editor_content.attrs.textAlign":["The value is not allowed."]},
			"schema":{"version":"2.73.2","hash":"abc"}}`)),
		}
	})

	client.SetHttpClient(testClient)

	block := &mailerlite.CreateEmailBlockOptions{
		Type:          mailerlite.EmailBlockTypeParagraph,
		EditorContent: map[string]interface{}{"attrs": map[string]interface{}{"textAlign": "sideways"}},
	}

	result, _, err := client.Email.ValidateBlock(context.TODO(), block)
	require.NoError(t, err)

	assert.False(t, result.Valid)
	assert.Equal(t, []string{"The value is not allowed."}, result.Errors["editor_content.attrs.textAlign"])
	assert.Equal(t, "2.73.2", result.Schema.Version)
	assert.Nil(t, result.Rejected)
}

func TestCanUpdateEmailSettings(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPatch, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/emails/123/settings", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"brand":{"mode":"custom","background_color":"#f4f4f4"},"footer":{"unsubscribe_text":""},"etag":"42"}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(`{"data":{"brand":{"mode":"custom","background_color":"#f4f4f4"},"footer":{"social_links":[]}},
			"etag":"43","rejected":{"message":"Some attributes were rejected.","keys":["footer.social_links[0].url"]}}`)),
		}
	})

	client.SetHttpClient(testClient)

	settings := &mailerlite.UpdateEmailSettingsOptions{
		EmailSettings: mailerlite.EmailSettings{
			Brand: &mailerlite.EmailBrandSettings{
				Mode:            mailerlite.String(mailerlite.EmailBrandModeCustom),
				BackgroundColor: mailerlite.String("#f4f4f4"),
			},
			Footer: &mailerlite.EmailFooterSettings{
				UnsubscribeText: mailerlite.String(""),
			},
		},
		Etag: "42",
	}

	root, _, err := client.Email.UpdateSettings(context.TODO(), "123", settings)
	require.NoError(t, err)

	assert.Equal(t, "43", root.Etag)
	assert.Equal(t, "#f4f4f4", *root.Data.Brand.BackgroundColor)
	assert.Empty(t, *root.Data.Footer.SocialLinks)
	require.NotNil(t, root.Rejected)
	assert.Equal(t, []string{"footer.social_links[0].url"}, root.Rejected.Keys)
}

func TestCanValidateEmailSettings(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/emails/123/settings/validate", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"button":{"border_radius":"16px"}}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(`{"valid":false,"errors":{"button.border_radius":["Must be a whole number."]},
			"schema":{"version":"2.73.2","hash":"abc"}}`)),
		}
	})

	client.SetHttpClient(testClient)

	settings := &mailerlite.UpdateEmailSettingsOptions{
		EmailSettings: mailerlite.EmailSettings{
			Button: &mailerlite.EmailButtonSettings{BorderRadius: mailerlite.String("16px")},
		},
	}

	result, _, err := client.Email.ValidateSettings(context.TODO(), "123", settings)
	require.NoError(t, err)

	assert.False(t, result.Valid)
	assert.Equal(t, []string{"Must be a whole number."}, result.Errors["button.border_radius"])
}

func TestCanCreateEmailTemplate(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/emails/123/templates", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"name":"Newsletter"}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusCreated,
			Body:       io.NopCloser(bytes.NewBufferString(`{"data":{"id":"789","name":"Newsletter","type":"email","created_at":"2026-10-28 12:00:00"}}`)),
		}
	})

	client.SetHttpClient(testClient)

	root, res, err := client.Email.CreateTemplate(context.TODO(), "123", &mailerlite.CreateEmailTemplateOptions{Name: "Newsletter"})
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.Equal(t, "789", root.Data.ID)
	assert.Equal(t, "Newsletter", root.Data.Name)
	assert.Equal(t, "email", root.Data.Type)
	assert.Equal(t, "2026-10-28 12:00:00", root.Data.CreatedAt)
}

func TestCanCreateEmailVersion(t *testing.T) {
	for _, status := range []int{http.StatusCreated, http.StatusOK} {
		client := mailerlite.NewClient(testKey)

		testClient := NewTestClient(func(req *http.Request) *http.Response {
			assert.Equal(t, http.MethodPost, req.Method)
			assert.Equal(t, "https://connect.mailerlite.com/api/emails/123/versions", req.URL.String())
			b, err := io.ReadAll(req.Body)
			assert.NoError(t, err)
			assert.Len(t, b, 0)
			return &http.Response{
				StatusCode: status,
				Body:       io.NopCloser(bytes.NewBufferString(`{"data":{"id":"5","created_at":"2026-10-28 12:00:00","current":true}}`)),
			}
		})

		client.SetHttpClient(testClient)

		root, res, err := client.Email.CreateVersion(context.TODO(), "123")
		require.NoError(t, err)

		assert.Equal(t, status, res.StatusCode)
		assert.Equal(t, "5", root.Data.ID)
		assert.True(t, root.Data.Current)
	}
}

func TestCanListEmailVersions(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/emails/123/versions?limit=10&page=2", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(`{"data":[{"id":"5","created_at":"2026-10-28 12:00:00","current":true},{"id":"4","created_at":"2026-10-27 12:00:00","current":false}],
			"links":{"first":"https://connect.mailerlite.com/api/emails/123/versions?page=1","last":null,"prev":null,"next":null},
			"meta":{"current_page":2,"from":11,"last_page":2,"links":[{"url":null,"label":"1","page":1,"active":false}],"path":"https://connect.mailerlite.com/api/emails/123/versions","per_page":10,"to":12,"total":12}}`)),
		}
	})

	client.SetHttpClient(testClient)

	options := &mailerlite.ListEmailVersionOptions{
		Page:  2,
		Limit: 10,
	}

	root, _, err := client.Email.ListVersions(context.TODO(), "123", options)
	require.NoError(t, err)

	require.Len(t, root.Data, 2)
	assert.Equal(t, "5", root.Data[0].ID)
	assert.True(t, root.Data[0].Current)
	assert.False(t, root.Data[1].Current)
	assert.Equal(t, 12, root.Meta.Total)
	assert.True(t, root.Links.IsLastPage())
}

func TestCanRestoreEmailVersion(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/emails/123/versions/4/restore", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"etag":"42"}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`{"data":` + testEmailDocument + `,"etag":"43"}`)),
		}
	})

	client.SetHttpClient(testClient)

	root, _, err := client.Email.RestoreVersion(context.TODO(), "123", "4", "42")
	require.NoError(t, err)

	assert.Equal(t, "43", root.Etag)
	assert.Equal(t, "123", root.Data.ID)
	assert.Len(t, root.Data.Blocks, 2)
}

func TestCanListEmailVersionsWithoutOptions(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, "https://connect.mailerlite.com/api/emails/123/versions", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`{"data":[]}`)),
		}
	})

	client.SetHttpClient(testClient)

	root, _, err := client.Email.ListVersions(context.TODO(), "123", nil)
	require.NoError(t, err)

	assert.Empty(t, root.Data)
}

func TestEmailWritesRejectEmptyEtag(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		t.Errorf("unexpected request: %v %v", req.Method, req.URL)
		return nil
	})

	client.SetHttpClient(testClient)

	ctx := context.TODO()
	writes := map[string]func() error{
		"CreateBlock": func() error {
			_, _, err := client.Email.CreateBlock(ctx, "123", &mailerlite.CreateEmailBlockOptions{Type: mailerlite.EmailBlockTypeDivider})
			return err
		},
		"CreateBlock nil": func() error {
			_, _, err := client.Email.CreateBlock(ctx, "123", nil)
			return err
		},
		"UpdateBlock": func() error {
			_, _, err := client.Email.UpdateBlock(ctx, "123", "57", &mailerlite.UpdateEmailBlockOptions{})
			return err
		},
		"MoveBlock": func() error {
			_, _, err := client.Email.MoveBlock(ctx, "123", "57", &mailerlite.MoveEmailBlockOptions{})
			return err
		},
		"DeleteBlock": func() error {
			_, _, err := client.Email.DeleteBlock(ctx, "123", "57", "")
			return err
		},
		"UpdateSettings": func() error {
			_, _, err := client.Email.UpdateSettings(ctx, "123", &mailerlite.UpdateEmailSettingsOptions{})
			return err
		},
		"RestoreVersion": func() error {
			_, _, err := client.Email.RestoreVersion(ctx, "123", "4", "")
			return err
		},
	}

	for name, write := range writes {
		assert.ErrorIs(t, write(), mailerlite.ErrEmptyEtag, name)
	}
}

func TestEmailMethodsRejectEmptyID(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		t.Errorf("unexpected request: %v %v", req.Method, req.URL)
		return nil
	})

	client.SetHttpClient(testClient)

	ctx := context.TODO()
	calls := map[string]func() error{
		"Get": func() error {
			_, _, err := client.Email.Get(ctx, "")
			return err
		},
		"GetBlock": func() error {
			_, _, err := client.Email.GetBlock(ctx, "123", "")
			return err
		},
		"DeleteBlock": func() error {
			_, _, err := client.Email.DeleteBlock(ctx, "123", "", "42")
			return err
		},
		"ListVersions": func() error {
			_, _, err := client.Email.ListVersions(ctx, "", nil)
			return err
		},
		"RestoreVersion": func() error {
			_, _, err := client.Email.RestoreVersion(ctx, "123", "", "42")
			return err
		},
	}

	for name, call := range calls {
		assert.ErrorIs(t, call(), mailerlite.ErrEmptyID, name)
	}
}

func TestEmailMethodsRejectNonNumericID(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		t.Errorf("unexpected request: %v %v", req.Method, req.URL)
		return nil
	})

	client.SetHttpClient(testClient)

	ctx := context.TODO()
	for _, id := range []string{"123/html", "123%2Fhtml", "../123", "123?a=b", " 123", "-1", "1e3"} {
		_, _, err := client.Email.Get(ctx, id)
		assert.ErrorIs(t, err, mailerlite.ErrInvalidID, "Get %q", id)

		_, _, err = client.Email.GetBlock(ctx, "123", id)
		assert.ErrorIs(t, err, mailerlite.ErrInvalidID, "GetBlock %q", id)

		_, _, err = client.Email.RestoreVersion(ctx, "123", id, "42")
		assert.ErrorIs(t, err, mailerlite.ErrInvalidID, "RestoreVersion %q", id)
	}
}

func TestEmailBlockPositionChecksBlockIDOnlyForBeforeAndAfter(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	var bodies []string
	testClient := NewTestClient(func(req *http.Request) *http.Response {
		b, _ := io.ReadAll(req.Body)
		bodies = append(bodies, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(testEmailBlock)),
		}
	})

	client.SetHttpClient(testClient)

	ctx := context.TODO()
	sends := map[string]func(mailerlite.EmailBlockPosition) (string, error){
		"CreateBlock": func(p mailerlite.EmailBlockPosition) (string, error) {
			_, _, err := client.Email.CreateBlock(ctx, "123", &mailerlite.CreateEmailBlockOptions{
				Type: mailerlite.EmailBlockTypeDivider, EditorContent: map[string]interface{}{}, Placement: &p, Etag: "42",
			})
			return `{"type":"divider","editor_content":{},"placement":%s,"etag":"42"}`, err
		},
		"MoveBlock": func(p mailerlite.EmailBlockPosition) (string, error) {
			_, _, err := client.Email.MoveBlock(ctx, "123", "57", &mailerlite.MoveEmailBlockOptions{Placement: p, Etag: "42"})
			return `{"placement":%s,"etag":"42"}`, err
		},
		"ValidateBlock": func(p mailerlite.EmailBlockPosition) (string, error) {
			_, _, err := client.Email.ValidateBlock(ctx, &mailerlite.CreateEmailBlockOptions{
				Type: mailerlite.EmailBlockTypeDivider, EditorContent: map[string]interface{}{}, Placement: &p,
			})
			return `{"type":"divider","editor_content":{},"placement":%s}`, err
		},
	}

	cases := []struct {
		position  string
		blockID   string
		err       error
		placement string
	}{
		{"start", "", nil, `{"position":"start"}`},
		{"start", "1/2", nil, `{"position":"start"}`},
		{"start", "56", nil, `{"position":"start"}`},
		{"end", "", nil, `{"position":"end"}`},
		{"end", "x", nil, `{"position":"end"}`},
		{"before", "56", nil, `{"position":"before","block_id":"56"}`},
		{"after", "56", nil, `{"position":"after","block_id":"56"}`},
		{"before", "", mailerlite.ErrEmptyID, ""},
		{"after", "", mailerlite.ErrEmptyID, ""},
		{"before", "1/2", mailerlite.ErrInvalidID, ""},
		{"after", "56\n", mailerlite.ErrInvalidID, ""},
		{"middle", "", nil, `{"position":"middle"}`},
		{"middle", "56", nil, `{"position":"middle","block_id":"56"}`},
		{"middle", "1/2", mailerlite.ErrInvalidID, ""},
	}

	for name, send := range sends {
		for _, c := range cases {
			bodies = nil
			p := mailerlite.EmailBlockPosition{Position: c.position, BlockID: c.blockID}
			body, err := send(p)
			label := name + " " + c.position + " " + c.blockID
			if c.err != nil {
				assert.ErrorIs(t, err, c.err, label)
				assert.Empty(t, bodies, label)
				continue
			}
			require.NoError(t, err, label)
			require.Len(t, bodies, 1, label)
			assert.Equal(t, strings.Replace(body, "%s", c.placement, 1), bodies[0], label)
		}
	}
}

func TestEmailBlockPositionDoesNotChangeTheCallerValue(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(testEmailBlock)),
		}
	})

	client.SetHttpClient(testClient)

	ctx := context.TODO()
	placement := &mailerlite.EmailBlockPosition{Position: mailerlite.EmailBlockPositionEnd, BlockID: "56"}
	block := &mailerlite.CreateEmailBlockOptions{Type: mailerlite.EmailBlockTypeDivider, EditorContent: map[string]interface{}{}, Placement: placement, Etag: "42"}
	move := &mailerlite.MoveEmailBlockOptions{Placement: *placement, Etag: "42"}

	_, _, err := client.Email.CreateBlock(ctx, "123", block)
	require.NoError(t, err)
	_, _, err = client.Email.ValidateBlock(ctx, block)
	require.NoError(t, err)
	_, _, err = client.Email.MoveBlock(ctx, "123", "57", move)
	require.NoError(t, err)

	assert.Same(t, placement, block.Placement)
	assert.Equal(t, "56", placement.BlockID)
	assert.Equal(t, "56", move.Placement.BlockID)
}

func TestEmailWriteReturnsVersionConflict(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusConflict,
			Request:    req,
			Body: io.NopCloser(strings.NewReader(`{"message":"The email changed since you read it. Fetch it again, then retry with its new etag.",
			"code":"version_conflict","etag":"43"}`)),
		}
	})

	client.SetHttpClient(testClient)

	settings := &mailerlite.UpdateEmailSettingsOptions{
		EmailSettings: mailerlite.EmailSettings{
			Layout: &mailerlite.EmailLayoutSettings{Alignment: mailerlite.String("left")},
		},
		Etag: "42",
	}

	_, res, err := client.Email.UpdateSettings(context.TODO(), "123", settings)

	assert.Equal(t, http.StatusConflict, res.StatusCode)
	var errResp *mailerlite.ErrorResponse
	require.ErrorAs(t, err, &errResp)
	assert.Equal(t, mailerlite.ErrorCodeVersionConflict, errResp.Code)
	assert.Equal(t, "43", errResp.Etag)
}
