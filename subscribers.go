package mailerlite

import (
	"context"
	"encoding/json"
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
	ActivityLog(ctx context.Context, options *ListActivityOptions) (*RootActivityLog, *Response, error)
	GetImport(ctx context.Context, importID string) (*RootImport, *Response, error)
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
	DeletedAt      string                 `json:"deleted_at,omitempty"`
	ForgetAt       string                 `json:"forget_at,omitempty"`
}

// UpdateSubscriber - modifies the behavior of SubscriberService.Update method. Resubscribe and Notes are ignored by PUT /subscribers/{subscriber}.
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
	Resubscribe    bool                   `json:"resubscribe,omitempty"`
	Notes          []string               `json:"notes,omitempty"`
}

// ListSubscriberOptions - modifies the behavior of SubscriberService.List method
type ListSubscriberOptions struct {
	Filters *[]Filter `json:"filters,omitempty"`
	Cursor  string    `url:"cursor,omitempty"`
	Limit   int       `url:"limit,omitempty"`
	// Page has no effect in the current API version, use Cursor instead.
	Page       int    `url:"page,omitempty"`
	Sort       string `url:"sort,omitempty"`
	Include    string `url:"include,omitempty"`
	StartDate  string `url:"start_date,omitempty"`
	EndDate    string `url:"end_date,omitempty"`
	WithFields bool   `url:"with_fields,omitempty,int"`
	// WithGroups has no effect in the current API version, use Include "groups" instead.
	WithGroups bool `url:"with_groups,omitempty,int"`
}

// GetSubscriberOptions - modifies the behavior of SubscriberService.Get method
type GetSubscriberOptions struct {
	SubscriberID string `json:"id,omitempty"`
	Email        string `json:"email,omitempty"`
}

// ListActivityOptions - modifies the behavior of SubscriberService.ActivityLog method
type ListActivityOptions struct {
	SubscriberID string    `url:"-"`
	Filters      *[]Filter `json:"filters,omitempty"`
	Page         int       `url:"page,omitempty"`
	Limit        int       `url:"limit,omitempty"`
}

type RootActivityLog struct {
	Data  []ActivityEntry `json:"data"`
	Links Links           `json:"links"`
	Meta  Meta            `json:"meta"`
}

type ActivityEntry struct {
	ID          string                 `json:"id"`
	LogName     string                 `json:"log_name"`
	SubjectID   string                 `json:"subject_id"`
	SubjectType string                 `json:"subject_type"`
	Properties  map[string]interface{} `json:"properties"`
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
}

func (a *ActivityEntry) UnmarshalJSON(b []byte) error {
	type entry ActivityEntry
	var raw struct {
		entry
		Properties json.RawMessage `json:"properties"`
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	*a = ActivityEntry(raw.entry)
	if len(raw.Properties) == 0 || isEmptyJSONArray(raw.Properties) {
		return nil
	}

	return json.Unmarshal(raw.Properties, &a.Properties)
}

type RootImport struct {
	Data Import `json:"data"`
}

type ImportEntry struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type Import struct {
	ID                                     string        `json:"id"`
	AccountID                              string        `json:"account_id"`
	Total                                  int           `json:"total"`
	Processed                              int           `json:"processed"`
	Imported                               int           `json:"imported"`
	Updated                                int           `json:"updated"`
	Errored                                int           `json:"errored"`
	Percent                                float64       `json:"percent"`
	Done                                   bool          `json:"done"`
	FilePath                               string        `json:"file_path"`
	Invalid                                []ImportEntry `json:"invalid"`
	InvalidCount                           int           `json:"invalid_count"`
	Mistyped                               []ImportEntry `json:"mistyped"`
	MistypedCount                          int           `json:"mistyped_count"`
	Changed                                []ImportEntry `json:"changed"`
	ChangedCount                           int           `json:"changed_count"`
	Unchanged                              []ImportEntry `json:"unchanged"`
	UnchangedCount                         int           `json:"unchanged_count"`
	Unsubscribed                           []ImportEntry `json:"unsubscribed"`
	UnsubscribedCount                      int           `json:"unsubscribed_count"`
	RoleBased                              []ImportEntry `json:"role_based"`
	RoleBasedCount                         int           `json:"role_based_count"`
	SuspiciousFormat                       []ImportEntry `json:"suspicious_format"`
	SuspiciousFormatCount                  int           `json:"suspicious_format_count"`
	Duplicate                              []ImportEntry `json:"duplicate"`
	DuplicateCount                         int           `json:"duplicate_count"`
	SubscriberLimit                        []ImportEntry `json:"subscriber_limit"`
	SubscriberLimitCount                   int           `json:"subscriber_limit_count"`
	BannedImportEmailsCount                int           `json:"banned_import_emails_count"`
	MatchRoute                             string        `json:"match_route"`
	SourceLabel                            string        `json:"source_label"`
	UpdatedAt                              string        `json:"updated_at"`
	UndoneAt                               interface{}   `json:"undone_at"`
	StoppedAt                              interface{}   `json:"stopped_at"`
	UndoStartedAt                          interface{}   `json:"undo_started_at"`
	FinishedAt                             string        `json:"finished_at"`
	IsStoppedDueToFreeAccountLimitExceeded bool          `json:"is_stopped_due_to_free_account_limit_exceeded"`
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

func (s *subscriberService) ActivityLog(ctx context.Context, options *ListActivityOptions) (*RootActivityLog, *Response, error) {
	path := fmt.Sprintf("%s/%s/activity-log", subscriberEndpoint, options.SubscriberID)

	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootActivityLog)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *subscriberService) GetImport(ctx context.Context, importID string) (*RootImport, *Response, error) {
	path := fmt.Sprintf("%s/import/%s", subscriberEndpoint, importID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootImport)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
