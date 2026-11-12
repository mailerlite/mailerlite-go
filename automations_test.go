package mailerlite_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/mailerlite/mailerlite-go"
	"github.com/stretchr/testify/assert"
)

func TestCanUpdateAutomation(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPut, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"name":"Welcome","trigger_data":{"repeatable":true},"triggers":[{"type":"manual"}]}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{}}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.Update(ctx, "1", &mailerlite.UpdateAutomationOptions{Name: "Welcome", TriggerData: map[string]interface{}{"repeatable": true}, Triggers: []map[string]interface{}{{"type": "manual"}}})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanCopyAutomation(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/copy", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{}}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.Copy(ctx, "1")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanEnableAutomation(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/enable", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{}}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.Enable(ctx, "1")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanDisableAutomation(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/disable", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{}}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.Disable(ctx, "1")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanTestAutomation(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/test", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"email":"test@example.com"}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusNoContent,
			Body:       io.NopCloser(strings.NewReader(``)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	res, err := client.Automation.Test(ctx, "1", &mailerlite.TestAutomationOptions{Email: "test@example.com"})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, res.StatusCode)
}

func TestCanAddSubscribersToAutomation(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/add-subscribers", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"step_id":"2","emails":["test@example.com"],"groups":["3"]}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusNoContent,
			Body:       io.NopCloser(strings.NewReader(``)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	res, err := client.Automation.AddSubscribers(ctx, "1", &mailerlite.AddSubscribersToAutomationOptions{StepID: "2", Emails: []string{"test@example.com"}, Groups: []string{"3"}})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, res.StatusCode)
}

func TestCanReaddSubscribersToAutomation(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/readd", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"readd_subscribers_to":"selected_step","selected_step_id":"2","filter":{"status":"failed","id":["3"],"search":"test@example.com","date_from":"2026-01-01","date_to":"2026-01-31","step":"4"}}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusNoContent,
			Body:       io.NopCloser(strings.NewReader(``)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	res, err := client.Automation.Readd(ctx, "1", &mailerlite.ReaddAutomationSubscribersOptions{ReaddSubscribersTo: "selected_step", SelectedStepID: "2", Filter: mailerlite.AutomationActivityFilter{Status: "failed", ID: []string{"3"}, Search: "test@example.com", DateFrom: "2026-01-01", DateTo: "2026-01-31", Step: "4"}})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, res.StatusCode)
}

func TestCanAddAutomationActivityToGroup(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/activity/add-to-group", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"destination_groups":["2"],"filter":{"status":"canceled","reason":"unsubscribed","scheduled_from":"2026-01-01","scheduled_to":"2026-01-31","next_step":"3","current_step":"4"}}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusNoContent,
			Body:       io.NopCloser(strings.NewReader(``)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	res, err := client.Automation.ActivityAddToGroup(ctx, "1", &mailerlite.AutomationActivityAddToGroupOptions{DestinationGroups: []string{"2"}, Filter: mailerlite.AutomationActivityFilter{Status: "canceled", Reason: "unsubscribed", ScheduledFrom: "2026-01-01", ScheduledTo: "2026-01-31", NextStep: "3", CurrentStep: "4"}})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, res.StatusCode)
}

func TestCanCancelAutomationActivity(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/activity/cancel-automation", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"filter":{"status":"active","id":["2"]}}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusNoContent,
			Body:       io.NopCloser(strings.NewReader(``)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	res, err := client.Automation.ActivityCancelAutomation(ctx, "1", &mailerlite.AutomationActivityOptions{Filter: mailerlite.AutomationActivityFilter{Status: "active", ID: []string{"2"}}})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, res.StatusCode)
}

func TestCanSkipAutomationActivityDelay(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/activity/skip-delay", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"filter":{"status":"active","next_step":"2"}}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusNoContent,
			Body:       io.NopCloser(strings.NewReader(``)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	res, err := client.Automation.ActivitySkipDelay(ctx, "1", &mailerlite.AutomationActivityOptions{Filter: mailerlite.AutomationActivityFilter{Status: "active", NextStep: "2"}})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, res.StatusCode)
}

