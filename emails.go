package mailerlite

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

const emailEndpoint = "/emails"

// ErrorCodeVersionConflict - the email changed after you read the etag that the request sent
const ErrorCodeVersionConflict = "version_conflict"

var (
	// ErrEmptyID - an email, block or version ID is empty. The SDK sends no request.
	ErrEmptyID = errors.New("mailerlite: the ID must not be empty")

	// ErrInvalidID - an email, block or version ID contains a character that is not a digit. The SDK sends no request.
	ErrInvalidID = errors.New("mailerlite: the ID must contain only digits")

	// ErrEmptyEtag - a write has no etag. The SDK sends no request.
	ErrEmptyEtag = errors.New("mailerlite: the etag must not be empty")
)

// EmailService defines an interface for operations on emails built with the simple editor.
//
// Every write must send the etag that the last read or write of the email returned.
// See https://developers.mailerlite.com/docs/emails.html
type EmailService interface {
	Get(ctx context.Context, emailID string) (*RootEmailDocument, *Response, error)
	GetHTML(ctx context.Context, emailID string) (*RootEmailHTML, *Response, error)
	GetBlock(ctx context.Context, emailID string, blockID string) (*RootEmailBlock, *Response, error)
	CreateBlock(ctx context.Context, emailID string, block *CreateEmailBlockOptions) (*RootEmailBlock, *Response, error)
	UpdateBlock(ctx context.Context, emailID string, blockID string, block *UpdateEmailBlockOptions) (*RootEmailBlock, *Response, error)
	MoveBlock(ctx context.Context, emailID string, blockID string, move *MoveEmailBlockOptions) (*RootEmailBlock, *Response, error)
	DeleteBlock(ctx context.Context, emailID string, blockID string, etag string) (*RootEmailEtag, *Response, error)
	ValidateBlock(ctx context.Context, block *CreateEmailBlockOptions) (*EmailValidation, *Response, error)
	UpdateSettings(ctx context.Context, emailID string, settings *UpdateEmailSettingsOptions) (*RootEmailSettings, *Response, error)
	ValidateSettings(ctx context.Context, emailID string, settings *UpdateEmailSettingsOptions) (*EmailValidation, *Response, error)
	CreateTemplate(ctx context.Context, emailID string, template *CreateEmailTemplateOptions) (*RootEmailTemplate, *Response, error)
	CreateVersion(ctx context.Context, emailID string) (*RootEmailVersion, *Response, error)
	ListVersions(ctx context.Context, emailID string, options *ListEmailVersionOptions) (*RootEmailVersions, *Response, error)
	RestoreVersion(ctx context.Context, emailID string, versionID string, etag string) (*RootEmailDocument, *Response, error)
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
// All fields are pointers so that an update sends only the keys that you set.
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

// RootEmailBlock - single email block response
type RootEmailBlock struct {
	Data EmailBlock `json:"data"`
	Etag string     `json:"etag"`
}

// CreateEmailBlockOptions - modifies the behavior of EmailService.CreateBlock method.
// EditorContent is required. Use an empty map for a block with no content, such as a divider.
type CreateEmailBlockOptions struct {
	Type          string                 `json:"type"`
	EditorContent map[string]interface{} `json:"editor_content"`
	Placement     *EmailBlockPosition    `json:"placement,omitempty"`
	Etag          string                 `json:"etag,omitempty"`
}

// EmailBlockPosition - where to put a block. BlockID is required for the positions "before" and "after".
// For "start" and "end", the SDK does not check BlockID and does not send it.
type EmailBlockPosition struct {
	Position string `json:"position"`
	BlockID  string `json:"block_id,omitempty"`
}

// UpdateEmailBlockOptions - modifies the behavior of EmailService.UpdateBlock method.
// EditorContent replaces the whole content of the block. The API removes the keys that you do not send.
type UpdateEmailBlockOptions struct {
	EditorContent map[string]interface{} `json:"editor_content"`
	Etag          string                 `json:"etag,omitempty"`
}

// MoveEmailBlockOptions - modifies the behavior of EmailService.MoveBlock method
type MoveEmailBlockOptions struct {
	Placement EmailBlockPosition `json:"placement"`
	Etag      string             `json:"etag,omitempty"`
}

// RootEmailEtag - response with only the new etag of the email
type RootEmailEtag struct {
	Etag string `json:"etag"`
}

// emailEtag - request body for writes that send only the etag
type emailEtag struct {
	Etag string `json:"etag"`
}

// EmailValidation - the result of a validation. Invalid input is still a successful request.
type EmailValidation struct {
	Valid    bool                `json:"valid"`
	Errors   map[string][]string `json:"errors"`
	Schema   EmailSchema         `json:"schema"`
	Rejected *EmailRejected      `json:"rejected,omitempty"`
}

// EmailSchema - the version of the rules that the validation used
type EmailSchema struct {
	Version string `json:"version"`
	Hash    string `json:"hash"`
}

// EmailRejected - the keys of unsafe values that the API removed or replaced
type EmailRejected struct {
	Message string   `json:"message"`
	Keys    []string `json:"keys"`
}

// UpdateEmailSettingsOptions - modifies the behavior of EmailService.UpdateSettings method.
// Set only the keys to change. The API stores each value that you send on the email, also if it is the same as the account value.
type UpdateEmailSettingsOptions struct {
	EmailSettings
	Etag string `json:"etag,omitempty"`
}

// RootEmailSettings - email settings response
type RootEmailSettings struct {
	Data     EmailSettings  `json:"data"`
	Etag     string         `json:"etag"`
	Rejected *EmailRejected `json:"rejected,omitempty"`
}

// CreateEmailTemplateOptions - modifies the behavior of EmailService.CreateTemplate method
type CreateEmailTemplateOptions struct {
	Name string `json:"name"`
}

// RootEmailTemplate - email template response
type RootEmailTemplate struct {
	Data EmailTemplate `json:"data"`
}

type EmailTemplate struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	CreatedAt string `json:"created_at"`
}

