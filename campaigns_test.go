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

func TestCanCreateCampaignWithSimpleEditor(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/campaigns", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"name":"Simple editor campaign","type":"regular","emails":[{"subject":"Subject","from_name":"Your Name","from":"your@domain.com","reply_to":"reply@domain.com","editor":"simple","template_id":40001}]}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusCreated,
			Body:       io.NopCloser(bytes.NewBufferString(`{"data":{"id":"1","name":"Simple editor campaign","type":"regular","emails":[{"id":"123","type":"builder_inline_email"}]}}`)),
		}
	})

	client.SetHttpClient(testClient)

	campaign := &mailerlite.CreateCampaign{
		Name: "Simple editor campaign",
		Type: mailerlite.CampaignTypeRegular,
		Emails: []mailerlite.Emails{{
			Subject:    "Subject",
			FromName:   "Your Name",
			From:       "your@domain.com",
			ReplyTo:    "reply@domain.com",
			Editor:     mailerlite.EmailEditorSimple,
			TemplateID: 40001,
		}},
	}

	root, _, err := client.Campaign.Create(context.TODO(), campaign)
	require.NoError(t, err)

	require.Len(t, root.Data.Emails, 1)
	assert.Equal(t, "123", root.Data.Emails[0].ID)
	assert.Equal(t, "builder_inline_email", root.Data.Emails[0].Type)
}

func TestCreateCampaignOmitsEmptyEmailFields(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"name":"Campaign Name","type":"regular","emails":[{"subject":"Subject","from_name":"Your Name","from":"your@domain.com"}]}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusCreated,
			Body:       io.NopCloser(bytes.NewBufferString(`{"data":{"id":"1"}}`)),
		}
	})

	client.SetHttpClient(testClient)

	campaign := &mailerlite.CreateCampaign{
		Name: "Campaign Name",
		Type: mailerlite.CampaignTypeRegular,
		Emails: []mailerlite.Emails{{
			Subject:  "Subject",
			FromName: "Your Name",
			From:     "your@domain.com",
		}},
	}

	_, _, err := client.Campaign.Create(context.TODO(), campaign)
	require.NoError(t, err)
}

func TestCanCancelCampaignWithoutBody(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/campaigns/1/cancel", req.URL.String())
		b, err := io.ReadAll(req.Body)
		assert.NoError(t, err)
		assert.Len(t, b, 0)
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(`{"data":{"id":"1"}}`)),
		}
	})

	client.SetHttpClient(testClient)

	_, _, err := client.Campaign.Cancel(context.TODO(), "1")
	require.NoError(t, err)
}

func TestCanDeleteCampaignWithoutBody(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodDelete, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/campaigns/1", req.URL.String())
		b, err := io.ReadAll(req.Body)
		assert.NoError(t, err)
		assert.Len(t, b, 0)
		return &http.Response{
			StatusCode: http.StatusNoContent,
			Body:       io.NopCloser(bytes.NewBufferString(``)),
		}
	})

	client.SetHttpClient(testClient)

	_, err := client.Campaign.Delete(context.TODO(), "1")
	require.NoError(t, err)
}
