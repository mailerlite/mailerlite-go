package mailerlite

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	automationEndpoint         = "/automations"
	automationTemplateEndpoint = "/automation-templates"
)

const (
	AutomationTriggerTypeSubscriberJoinsGroup     = "subscriber_joins_group"
	AutomationTriggerTypeSubscriberJoinsSegment   = "subscriber_joins_segment"
	AutomationTriggerTypeClicksLink               = "clicks_link"
	AutomationTriggerTypeFieldUpdated             = "field_updated"
	AutomationTriggerTypeAnniversary              = "anniversary"
	AutomationTriggerTypeExactDate                = "exact_date"
	AutomationTriggerTypeFormCompleted            = "form_completed"
	AutomationTriggerTypeManual                   = "manual"
	AutomationTriggerTypePurchasedAnyProduct      = "purchased_any_product"
	AutomationTriggerTypePurchasedSpecificProduct = "purchased_specific_product"
	AutomationTriggerTypePurchasedFromCategory    = "purchased_from_category"
	AutomationTriggerTypeAbandonedCart            = "abandoned_cart"
	AutomationTriggerTypeAbandonedCheckout        = "abandoned_checkout"
	AutomationTriggerTypePurchaseFrequency        = "purchase_frequency"

	AutomationStepTypeEmail     = "email"
	AutomationStepTypeCondition = "condition"
	AutomationStepTypeDelay     = "delay"
	AutomationStepTypeAction    = "action"
	AutomationStepTypeAbTest    = "abtest"
	AutomationStepTypeWebhook   = "webhook"

	AutomationBranchYes   = "yes"
	AutomationBranchNo    = "no"
	AutomationBranchAll   = "all"
	AutomationBranchAPath = "a_path"
	AutomationBranchBPath = "b_path"
	AutomationBranchCPath = "c_path"

	AutomationActionTypeUpdateCustomField             = "update_custom_field"
	AutomationActionTypeMarkAsUnsubscribed            = "mark_as_unsubscribed"
	AutomationActionTypeRemoveFromGroup               = "remove_from_group"
	AutomationActionTypeGoToStep                      = "go_to_step"
	AutomationActionTypeCopyToGroup                   = "copy_to_group"
	AutomationActionTypeMoveToGroup                   = "move_to_group"
	AutomationActionTypeSendInternalEmailNotification = "send_internal_email_notification"
)