func TestCanListAutomationsWithSortAndSteps(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations?filter%5Bautomation_state%5D=drafts&limit=10&page=2&sort=-name&with_steps=0", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":[]}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.List(ctx, &mailerlite.ListAutomationOptions{Page: 2, Limit: 10, Sort: "-name", WithSteps: mailerlite.Bool(false), Filters: &[]mailerlite.Filter{{Name: "automation_state", Value: "drafts"}}})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanListAutomationActivityWithSort(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/activity?filter%5Bsearch%5D=test%40example.com&filter%5Bstatus%5D=completed&limit=10&page=2&sort=-updated_at", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":[]}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.Subscribers(ctx, &mailerlite.ListAutomationSubscriberOptions{AutomationID: "1", Page: 2, Limit: 10, Sort: "-updated_at", Filters: &[]mailerlite.Filter{{Name: "status", Value: "completed"}, {Name: "search", Value: "test@example.com"}}})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanListAutomationTriggers(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/triggers", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":[]}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.ListTriggers(ctx, "1")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanCreateAutomationTrigger(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/triggers", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"type":"subscriber_joins_group","data":{"group_ids":[2]}}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{}}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.CreateTrigger(ctx, "1", &mailerlite.CreateAutomationTriggerOptions{Type: mailerlite.AutomationTriggerTypeSubscriberJoinsGroup, Data: map[string]interface{}{"group_ids": []int{2}}})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanGetAutomationTrigger(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/triggers/2", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{}}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.GetTrigger(ctx, "1", "2")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanUpdateAutomationTrigger(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPut, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/triggers/2", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"data":{"segment_id":3}}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{"trigger":{"id":"2"},"trigger_data":{},"qualified_subscribers_count":5}}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.UpdateTrigger(ctx, "1", "2", &mailerlite.UpdateAutomationTriggerOptions{Data: map[string]interface{}{"segment_id": 3}})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanDeleteAutomationTrigger(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodDelete, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/triggers/2", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusNoContent,
			Body:       io.NopCloser(strings.NewReader(``)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	res, err := client.Automation.DeleteTrigger(ctx, "1", "2")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, res.StatusCode)
}

func TestCanListAutomationSteps(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/steps/index", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":[]}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.ListSteps(ctx, "1")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanCreateAutomationStep(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/steps", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"type":"action","parent_id":"2","branch":"yes","move_to_branch":"no","data":{"action_type":"copy_to_group"}}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{}}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.CreateStep(ctx, "1", &mailerlite.CreateAutomationStepOptions{Type: mailerlite.AutomationStepTypeAction, ParentID: "2", Branch: mailerlite.AutomationBranchYes, MoveToBranch: mailerlite.AutomationBranchNo, Data: map[string]interface{}{"action_type": mailerlite.AutomationActionTypeCopyToGroup}})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanGetAutomationStep(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/steps/2", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{}}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.GetStep(ctx, "1", "2")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanUpdateAutomationStep(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPut, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/steps/2", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"parent_id":"3","data":{"unit":"days","value":1},"reschedule_waiting_subscribers":true}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{}}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.UpdateStep(ctx, "1", "2", &mailerlite.UpdateAutomationStepOptions{ParentID: "3", Data: map[string]interface{}{"unit": "days", "value": 1}, RescheduleWaitingSubscribers: true})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanDeleteAutomationStep(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodDelete, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/steps/2?branch_to_delete=all", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusNoContent,
			Body:       io.NopCloser(strings.NewReader(``)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	res, err := client.Automation.DeleteStep(ctx, "1", "2", &mailerlite.DeleteAutomationStepOptions{BranchToDelete: mailerlite.AutomationBranchAll})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, res.StatusCode)
}