// RootEmailVersion - single email version response
type RootEmailVersion struct {
	Data EmailVersion `json:"data"`
}

// EmailVersion - a saved copy of the content and the settings of an email
type EmailVersion struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	Current   bool   `json:"current"`
}

// RootEmailVersions - email versions response
type RootEmailVersions struct {
	Data  []EmailVersion `json:"data"`
	Links Links          `json:"links"`
	Meta  Meta           `json:"meta"`
}

// ListEmailVersionOptions - modifies the behavior of EmailService.ListVersions method
type ListEmailVersionOptions struct {
	Page  int `url:"page,omitempty"`
	Limit int `url:"limit,omitempty"`
}

// Get - get a single email with its blocks and settings
func (s *emailService) Get(ctx context.Context, emailID string) (*RootEmailDocument, *Response, error) {
	path, err := emailPath("/%s", emailID)
	if err != nil {
		return nil, nil, err
	}

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
	path, err := emailPath("/%s/html", emailID)
	if err != nil {
		return nil, nil, err
	}

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

// GetBlock - get a single block of an email, with its placement
func (s *emailService) GetBlock(ctx context.Context, emailID string, blockID string) (*RootEmailBlock, *Response, error) {
	path, err := emailPath("/%s/blocks/%s", emailID, blockID)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootEmailBlock)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// CreateBlock - insert a new block into an email. Without a placement, the block goes at the end.
func (s *emailService) CreateBlock(ctx context.Context, emailID string, block *CreateEmailBlockOptions) (*RootEmailBlock, *Response, error) {
	if block == nil || block.Etag == "" {
		return nil, nil, ErrEmptyEtag
	}

	path, err := emailPath("/%s/blocks", emailID)
	if err != nil {
		return nil, nil, err
	}

	block, err = blockToSend(block)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.newRequest(http.MethodPost, path, block)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootEmailBlock)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// UpdateBlock - replace the content of a block. The block type and position do not change.
func (s *emailService) UpdateBlock(ctx context.Context, emailID string, blockID string, block *UpdateEmailBlockOptions) (*RootEmailBlock, *Response, error) {
	if block == nil || block.Etag == "" {
		return nil, nil, ErrEmptyEtag
	}

	path, err := emailPath("/%s/blocks/%s", emailID, blockID)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.newRequest(http.MethodPatch, path, block)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootEmailBlock)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// MoveBlock - move a block among its siblings. The nesting of the block does not change.
func (s *emailService) MoveBlock(ctx context.Context, emailID string, blockID string, move *MoveEmailBlockOptions) (*RootEmailBlock, *Response, error) {
	if move == nil || move.Etag == "" {
		return nil, nil, ErrEmptyEtag
	}

	path, err := emailPath("/%s/blocks/%s/move", emailID, blockID)
	if err != nil {
		return nil, nil, err
	}

	placement, err := positionToSend(move.Placement)
	if err != nil {
		return nil, nil, err
	}
	sent := *move
	sent.Placement = placement

	req, err := s.client.newRequest(http.MethodPost, path, &sent)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootEmailBlock)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// DeleteBlock - delete a block and the blocks nested in it
func (s *emailService) DeleteBlock(ctx context.Context, emailID string, blockID string, etag string) (*RootEmailEtag, *Response, error) {
	if etag == "" {
		return nil, nil, ErrEmptyEtag
	}

	path, err := emailPath("/%s/blocks/%s", emailID, blockID)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.newRequest(http.MethodDelete, path, &emailEtag{Etag: etag})
	if err != nil {
		return nil, nil, err
	}

	root := new(RootEmailEtag)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// ValidateBlock - check a block against the rules of EmailService.CreateBlock. The API does not save the block.
func (s *emailService) ValidateBlock(ctx context.Context, block *CreateEmailBlockOptions) (*EmailValidation, *Response, error) {
	path, err := emailPath("/blocks/validate")
	if err != nil {
		return nil, nil, err
	}

	block, err = blockToSend(block)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.newRequest(http.MethodPost, path, block)
	if err != nil {
		return nil, nil, err
	}

	root := new(EmailValidation)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// UpdateSettings - change the design settings of an email
func (s *emailService) UpdateSettings(ctx context.Context, emailID string, settings *UpdateEmailSettingsOptions) (*RootEmailSettings, *Response, error) {
	if settings == nil || settings.Etag == "" {
		return nil, nil, ErrEmptyEtag
	}

	path, err := emailPath("/%s/settings", emailID)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.newRequest(http.MethodPatch, path, settings)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootEmailSettings)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// ValidateSettings - check a settings change against the rules of EmailService.UpdateSettings. The API does not save the settings.
func (s *emailService) ValidateSettings(ctx context.Context, emailID string, settings *UpdateEmailSettingsOptions) (*EmailValidation, *Response, error) {
	path, err := emailPath("/%s/settings/validate", emailID)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.newRequest(http.MethodPost, path, settings)
	if err != nil {
		return nil, nil, err
	}

	root := new(EmailValidation)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// CreateTemplate - save an email as a template
func (s *emailService) CreateTemplate(ctx context.Context, emailID string, template *CreateEmailTemplateOptions) (*RootEmailTemplate, *Response, error) {
	path, err := emailPath("/%s/templates", emailID)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.newRequest(http.MethodPost, path, template)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootEmailTemplate)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// CreateVersion - save the email as a new version. If nothing changed since the last version,
// the API returns that version with status 200 instead of 201.
func (s *emailService) CreateVersion(ctx context.Context, emailID string) (*RootEmailVersion, *Response, error) {
	path, err := emailPath("/%s/versions", emailID)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.newRequest(http.MethodPost, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootEmailVersion)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// ListVersions - list the saved versions of an email, newest first
func (s *emailService) ListVersions(ctx context.Context, emailID string, options *ListEmailVersionOptions) (*RootEmailVersions, *Response, error) {
	path, err := emailPath("/%s/versions", emailID)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootEmailVersions)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// RestoreVersion - put an email back to a saved version. The API saves the replaced state as a version first.
func (s *emailService) RestoreVersion(ctx context.Context, emailID string, versionID string, etag string) (*RootEmailDocument, *Response, error) {
	if etag == "" {
		return nil, nil, ErrEmptyEtag
	}

	path, err := emailPath("/%s/versions/%s/restore", emailID, versionID)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.newRequest(http.MethodPost, path, &emailEtag{Etag: etag})
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

// emailPath - build a path under emailEndpoint from format. Each ID must contain only digits,
// because the API decodes an escaped "/" in the path and then routes the request to a different endpoint.
func emailPath(format string, ids ...string) (string, error) {
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		if err := checkEmailID(id); err != nil {
			return "", err
		}
		args[i] = id
	}

	return emailEndpoint + fmt.Sprintf(format, args...), nil
}

// checkEmailID - return ErrEmptyID if id is empty, and ErrInvalidID if id contains a character that is not a digit
func checkEmailID(id string) error {
	if id == "" {
		return ErrEmptyID
	}
	for _, r := range id {
		if r < '0' || r > '9' {
			return ErrInvalidID
		}
	}

	return nil
}

// blockToSend - return block, or a copy of block with the placement from positionToSend. block does not change.
func blockToSend(block *CreateEmailBlockOptions) (*CreateEmailBlockOptions, error) {
	if block == nil || block.Placement == nil {
		return block, nil
	}

	placement, err := positionToSend(*block.Placement)
	if err != nil {
		return nil, err
	}
	sent := *block
	sent.Placement = &placement

	return &sent, nil
}

// positionToSend - check the BlockID of position and return the position to send.
// The API uses BlockID only for "before" and "after", and discards it without a check for "start" and "end".
// For these two positions, the SDK does not check BlockID and removes it.
// For other positions, the SDK checks BlockID if it is set, and sends it.
func positionToSend(position EmailBlockPosition) (EmailBlockPosition, error) {
	switch position.Position {
	case EmailBlockPositionStart, EmailBlockPositionEnd:
		position.BlockID = ""
		return position, nil
	case EmailBlockPositionBefore, EmailBlockPositionAfter:
		return position, checkEmailID(position.BlockID)
	}

	if position.BlockID != "" {
		return position, checkEmailID(position.BlockID)
	}

	return position, nil
}