// AutomationService defines an interface for automation-related operations.
type AutomationService interface {
	List(ctx context.Context, options *ListAutomationOptions) (*RootAutomations, *Response, error)
	Get(ctx context.Context, automationID string) (*RootAutomation, *Response, error)
	Subscribers(ctx context.Context, options *ListAutomationSubscriberOptions) (*RootAutomationsSubscriber, *Response, error)
	Create(ctx context.Context, automationName string) (*RootAutomation, *Response, error)
	Delete(ctx context.Context, automationID string) (*Response, error)
	Update(ctx context.Context, automationID string, options *UpdateAutomationOptions) (*RootAutomation, *Response, error)
	Copy(ctx context.Context, automationID string) (*RootAutomation, *Response, error)
	Enable(ctx context.Context, automationID string) (*RootAutomation, *Response, error)
	Disable(ctx context.Context, automationID string) (*RootAutomation, *Response, error)
	Test(ctx context.Context, automationID string, options *TestAutomationOptions) (*Response, error)
	AddSubscribers(ctx context.Context, automationID string, options *AddSubscribersToAutomationOptions) (*Response, error)
	Readd(ctx context.Context, automationID string, options *ReaddAutomationSubscribersOptions) (*Response, error)
	ActivityAddToGroup(ctx context.Context, automationID string, options *AutomationActivityAddToGroupOptions) (*Response, error)
	ActivityCancelAutomation(ctx context.Context, automationID string, options *AutomationActivityOptions) (*Response, error)
	ActivitySkipDelay(ctx context.Context, automationID string, options *AutomationActivityOptions) (*Response, error)
	ListTriggers(ctx context.Context, automationID string) (*RootAutomationTriggers, *Response, error)
	CreateTrigger(ctx context.Context, automationID string, options *CreateAutomationTriggerOptions) (*RootAutomationTrigger, *Response, error)
	GetTrigger(ctx context.Context, automationID, triggerID string) (*RootAutomationTrigger, *Response, error)
	UpdateTrigger(ctx context.Context, automationID, triggerID string, options *UpdateAutomationTriggerOptions) (*RootUpdatedAutomationTrigger, *Response, error)
	DeleteTrigger(ctx context.Context, automationID, triggerID string) (*Response, error)
	ListSteps(ctx context.Context, automationID string) (*RootAutomationSteps, *Response, error)
	CreateStep(ctx context.Context, automationID string, options *CreateAutomationStepOptions) (*RootAutomationStep, *Response, error)
	GetStep(ctx context.Context, automationID, stepID string) (*RootAutomationStep, *Response, error)
	UpdateStep(ctx context.Context, automationID, stepID string, options *UpdateAutomationStepOptions) (*RootAutomationStep, *Response, error)
	DeleteStep(ctx context.Context, automationID, stepID string, options *DeleteAutomationStepOptions) (*Response, error)
	CopyStep(ctx context.Context, automationID, stepID string, options *CopyAutomationStepOptions) (*RootAutomation, *Response, error)
	ReorderSteps(ctx context.Context, automationID string, options *ReorderAutomationStepsOptions) (*RootAutomation, *Response, error)
	SelectAbTestWinner(ctx context.Context, automationID, stepID string, options *SelectAbTestWinnerOptions) (*RootAutomationAbStepStats, *Response, error)
	GetEmailContent(ctx context.Context, automationID, emailID string) (*RootAutomationEmailContent, *Response, error)
	UpdateEmailContent(ctx context.Context, automationID, emailID string, options *UpdateAutomationEmailContentOptions) (*RootAutomationEmail, *Response, error)
	EmailReportStats(ctx context.Context, options *AutomationEmailStatsOptions) (*RootAutomationEmailStats, *Response, error)
	EmailEcommerceOverviewStats(ctx context.Context, options *AutomationEmailStatsOptions) (*RootAutomationEcommerceStats, *Response, error)
	EmailUnsubscribeReasons(ctx context.Context, automationID, emailID string) (*RootAutomationUnsubscribeReasons, *Response, error)
	ListTemplates(ctx context.Context, options *ListAutomationTemplateOptions) (*RootAutomationTemplates, *Response, error)
	SelectTemplate(ctx context.Context, templateID string) (*RootAutomation, *Response, error)
	StepRunCounts(ctx context.Context, automationID string) (*RootAutomationStepRunCounts, *Response, error)
	StepRunCount(ctx context.Context, automationID, stepID string) (*RootAutomationStepRunCount, *Response, error)
	TriggerRunCounts(ctx context.Context, automationID string) (*RootAutomationTriggersRunCounts, *Response, error)
	TriggerRunCount(ctx context.Context, automationID, triggerID string) (*RootAutomationTriggerRunCount, *Response, error)
	AbStepsStats(ctx context.Context, automationID string) (*RootAutomationAbStepsStats, *Response, error)
	AbStepStats(ctx context.Context, automationID, stepID string) (*RootAutomationAbStepStats, *Response, error)
	StepSubscribers(ctx context.Context, options *ListAutomationStepSubscriberOptions) (*RootAutomationStepSubscribers, *Response, error)
	TriggerSubscribers(ctx context.Context, options *ListAutomationTriggerSubscriberOptions) (*RootAutomationStepSubscribers, *Response, error)
	OrderActivity(ctx context.Context, options *ListAutomationOrderActivityOptions) (*RootAutomationOrderActivity, *Response, error)
	MonthlyEmailTotals(ctx context.Context, options *AutomationStatsOptions) (*RootAutomationEmailStats, *Response, error)
	MonthlyEcommerceTotals(ctx context.Context, options *AutomationStatsOptions) (*RootAutomationEcommerceTotals, *Response, error)
}

type automationService struct {
	*service
}

type RootAutomation struct {
	Data Automation `json:"data"`
}

type RootAutomations struct {
	Data  []Automation `json:"data"`
	Links Links        `json:"links"`
	Meta  Meta         `json:"meta"`
}

type RootAutomationsSubscriber struct {
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
	AutomationTemplateID      interface{}     `json:"automation_template_id"`
	AutomationTemplateName    string          `json:"automation_template_name"`
	BuildingInProgress        bool            `json:"building_in_progress"`
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
	Value               interface{}  `json:"value,omitempty"`
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
	ActionType          string       `json:"action_type,omitempty"`
	TypeCount           int          `json:"type_count,omitempty"`
	// Data is the full raw object as the API returned it.
	Data map[string]interface{} `json:"-"`
}

func (s *Step) UnmarshalJSON(b []byte) error {
	type step Step
	return decodeWithRaw(b, (*step)(s), &s.Data)
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
	ID                      string                `json:"id"`
	Type                    string                `json:"type"`
	GroupID                 string                `json:"group_id"`
	Group                   AutomationGroupMeta   `json:"group"`
	ExcludeGroupIds         []interface{}         `json:"exclude_group_ids"`
	ExcludedGroups          []interface{}         `json:"excluded_groups"`
	Broken                  bool                  `json:"broken"`
	Complete                bool                  `json:"complete"`
	GroupIDs                []string              `json:"group_ids,omitempty"`
	Groups                  []AutomationGroupMeta `json:"groups,omitempty"`
	SegmentID               string                `json:"segment_id,omitempty"`
	Link                    string                `json:"link,omitempty"`
	FormID                  string                `json:"form_id,omitempty"`
	FormType                string                `json:"form_type,omitempty"`
	FieldID                 string                `json:"field_id,omitempty"`
	FieldOperator           string                `json:"field_operator,omitempty"`
	EcommerceShopID         string                `json:"ecommerce_shop_id,omitempty"`
	ExitWhenNoLongerMatches bool                  `json:"exit_when_no_longer_matches,omitempty"`
	// Data is the full raw object as the API returned it.
	Data map[string]interface{} `json:"-"`
}

