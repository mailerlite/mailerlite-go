package mailerlite

import (
	"context"
	"fmt"
	"net/http"
)

const webhookEndpoint = "/webhooks"

// WebhookService defines an interface for webhook-related operations.
type WebhookService interface {
	List(ctx context.Context, options *ListWebhookOptions) (*RootWebhooks, *Response, error)
	Get(ctx context.Context, webhookID string) (*RootWebhook, *Response, error)
	Create(ctx context.Context, webhook *CreateWebhookOptions) (*RootWebhook, *Response, error)
	Update(ctx context.Context, webhook *UpdateWebhookOptions) (*RootWebhook, *Response, error)
	Delete(ctx context.Context, webhookID string) (*Response, error)
}

// webhookService implements WebhookService.
type webhookService struct {
	*service
}

type RootWebhook struct {
	Data Webhook `json:"data"`
}

type RootWebhooks struct {
	Data  []Webhook `json:"data"`
	Links Links     `json:"links"`
	Meta  Meta      `json:"meta"`
}

type Webhook struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Url       string   `json:"url"`
	Events    []string `json:"events"`
	Enabled   bool     `json:"enabled"`
	Secret    string   `json:"secret"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

// ListWebhookOptions - modifies the behavior of WebhookService.List method
type ListWebhookOptions struct {
	Sort  string `url:"sort,omitempty"`
	Page  int    `url:"page,omitempty"`
	Limit int    `url:"limit,omitempty"`
}

// CreateWebhookOptions - modifies the behavior of WebhookService.Create method
type CreateWebhookOptions struct {
	Name   string   `json:"name,omitempty"`
	Events []string `json:"events"`
	Url    string   `json:"url"`
}

// UpdateWebhookOptions - modifies the behavior of WebhookService.Create method
type UpdateWebhookOptions struct {
	WebhookID string   `json:"-"`
	Name      string   `json:"name,omitempty"`
	Events    []string `json:"events,omitempty"`
	Url       string   `json:"url,omitempty"`
	Enabled   string   `json:"enabled,omitempty"`
}

func (s *webhookService) List(ctx context.Context, options *ListWebhookOptions) (*RootWebhooks, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, webhookEndpoint, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootWebhooks)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *webhookService) Get(ctx context.Context, webhookID string) (*RootWebhook, *Response, error) {
	path := fmt.Sprintf("%s/%s", webhookEndpoint, webhookID)
	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootWebhook)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *webhookService) Create(ctx context.Context, options *CreateWebhookOptions) (*RootWebhook, *Response, error) {
	req, err := s.client.newRequest(http.MethodPost, webhookEndpoint, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootWebhook)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *webhookService) Update(ctx context.Context, options *UpdateWebhookOptions) (*RootWebhook, *Response, error) {
	path := fmt.Sprintf("%s/%s", webhookEndpoint, options.WebhookID)

	req, err := s.client.newRequest(http.MethodPut, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootWebhook)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *webhookService) Delete(ctx context.Context, webhookID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", webhookEndpoint, webhookID)

	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	res, err := s.client.do(ctx, req, nil)
	if err != nil {
		return res, err
	}

	return res, nil
}
