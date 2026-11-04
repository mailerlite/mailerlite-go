<a href="https://www.mailerlite.com"><img src="https://app.mailerlite.com/assets/images/logo-color.png" width="200px"/></a>

MailerLite Golang SDK

[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE) [![Go Reference](https://pkg.go.dev/badge/github.com/mailerlite/mailerlite-go.svg)](https://pkg.go.dev/github.com/mailerlite/mailerlite-go)

# Table of Contents

- [Installation](#installation)
- [Usage](#usage)
    - [Subscribers](#subscribers)
        - [Get a list of subscribers](#get-a-list-of-subscribers)
        - [Get a single subscriber](#get-a-single-subscriber)
        - [Count all subscribers](#count-all-subscribers)
        - [Create a subscriber](#createupsert-a-subscriber)
        - [Update a subscriber](#update-a-subscriber)
        - [Delete a subscriber](#delete-a-subscriber)
        - [Fetch subscriber activity](#fetch-subscriber-activity)
        - [Get single import](#get-single-import)
    - [Groups](#groups)
        - [Get a list of groups](#get-a-list-of-groups)
        - [Create a group](#create-a-group)
        - [Update a group](#update-a-group)
        - [Delete a group](#delete-a-group)
        - [Get subscribers belonging to a group](#get-subscribers-belonging-to-a-group)
        - [Assign subscriber to a group](#assign-subscribers-to-a-group)
        - [Unassign subscriber from a group](#unassign-subscriber-from-a-group)
        - [Import subscribers to a group](#import-subscribers-to-a-group)
    - [Segments](#segments)
        - [Get a list of segments](#get-a-list-of-segments)
        - [Update a segment](#update-a-segment)
        - [Update a segment name and filter](#update-a-segment-name-and-filter)
        - [Delete a segment](#delete-a-segment)
        - [Get subscribers belonging to a segment](#get-subscribers-belonging-to-a-segment)
    - [Fields](#fields)
        - [Get a list of fields](#get-a-list-of-fields)
        - [Create a field](#create-a-field)
        - [Update a field](#update-a-field)
        - [Delete a field](#delete-a-field)
    - [Automations](#automations)
        - [Get a list of automations](#get-a-list-of-automations)
        - [Get an automation](#get-an-automation)
        - [Get subscribers activity for an automation](#get-subscribers-activity-for-an-automation)
        - [Create a draft automation](#create-a-draft-automation)
        - [Delete an automation](#delete-an-automation)
        - [Update an automation](#update-an-automation)
        - [Copy an automation](#copy-an-automation)
        - [Enable an automation](#enable-an-automation)
        - [Disable an automation](#disable-an-automation)
        - [Send a test of an automation](#send-a-test-of-an-automation)
        - [Add subscribers to an automation](#add-subscribers-to-an-automation)
        - [Re-add subscribers to an automation](#re-add-subscribers-to-an-automation)
        - [Copy automation activity subscribers to a group](#copy-automation-activity-subscribers-to-a-group)
        - [Cancel active automation runs](#cancel-active-automation-runs)
        - [Skip the delay of active automation runs](#skip-the-delay-of-active-automation-runs)
        - [Get a list of automation triggers](#get-a-list-of-automation-triggers)
        - [Create an automation trigger](#create-an-automation-trigger)
        - [Get an automation trigger](#get-an-automation-trigger)
        - [Update an automation trigger](#update-an-automation-trigger)
        - [Delete an automation trigger](#delete-an-automation-trigger)
        - [Get a list of automation steps](#get-a-list-of-automation-steps)
        - [Create an automation step](#create-an-automation-step)
        - [Get an automation step](#get-an-automation-step)
        - [Update an automation step](#update-an-automation-step)
        - [Delete an automation step](#delete-an-automation-step)
        - [Copy an automation step](#copy-an-automation-step)
        - [Reorder automation steps](#reorder-automation-steps)
        - [Select an automation A/B test winner](#select-an-automation-ab-test-winner)
        - [Get an automation email content](#get-an-automation-email-content)
        - [Update an automation email content](#update-an-automation-email-content)
        - [Get an automation email report stats](#get-an-automation-email-report-stats)
        - [Get an automation email e-commerce stats](#get-an-automation-email-e-commerce-stats)
        - [Get an automation email unsubscribe reasons](#get-an-automation-email-unsubscribe-reasons)
        - [Get a list of automation templates](#get-a-list-of-automation-templates)
        - [Create an automation from a template](#create-an-automation-from-a-template)
        - [Get automation step run counts](#get-automation-step-run-counts)
        - [Get an automation step run count](#get-an-automation-step-run-count)
        - [Get automation trigger run counts](#get-automation-trigger-run-counts)
        - [Get an automation trigger run count](#get-an-automation-trigger-run-count)
        - [Get automation A/B test stats](#get-automation-ab-test-stats)
        - [Get an automation A/B test step stats](#get-an-automation-ab-test-step-stats)
        - [Get subscribers of an automation step](#get-subscribers-of-an-automation-step)
        - [Get subscribers of an automation trigger](#get-subscribers-of-an-automation-trigger)
        - [Get automation order activity](#get-automation-order-activity)
        - [Get automation email totals](#get-automation-email-totals)
        - [Get automation e-commerce totals](#get-automation-e-commerce-totals)
    - [Campaigns](#campaigns)
        - [Get a list of campaigns](#get-a-list-of-campaigns)
        - [Get a campaign](#get-a-campaign)
        - [Create a campaign](#create-a-campaign)
        - [Update a campaign](#update-a-campaign)
        - [Schedule a campaign](#schedule-a-campaign)
        - [Cancel a ready campaign](#cancel-a-ready-campaign)
        - [Delete a campaign](#delete-a-campaign)
        - [Get subscribers activity for a campaign](#get-subscribers-activity-for-an-campaign)
    - [Forms](#forms)
        - [Get a list of forms](#get-a-list-of-forms)
        - [Get a form](#get-a-form)
        - [Update a form](#update-a-form)
        - [Delete a form](#delete-a-form)
        - [Get subscribers of a form](#get-subscribers-of-a-form)
    - [Batching](#batching)
        - [Create a new batch](#create-a-new-batch)
    - [Webhooks](#webhooks)
        - [Get a list of webhooks](#get-a-list-of-webhooks)
        - [Get a webhook](#get-a-webhook)
        - [Create a webhook](#update-a-webhook)
        - [Update a webhook](#update-a-webhook)
        - [Delete a webhook](#delete-a-webhook)
    - [Timezones](#timezones)
        - [Get a list of timezones](#get-a-list-of-timezones)
    - [Campaign languages](#languages)
        - [Get a list of languages](#get-a-list-of-languages)

# Installation
We recommend using this package with golang [modules](https://github.com/golang/go/wiki/Modules)

```
$ go get github.com/mailerlite/mailerlite-go
```

# Usage

## Subscribers

### Get a list of subscribers

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	listOptions := &mailerlite.ListSubscriberOptions{
		Limit:  200,
		Page:   0, 
		Filters: &[]mailerlite.Filter{{
			Name:  "status", 
			Value: "active",
		}},
	}

	subscribers, _, err := client.Subscriber.List(ctx, listOptions)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(subscribers.Meta.Total)
}
```

### Get a single subscriber

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	getOptions := &mailerlite.GetSubscriberOptions{
		SubscriberID: "subscriber-id",
		//Email: "client@example.com"
	}

	subscriber, _, err := client.Subscriber.Get(ctx, getOptions)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(subscriber.Data.Email)
}
```

### Count all subscribers

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	count, _, err := client.Subscriber.Count(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(count.Total)
}
```

### Create/Upsert a subscriber

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	subscriber := &mailerlite.UpsertSubscriber{
		Email: "example@example.com",
		Fields: map[string]interface{}{
			"city": "Vilnius",
		},
	}

	newSubscriber, _, err := client.Subscriber.Upsert(ctx, subscriber)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(newSubscriber.Data.Email)
}
```

### Update a subscriber

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	subscriber := &mailerlite.UpdateSubscriber{
		ID: "1",
		Email: "example@example.com",
		Fields: map[string]interface{}{
			"company": "MailerLite",
		},
	}

	newSubscriber, _, err := client.Subscriber.Update(ctx, subscriber)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(newSubscriber.Data.Email)
}
```

### Delete a subscriber

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, err := client.Subscriber.Delete(ctx, "subscriber-id")
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Subscriber Deleted")
}
```

### Fetch subscriber activity

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	listOptions := &mailerlite.ListActivityOptions{
		SubscriberID: "subscriber-id",
		Filters: &[]mailerlite.Filter{{
			Name:  "log_name",
			Value: "email_open",
		}},
		Page:  1,
		Limit: 25,
	}

	_, _, err := client.Subscriber.ActivityLog(ctx, listOptions)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get single import

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	importReport, _, err := client.Subscriber.GetImport(ctx, "import-id")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(importReport.Data.Done)
}
```

## Groups

### Get a list of groups

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	listOptions := &mailerlite.ListGroupOptions{
		Page:  1,
		Limit: 10,
		Sort: mailerlite.SortByName,
	}

	groups, _, err := client.Group.List(ctx, listOptions)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(groups.Meta.Total)
}
```

### Create a group

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Group.Create(ctx, "group-name")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Update a group

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Group.Update(ctx, "group-id", "Group Name")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Delete a group

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, err := client.Group.Delete(ctx, "69861610909337422")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get subscribers belonging to a group

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	listSubscriberOptions := &mailerlite.ListGroupSubscriberOptions{
		GroupID: "group-id",
		Filters: &[]mailerlite.Filter{{
			Name:  "status",
			Value: "active",
		}},
	}
		
	subscribers, _, err := client.Group.Subscribers(ctx, listSubscriberOptions)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(subscribers.Meta.Total)
}
```

### Assign subscriber to a group

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Group.Assign(ctx, "group-id", "subscriber-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Unassign subscriber from a group

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, err := client.Group.UnAssign(ctx, "group-id", "subscriber-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Import subscribers to a group

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	options := &mailerlite.ImportSubscribersOptions{
		Subscribers: []mailerlite.ImportSubscriber{
			{
				Email: "example@example.com",
				Fields: map[string]interface{}{
					"name": "Example",
				},
			},
			{
				Email: "another@example.com",
			},
		},
	}

	_, _, err := client.Group.ImportSubscribers(ctx, "group-id", options)
	if err != nil {
		log.Fatal(err)
	}
}
```

## Segments

### Get a list of segments

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	listOptions := &mailerlite.ListSegmentOptions{
		Page:  1,
		Limit: 10,
	}

	_, _, err := client.Segment.List(ctx, listOptions)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Update a segment

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Segment.Update(ctx, "segment-id", "Segment Name")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Update a segment name and filter

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	options := &mailerlite.UpdateSegmentOptions{
		Name: "Segment Name",
	}

	_, _, err := client.Segment.UpdateWithOptions(ctx, "segment-id", options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Delete a segment

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, err := client.Segment.Delete(ctx, "segment-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get subscribers belonging to a segment

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	listSubscriberOptions := &mailerlite.ListSegmentSubscriberOptions{
		SegmentID: "segment-id",
		Filters: &[]mailerlite.Filter{{
			Name:  "status",
			Value: "active",
		}},
	}
	
	_, _, err := client.Segment.Subscribers(ctx, listSubscriberOptions)
	if err != nil {
		log.Fatal(err)
	}
}
```

## Fields

### Get a list of fields

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	listOptions := &mailerlite.ListFieldOptions{
		Page:   1,
		Limit:  10,
		Filters: &[]mailerlite.Filter{{
			Name:  "keyword",
			Value: "name",
		}},
		Sort:   mailerlite.SortByName,
	}
	
	_, _, err := client.Field.List(ctx, listOptions)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Create a field

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	// text, number or data
	_, _, err := client.Field.Create(ctx, "field-name", "field-type")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Update a field

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Field.Update(ctx, "field-id", "Field name")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Delete a field

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, err := client.Field.Delete(ctx, "field-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

## Automations

### Get a list of automations

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	listOptions := &mailerlite.ListAutomationOptions{
		Filters: &[]mailerlite.Filter{{
			Name:  "status",
			Value: true,
		}},
		Page:  1,
		Limit: 10,
	}
	
	_, _, err := client.Automation.List(ctx, listOptions)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get an automation

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Automation.Get(ctx, "automation-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get subscribers activity for an automation

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	listOptions := &mailerlite.ListAutomationSubscriberOptions{
		Filters: &[]mailerlite.Filter{{
			Name:  "status",
			Value: "active",
		}},
		AutomationID: "automation-id",
		Page:         1,
		Limit:        10,
	}

	_, _, err := client.Automation.Subscribers(ctx, listOptions)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Create a draft automation

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Automation.Create(ctx, "Automation Name")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Delete an automation

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, err := client.Automation.Delete(ctx, "automation-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Update an automation

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	options := &mailerlite.UpdateAutomationOptions{
		Name: "Automation Name",
	}

	_, _, err := client.Automation.Update(ctx, "automation-id", options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Copy an automation

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Automation.Copy(ctx, "automation-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Enable an automation

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Automation.Enable(ctx, "automation-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Disable an automation

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Automation.Disable(ctx, "automation-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Send a test of an automation

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	options := &mailerlite.TestAutomationOptions{
		Email: "test@example.com",
	}

	_, err := client.Automation.Test(ctx, "automation-id", options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Add subscribers to an automation

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	options := &mailerlite.AddSubscribersToAutomationOptions{
		Emails: []string{"test@example.com"},
	}

	_, err := client.Automation.AddSubscribers(ctx, "automation-id", options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Re-add subscribers to an automation

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	options := &mailerlite.ReaddAutomationSubscribersOptions{
		ReaddSubscribersTo: "last_failed_step",
		Filter: mailerlite.AutomationActivityFilter{
			Status: "failed",
		},
	}

	_, err := client.Automation.Readd(ctx, "automation-id", options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Copy automation activity subscribers to a group

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	options := &mailerlite.AutomationActivityAddToGroupOptions{
		DestinationGroups: []string{"group-id"},
		Filter: mailerlite.AutomationActivityFilter{
			Status: "completed",
		},
	}

	_, err := client.Automation.ActivityAddToGroup(ctx, "automation-id", options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Cancel active automation runs

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	options := &mailerlite.AutomationActivityOptions{
		Filter: mailerlite.AutomationActivityFilter{
			Status: "active",
		},
	}

	_, err := client.Automation.ActivityCancelAutomation(ctx, "automation-id", options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Skip the delay of active automation runs

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	options := &mailerlite.AutomationActivityOptions{
		Filter: mailerlite.AutomationActivityFilter{
			Status: "active",
		},
	}

	_, err := client.Automation.ActivitySkipDelay(ctx, "automation-id", options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get a list of automation triggers

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Automation.ListTriggers(ctx, "automation-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Create an automation trigger

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	options := &mailerlite.CreateAutomationTriggerOptions{
		Type: mailerlite.AutomationTriggerTypeSubscriberJoinsGroup,
		Data: map[string]interface{}{
			"group_ids": []string{"group-id"},
		},
	}

	_, _, err := client.Automation.CreateTrigger(ctx, "automation-id", options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get an automation trigger

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Automation.GetTrigger(ctx, "automation-id", "trigger-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Update an automation trigger

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	options := &mailerlite.UpdateAutomationTriggerOptions{
		Data: map[string]interface{}{
			"group_ids": []string{"group-id"},
		},
	}

	_, _, err := client.Automation.UpdateTrigger(ctx, "automation-id", "trigger-id", options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Delete an automation trigger

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, err := client.Automation.DeleteTrigger(ctx, "automation-id", "trigger-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get a list of automation steps

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Automation.ListSteps(ctx, "automation-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Create an automation step

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	options := &mailerlite.CreateAutomationStepOptions{
		Type:     mailerlite.AutomationStepTypeDelay,
		ParentID: "parent-step-id",
		Data: map[string]interface{}{
			"unit":  "days",
			"value": 1,
		},
	}

	_, _, err := client.Automation.CreateStep(ctx, "automation-id", options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get an automation step

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Automation.GetStep(ctx, "automation-id", "step-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Update an automation step

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	options := &mailerlite.UpdateAutomationStepOptions{
		Data: map[string]interface{}{
			"unit":  "days",
			"value": 2,
		},
	}

	_, _, err := client.Automation.UpdateStep(ctx, "automation-id", "step-id", options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Delete an automation step

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	options := &mailerlite.DeleteAutomationStepOptions{
		BranchToDelete: mailerlite.AutomationBranchAll,
	}

	_, err := client.Automation.DeleteStep(ctx, "automation-id", "step-id", options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Copy an automation step

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	options := &mailerlite.CopyAutomationStepOptions{
		ParentID: "parent-step-id",
	}

	_, _, err := client.Automation.CopyStep(ctx, "automation-id", "step-id", options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Reorder automation steps

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	options := &mailerlite.ReorderAutomationStepsOptions{
		DraggedStepID:   "step-id",
		NewParentStepID: "parent-step-id",
	}

	_, _, err := client.Automation.ReorderSteps(ctx, "automation-id", options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Select an automation A/B test winner

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	options := &mailerlite.SelectAbTestWinnerOptions{
		AbTestWinner: mailerlite.AutomationBranchAPath,
	}

	_, _, err := client.Automation.SelectAbTestWinner(ctx, "automation-id", "step-id", options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get an automation email content

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Automation.GetEmailContent(ctx, "automation-id", "email-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Update an automation email content

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	options := &mailerlite.UpdateAutomationEmailContentOptions{
		HTML: "<p>Hello</p>",
	}

	_, _, err := client.Automation.UpdateEmailContent(ctx, "automation-id", "email-id", options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get an automation email report stats

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	options := &mailerlite.AutomationEmailStatsOptions{
		AutomationID: "automation-id",
		EmailID:      "email-id",
		StartDate:    "2026-01-01",
		EndDate:      "2026-01-31",
	}

	_, _, err := client.Automation.EmailReportStats(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get an automation email e-commerce stats

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	options := &mailerlite.AutomationEmailStatsOptions{
		AutomationID: "automation-id",
		EmailID:      "email-id",
	}

	_, _, err := client.Automation.EmailEcommerceOverviewStats(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get an automation email unsubscribe reasons

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Automation.EmailUnsubscribeReasons(ctx, "automation-id", "email-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get a list of automation templates

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	listOptions := &mailerlite.ListAutomationTemplateOptions{
		Filters: &[]mailerlite.Filter{{
			Name:  "type",
			Value: "welcome",
		}},
		Page:  1,
		Limit: 10,
	}

	_, _, err := client.Automation.ListTemplates(ctx, listOptions)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Create an automation from a template

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Automation.SelectTemplate(ctx, "template-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get automation step run counts

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Automation.StepRunCounts(ctx, "automation-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get an automation step run count

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Automation.StepRunCount(ctx, "automation-id", "step-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get automation trigger run counts

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Automation.TriggerRunCounts(ctx, "automation-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get an automation trigger run count

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Automation.TriggerRunCount(ctx, "automation-id", "trigger-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get automation A/B test stats

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Automation.AbStepsStats(ctx, "automation-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get an automation A/B test step stats

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Automation.AbStepStats(ctx, "automation-id", "step-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get subscribers of an automation step

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	listOptions := &mailerlite.ListAutomationStepSubscriberOptions{
		AutomationID: "automation-id",
		StepID:       "step-id",
		Status:       "queued",
		Page:         1,
		Limit:        10,
	}

	_, _, err := client.Automation.StepSubscribers(ctx, listOptions)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get subscribers of an automation trigger

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	listOptions := &mailerlite.ListAutomationTriggerSubscriberOptions{
		AutomationID: "automation-id",
		TriggerID:    "trigger-id",
		Status:       "completed",
		Page:         1,
		Limit:        10,
	}

	_, _, err := client.Automation.TriggerSubscribers(ctx, listOptions)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get automation order activity

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	listOptions := &mailerlite.ListAutomationOrderActivityOptions{
		AutomationID: "automation-id",
		TimeFrame:    "month",
		Page:         1,
		Limit:        10,
	}

	_, _, err := client.Automation.OrderActivity(ctx, listOptions)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get automation email totals

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	options := &mailerlite.AutomationStatsOptions{
		AutomationID: "automation-id",
	}

	_, _, err := client.Automation.MonthlyEmailTotals(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get automation e-commerce totals

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	options := &mailerlite.AutomationStatsOptions{
		AutomationID: "automation-id",
	}

	_, _, err := client.Automation.MonthlyEcommerceTotals(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

## Campaigns

### Get a list of campaigns

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	listOptions := &mailerlite.ListCampaignOptions{
		Filters: &[]mailerlite.Filter{{
			Name:  "status",
			Value: "draft",
		}},
		Page:  1,
		Limit: 10,
	}
	
	_, _, err := client.Campaign.List(ctx, listOptions)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get a campaign

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Campaign.Get(ctx, "campaign-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Create a campaign

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	emails := &[]mailerlite.Emails{
		{
			Subject:  "Subject",
			FromName: "Your Name",
			From:     "your@domain.com",
			Content:  "<p>This is the HTML content</p>",
		},
	}
	
	campaign := &mailerlite.CreateCampaign{
		Name:   "Campaign Name",
		Type:   mailerlite.CampaignTypeRegular,
		Emails: *emails,
	}
	
	_, _, err := client.Campaign.Create(ctx, campaign)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Update a campaign

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	emails := &[]mailerlite.Emails{
		{
			Subject:  "Subject",
			FromName: "Your Name",
			From:     "your@domain.com",
			Content:  "<p>This is the HTML content</p>",
		},
	}
	
	campaign := &mailerlite.UpdateCampaign{
		Name:   "Campaign Name",
		Type:   mailerlite.CampaignTypeRegular,
		Emails: *emails,
	}
	
	_, _, err := client.Campaign.Update(ctx, "campaign-id", campaign)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Schedule a campaign

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	schedule := &mailerlite.ScheduleCampaign{
		Delivery: mailerlite.CampaignScheduleTypeInstant,
	}
	
	_, _, err := client.Campaign.Schedule(ctx, "campaign-id", schedule)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Cancel a ready campaign

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Campaign.Cancel(ctx, "campaign-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Delete a campaign

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, err := client.Campaign.Delete(ctx, "campaign-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get subscribers activity for a campaign

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	listOptions := &mailerlite.ListCampaignSubscriberOptions{
		CampaignID: "campaign-id",
		Filters: &[]mailerlite.Filter{
			{
				Name:  "type",
				Value: "opened",
			},
			{
				Name:  "search",
				Value: "subscriber@example.com",
			},
		},
		Page:  1,
		Limit: 10,
	}

	_, _, err := client.Campaign.Subscribers(ctx, listOptions)
	if err != nil {
		log.Fatal(err)
	}
}
```

## Forms

### Get a list of forms

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	listOptions := &mailerlite.ListFormOptions{
		Type:   mailerlite.FormTypePopup,
		Page:   1,
		Limit:  10,
      	Filters: &[]mailerlite.Filter{{
      		Name:  "name",
      		Value: "Form Name",
      	}},
		Sort:   mailerlite.SortByName,
	}
	
	_, _, err := client.Form.List(ctx, listOptions)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get a form

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	form, _, err := client.Form.Get(ctx, "form-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Update a form

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Form.Update(ctx, "form-id", "Form Name")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Delete a form

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, err := client.Form.Delete(ctx, "form-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get subscribers of a form

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	listOptions := &mailerlite.ListFormSubscriberOptions{
		Page:  1,
		Limit: 10,
		Filters: &[]mailerlite.Filter{{
			Name:  "status",
			Value: "active",
		}},
	}

	_, _, err := client.Form.Subscribers(ctx, listOptions)
	if err != nil {
		log.Fatal(err)
	}
}
```

## Batching

### Create a new batch

TBC

## Webhooks

### Get a list of webhooks

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	options := &mailerlite.ListWebhookOptions{
		Sort:  mailerlite.SortByName,
		Page:  1,
		Limit: 10,
	}

	_, _, err := client.Webhook.List(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Get a webhook

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Webhook.Get(ctx, "webhook-id")
	if err != nil {
		log.Fatal(err)
	}
}
```

### Create a webhook

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	options := &mailerlite.CreateWebhookOptions{
		Name:   "",
		Events: []string{"subscriber.bounced"},
		Url:    "https://example.com/webhook",
	}
	
	_, _, err := client.Webhook.Create(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Update a webhook

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	options := &mailerlite.UpdateWebhookOptions{
		WebhookID: "webhook-id",
		Events:    []string{"subscriber.bounced", "subscriber.unsubscribed"},
		Name:      "Update",
	}
	
	_, _, err := client.Webhook.Update(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Delete a webhook

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, err := client.Webhook.Delete(ctx, "75000728688526795")
	if err != nil {
		log.Fatal(err)
	}
}
```

## Timezones

### Get a list of timezones

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Timezone.List(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
```

## Campaign languages

### Get a list of languages

```go
package main

import (
	"context"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	_, _, err := client.Campaign.Languages(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
```

# Testing

We provide interfaces for all services to help with testing

```go
type mockSubscriberService struct {
	mailerlite.SubscriberService
}

func (m *mockSubscriberService) List(ctx context.Context, options *mailerlite.ListSubscriberOptions) (*mailerlite.RootSubscribers, *mailerlite.Response, error) {
	return &mailerlite.RootSubscribers{Data: []mailerlite.Subscriber{{Email: "example@example.com"}}}, nil, nil
}

func TestListSubscribers(t *testing.T) {
	client := &mailerlite.Client{}
	client.Subscriber = &mockSubscriberService{}

	ctx := context.Background()
	result, _, err := client.Subscriber.List(ctx, nil)
	if err != nil || len(result.Data) == 0 || result.Data[0].Email != "example@example.com" {
		t.Fatalf("mock failed")
	}
}
```

[pkg/testing](https://golang.org/pkg/testing/)

```
$ go test
```

<a name="support-and-feedback"></a>

# Support and Feedback

In case you find any bugs, submit an issue directly here in GitHub.

You are welcome to create SDK for any other programming language.

If you have any trouble using our API or SDK feel free to contact our support by
email [info@mailerlite.com](mailto:info@mailerlite.com)

The official API documentation is at [https://developers.mailerlite.com](https://developers.mailerlite.com)

<a name="license"></a>

# License

[The MIT License (MIT)](LICENSE)