func (t *Triggers) UnmarshalJSON(b []byte) error {
	type trigger Triggers
	return decodeWithRaw(b, (*trigger)(t), &t.Data)
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
	StartedRunsCount          int              `json:"started_runs_count,omitempty"`
	CancelledRunsCount        int              `json:"cancelled_runs_count,omitempty"`
	FailedRunsCount           int              `json:"failed_runs_count,omitempty"`
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
	Name  string `json:"name,omitempty"`
}

type RootAutomationTrigger struct {
	Data Triggers `json:"data"`
}

type RootAutomationTriggers struct {
	Data []Triggers `json:"data"`
}

type RootUpdatedAutomationTrigger struct {
	Data UpdatedAutomationTrigger `json:"data"`
}

type UpdatedAutomationTrigger struct {
	Trigger                   Triggers    `json:"trigger"`
	TriggerData               TriggerData `json:"trigger_data"`
	QualifiedSubscribersCount int         `json:"qualified_subscribers_count"`
}

type RootAutomationStep struct {
	Data Step `json:"data"`
}

type RootAutomationSteps struct {
	Data []Step `json:"data"`
}

type RootAutomationEmail struct {
	Data Email `json:"data"`
}

type RootAutomationEmailContent struct {
	Data AutomationEmailContent `json:"data"`
}

type AutomationEmailContent struct {
	HTML string `json:"html"`
}

type RootAutomationEmailStats struct {
	Data AutomationEmailStats `json:"data"`
}

type AutomationEmailStats struct {
	Engagement     AutomationEngagementStats     `json:"engagement"`
	Deliverability AutomationDeliverabilityStats `json:"deliverability"`
}

type AutomationEngagementStats struct {
	Sent                    int                    `json:"sent"`
	OpensCount              int                    `json:"opens_count"`
	ClicksCount             int                    `json:"clicks_count"`
	TotalOpensCount         int                    `json:"total_opens_count"`
	TotalClicksCount        int                    `json:"total_clicks_count"`
	OpenRate                *OpenRate              `json:"open_rate"`
	ClickRate               *ClickRate             `json:"click_rate"`
	ClickToOpenRate         *ClickToOpenRate       `json:"click_to_open_rate"`
	ForwardsCount           int                    `json:"forwards_count"`
	ForwardRate             *ForwardRate           `json:"forward_rate"`
	SocialInteractionsCount int                    `json:"social_interactions_count"`
	SocialInteractionRate   *SocialInteractionRate `json:"social_interaction_rate"`
}

type AutomationDeliverabilityStats struct {
	Sent              int              `json:"sent"`
	DeliveredCount    int              `json:"delivered_count"`
	DeliveredRate     *DeliveredRate   `json:"delivered_rate"`
	HardBouncesCount  int              `json:"hard_bounces_count"`
	HardBounceRate    *HardBounceRate  `json:"hard_bounce_rate"`
	SoftBouncesCount  int              `json:"soft_bounces_count"`
	SoftBounceRate    *SoftBounceRate  `json:"soft_bounce_rate"`
	SpamCount         int              `json:"spam_count"`
	SpamRate          *SpamRate        `json:"spam_rate"`
	UnsubscribesCount int              `json:"unsubscribes_count"`
	UnsubscribeRate   *UnsubscribeRate `json:"unsubscribe_rate"`
}

type RootAutomationEcommerceStats struct {
	Data AutomationEcommerceStats `json:"data"`
}

type RootAutomationEcommerceTotals struct {
	Data AutomationEcommerceTotals `json:"data"`
}

type AutomationEcommerceStats struct {
	Ecommerce AutomationEcommerce `json:"ecommerce"`
}

type AutomationEcommerceTotals struct {
	Ecommerce           AutomationEcommerce `json:"ecommerce"`
	BestPerformingEmail interface{}         `json:"best_performing_email"`
}

type AutomationEcommerce struct {
	OrdersCount       int         `json:"orders_count"`
	Revenue           interface{} `json:"revenue"`
	AverageOrderValue interface{} `json:"average_order_value"`
	ConversionRate    interface{} `json:"conversion_rate,omitempty"`
}

type RootAutomationUnsubscribeReasons struct {
	Data []AutomationUnsubscribeReason `json:"data"`
}

type AutomationUnsubscribeReason struct {
	Label      string      `json:"label"`
	Count      interface{} `json:"count"`
	Percentage interface{} `json:"percentage"`
}

type RootAutomationTemplates struct {
	Data  []AutomationTemplate `json:"data"`
	Links Links                `json:"links"`
	Meta  Meta                 `json:"meta"`
}

