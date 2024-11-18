package mailerlite

import (
	"context"
	"fmt"
	"net/http"
)

const groupEndpoint = "/groups"

// GroupService defines an interface for group-related operations.
type GroupService interface {
	List(ctx context.Context, options *ListGroupOptions) (*RootGroups, *Response, error)
	Create(ctx context.Context, groupName string) (*RootGroup, *Response, error)
	Update(ctx context.Context, groupID string, groupName string) (*RootGroup, *Response, error)
	Delete(ctx context.Context, groupID string) (*Response, error)
	Subscribers(ctx context.Context, options *ListGroupSubscriberOptions) (*RootSubscribers, *Response, error)
	Assign(ctx context.Context, groupID, subscriberID string) (*RootGroup, *Response, error)
	UnAssign(ctx context.Context, groupID, subscriberID string) (*Response, error)
}

// groupService implements GroupService.
type groupService struct {
	*service
}

type RootGroup struct {
	Data Group `json:"data"`
}

type RootGroups struct {
	Data  []Group `json:"data"`
	Links Links   `json:"links"`
	Meta  Meta    `json:"meta"`
}

type Group struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	ActiveCount       int       `json:"active_count"`
	SentCount         int       `json:"sent_count"`
	OpensCount        int       `json:"opens_count"`
	OpenRate          OpenRate  `json:"open_rate"`
	ClicksCount       int       `json:"clicks_count"`
	ClickRate         ClickRate `json:"click_rate"`
	UnsubscribedCount int       `json:"unsubscribed_count"`
	UnconfirmedCount  int       `json:"unconfirmed_count"`
	BouncedCount      int       `json:"bounced_count"`
	JunkCount         int       `json:"junk_count"`
	CreatedAt         string    `json:"created_at"`
}

// ListGroupOptions - modifies the behavior of GroupService.List method
type ListGroupOptions struct {
	Filters *[]Filter `json:"filters,omitempty"`
	Page    int       `url:"page,omitempty"`
	Limit   int       `url:"limit,omitempty"`
	Sort    string    `url:"sort,omitempty"`
}

type ListGroupSubscriberOptions struct {
	GroupID string    `url:"-"`
	Filters *[]Filter `json:"filters,omitempty"`
	Page    int       `url:"page,omitempty"`
	Limit   int       `url:"limit,omitempty"`
}

func (s *groupService) List(ctx context.Context, options *ListGroupOptions) (*RootGroups, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, groupEndpoint, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootGroups)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *groupService) Create(ctx context.Context, groupName string) (*RootGroup, *Response, error) {
	body := map[string]interface{}{"name": groupName}
	req, err := s.client.newRequest(http.MethodPost, groupEndpoint, body)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootGroup)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *groupService) Update(ctx context.Context, groupID, groupName string) (*RootGroup, *Response, error) {
	body := map[string]interface{}{"name": groupName}
	path := fmt.Sprintf("%s/%s", groupEndpoint, groupID)

	req, err := s.client.newRequest(http.MethodPut, path, body)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootGroup)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *groupService) Delete(ctx context.Context, groupID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", groupEndpoint, groupID)

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

func (s *groupService) Subscribers(ctx context.Context, options *ListGroupSubscriberOptions) (*RootSubscribers, *Response, error) {
	path := fmt.Sprintf("%s/%s/subscribers", groupEndpoint, options.GroupID)

	req, err := s.client.newRequest(http.MethodGet, path, options)
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

func (s *groupService) Assign(ctx context.Context, groupID, subscriberID string) (*RootGroup, *Response, error) {
	path := fmt.Sprintf("%s/%s/groups/%s", subscriberEndpoint, subscriberID, groupID)

	req, err := s.client.newRequest(http.MethodPost, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootGroup)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *groupService) UnAssign(ctx context.Context, groupID, subscriberID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s/groups/%s", subscriberEndpoint, subscriberID, groupID)

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