func TestCanDeleteAutomationStepWithoutOptions(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodDelete, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/steps/2", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusNoContent,
			Body:       io.NopCloser(strings.NewReader(``)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	res, err := client.Automation.DeleteStep(ctx, "1", "2", nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, res.StatusCode)
}

func TestCanCopyAutomationStep(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/steps/2/copy", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"parent_id":"3","branch":"a_path","move_to_branch":"yes"}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{}}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.CopyStep(ctx, "1", "2", &mailerlite.CopyAutomationStepOptions{ParentID: "3", Branch: mailerlite.AutomationBranchAPath, MoveToBranch: mailerlite.AutomationBranchYes})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanReorderAutomationSteps(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/steps/reorder", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"dragged_step_id":"2","new_parent_step_id":"3","branch":"no","branch_for_child":"yes"}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{}}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.ReorderSteps(ctx, "1", &mailerlite.ReorderAutomationStepsOptions{DraggedStepID: "2", NewParentStepID: "3", Branch: mailerlite.AutomationBranchNo, BranchForChild: mailerlite.AutomationBranchYes})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanSelectAutomationAbTestWinner(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPut, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/steps/2/select-ab-test-winner", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"ab_test_winner":"b_path"}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{}}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.SelectAbTestWinner(ctx, "1", "2", &mailerlite.SelectAbTestWinnerOptions{AbTestWinner: mailerlite.AutomationBranchBPath})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanGetAutomationEmailContent(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/emails/2/content", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{"html":"<p>Hi</p>"}}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.GetEmailContent(ctx, "1", "2")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanUpdateAutomationEmailContent(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPut, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/emails/2/content", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{"html":"\u003cp\u003eHi\u003c/p\u003e"}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{}}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.UpdateEmailContent(ctx, "1", "2", &mailerlite.UpdateAutomationEmailContentOptions{HTML: "<p>Hi</p>"})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanDecodeAutomationEmailReportStatsRates(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{"engagement":{"sent":2,"open_rate":{"string":"50%"},"click_rate":{"float":0.25,"string":"25%"},"forward_rate":{"string":"0%"}},"deliverability":{"sent":2,"delivered_count":2,"delivered_rate":{"string":"100%"},"unsubscribe_rate":null}}}`)),
		}
	})

	client.SetHttpClient(testClient)

	stats, _, err := client.Automation.EmailReportStats(context.TODO(), &mailerlite.AutomationEmailStatsOptions{AutomationID: "1", EmailID: "2"})
	assert.NoError(t, err)
	assert.Equal(t, "50%", stats.Data.Engagement.OpenRate.String)
	assert.Equal(t, 0.25, stats.Data.Engagement.ClickRate.Float)
	assert.Equal(t, "0%", stats.Data.Engagement.ForwardRate.String)
	assert.Equal(t, "100%", stats.Data.Deliverability.DeliveredRate.String)
	assert.Nil(t, stats.Data.Deliverability.UnsubscribeRate)
}