type AutomationTemplate struct {
	ID                       string      `json:"id"`
	Name                     string      `json:"name"`
	Type                     string      `json:"type"`
	Workflow                 interface{} `json:"workflow"`
	Description              string      `json:"description"`
	Excerpt                  string      `json:"excerpt"`
	WorkflowStepsDescription interface{} `json:"workflow_steps_description"`
	URL                      string      `json:"url"`
}

type RootAutomationStepRunCounts struct {
	Data []AutomationStepRunCount `json:"data"`
}

type RootAutomationStepRunCount struct {
	Data AutomationStepRunCount `json:"data"`
}

type AutomationStepRunCount struct {
	StepID            string `json:"step_id"`
	CompletedStepRuns int    `json:"completed_step_runs"`
	InQueueStepRuns   int    `json:"in_queue_step_runs"`
}

type RootAutomationTriggersRunCounts struct {
	Data AutomationTriggersRunCounts `json:"data"`
}

type AutomationTriggersRunCounts struct {
	Triggers   []AutomationTriggerRunCount `json:"triggers"`
	ManualRuns *AutomationManualRunCount   `json:"manual_runs"`
}

type AutomationManualRunCount struct {
	InProgressRuns int `json:"in_progress_runs"`
	CompletedRuns  int `json:"completed_runs"`
}

type RootAutomationTriggerRunCount struct {
	Data AutomationTriggerRunCount `json:"data"`
}

type AutomationTriggerRunCount struct {
	ID            string `json:"id,omitempty"`
	CompletedRuns int    `json:"completed_runs"`
	QueuedRuns    int    `json:"queued_runs"`
}

type RootAutomationAbStepsStats struct {
	Data []AutomationAbStepStats `json:"data"`
}

type RootAutomationAbStepStats struct {
	Data AutomationAbStepStats `json:"data"`
}

type AutomationAbStepStats struct {
	StepID           string                `json:"step_id"`
	Name             string                `json:"name"`
	AbTestWinner     string                `json:"ab_test_winner"`
	ParentID         string                `json:"parent_id"`
	CanSelectWinner  bool                  `json:"can_select_winner"`
	APathStepID      string                `json:"a_path_step_id"`
	BPathStepID      string                `json:"b_path_step_id"`
	CPathStepID      string                `json:"c_path_step_id"`
	SplitPercentages interface{}           `json:"split_percentages"`
	Stats            AutomationAbTestStats `json:"stats"`
}

type AutomationAbTestStats map[string]AutomationAbPathStats

type AutomationAbPathStats struct {
	Subscribers     int             `json:"subscribers"`
	Sent            int             `json:"sent"`
	Opens           int             `json:"opens"`
	Clicks          int             `json:"clicks"`
	OpenRate        OpenRate        `json:"open_rate"`
	ClickRate       ClickRate       `json:"click_rate"`
	ClickToOpenRate ClickToOpenRate `json:"click_to_open_rate"`
}

type RootAutomationStepSubscribers struct {
	Data  []AutomationStepSubscriber `json:"data"`
	Links Links                      `json:"links"`
	Meta  Meta                       `json:"meta"`
}

type AutomationStepSubscriber struct {
	ID         string                   `json:"id"`
	Subscriber AutomationSubscriberMeta `json:"subscriber"`
	Date       string                   `json:"date"`
}

type RootAutomationOrderActivity struct {
	Data  []AutomationOrder `json:"data"`
	Links Links             `json:"links"`
	Meta  Meta              `json:"meta"`
}

type AutomationOrder struct {
	Subscriber AutomationSubscriberMeta `json:"subscriber"`
	Total      float64                  `json:"total"`
	UpdatedAt  string                   `json:"updated_at"`
	Order      string                   `json:"order"`
	Email      *Email                   `json:"email"`
}

func (s *AutomationAbTestStats) UnmarshalJSON(b []byte) error {
	if isEmptyJSONArray(b) {
		*s = AutomationAbTestStats{}
		return nil
	}

	return json.Unmarshal(b, (*map[string]AutomationAbPathStats)(s))
}

type StepRun struct {
	ID           string `json:"id"`
	StepID       string `json:"step_id"`
	Description  string `json:"description"`
	ScheduledFor string `json:"scheduled_for"`
}

// ListAutomationOptions - modifies the behavior of AutomationService.List method
type ListAutomationOptions struct {
	Filters   *[]Filter `json:"filters,omitempty"`
	Page      int       `url:"page,omitempty"`
	Limit     int       `url:"limit,omitempty"`
	Sort      string    `url:"sort,omitempty"`
	WithSteps *bool     `url:"with_steps,omitempty,int"`
}

// ListAutomationSubscriberOptions - modifies the behavior of AutomationService.Subscribers method
type ListAutomationSubscriberOptions struct {
	AutomationID    string    `url:"-"`
	Filters         *[]Filter `json:"filters,omitempty"`
	Page            int       `url:"page,omitempty"`
	Limit           int       `url:"limit,omitempty"`
	Sort            string    `url:"sort,omitempty"`
	WithSubscribers bool      `url:"with_subscribers,omitempty,int"`
}

