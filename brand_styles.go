package mailerlite

import (
	"context"
	"net/http"
)

const brandStyleEndpoint = "/brand-styles"

// BrandStyleService defines an interface for brand style operations.
type BrandStyleService interface {
	Get(ctx context.Context) (*RootBrandStyles, *Response, error)
}

// brandStyleService implements BrandStyleService.
type brandStyleService struct {
	*service
}

// RootBrandStyles - brand styles response
type RootBrandStyles struct {
	Data BrandStyles `json:"data"`
}

// BrandStyles - all palettes available to the account.
// Brand is nil when the account has no brand styles of its own.
type BrandStyles struct {
	Brand   *BrandStyle  `json:"brand"`
	Default BrandStyle   `json:"default"`
	Presets []BrandStyle `json:"presets"`
}

// BrandStyle - a palette of one font and six colours.
// Identifier is stable. Name is for display only and can change.
type BrandStyle struct {
	Identifier      string `json:"identifier"`
	Name            string `json:"name"`
	FontFamily      string `json:"font_family"`
	PrimaryColor    string `json:"primary_color"`
	SecondaryColor  string `json:"secondary_color"`
	HeadingColor    string `json:"heading_color"`
	TextColor       string `json:"text_color"`
	BorderColor     string `json:"border_color"`
	BackgroundColor string `json:"background_color"`
}

// Get - get the brand styles of the account, the default palette and the presets
func (s *brandStyleService) Get(ctx context.Context) (*RootBrandStyles, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, brandStyleEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootBrandStyles)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
