package mailerlite_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/mailerlite/mailerlite-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testBrandStylePalette = `{"identifier":"%s","name":"%s","font_family":"Inter, sans-serif","primary_color":"#2cb191","secondary_color":"#111111",
"heading_color":"#111111","text_color":"#333333","border_color":"#e5e7eb","background_color":"#ffffff"}`

func TestCanGetBrandStyles(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/brand-styles", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(`{"data":{"brand":null,"default":` +
				fmt.Sprintf(testBrandStylePalette, "default-brand", "Default") + `,"presets":[` +
				fmt.Sprintf(testBrandStylePalette, "preset-retro", "Retro") + `]}}`)),
		}
	})

	client.SetHttpClient(testClient)

	root, _, err := client.BrandStyle.Get(context.TODO())
	require.NoError(t, err)

	assert.Nil(t, root.Data.Brand)
	assert.Equal(t, "default-brand", root.Data.Default.Identifier)
	assert.Equal(t, "#2cb191", root.Data.Default.PrimaryColor)
	require.Len(t, root.Data.Presets, 1)
	assert.Equal(t, "preset-retro", root.Data.Presets[0].Identifier)
	assert.Equal(t, "Retro", root.Data.Presets[0].Name)
}

func TestCanGetBrandStylesWithAccountBrand(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(bytes.NewBufferString(`{"data":{"brand":` +
				fmt.Sprintf(testBrandStylePalette, "account", "Your brand styles") + `,"default":` +
				fmt.Sprintf(testBrandStylePalette, "default-brand", "Default") + `,"presets":[]}}`)),
		}
	})

	client.SetHttpClient(testClient)

	root, _, err := client.BrandStyle.Get(context.TODO())
	require.NoError(t, err)

	require.NotNil(t, root.Data.Brand)
	assert.Equal(t, "account", root.Data.Brand.Identifier)
	assert.Equal(t, "Inter, sans-serif", root.Data.Brand.FontFamily)
}
