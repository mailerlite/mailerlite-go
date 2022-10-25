package mailerlite

import (
	"context"
	"fmt"
	"net/http"
)

const groupEndpoint = "/groups"

type GroupService service

type rootGroup struct {
	Data Group `json:"data"`
}

type rootGroups struct {
	Data  []Group `json:"data"`
	Links Links   `json:"links"`
	Meta  Meta    `json:"meta"`
}

type OpenRate struct {
	Float  int    `json:"float"`
	String string `json:"string"`
}

type ClickRate struct {
	Float  int    `json:"float"`
	String string `json:"string"`
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
	Filter *Filter `json:"filter,omitempty"`
	Page   int     `url:"page,omitempty"`
	Limit  int     `url:"limit,omitempty"`
	Sort   string  `url:"sort,omitempty"`
}

type ListGroupSubscriberOptions struct {
	GroupID string  `url:"-"`
	Filter  *Filter `json:"filter,omitempty"`
	Page    int     `url:"page,omitempty"`
	Limit   int     `url:"limit,omitempty"`
}

func (s *GroupService) List(ctx context.Context, options *ListGroupOptions) (*rootGroups, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, groupEndpoint, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootGroups)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *GroupService) Create(ctx context.Context, groupName string) (*rootGroup, *Response, error) {
	body := map[string]interface{}{"name": groupName}
	req, err := s.client.newRequest(http.MethodPost, groupEndpoint, body)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootGroup)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *GroupService) Update(ctx context.Context, groupID, groupName string) (*rootGroup, *Response, error) {
	body := map[string]interface{}{"name": groupName}
	path := fmt.Sprintf("%s/%s", groupEndpoint, groupID)

	req, err := s.client.newRequest(http.MethodPut, path, body)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootGroup)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *GroupService) Delete(ctx context.Context, groupID string) (*rootGroup, *Response, error) {
	path := fmt.Sprintf("%s/%s", groupEndpoint, groupID)

	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootGroup)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// TODO: is this required we could use subscribers endpoint to filter by groupID
func (s *GroupService) Subscribers(ctx context.Context, options *ListGroupSubscriberOptions) (*rootSubscribers, *Response, error) {
	path := fmt.Sprintf("%s/%s/subscribers", groupEndpoint, options.GroupID)

	req, err := s.client.newRequest(http.MethodGet, path, options)
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

func (s *GroupService) Assign(ctx context.Context, groupID, subscriberID string) (*rootGroup, *Response, error) {
	path := fmt.Sprintf("%s/%s/groups/%s", subscriberEndpoint, subscriberID, groupID)

	req, err := s.client.newRequest(http.MethodPost, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootGroup)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *GroupService) UnAssign(ctx context.Context, groupID, subscriberID string) (*Response, error) {
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
