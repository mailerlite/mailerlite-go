package mailerlite

import (
	"context"
	"fmt"
	"net/http"
)

const automationEndpoint = "/automations"

type AutomationService service

type rootAutomation struct {
	Data Automation `json:"data"`
}

type rootAutomations struct {
	Data  []Automation `json:"data"`
	Links Links        `json:"links"`
	Meta  Meta         `json:"meta"`
}

type rootAutomationsSubscriber struct {
	Data  []AutomationSubscriber `json:"data"`
	Links Links                  `json:"links"`
	Meta  Meta                   `json:"meta"`
}

type Automation struct {
	ID                        string          `json:"id"`
	Name                      string          `json:"name"`
	Enabled                   bool            `json:"enabled"`
	TriggerData               TriggerData     `json:"trigger_data"`
	Steps                     []Step          `json:"steps"`
	Triggers                  []Triggers      `json:"triggers"`
	Complete                  bool            `json:"complete"`
	Broken                    bool            `json:"broken"`
	Warnings                  []interface{}   `json:"warnings"`
	EmailsCount               int             `json:"emails_count"`
	FirstEmailScreenshotURL   interface{}     `json:"first_email_screenshot_url"`
	Stats                     AutomationStats `json:"stats"`
	CreatedAt                 string          `json:"created_at"`
	HasBannedContent          bool            `json:"has_banned_content"`
	QualifiedSubscribersCount int             `json:"qualified_subscribers_count"`
}

type TriggerData struct {
	TrackEcommerce bool `json:"track_ecommerce"`
	Repeatable     bool `json:"repeatable"`
	Valid          bool `json:"valid"`
}

type Step struct {
	ID                  string       `json:"id"`
	Type                string       `json:"type"`
	ParentID            string       `json:"parent_id"`
	Unit                string       `json:"unit,omitempty"`
	Complete            bool         `json:"complete,omitempty"`
	CreatedAt           string       `json:"created_at"`
	YesStepId           string       `json:"yes_step_id,omitempty"`
	NoStepId            string       `json:"no_step_id,omitempty"`
	Broken              bool         `json:"broken"`
	UpdatedAt           string       `json:"updated_at"`
	Value               string       `json:"value,omitempty"`
	MatchingType        string       `json:"matching_type,omitempty"`
	Description         string       `json:"description"`
	Name                string       `json:"name,omitempty"`
	Subject             string       `json:"subject,omitempty"`
	From                string       `json:"from,omitempty"`
	FromName            string       `json:"from_name,omitempty"`
	EmailID             string       `json:"email_id,omitempty"`
	Email               *Email       `json:"email,omitempty"`
	Conditions          *[]Condition `json:"conditions,omitempty"`
	LanguageID          int          `json:"language_id,omitempty"`
	TrackOpens          bool         `json:"track_opens,omitempty"`
	GoogleAnalytics     interface{}  `json:"google_analytics,omitempty"`
	TrackingWasDisabled bool         `json:"tracking_was_disabled,omitempty"`
}

type Condition struct {
	Type    string              `json:"type"`
	EmailID string              `json:"email_id"`
	Action  string              `json:"action"`
	LinkID  interface{}         `json:"link_id"`
	Email   AutomationEmailMeta `json:"email"`
}

type AutomationEmailMeta struct {
	ID   string      `json:"id"`
	Name string      `json:"name"`
	URL  interface{} `json:"url"`
}

type Triggers struct {
	ID              string              `json:"id"`
	Type            string              `json:"type"`
	GroupID         string              `json:"group_id"`
	Group           AutomationGroupMeta `json:"group"`
	ExcludeGroupIds []interface{}       `json:"exclude_group_ids"`
	ExcludedGroups  []interface{}       `json:"excluded_groups"`
	Broken          bool                `json:"broken"`
}

type AutomationGroupMeta struct {
	ID   string      `json:"id"`
	Name string      `json:"name"`
	URL  interface{} `json:"url"`
}

type AutomationStats struct {
	CompletedSubscribersCount int              `json:"completed_subscribers_count"`
	SubscribersInQueueCount   int              `json:"subscribers_in_queue_count"`
	BounceRate                *BounceRate      `json:"bounce_rate,omitempty"`
	ClickToOpenRate           *ClickToOpenRate `json:"click_to_open_rate,omitempty"`
	Sent                      int              `json:"sent"`
	OpensCount                int              `json:"opens_count"`
	UniqueOpensCount          interface{}      `json:"unique_opens_count"`
	OpenRate                  *OpenRate        `json:"open_rate,omitempty"`
	ClicksCount               int              `json:"clicks_count"`
	UniqueClicksCount         interface{}      `json:"unique_clicks_count"`
	ClickRate                 *ClickRate       `json:"click_rate,omitempty"`
	UnsubscribesCount         int              `json:"unsubscribes_count"`
	UnsubscribeRate           *UnsubscribeRate `json:"unsubscribe_rate,omitempty"`
	SpamCount                 int              `json:"spam_count"`
	SpamRate                  *SpamRate        `json:"spam_rate,omitempty"`
	HardBouncesCount          int              `json:"hard_bounces_count"`
	HardBounceRate            *HardBounceRate  `json:"hard_bounce_rate,omitempty"`
	SoftBouncesCount          int              `json:"soft_bounces_count"`
	SoftBounceRate            *SoftBounceRate  `json:"soft_bounce_rate,omitempty"`
}

type AutomationSubscriber struct {
	ID                string                   `json:"id"`
	Status            string                   `json:"status"`
	Date              string                   `json:"date"`
	Reason            interface{}              `json:"reason"`
	ReasonDescription string                   `json:"reason_description"`
	Subscriber        AutomationSubscriberMeta `json:"subscriber"`
	StepRuns          []StepRun                `json:"stepRuns"`
	NextStep          Step                     `json:"nextStep"`
	CurrentStep       Step                     `json:"currentStep"`
}

type AutomationSubscriberMeta struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type StepRun struct {
	ID           string `json:"id"`
	StepID       string `json:"step_id"`
	Description  string `json:"description"`
	ScheduledFor string `json:"scheduled_for"`
}

// ListAutomationOptions - modifies the behavior of AutomationService.List method
type ListAutomationOptions struct {
	Filters *[]Filter `json:"filters,omitempty"`
	Page    int       `url:"page,omitempty"`
	Limit   int       `url:"limit,omitempty"`
}

// ListAutomationSubscriberOptions - modifies the behavior of AutomationService.Subscribers method
type ListAutomationSubscriberOptions struct {
	AutomationID string    `url:"-"`
	Filters      *[]Filter `json:"filters,omitempty"`
	Page         int       `url:"page,omitempty"`
	Limit        int       `url:"limit,omitempty"`
}

func (s *AutomationService) List(ctx context.Context, options *ListAutomationOptions) (*rootAutomations, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, automationEndpoint, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootAutomations)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *AutomationService) Get(ctx context.Context, automationID string) (*rootAutomation, *Response, error) {
	path := fmt.Sprintf("%s/%s", automationEndpoint, automationID)
	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootAutomation)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *AutomationService) Subscribers(ctx context.Context, options *ListAutomationSubscriberOptions) (*rootAutomationsSubscriber, *Response, error) {
	path := fmt.Sprintf("%s/%s/activity", automationEndpoint, options.AutomationID)

	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootAutomationsSubscriber)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}