// UpdateAutomationOptions - modifies the behavior of AutomationService.Update method
type UpdateAutomationOptions struct {
	Name        string                   `json:"name"`
	TriggerData map[string]interface{}   `json:"trigger_data,omitempty"`
	Triggers    []map[string]interface{} `json:"triggers,omitempty"`
}

// TestAutomationOptions - modifies the behavior of AutomationService.Test method
type TestAutomationOptions struct {
	Email string `json:"email"`
}

// AddSubscribersToAutomationOptions - modifies the behavior of AutomationService.AddSubscribers method
type AddSubscribersToAutomationOptions struct {
	StepID string   `json:"step_id,omitempty"`
	Emails []string `json:"emails,omitempty"`
	Groups []string `json:"groups,omitempty"`
}

// AutomationActivityFilter - selects the automation runs an activity action applies to
type AutomationActivityFilter struct {
	Status        string   `json:"status"`
	ID            []string `json:"id,omitempty"`
	Search        string   `json:"search,omitempty"`
	Reason        string   `json:"reason,omitempty"`
	DateFrom      string   `json:"date_from,omitempty"`
	DateTo        string   `json:"date_to,omitempty"`
	ScheduledFrom string   `json:"scheduled_from,omitempty"`
	ScheduledTo   string   `json:"scheduled_to,omitempty"`
	Step          string   `json:"step,omitempty"`
	NextStep      string   `json:"next_step,omitempty"`
	CurrentStep   string   `json:"current_step,omitempty"`
}

// ReaddAutomationSubscribersOptions - modifies the behavior of AutomationService.Readd method
type ReaddAutomationSubscribersOptions struct {
	ReaddSubscribersTo string                   `json:"readd_subscribers_to"`
	SelectedStepID     string                   `json:"selected_step_id,omitempty"`
	Filter             AutomationActivityFilter `json:"filter"`
}

// AutomationActivityAddToGroupOptions - modifies the behavior of AutomationService.ActivityAddToGroup method
type AutomationActivityAddToGroupOptions struct {
	DestinationGroups []string                 `json:"destination_groups,omitempty"`
	Filter            AutomationActivityFilter `json:"filter"`
}

// AutomationActivityOptions - modifies the behavior of AutomationService.ActivityCancelAutomation and AutomationService.ActivitySkipDelay methods
type AutomationActivityOptions struct {
	Filter AutomationActivityFilter `json:"filter"`
}

// CreateAutomationTriggerOptions - modifies the behavior of AutomationService.CreateTrigger method
type CreateAutomationTriggerOptions struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data,omitempty"`
}

// UpdateAutomationTriggerOptions - modifies the behavior of AutomationService.UpdateTrigger method
type UpdateAutomationTriggerOptions struct {
	Data map[string]interface{} `json:"data,omitempty"`
}

// CreateAutomationStepOptions - modifies the behavior of AutomationService.CreateStep method
type CreateAutomationStepOptions struct {
	Type         string                 `json:"type"`
	ParentID     string                 `json:"parent_id,omitempty"`
	Branch       string                 `json:"branch,omitempty"`
	MoveToBranch string                 `json:"move_to_branch,omitempty"`
	Data         map[string]interface{} `json:"data,omitempty"`
}

// UpdateAutomationStepOptions - modifies the behavior of AutomationService.UpdateStep method
type UpdateAutomationStepOptions struct {
	ParentID                     string                 `json:"parent_id,omitempty"`
	Data                         map[string]interface{} `json:"data,omitempty"`
	RescheduleWaitingSubscribers bool                   `json:"reschedule_waiting_subscribers,omitempty"`
}

// DeleteAutomationStepOptions - modifies the behavior of AutomationService.DeleteStep method
type DeleteAutomationStepOptions struct {
	BranchToDelete string `url:"branch_to_delete,omitempty"`
}

// CopyAutomationStepOptions - modifies the behavior of AutomationService.CopyStep method
type CopyAutomationStepOptions struct {
	ParentID     string `json:"parent_id,omitempty"`
	Branch       string `json:"branch,omitempty"`
	MoveToBranch string `json:"move_to_branch,omitempty"`
}

// ReorderAutomationStepsOptions - modifies the behavior of AutomationService.ReorderSteps method
type ReorderAutomationStepsOptions struct {
	DraggedStepID   string `json:"dragged_step_id"`
	NewParentStepID string `json:"new_parent_step_id,omitempty"`
	Branch          string `json:"branch,omitempty"`
	BranchForChild  string `json:"branch_for_child,omitempty"`
}

// SelectAbTestWinnerOptions - modifies the behavior of AutomationService.SelectAbTestWinner method
type SelectAbTestWinnerOptions struct {
	AbTestWinner string `json:"ab_test_winner"`
}

