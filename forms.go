package mailerlite

import (
	"context"
	"fmt"
	"net/http"
)

const formEndpoint = "/forms"

type FormService service

type rootForm struct {
	Data Form `json:"data"`
}

type rootForms struct {
	Data  []Form `json:"data"`
	Links Links  `json:"links"`
	Meta  Meta   `json:"meta"`
}

type Form struct {
	Id                 string                 `json:"id"`
	Type               string                 `json:"type"`
	Slug               string                 `json:"slug"`
	Name               string                 `json:"name"`
	CreatedAt          string                 `json:"created_at"`
	ConversionsCount   int                    `json:"conversions_count"`
	ConversionsRate    ConversionRate         `json:"conversions_rate"`
	OpensCount         int                    `json:"opens_count"`
	Settings           map[string]interface{} `json:"settings"`
	LastRegistrationAt interface{}            `json:"last_registration_at"`
	Active             bool                   `json:"active"`
	IsBroken           bool                   `json:"is_broken"`
	HasContent         bool                   `json:"has_content"`
	Can                Can                    `json:"can"`
	UsedInAutomations  bool                   `json:"used_in_automations"`
	Warnings           []interface{}          `json:"warnings"`
	DoubleOptin        interface{}            `json:"double_optin"`
	ScreenshotUrl      interface{}            `json:"screenshot_url"`
}

type ConversionRate struct {
	Float  int    `json:"float"`
	String string `json:"string"`
}

type Can struct {
	Update bool `json:"update"`
}

// ListFormOptions - modifies the behavior of FormService.List method
type ListFormOptions struct {
	Type   string  `url:"-"`
	Filter *Filter `json:"filter,omitempty"`
	Page   int     `url:"page,omitempty"`
	Limit  int     `url:"limit,omitempty"`
	Sort   string  `url:"sort,omitempty"`
}

func (s *FormService) List(ctx context.Context, options *ListFormOptions) (*rootForms, *Response, error) {
	path := fmt.Sprintf("%s/%s", formEndpoint, options.Type)
	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootForms)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *FormService) Get(ctx context.Context, formID string) (*rootForm, *Response, error) {
	path := fmt.Sprintf("%s/%s", formEndpoint, formID)
	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootForm)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *FormService) Update(ctx context.Context, formID, formName string) (*rootForm, *Response, error) {
	body := map[string]interface{}{"name": formName}
	path := fmt.Sprintf("%s/%s", formEndpoint, formID)

	req, err := s.client.newRequest(http.MethodPut, path, body)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootForm)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *FormService) Delete(ctx context.Context, formID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", formEndpoint, formID)

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

// TODO: Fix this in docs
//https://connect.mailerlite.com/api/subscribers?filter[form]={form_id}