func TestCanGetAutomationEmailReportStats(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/emails/2/report-stats?end_date=2026-01-31&shop_id=3&start_date=2026-01-01", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{}}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.EmailReportStats(ctx, &mailerlite.AutomationEmailStatsOptions{AutomationID: "1", EmailID: "2", StartDate: "2026-01-01", EndDate: "2026-01-31", ShopID: "3"})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanGetAutomationEmailEcommerceOverviewStats(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/emails/2/ecommerce-overview-stats?end_date=2026-01-31&shop_id=3&start_date=2026-01-01", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{}}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.EmailEcommerceOverviewStats(ctx, &mailerlite.AutomationEmailStatsOptions{AutomationID: "1", EmailID: "2", StartDate: "2026-01-01", EndDate: "2026-01-31", ShopID: "3"})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanGetAutomationEmailUnsubscribeReasons(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/emails/2/unsubscribe-reasons", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":[]}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.EmailUnsubscribeReasons(ctx, "1", "2")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanListAutomationTemplates(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automation-templates?filter%5Bcategory%5D=3&filter%5Btype%5D=welcome&limit=10&page=2", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":[]}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.ListTemplates(ctx, &mailerlite.ListAutomationTemplateOptions{Page: 2, Limit: 10, Filters: &[]mailerlite.Filter{{Name: "type", Value: "welcome"}, {Name: "category", Value: 3}}})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanSelectAutomationTemplate(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPost, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/select-template/5", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{}}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.SelectTemplate(ctx, "5")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanGetAutomationStepRunCounts(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/steps/run-counts", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":[]}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.StepRunCounts(ctx, "1")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanGetAutomationStepRunCount(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/steps/run-counts/2", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{}}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.StepRunCount(ctx, "1", "2")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanGetAutomationTriggerRunCounts(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/triggers/run-counts", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{"triggers":[{"id":"2","completed_runs":3,"queued_runs":4}],"manual_runs":null}}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.TriggerRunCounts(ctx, "1")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanGetAutomationTriggerRunCount(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/triggers/run-counts/2", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{}}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.TriggerRunCount(ctx, "1", "2")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanGetAutomationAbStepsStats(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/ab-steps-stats", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":[]}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.AbStepsStats(ctx, "1")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanGetAutomationAbStepStats(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/ab-steps-stats/2", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{}}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.AbStepStats(ctx, "1", "2")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanListAutomationStepSubscribers(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/activity/steps/2/subscribers?filter%5Bdate_from%5D=2026-01-01&filter%5Bsearch%5D=test%40example.com&limit=10&page=2&sort=-date&status=queued", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":[]}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.StepSubscribers(ctx, &mailerlite.ListAutomationStepSubscriberOptions{AutomationID: "1", StepID: "2", Status: "queued", Page: 2, Limit: 10, Sort: "-date", Filters: &[]mailerlite.Filter{{Name: "search", Value: "test@example.com"}, {Name: "date_from", Value: "2026-01-01"}}})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanListAutomationTriggerSubscribers(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/activity/triggers/2/subscribers?filter%5Bdate_to%5D=2026-01-31&limit=10&page=2&sort=email&status=completed", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":[]}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.TriggerSubscribers(ctx, &mailerlite.ListAutomationTriggerSubscriberOptions{AutomationID: "1", TriggerID: "2", Status: "completed", Page: 2, Limit: 10, Sort: "email", Filters: &[]mailerlite.Filter{{Name: "date_to", Value: "2026-01-31"}}})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanListAutomationOrderActivity(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/order-activity?ecommerce_shop_id=3&limit=10&page=2&time_frame=custom&time_frame_end=2026-01-31&time_frame_start=2026-01-01", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":[]}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.OrderActivity(ctx, &mailerlite.ListAutomationOrderActivityOptions{AutomationID: "1", TimeFrame: "custom", TimeFrameStart: "2026-01-01", TimeFrameEnd: "2026-01-31", EcommerceShopID: "3", Page: 2, Limit: 10})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanGetAutomationMonthlyEmailTotals(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/monthly-totals/email?end_date=2026-01-31&start_date=2026-01-01", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{}}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.MonthlyEmailTotals(ctx, &mailerlite.AutomationStatsOptions{AutomationID: "1", StartDate: "2026-01-01", EndDate: "2026-01-31"})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanGetAutomationMonthlyEcommerceTotals(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/monthly-totals/ecommerce?end_date=2026-01-31&shop_id=3&start_date=2026-01-01", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{}}`)),
		}
	})

	ctx := context.TODO()

	client.SetHttpClient(testClient)

	_, res, err := client.Automation.MonthlyEcommerceTotals(ctx, &mailerlite.AutomationStatsOptions{AutomationID: "1", StartDate: "2026-01-01", EndDate: "2026-01-31", ShopID: "3"})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCanDecodeGenericAutomationStepAndTrigger(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(strings.NewReader(`{"data":{"id":"1","name":"Welcome","triggers":[
				{"id":"2","type":"abandoned_checkout","ecommerce_shop_id":"3","unit":"hours","value":4,"broken":false,"complete":true}
			],"steps":[
				{"id":"5","type":"webhook","parent_id":null,"url":"https://example.com","complete":true,"broken":false},
				{"id":"6","type":"action","parent_id":"5","action_type":"send_internal_email_notification","complete":false,"broken":false}
			]}}`)),
		}
	})

	client.SetHttpClient(testClient)

	automation, _, err := client.Automation.Get(context.TODO(), "1")
	assert.NoError(t, err)

	trigger := automation.Data.Triggers[0]
	assert.Equal(t, mailerlite.AutomationTriggerTypeAbandonedCheckout, trigger.Type)
	assert.True(t, trigger.Complete)
	assert.Equal(t, "3", trigger.Data["ecommerce_shop_id"])
	assert.Equal(t, float64(4), trigger.Data["value"])

	webhook := automation.Data.Steps[0]
	assert.Equal(t, mailerlite.AutomationStepTypeWebhook, webhook.Type)
	assert.Equal(t, "https://example.com", webhook.Data["url"])

	action := automation.Data.Steps[1]
	assert.Equal(t, mailerlite.AutomationActionTypeSendInternalEmailNotification, action.ActionType)
	assert.Equal(t, "5", action.ParentID)
}

func TestCanDecodeAutomationAbStepStats(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(strings.NewReader(`{"data":[
				{"step_id":"1","stats":{"a_path":{"subscribers":10,"sent":8,"opens":4,"clicks":2,"open_rate":{"float":0.5,"string":"50%"}}}},
				{"step_id":"2","stats":[]}
			]}`)),
		}
	})

	client.SetHttpClient(testClient)

	stats, _, err := client.Automation.AbStepsStats(context.TODO(), "1")
	assert.NoError(t, err)
	assert.Equal(t, 8, stats.Data[0].Stats[mailerlite.AutomationBranchAPath].Sent)
	assert.Equal(t, "50%", stats.Data[0].Stats[mailerlite.AutomationBranchAPath].OpenRate.String)
	assert.Empty(t, stats.Data[1].Stats)
}

