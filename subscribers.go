package mailerlite

import (
	"context"
	"fmt"
	"net/http"
)

const subscriberEndpoint = "/subscribers"

type SubscriberService service

// subscribers - subscribers response
type rootSubscribers struct {
	Data  []Subscriber `json:"data"`
	Links Links        `json:"links"`
	Meta  Meta         `json:"meta"`
}

// subscribers - subscribers response
type rootSubscriber struct {
	Data Subscriber `json:"data"`
}

type count struct {
	Total int `json:"total"`
}

type Subscriber struct {
	ID             string                 `json:"id,omitempty"`
	Email          string                 `json:"email,omitempty"`
	Status         string                 `json:"status,omitempty"`
	Source         string                 `json:"source,omitempty"`
	Sent           int                    `json:"sent,omitempty"`
	OpensCount     int                    `json:"opens_count,omitempty"`
	ClicksCount    int                    `json:"clicks_count,omitempty"`
	OpenRate       float64                `json:"open_rate,omitempty"`
	ClickRate      float64                `json:"click_rate,omitempty"`
	IPAddress      interface{}            `json:"ip_address,omitempty"`
	SubscribedAt   string                 `json:"subscribed_at,omitempty"`
	UnsubscribedAt interface{}            `json:"unsubscribed_at,omitempty"`
	CreatedAt      string                 `json:"created_at,omitempty"`
	UpdatedAt      string                 `json:"updated_at,omitempty"`
	Fields         map[string]interface{} `json:"fields,omitempty"`
	Groups         []Group                `json:"groups,omitempty"`
	OptedInAt      string                 `json:"opted_in_at,omitempty"`
	OptinIP        string                 `json:"optin_ip,omitempty"`
}

type Fields struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

// ListSubscriberOptions - modifies the behavior of SubscriberService.List method
type ListSubscriberOptions struct {
	Filters *[]Filter `json:"filters,omitempty"`
	Page    int       `url:"page,omitempty"`
	Limit   int       `url:"limit,omitempty"`
}

// GetSubscriberOptions - modifies the behavior of SubscriberService.Get method
type GetSubscriberOptions struct {
	SubscriberID string `json:"id,omitempty"`
	Email        string `json:"email,omitempty"`
}

func (s *SubscriberService) List(ctx context.Context, options *ListSubscriberOptions) (*rootSubscribers, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, subscriberEndpoint, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootSubscribers)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// Count - get a count of subscribers
func (s *SubscriberService) Count(ctx context.Context) (*count, *Response, error) {
	path := fmt.Sprintf("%s?limit=0", subscriberEndpoint)
	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(count)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// Get - get a single subscriber by email or ID
func (s *SubscriberService) Get(ctx context.Context, options *GetSubscriberOptions) (*rootSubscriber, *Response, error) {
	param := options.SubscriberID
	if options.Email != "" {
		param = options.Email
	}
	path := fmt.Sprintf("%s/%s", subscriberEndpoint, param)
	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootSubscriber)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *SubscriberService) Create(ctx context.Context, subscriber *Subscriber) (*rootSubscriber, *Response, error) {
	req, err := s.client.newRequest(http.MethodPost, subscriberEndpoint, subscriber)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootSubscriber)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *SubscriberService) Update(ctx context.Context, subscriber *Subscriber) (*rootSubscriber, *Response, error) {
	req, err := s.client.newRequest(http.MethodPost, subscriberEndpoint, subscriber)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootSubscriber)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *SubscriberService) Delete(ctx context.Context, subscriberID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", subscriberEndpoint, subscriberID)

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

func (s *SubscriberService) Forget(ctx context.Context, subscriberID string) (*rootSubscriber, *Response, error) {
	path := fmt.Sprintf("%s/%s/forget", subscriberEndpoint, subscriberID)

	req, err := s.client.newRequest(http.MethodPost, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootSubscriber)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
