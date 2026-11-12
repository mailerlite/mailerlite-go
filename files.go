package mailerlite

import (
	"context"
	"net/http"
)

const fileEndpoint = "/files"

// FileService defines an interface for file operations.
type FileService interface {
	List(ctx context.Context, options *ListFileOptions) (*RootFiles, *Response, error)
}

// fileService implements FileService.
type fileService struct {
	*service
}

// RootFiles - files response
type RootFiles struct {
	Data  []File `json:"data"`
	Links Links  `json:"links"`
	Meta  Meta   `json:"meta"`
}

// File - a file in the account media library. Width and Height are nil for files that are not images.
type File struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	Size      int    `json:"size"`
	Width     *int   `json:"width"`
	Height    *int   `json:"height"`
	CreatedAt string `json:"created_at"`
}

// ListFileOptions - modifies the behavior of FileService.List method.
// Filters accepts "name" (partial match) and "type" (FileTypeImage, FileTypePDF or FileTypeEpub).
type ListFileOptions struct {
	Filters *[]Filter `json:"filters,omitempty"`
	Sort    string    `url:"sort,omitempty"`
	Page    int       `url:"page,omitempty"`
	Limit   int       `url:"limit,omitempty"`
}

// List - list the files of the account
func (s *fileService) List(ctx context.Context, options *ListFileOptions) (*RootFiles, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, fileEndpoint, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootFiles)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
