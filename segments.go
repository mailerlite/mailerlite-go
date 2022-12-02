package mailerlite

import (
	"context"
	"fmt"
	"net/http"
)

const segmentEndpoint = "/segments"

type SegmentService service

type rootSegment struct {
	Data Segment `json:"data"`
}

type rootSegments struct {
	Data  []Segment `json:"data"`
	Links Links     `json:"links"`
	Meta  Meta      `json:"meta"`
}

type Segment struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Total     int       `json:"total"`
	OpenRate  OpenRate  `json:"open_rate"`
	ClickRate ClickRate `json:"click_rate"`
	CreatedAt string    `json:"created_at"`
}

// ListSegmentOptions - modifies the behavior of SegmentService.List method
type ListSegmentOptions struct {
	Page  int `url:"page,omitempty"`
	Limit int `url:"limit,omitempty"`
}

// ListSegmentSubscriberOptions - modifies the behavior of SegmentService.Subscribers method
type ListSegmentSubscriberOptions struct {
	SegmentID string  `url:"-"`
	Filter    *Filter `json:"filter,omitempty"`
	Limit     int     `url:"limit,omitempty"`
	After     int     `url:"after,omitempty"`
}

func (s *SegmentService) List(ctx context.Context, options *ListSegmentOptions) (*rootSegments, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, segmentEndpoint, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootSegments)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *SegmentService) Update(ctx context.Context, segmentID, segmentName string) (*rootSegment, *Response, error) {
	body := map[string]interface{}{"name": segmentName}
	path := fmt.Sprintf("%s/%s", segmentEndpoint, segmentID)

	req, err := s.client.newRequest(http.MethodPut, path, body)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootSegment)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *SegmentService) Delete(ctx context.Context, segmentID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", segmentEndpoint, segmentID)

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

func (s *SegmentService) Subscribers(ctx context.Context, options *ListSegmentSubscriberOptions) (*rootSubscribers, *Response, error) {
	path := fmt.Sprintf("%s/%s/subscribers", segmentEndpoint, options.SegmentID)

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
