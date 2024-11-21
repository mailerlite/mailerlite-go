package mailerlite

import (
	"context"
	"fmt"
	"net/http"
)

const fieldEndpoint = "/fields"

// FieldService defines an interface for field-related operations.
type FieldService interface {
	List(ctx context.Context, options *ListFieldOptions) (*RootFields, *Response, error)
	Create(ctx context.Context, fieldName, fieldType string) (*RootField, *Response, error)
	Update(ctx context.Context, fieldID, fieldName string) (*RootField, *Response, error)
	Delete(ctx context.Context, fieldID string) (*Response, error)
}

// fieldService implements FieldService.
type fieldService struct {
	*service
}

type RootField struct {
	Data Field `json:"data"`
}

type RootFields struct {
	Data  []Field `json:"data"`
	Links Links   `json:"links"`
	Meta  Meta    `json:"meta"`
}

type Field struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
	Type string `json:"type"`
}

// ListFieldOptions - modifies the behavior of FieldService.List method
type ListFieldOptions struct {
	Filters *[]Filter `json:"filters,omitempty"`
	Page    int       `url:"page,omitempty"`
	Limit   int       `url:"limit,omitempty"`
	Sort    string    `url:"sort,omitempty"`
}

func (s *fieldService) List(ctx context.Context, options *ListFieldOptions) (*RootFields, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, fieldEndpoint, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootFields)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *fieldService) Create(ctx context.Context, fieldName, fieldType string) (*RootField, *Response, error) {
	body := map[string]interface{}{
		"name": fieldName,
		"type": fieldType,
	}
	req, err := s.client.newRequest(http.MethodPost, fieldEndpoint, body)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootField)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *fieldService) Update(ctx context.Context, fieldID, fieldName string) (*RootField, *Response, error) {
	body := map[string]interface{}{"name": fieldName}
	path := fmt.Sprintf("%s/%s", fieldEndpoint, fieldID)

	req, err := s.client.newRequest(http.MethodPut, path, body)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootField)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *fieldService) Delete(ctx context.Context, fieldID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", fieldEndpoint, fieldID)

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