func TestCanListAutomationActivityWithSubscribers(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/activity?filter%5Bstatus%5D=active&with_subscribers=1", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":[{"id":"2","status":"In progress","subscriber":{"id":"3","email":"test@example.com","name":"Test"}}]}`)),
		}
	})

	client.SetHttpClient(testClient)

	activity, _, err := client.Automation.Subscribers(context.TODO(), &mailerlite.ListAutomationSubscriberOptions{AutomationID: "1", WithSubscribers: true, Filters: &[]mailerlite.Filter{{Name: "status", Value: "active"}}})
	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", activity.Data[0].Subscriber.Email)
	assert.Equal(t, "Test", activity.Data[0].Subscriber.Name)
}

func TestCanListAutomationActivityWithoutSubscribersOmitsParam(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/activity?filter%5Bstatus%5D=active", req.URL.String())
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":[]}`)),
		}
	})

	client.SetHttpClient(testClient)

	_, _, err := client.Automation.Subscribers(context.TODO(), &mailerlite.ListAutomationSubscriberOptions{AutomationID: "1", WithSubscribers: false, Filters: &[]mailerlite.Filter{{Name: "status", Value: "active"}}})
	assert.NoError(t, err)
}

func TestCanUpdateAutomationTriggerWithoutData(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodPut, req.Method)
		assert.Equal(t, "https://connect.mailerlite.com/api/automations/1/triggers/2", req.URL.String())
		b, _ := io.ReadAll(req.Body)
		assert.Equal(t, `{}`, strings.TrimRight(string(b), "\r\n"))
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"data":{"trigger":{"id":"2"},"trigger_data":{}}}`)),
		}
	})

	client.SetHttpClient(testClient)

	_, _, err := client.Automation.UpdateTrigger(context.TODO(), "1", "2", &mailerlite.UpdateAutomationTriggerOptions{})
	assert.NoError(t, err)
}

func TestCanDecodeAutomationTriggerFields(t *testing.T) {
	client := mailerlite.NewClient(testKey)

	testClient := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body: io.NopCloser(strings.NewReader(`{"data":[
				{"id":"1","type":"subscriber_joins_group","group_ids":["10","11"],"groups":[{"id":"10","name":"VIP","url":"https://example.com"}],"exclude_group_ids":["12"],"excluded_groups":[],"broken":false,"complete":true,"exit_when_no_longer_matches":true},
				{"id":"2","type":"subscriber_joins_segment","segment_id":"20","segment":[],"broken":false,"complete":true},
				{"id":"3","type":"clicks_link","link":"https://example.com/a","broken":false,"complete":true},
				{"id":"4","type":"form_completed","form_id":"30","form_type":"popup","form":null,"broken":true,"complete":true},
				{"id":"5","type":"field_updated","field_id":"40","field_operator":"is_set","value":null,"broken":false,"complete":true},
				{"id":"6","type":"purchased_any_product","ecommerce_shop_id":"50","ecommerce_shop":null,"broken":true,"complete":true}
			]}`)),
		}
	})

	client.SetHttpClient(testClient)

	triggers, _, err := client.Automation.ListTriggers(context.TODO(), "1")
	assert.NoError(t, err)
	assert.Equal(t, []string{"10", "11"}, triggers.Data[0].GroupIDs)
	assert.Equal(t, "VIP", triggers.Data[0].Groups[0].Name)
	assert.Equal(t, []interface{}{"12"}, triggers.Data[0].ExcludeGroupIds)
	assert.True(t, triggers.Data[0].ExitWhenNoLongerMatches)
	assert.Equal(t, "20", triggers.Data[1].SegmentID)
	assert.Equal(t, "https://example.com/a", triggers.Data[2].Link)
	assert.Equal(t, "30", triggers.Data[3].FormID)
	assert.Equal(t, "popup", triggers.Data[3].FormType)
	assert.Equal(t, "40", triggers.Data[4].FieldID)
	assert.Equal(t, "is_set", triggers.Data[4].FieldOperator)
	assert.Equal(t, "50", triggers.Data[5].EcommerceShopID)
}
