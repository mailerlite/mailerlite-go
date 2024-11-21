package mailerlite

import (
	"context"
	"net/http"
)

const timezoneEndpoint = "/timezones"

// TimezoneService defines an interface for timezone-related operations.
type TimezoneService interface {
	List(ctx context.Context) (*RootTimezones, *Response, error)
}

type timezoneService struct {
	*service
}

type RootTimezones struct {
	Data []Timezone `json:"data"`
}

type Timezone struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	NameForHumans string `json:"name_for_humans"`
	OffsetName    string `json:"offset_name"`
	Offset        int    `json:"offset"`
}

func (s *timezoneService) List(ctx context.Context) (*RootTimezones, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, timezoneEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootTimezones)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
