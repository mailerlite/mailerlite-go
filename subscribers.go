package mailerlite

import (
	"context"
	"fmt"
	"net/http"
)

const subscriberEndpoint = "/subscribers"

// SubscriberService defines an interface for subscriber-related operations.
type SubscriberService interface {
	List(ctx context.Context, options *ListSubscriberOptions) (*RootSubscribers, *Response, error)
	Count(ctx context.Context) (*Count, *Response, error)
	Get(ctx context.Context, options *GetSubscriberOptions) (*RootSubscriber, *Response, error)
	// Deprecated: use Upsert instead (https://github.com/mailerlite/mailerlite-go/issues/17)
	Create(ctx context.Context, subscriber *Subscriber) (*RootSubscriber, *Response, error)
	Upsert(ctx context.Context, subscriber *UpsertSubscriber) (*RootSubscriber, *Response, error)
	Update(ctx context.Context, subscriber *UpdateSubscriber) (*RootSubscriber, *Response, error)
	Delete(ctx context.Context, subscriberID string) (*Response, error)
	Forget(ctx context.Context, subscriberID string) (*RootSubscriber, *Response, error)
}

// subscriberService implements SubscriberService.
type subscriberService struct {
	*service
}

// subscribers - subscribers response
type RootSubscribers struct {
	Data  []Subscriber `json:"data"`
	Links Links        `json:"links"`
	Meta  Meta         `json:"meta"`
}

// subscribers - subscribers response
type RootSubscriber struct {
	Data Subscriber `json:"data"`
}

type Count struct {
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

type UpdateSubscriber UpsertSubscriber

type UpsertSubscriber struct {
	ID             string                 `json:"id,omitempty"`
	Email          string                 `json:"email,omitempty"`
	Status         string                 `json:"status,omitempty"`
	IPAddress      interface{}            `json:"ip_address,omitempty"`
	SubscribedAt   string                 `json:"subscribed_at,omitempty"`
	UnsubscribedAt interface{}            `json:"unsubscribed_at,omitempty"`
	Fields         map[string]interface{} `json:"fields,omitempty"`
	Groups         []string               `json:"groups,omitempty"`
	OptedInAt      string                 `json:"opted_in_at,omitempty"`
	OptinIP        string                 `json:"optin_ip,omitempty"`
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

func (s *subscriberService) List(ctx context.Context, options *ListSubscriberOptions) (*RootSubscribers, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, subscriberEndpoint, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootSubscribers)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// Count - get a count of subscribers
func (s *subscriberService) Count(ctx context.Context) (*Count, *Response, error) {
	path := fmt.Sprintf("%s?limit=0", subscriberEndpoint)
	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(Count)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// Get - get a single subscriber by email or ID
func (s *subscriberService) Get(ctx context.Context, options *GetSubscriberOptions) (*RootSubscriber, *Response, error) {
	param := options.SubscriberID
	if options.Email != "" {
		param = options.Email
	}
	path := fmt.Sprintf("%s/%s", subscriberEndpoint, param)
	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootSubscriber)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// Deprecated: use Upsert instead (https://github.com/mailerlite/mailerlite-go/issues/17)
func (s *subscriberService) Create(ctx context.Context, subscriber *Subscriber) (*RootSubscriber, *Response, error) {
	req, err := s.client.newRequest(http.MethodPost, subscriberEndpoint, subscriber)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootSubscriber)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *subscriberService) Upsert(ctx context.Context, subscriber *UpsertSubscriber) (*RootSubscriber, *Response, error) {
	req, err := s.client.newRequest(http.MethodPost, subscriberEndpoint, subscriber)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootSubscriber)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *subscriberService) Update(ctx context.Context, subscriber *UpdateSubscriber) (*RootSubscriber, *Response, error) {
	path := fmt.Sprintf("%s/%s", subscriberEndpoint, subscriber.ID)

	req, err := s.client.newRequest(http.MethodPut, path, subscriber)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootSubscriber)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *subscriberService) Delete(ctx context.Context, subscriberID string) (*Response, error) {
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

func (s *subscriberService) Forget(ctx context.Context, subscriberID string) (*RootSubscriber, *Response, error) {
	path := fmt.Sprintf("%s/%s/forget", subscriberEndpoint, subscriberID)

	req, err := s.client.newRequest(http.MethodPost, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootSubscriber)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
