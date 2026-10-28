package mailerlite

import (
	"context"
	"fmt"
	"net/http"
)

const emailEndpoint = "/emails"

// EmailService defines an interface for operations on emails built with the simple editor.
//
// Every write must send the etag that the last read or write of the email returned.
// See https://developers.mailerlite.com/docs/emails.html
type EmailService interface {
	Get(ctx context.Context, emailID string) (*RootEmailDocument, *Response, error)
	GetHTML(ctx context.Context, emailID string) (*RootEmailHTML, *Response, error)
}

// emailService implements EmailService.
type emailService struct {
	*service
}

// RootEmailDocument - single email response
type RootEmailDocument struct {
	Data EmailDocument `json:"data"`
	Etag string        `json:"etag"`
}

// EmailDocument - an email built with the simple editor: its blocks and its design settings
type EmailDocument struct {
	ID            string        `json:"id"`
	Type          string        `json:"type"`
	EmailableType string        `json:"emailable_type"`
	EmailableID   string        `json:"emailable_id"`
	Settings      EmailSettings `json:"settings"`
	Blocks        []EmailBlock  `json:"blocks"`
}

// EmailBlock - a content block of an email.
// EditorContent is the editor's document for the block, so its keys are camelCase.
type EmailBlock struct {
	ID            string                 `json:"id"`
	Type          string                 `json:"type"`
	EditorContent map[string]interface{} `json:"editor_content"`
	Children      []EmailBlock           `json:"children,omitempty"`
	Placement     *EmailBlockPlacement   `json:"placement,omitempty"`
}

// EmailBlockPlacement - the IDs of the blocks before and after a block, nil when there is none
type EmailBlockPlacement struct {
	After  *string `json:"after"`
	Before *string `json:"before"`
}

// EmailSettings - the design settings of an email.
//
// All fields are pointers so that an update sends only the keys that are set.
// An empty string unsets a value, and the email then uses the account value.
type EmailSettings struct {
	Brand      *EmailBrandSettings      `json:"brand,omitempty"`
	Layout     *EmailLayoutSettings     `json:"layout,omitempty"`
	Typography *EmailTypographySettings `json:"typography,omitempty"`
	Button     *EmailButtonSettings     `json:"button,omitempty"`
	Footer     *EmailFooterSettings     `json:"footer,omitempty"`
}

type EmailBrandSettings struct {
	Mode            *string `json:"mode,omitempty"`
	Style           *string `json:"style,omitempty"` // write-only: a preset identifier from BrandStyleService.Get
	FontFamily      *string `json:"font_family,omitempty"`
	BackgroundColor *string `json:"background_color,omitempty"`
	TextColor       *string `json:"text_color,omitempty"`
	BorderColor     *string `json:"border_color,omitempty"`
}

type EmailLayoutSettings struct {
	WidthMode *string `json:"width_mode,omitempty"`
	Alignment *string `json:"alignment,omitempty"`
}

type EmailTypographySettings struct {
	Link      *EmailLinkStyle      `json:"link,omitempty"`
	Heading   *EmailHeadingStyle   `json:"heading,omitempty"`
	Paragraph *EmailParagraphStyle `json:"paragraph,omitempty"`
}

type EmailLinkStyle struct {
	Color *string `json:"color,omitempty"`
}

type EmailHeadingStyle struct {
	FontFamily    *string         `json:"font_family,omitempty"`
	FontWeight    *string         `json:"font_weight,omitempty"`
	LineHeight    *string         `json:"line_height,omitempty"`
	LetterSpacing *string         `json:"letter_spacing,omitempty"`
	H1            *EmailTextStyle `json:"h1,omitempty"`
	H2            *EmailTextStyle `json:"h2,omitempty"`
	H3            *EmailTextStyle `json:"h3,omitempty"`
}

type EmailParagraphStyle struct {
	FontFamily    *string         `json:"font_family,omitempty"`
	FontWeight    *string         `json:"font_weight,omitempty"`
	LineHeight    *string         `json:"line_height,omitempty"`
	LetterSpacing *string         `json:"letter_spacing,omitempty"`
	P1            *EmailTextStyle `json:"p1,omitempty"`
}

type EmailTextStyle struct {
	Color *string `json:"color,omitempty"`
	Size  *string `json:"size,omitempty"`
}

type EmailButtonSettings struct {
	FontFamily      *string `json:"font_family,omitempty"`
	FontWeight      *string `json:"font_weight,omitempty"`
	LetterSpacing   *string `json:"letter_spacing,omitempty"`
	BackgroundColor *string `json:"background_color,omitempty"`
	TextColor       *string `json:"text_color,omitempty"`
	BorderRadius    *string `json:"border_radius,omitempty"`
}

type EmailFooterSettings struct {
	ShowPreferences *bool              `json:"show_preferences,omitempty"`
	CompanyName     *EmailEditorText   `json:"company_name,omitempty"`
	Address         *EmailEditorText   `json:"address,omitempty"`
	InfoText        *EmailEditorText   `json:"info_text,omitempty"`
	UnsubscribeText *string            `json:"unsubscribe_text,omitempty"`
	PreferencesText *string            `json:"preferences_text,omitempty"`
	SocialLinks     *[]EmailSocialLink `json:"social_links,omitempty"` // an update replaces the whole list
}

// EmailEditorText - a footer text as an editor document. An empty EditorContent map unsets the text.
type EmailEditorText struct {
	EditorContent map[string]interface{} `json:"editor_content"`
}

// EmailSocialLink - a footer social link. Send LinkID back unchanged to keep the click history of the link.
type EmailSocialLink struct {
	URL    string `json:"url,omitempty"`
	Type   string `json:"type,omitempty"`
	Label  string `json:"label,omitempty"`
	LinkID *int   `json:"link_id,omitempty"`
}

// RootEmailHTML - email HTML response
type RootEmailHTML struct {
	Data EmailHTML `json:"data"`
	Etag string    `json:"etag"`
}

// EmailHTML - a preview of the email HTML. It is not ready to send.
type EmailHTML struct {
	HTML string `json:"html"`
}

// Get - get a single email with its blocks and settings
func (s *emailService) Get(ctx context.Context, emailID string) (*RootEmailDocument, *Response, error) {
	path := fmt.Sprintf("%s/%s", emailEndpoint, emailID)
	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootEmailDocument)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// GetHTML - get the HTML preview of an email
func (s *emailService) GetHTML(ctx context.Context, emailID string) (*RootEmailHTML, *Response, error) {
	path := fmt.Sprintf("%s/%s/html", emailEndpoint, emailID)
	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootEmailHTML)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