// UpdateAutomationEmailContentOptions - modifies the behavior of AutomationService.UpdateEmailContent method
type UpdateAutomationEmailContentOptions struct {
	HTML string `json:"html"`
}

// AutomationEmailStatsOptions - modifies the behavior of AutomationService.EmailReportStats and AutomationService.EmailEcommerceOverviewStats methods
type AutomationEmailStatsOptions struct {
	AutomationID string `url:"-"`
	EmailID      string `url:"-"`
	StartDate    string `url:"start_date,omitempty"`
	EndDate      string `url:"end_date,omitempty"`
	ShopID       string `url:"shop_id,omitempty"`
}

// AutomationStatsOptions - modifies the behavior of AutomationService.MonthlyEmailTotals and AutomationService.MonthlyEcommerceTotals methods
type AutomationStatsOptions struct {
	AutomationID string `url:"-"`
	StartDate    string `url:"start_date,omitempty"`
	EndDate      string `url:"end_date,omitempty"`
	ShopID       string `url:"shop_id,omitempty"`
}

// ListAutomationTemplateOptions - modifies the behavior of AutomationService.ListTemplates method
type ListAutomationTemplateOptions struct {
	Filters *[]Filter `json:"filters,omitempty"`
	Page    int       `url:"page,omitempty"`
	Limit   int       `url:"limit,omitempty"`
}

// ListAutomationStepSubscriberOptions - modifies the behavior of AutomationService.StepSubscribers method
type ListAutomationStepSubscriberOptions struct {
	AutomationID string    `url:"-"`
	StepID       string    `url:"-"`
	Status       string    `url:"status,omitempty"`
	Filters      *[]Filter `json:"filters,omitempty"`
	Page         int       `url:"page,omitempty"`
	Limit        int       `url:"limit,omitempty"`
	Sort         string    `url:"sort,omitempty"`
}

// ListAutomationTriggerSubscriberOptions - modifies the behavior of AutomationService.TriggerSubscribers method
type ListAutomationTriggerSubscriberOptions struct {
	AutomationID string    `url:"-"`
	TriggerID    string    `url:"-"`
	Status       string    `url:"status,omitempty"`
	Filters      *[]Filter `json:"filters,omitempty"`
	Page         int       `url:"page,omitempty"`
	Limit        int       `url:"limit,omitempty"`
	Sort         string    `url:"sort,omitempty"`
}

// ListAutomationOrderActivityOptions - modifies the behavior of AutomationService.OrderActivity method
type ListAutomationOrderActivityOptions struct {
	AutomationID    string `url:"-"`
	TimeFrame       string `url:"time_frame,omitempty"`
	TimeFrameStart  string `url:"time_frame_start,omitempty"`
	TimeFrameEnd    string `url:"time_frame_end,omitempty"`
	EcommerceShopID string `url:"ecommerce_shop_id,omitempty"`
	Page            int    `url:"page,omitempty"`
	Limit           int    `url:"limit,omitempty"`
}

