package mailerlite

import (
	"context"
	"fmt"
	"net/http"
)

const webhookEndpoint = "/webhooks"

type WebhookService service

type rootWebhook struct {
	Data Webhook `json:"data"`
}

type rootWebhooks struct {
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

func (s *WebhookService) List(ctx context.Context, options *ListWebhookOptions) (*rootWebhooks, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, webhookEndpoint, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootWebhooks)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *WebhookService) Get(ctx context.Context, webhookID string) (*rootWebhook, *Response, error) {
	path := fmt.Sprintf("%s/%s", webhookEndpoint, webhookID)
	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootWebhook)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *WebhookService) Create(ctx context.Context, options *CreateWebhookOptions) (*rootWebhook, *Response, error) {
	req, err := s.client.newRequest(http.MethodPost, webhookEndpoint, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootWebhook)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *WebhookService) Update(ctx context.Context, options *UpdateWebhookOptions) (*rootWebhook, *Response, error) {
	path := fmt.Sprintf("%s/%s", webhookEndpoint, options.WebhookID)

	req, err := s.client.newRequest(http.MethodPut, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootWebhook)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *WebhookService) Delete(ctx context.Context, webhookID string) (*Response, error) {
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
