package mailerlite

import (
	"context"
	"net/http"
)

const timezoneEndpoint = "/timezones"

type TimezoneService service

type rootTimezones struct {
	Data []Timezone `json:"data"`
}

type Timezone struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	NameForHumans string `json:"name_for_humans"`
	OffsetName    string `json:"offset_name"`
	Offset        int    `json:"offset"`
}

func (s *TimezoneService) List(ctx context.Context) (*rootTimezones, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, timezoneEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootTimezones)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