func (s *automationService) List(ctx context.Context, options *ListAutomationOptions) (*RootAutomations, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, automationEndpoint, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomations)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) Get(ctx context.Context, automationID string) (*RootAutomation, *Response, error) {
	path := fmt.Sprintf("%s/%s", automationEndpoint, automationID)
	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomation)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) Subscribers(ctx context.Context, options *ListAutomationSubscriberOptions) (*RootAutomationsSubscriber, *Response, error) {
	path := fmt.Sprintf("%s/%s/activity", automationEndpoint, options.AutomationID)

	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomationsSubscriber)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) Create(ctx context.Context, automationName string) (*RootAutomation, *Response, error) {
	body := map[string]interface{}{"name": automationName}
	req, err := s.client.newRequest(http.MethodPost, automationEndpoint, body)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomation)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) Delete(ctx context.Context, automationID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", automationEndpoint, automationID)

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

func (s *automationService) Update(ctx context.Context, automationID string, options *UpdateAutomationOptions) (*RootAutomation, *Response, error) {
	path := automationPath(automationID)

	req, err := s.client.newRequest(http.MethodPut, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomation)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) Copy(ctx context.Context, automationID string) (*RootAutomation, *Response, error) {
	path := automationPath(automationID, "copy")

	req, err := s.client.newRequest(http.MethodPost, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomation)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) Enable(ctx context.Context, automationID string) (*RootAutomation, *Response, error) {
	path := automationPath(automationID, "enable")

	req, err := s.client.newRequest(http.MethodPost, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomation)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) Disable(ctx context.Context, automationID string) (*RootAutomation, *Response, error) {
	path := automationPath(automationID, "disable")

	req, err := s.client.newRequest(http.MethodPost, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomation)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) Test(ctx context.Context, automationID string, options *TestAutomationOptions) (*Response, error) {
	path := automationPath(automationID, "test")

	req, err := s.client.newRequest(http.MethodPost, path, options)
	if err != nil {
		return nil, err
	}

	res, err := s.client.do(ctx, req, nil)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *automationService) AddSubscribers(ctx context.Context, automationID string, options *AddSubscribersToAutomationOptions) (*Response, error) {
	path := automationPath(automationID, "add-subscribers")

	req, err := s.client.newRequest(http.MethodPost, path, options)
	if err != nil {
		return nil, err
	}

	res, err := s.client.do(ctx, req, nil)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *automationService) Readd(ctx context.Context, automationID string, options *ReaddAutomationSubscribersOptions) (*Response, error) {
	path := automationPath(automationID, "readd")

	req, err := s.client.newRequest(http.MethodPost, path, options)
	if err != nil {
		return nil, err
	}

	res, err := s.client.do(ctx, req, nil)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *automationService) ActivityAddToGroup(ctx context.Context, automationID string, options *AutomationActivityAddToGroupOptions) (*Response, error) {
	path := automationPath(automationID, "activity", "add-to-group")

	req, err := s.client.newRequest(http.MethodPost, path, options)
	if err != nil {
		return nil, err
	}

	res, err := s.client.do(ctx, req, nil)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *automationService) ActivityCancelAutomation(ctx context.Context, automationID string, options *AutomationActivityOptions) (*Response, error) {
	path := automationPath(automationID, "activity", "cancel-automation")

	req, err := s.client.newRequest(http.MethodPost, path, options)
	if err != nil {
		return nil, err
	}

	res, err := s.client.do(ctx, req, nil)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *automationService) ActivitySkipDelay(ctx context.Context, automationID string, options *AutomationActivityOptions) (*Response, error) {
	path := automationPath(automationID, "activity", "skip-delay")

	req, err := s.client.newRequest(http.MethodPost, path, options)
	if err != nil {
		return nil, err
	}

	res, err := s.client.do(ctx, req, nil)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *automationService) ListTriggers(ctx context.Context, automationID string) (*RootAutomationTriggers, *Response, error) {
	path := automationPath(automationID, "triggers")

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomationTriggers)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) CreateTrigger(ctx context.Context, automationID string, options *CreateAutomationTriggerOptions) (*RootAutomationTrigger, *Response, error) {
	path := automationPath(automationID, "triggers")

	req, err := s.client.newRequest(http.MethodPost, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomationTrigger)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) GetTrigger(ctx context.Context, automationID, triggerID string) (*RootAutomationTrigger, *Response, error) {
	path := automationPath(automationID, "triggers", triggerID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomationTrigger)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) UpdateTrigger(ctx context.Context, automationID, triggerID string, options *UpdateAutomationTriggerOptions) (*RootUpdatedAutomationTrigger, *Response, error) {
	path := automationPath(automationID, "triggers", triggerID)

	req, err := s.client.newRequest(http.MethodPut, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootUpdatedAutomationTrigger)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) DeleteTrigger(ctx context.Context, automationID, triggerID string) (*Response, error) {
	path := automationPath(automationID, "triggers", triggerID)

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

func (s *automationService) ListSteps(ctx context.Context, automationID string) (*RootAutomationSteps, *Response, error) {
	path := automationPath(automationID, "steps", "index")

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomationSteps)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) CreateStep(ctx context.Context, automationID string, options *CreateAutomationStepOptions) (*RootAutomationStep, *Response, error) {
	path := automationPath(automationID, "steps")

	req, err := s.client.newRequest(http.MethodPost, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomationStep)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) GetStep(ctx context.Context, automationID, stepID string) (*RootAutomationStep, *Response, error) {
	path := automationPath(automationID, "steps", stepID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomationStep)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) UpdateStep(ctx context.Context, automationID, stepID string, options *UpdateAutomationStepOptions) (*RootAutomationStep, *Response, error) {
	path := automationPath(automationID, "steps", stepID)

	req, err := s.client.newRequest(http.MethodPut, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomationStep)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) DeleteStep(ctx context.Context, automationID, stepID string, options *DeleteAutomationStepOptions) (*Response, error) {
	path, err := addOptions(automationPath(automationID, "steps", stepID), options)
	if err != nil {
		return nil, err
	}

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

func (s *automationService) CopyStep(ctx context.Context, automationID, stepID string, options *CopyAutomationStepOptions) (*RootAutomation, *Response, error) {
	path := automationPath(automationID, "steps", stepID, "copy")

	req, err := s.client.newRequest(http.MethodPost, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomation)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) ReorderSteps(ctx context.Context, automationID string, options *ReorderAutomationStepsOptions) (*RootAutomation, *Response, error) {
	path := automationPath(automationID, "steps", "reorder")

	req, err := s.client.newRequest(http.MethodPost, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomation)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) SelectAbTestWinner(ctx context.Context, automationID, stepID string, options *SelectAbTestWinnerOptions) (*RootAutomationAbStepStats, *Response, error) {
	path := automationPath(automationID, "steps", stepID, "select-ab-test-winner")

	req, err := s.client.newRequest(http.MethodPut, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomationAbStepStats)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) GetEmailContent(ctx context.Context, automationID, emailID string) (*RootAutomationEmailContent, *Response, error) {
	path := automationPath(automationID, "emails", emailID, "content")

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomationEmailContent)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) UpdateEmailContent(ctx context.Context, automationID, emailID string, options *UpdateAutomationEmailContentOptions) (*RootAutomationEmail, *Response, error) {
	path := automationPath(automationID, "emails", emailID, "content")

	req, err := s.client.newRequest(http.MethodPut, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomationEmail)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) EmailReportStats(ctx context.Context, options *AutomationEmailStatsOptions) (*RootAutomationEmailStats, *Response, error) {
	path := automationPath(options.AutomationID, "emails", options.EmailID, "report-stats")

	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomationEmailStats)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) EmailEcommerceOverviewStats(ctx context.Context, options *AutomationEmailStatsOptions) (*RootAutomationEcommerceStats, *Response, error) {
	path := automationPath(options.AutomationID, "emails", options.EmailID, "ecommerce-overview-stats")

	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomationEcommerceStats)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) EmailUnsubscribeReasons(ctx context.Context, automationID, emailID string) (*RootAutomationUnsubscribeReasons, *Response, error) {
	path := automationPath(automationID, "emails", emailID, "unsubscribe-reasons")

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomationUnsubscribeReasons)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) ListTemplates(ctx context.Context, options *ListAutomationTemplateOptions) (*RootAutomationTemplates, *Response, error) {
	req, err := s.client.newRequest(http.MethodGet, automationTemplateEndpoint, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomationTemplates)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) SelectTemplate(ctx context.Context, templateID string) (*RootAutomation, *Response, error) {
	path := fmt.Sprintf("%s/select-template/%s", automationEndpoint, templateID)

	req, err := s.client.newRequest(http.MethodPost, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomation)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) StepRunCounts(ctx context.Context, automationID string) (*RootAutomationStepRunCounts, *Response, error) {
	path := automationPath(automationID, "steps", "run-counts")

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomationStepRunCounts)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) StepRunCount(ctx context.Context, automationID, stepID string) (*RootAutomationStepRunCount, *Response, error) {
	path := automationPath(automationID, "steps", "run-counts", stepID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomationStepRunCount)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) TriggerRunCounts(ctx context.Context, automationID string) (*RootAutomationTriggersRunCounts, *Response, error) {
	path := automationPath(automationID, "triggers", "run-counts")

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomationTriggersRunCounts)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) TriggerRunCount(ctx context.Context, automationID, triggerID string) (*RootAutomationTriggerRunCount, *Response, error) {
	path := automationPath(automationID, "triggers", "run-counts", triggerID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomationTriggerRunCount)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) AbStepsStats(ctx context.Context, automationID string) (*RootAutomationAbStepsStats, *Response, error) {
	path := automationPath(automationID, "ab-steps-stats")

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomationAbStepsStats)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) AbStepStats(ctx context.Context, automationID, stepID string) (*RootAutomationAbStepStats, *Response, error) {
	path := automationPath(automationID, "ab-steps-stats", stepID)

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomationAbStepStats)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) StepSubscribers(ctx context.Context, options *ListAutomationStepSubscriberOptions) (*RootAutomationStepSubscribers, *Response, error) {
	path := automationPath(options.AutomationID, "activity", "steps", options.StepID, "subscribers")

	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomationStepSubscribers)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) TriggerSubscribers(ctx context.Context, options *ListAutomationTriggerSubscriberOptions) (*RootAutomationStepSubscribers, *Response, error) {
	path := automationPath(options.AutomationID, "activity", "triggers", options.TriggerID, "subscribers")

	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomationStepSubscribers)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) OrderActivity(ctx context.Context, options *ListAutomationOrderActivityOptions) (*RootAutomationOrderActivity, *Response, error) {
	path := automationPath(options.AutomationID, "order-activity")

	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomationOrderActivity)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) MonthlyEmailTotals(ctx context.Context, options *AutomationStatsOptions) (*RootAutomationEmailStats, *Response, error) {
	path := automationPath(options.AutomationID, "monthly-totals", "email")

	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomationEmailStats)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func (s *automationService) MonthlyEcommerceTotals(ctx context.Context, options *AutomationStatsOptions) (*RootAutomationEcommerceTotals, *Response, error) {
	path := automationPath(options.AutomationID, "monthly-totals", "ecommerce")

	req, err := s.client.newRequest(http.MethodGet, path, options)
	if err != nil {
		return nil, nil, err
	}

	root := new(RootAutomationEcommerceTotals)
	res, err := s.client.do(ctx, req, root)
	if err != nil {
		return nil, res, err
	}

	return root, res, nil
}

func automationPath(automationID string, parts ...string) string {
	return strings.Join(append([]string{automationEndpoint, automationID}, parts...), "/")
}
