package mailerlite

import (
	"context"
	"fmt"
	"net/http"
)

const campaignEndpoint = "/campaigns"

type CampaignService service

// rootCampaigns - campaigns response
type rootCampaigns struct {
	Data  []Campaign `json:"data"`
	Links Links      `json:"links"`
	Meta  Meta       `json:"meta"`
}

// rootCampaign - single campaign response
type rootCampaign struct {
	Data Campaign `json:"data"`
}

type rootCampaignSubscribers struct {
	Data  []CampaignSubscriber `json:"data"`
	Links Links                `json:"links"`
	Meta  Meta                 `json:"meta"`
}

type rootCampaignLanguages struct {
	Data []CampaignLanguage
}

type Campaign struct {
	ID                         string             `json:"id"`
	AccountID                  string             `json:"account_id"`
	Name                       string             `json:"name"`
	Type                       string             `json:"type"`
	Status                     string             `json:"status"`
	MissingData                []interface{}      `json:"missing_data"`
	Settings                   CampaignSettings   `json:"settings"`
	Filter                     [][]CampaignFilter `json:"filter"`
	FilterForHumans            [][]string         `json:"filter_for_humans"`
	DeliverySchedule           string             `json:"delivery_schedule"`
	LanguageID                 string             `json:"language_id"`
	CreatedAt                  string             `json:"created_at"`
	UpdatedAt                  string             `json:"updated_at"`
	ScheduledFor               string             `json:"scheduled_for"`
	QueuedAt                   string             `json:"queued_at"`
	StartedAt                  string             `json:"started_at"`
	FinishedAt                 string             `json:"finished_at"`
	StoppedAt                  interface{}        `json:"stopped_at"`
	DefaultEmailID             string             `json:"default_email_id"`
	Emails                     []Email            `json:"emails"`
	UsedInAutomations          bool               `json:"used_in_automations"`
	TypeForHumans              string             `json:"type_for_humans"`
	Stats                      Stats              `json:"stats"`
	IsStopped                  bool               `json:"is_stopped"`
	HasWinner                  interface{}        `json:"has_winner"`
	WinnerVersionForHuman      interface{}        `json:"winner_version_for_human"`
	WinnerSendingTimeForHumans interface{}        `json:"winner_sending_time_for_humans"`
	WinnerSelectedManuallyAt   interface{}        `json:"winner_selected_manually_at"`
	UsesEcommerce              bool               `json:"uses_ecommerce"`
	UsesSurvey                 bool               `json:"uses_survey"`
	CanBeScheduled             bool               `json:"can_be_scheduled"`
	Warnings                   []interface{}      `json:"warnings"`
	InitialCreatedAt           interface{}        `json:"initial_created_at"`
	IsCurrentlySendingOut      bool               `json:"is_currently_sending_out"`
}

type CampaignSettings struct {
	TrackOpens         string `json:"track_opens"`
	UseGoogleAnalytics string `json:"use_google_analytics"`
	EcommerceTracking  string `json:"ecommerce_tracking"`
}

type CampaignFilter struct {
	Operator string        `json:"operator"`
	Args     []interface{} `json:"args"`
}

type Email struct {
	ID            string      `json:"id"`
	AccountID     string      `json:"account_id"`
	EmailableID   string      `json:"emailable_id"`
	EmailableType string      `json:"emailable_type"`
	Type          string      `json:"type"`
	From          string      `json:"from"`
	FromName      string      `json:"from_name"`
	Name          string      `json:"name"`
	Subject       string      `json:"subject"`
	PlainText     string      `json:"plain_text"`
	ScreenshotURL string      `json:"screenshot_url"`
	PreviewURL    string      `json:"preview_url"`
	CreatedAt     string      `json:"created_at"`
	UpdatedAt     string      `json:"updated_at"`
	IsDesigned    bool        `json:"is_designed"`
	LanguageID    float64     `json:"language_id"`
	IsWinner      bool        `json:"is_winner"`
	Stats         Stats       `json:"stats"`
	SendAfter     interface{} `json:"send_after"`
	TrackOpens    bool        `json:"track_opens"`
}

type Stats struct {
	Sent              int             `json:"sent"`
	OpensCount        int             `json:"opens_count"`
	UniqueOpensCount  int             `json:"unique_opens_count"`
	OpenRate          OpenRate        `json:"open_rate"`
	ClicksCount       int             `json:"clicks_count"`
	UniqueClicksCount int             `json:"unique_clicks_count"`
	ClickRate         ClickRate       `json:"click_rate"`
	UnsubscribesCount int             `json:"unsubscribes_count"`
	UnsubscribeRate   UnsubscribeRate `json:"unsubscribe_rate"`
	SpamCount         int             `json:"spam_count"`
	SpamRate          SpamRate        `json:"spam_rate"`
	HardBouncesCount  int             `json:"hard_bounces_count"`
	HardBounceRate    HardBounceRate  `json:"hard_bounce_rate"`
	SoftBouncesCount  int             `json:"soft_bounces_count"`
	SoftBounceRate    SoftBounceRate  `json:"soft_bounce_rate"`
	ForwardsCount     int             `json:"forwards_count"`
	ClickToOpenRate   ClickToOpenRate `json:"click_to_open_rate"`
}

// ListCampaignOptions - modifies the behavior of CampaignService.List method
type ListCampaignOptions struct {
	Filters *[]Filter `json:"filters,omitempty"`
	Page    int       `url:"page,omitempty"`
	Limit   int       `url:"limit,omitempty"`
}

// GetCampaignOptions - modifies the behavior of CampaignService.Get method
type GetCampaignOptions struct {
	ID int `json:"id,omitempty"`
}

type UpdateCampaign CreateCampaign

// CreateCampaign - modifies the behavior of CampaignService.Create method
type CreateCampaign struct {
	Name           string          `json:"name"`
	LanguageID     int             `json:"language_id,omitempty"`
	Type           string          `json:"type"`
	Emails         []Emails        `json:"emails"`
	Groups         []string        `json:"groups,omitempty"`
	Segments       []string        `json:"segments,omitempty"`
	AbSettings     *AbSettings     `json:"ab_settings,omitempty"`
	ResendSettings *ResendSettings `json:"resend_settings"`
}

type Emails struct {
	Subject  string `json:"subject"`
	FromName string `json:"from_name"`
	From     string `json:"from"`
	Content  string `json:"content"`
}

type AbSettings struct {
	TestType        string `json:"test_type"`
	SelectWinnerBy  string `json:"select_winner_by"`
	AfterTimeAmount int    `json:"after_time_amount"`
	AfterTimeUnit   string `json:"after_time_unit"`
	TestSplit       int    `json:"test_split"`
	BValue          BValue `json:"b_value"`
}

type BValue struct {
	Subject  string `json:"subject,omitempty"`
	FromName string `json:"from_name,omitempty"`
	From     string `json:"from,omitempty"`
}

type ResendSettings struct {
	TestType       string `json:"test_type"`
	SelectWinnerBy string `json:"select_winner_by"`
	BValue         BValue `json:"b_value"`
}

// ScheduleCampaign - modifies the behavior of CampaignService.Schedule method
type ScheduleCampaign struct {
	Delivery string    `json:"delivery"`
	Schedule *Schedule `json:"schedule,omitempty"`
	Resend   *Resend   `json:"resend,omitempty"`
}

type Schedule struct {
	Date       string `json:"date"`
	Hours      string `json:"hours"`
	Minutes    string `json:"minutes"`
	TimezoneID int    `json:"timezone_id,omitempty"`
}

type Resend struct {
	Delivery   string `json:"delivery"`
	Date       string `json:"date"`
	Hours      string `json:"hours"`
	Minutes    string `json:"minutes"`
	TimezoneID int    `json:"timezone_id,omitempty"`
}

type CampaignSubscriber struct {
	ID          string     `json:"id"`
	OpensCount  int        `json:"opens_count"`
	ClicksCount int        `json:"clicks_count"`
	Subscriber  Subscriber `json:"subscriber"`
}

type CampaignLanguage struct {
	Id        string `json:"id"`
	Shortcode string `json:"shortcode"`
	Iso639    string `json:"iso639"`
	Name      string `json:"name"`
	Direction string `json:"direction"`
}

type ListCampaignSubscriberOptions struct {
	CampaignID string    `url:"-"`
	Filters    *[]Filter `json:"filters,omitempty"`
	Page       int       `url:"page,omitempty"`
	Sort       string    `url:"sort,omitempty"`
	Limit      int       `url:"limit,omitempty"`
}

// List - list of campaigns
func (s *CampaignService) List(ctx context.Context, options *ListCampaignOptions) (*rootCampaigns, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, campaignEndpoint, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootCampaigns)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// Get - get a single campaign ID
func (s *CampaignService) Get(ctx context.Context, campaignID string) (*rootCampaign, *Response, error) {
	path := fmt.Sprintf("%s/%s", campaignEndpoint, campaignID)
	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootCampaign)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *CampaignService) Create(ctx context.Context, campaign *CreateCampaign) (*rootCampaign, *Response, error) {
	req, err := s.client.newRequest(http.MethodPost, campaignEndpoint, campaign)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootCampaign)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *CampaignService) Update(ctx context.Context, campaignID string, campaign *UpdateCampaign) (*rootCampaign, *Response, error) {
	path := fmt.Sprintf("%s/%s", campaignEndpoint, campaignID)
	req, err := s.client.newRequest(http.MethodPut, path, campaign)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootCampaign)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *CampaignService) Schedule(ctx context.Context, campaignID string, campaign *ScheduleCampaign) (*rootCampaign, *Response, error) {
	path := fmt.Sprintf("%s/%s/schedule", campaignEndpoint, campaignID)
	req, err := s.client.newRequest(http.MethodPost, path, campaign)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootCampaign)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// Cancel - cancel a single campaign
func (s *CampaignService) Cancel(ctx context.Context, campaignID string) (*rootCampaign, *Response, error) {
	path := fmt.Sprintf("%s/%s/cancel", campaignEndpoint, campaignID)
	req, err := s.client.newRequest(http.MethodPost, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootCampaign)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

// Subscribers - get subscribers activity of a campaign
func (s *CampaignService) Subscribers(ctx context.Context, options *ListCampaignSubscriberOptions) (*rootCampaignSubscribers, *Response, error) {
	path := fmt.Sprintf("%s/%s/reports/subscriber-activity", campaignEndpoint, options.CampaignID)

	req, err := s.client.newRequest(http.MethodPost, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootCampaignSubscribers)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *CampaignService) Languages(ctx context.Context) (*rootCampaignLanguages, *Response, error) {
	path := fmt.Sprintf("%s/languages", campaignEndpoint)
	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootCampaignLanguages)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *CampaignService) Delete(ctx context.Context, campaignID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", campaignEndpoint, campaignID)

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
