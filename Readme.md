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
    - [Emails](#emails)
        - [Create a campaign with the simple editor](#create-a-campaign-with-the-simple-editor)
        - [Get an email](#get-an-email)
        - [Get the HTML of an email](#get-the-html-of-an-email)
        - [Get a block of an email](#get-a-block-of-an-email)
        - [Insert a block into an email](#insert-a-block-into-an-email)
        - [Update the content of a block](#update-the-content-of-a-block)
        - [Move a block](#move-a-block)
        - [Delete a block](#delete-a-block)
        - [Validate a block](#validate-a-block)
        - [Update the design settings of an email](#update-the-design-settings-of-an-email)
        - [Validate the design settings of an email](#validate-the-design-settings-of-an-email)
        - [Save an email as a template](#save-an-email-as-a-template)
        - [Get a list of versions of an email](#get-a-list-of-versions-of-an-email)
        - [Save a version of an email](#save-a-version-of-an-email)
        - [Restore a version of an email](#restore-a-version-of-an-email)
        - [Retry a write with an idempotency key](#retry-a-write-with-an-idempotency-key)
    - [Brand styles](#brand-styles)
        - [Get the brand styles](#get-the-brand-styles)
    - [Files](#files)
        - [Get a list of files](#get-a-list-of-files)
    - [E-commerce](#e-commerce)
        - [Shops](#shops)
        - [Products and variants](#products-and-variants)
        - [Categories](#categories)
        - [Customers, orders and carts](#customers-orders-and-carts)
        - [Imports](#imports)
        - [Revenue reports and settings](#revenue-reports-and-settings)

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

## Emails

These endpoints work on emails built with the simple editor. Get the email ID from a campaign's `emails` array.

To get a simple editor email, create a campaign with `Editor: mailerlite.EmailEditorSimple`. You can set `Editor` only when you create the campaign. Do not set `Content`. If the email does not use the simple editor, the API returns a `422` error.

`TemplateID` is optional, and works only with `Editor: mailerlite.EmailEditorSimple`. Use a simple editor template from the shared gallery, or a template that you saved from your own email with `client.Email.CreateTemplate`. The API copies the template into the new email. Without `TemplateID`, the email starts from the default blank template. Use the same `Editor` and the same `TemplateID` on every email of the campaign.

Every write must send the `etag` that the last read or write of the email returned. If the etag is empty, the API returns a `422` error. If the email changed after that read or write, the API returns a `409` error with the code `mailerlite.ErrorCodeVersionConflict`. The `Etag` field of the `*mailerlite.ErrorResponse` contains the current etag of the email.

If an email, block or version ID is empty, the SDK returns `mailerlite.ErrEmptyID` and sends no request. If an ID contains a character that is not a digit, the SDK returns `mailerlite.ErrInvalidID` and sends no request.

The SDK does not check the other values, for example the placement. It sends them as you give them. The API checks them and returns an error. The validate methods return `Valid: false` with the errors.

### Create a campaign with the simple editor

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

	campaign := &mailerlite.CreateCampaign{
		Name: "Campaign Name",
		Type: mailerlite.CampaignTypeRegular,
		Emails: []mailerlite.Emails{
			{
				Subject:  "Subject",
				FromName: "Your Name",
				From:     "your@domain.com",
				Editor:   mailerlite.EmailEditorSimple,
			},
		},
	}

	root, _, err := client.Campaign.Create(ctx, campaign)
	if err != nil {
		log.Fatal(err)
	}

	email, _, err := client.Email.Get(ctx, root.Data.Emails[0].ID)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(email.Etag)
}
```

### Get an email

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

	email, _, err := client.Email.Get(ctx, "email-id")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(email.Etag)
}
```

### Get the HTML of an email

The HTML is a preview. Do not send it yourself, because it contains personalization tokens and no tracking URLs.

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

	html, _, err := client.Email.GetHTML(ctx, "email-id")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(html.Data.HTML)
}
```

### Get a block of an email

`Placement` gives the IDs of the blocks before and after the block.

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

	block, _, err := client.Email.GetBlock(ctx, "email-id", "block-id")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(block.Data.Type)
}
```

### Insert a block into an email

Without `Placement`, the block goes at the end of the email. `EditorContent` is required. Use an empty map for a block with no content, such as a divider.

If the API removes attributes or `mobile` keys that the block type does not declare, the response has `Rejected` with the keys of the removed values. The API saves the rest of the block.

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

	email, _, err := client.Email.Get(ctx, "email-id")
	if err != nil {
		log.Fatal(err)
	}

	block := &mailerlite.CreateEmailBlockOptions{
		Type: mailerlite.EmailBlockTypeParagraph,
		EditorContent: map[string]interface{}{
			"content": []interface{}{
				map[string]interface{}{"type": "text", "text": "Hello"},
			},
		},
		Placement: &mailerlite.EmailBlockPosition{
			Position: mailerlite.EmailBlockPositionAfter,
			BlockID:  "block-id",
		},
		Etag: email.Etag,
	}

	created, _, err := client.Email.CreateBlock(ctx, "email-id", block)
	if err != nil {
		log.Fatal(err)
	}

	// Use created.Etag for the next write to this email.
	log.Print(created.Data.ID)
}
```

### Update the content of a block

`EditorContent` replaces the whole content of the block. Get the block first, change it, and send all of it back.

If the API removes attributes or `mobile` keys that the block type does not declare, the response has `Rejected` with the keys of the removed values.

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

	block, _, err := client.Email.GetBlock(ctx, "email-id", "block-id")
	if err != nil {
		log.Fatal(err)
	}

	content := block.Data.EditorContent
	content["attrs"] = map[string]interface{}{"textAlign": "center"}

	updated, _, err := client.Email.UpdateBlock(ctx, "email-id", "block-id", &mailerlite.UpdateEmailBlockOptions{
		EditorContent: content,
		Etag:          block.Etag,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Print(updated.Etag)
}
```

### Move a block

The position is `start`, `end`, `before` or `after`. `BlockID` is required for `before` and `after`, and must be a sibling of the block.

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

	moved, _, err := client.Email.MoveBlock(ctx, "email-id", "block-id", &mailerlite.MoveEmailBlockOptions{
		Placement: mailerlite.EmailBlockPosition{
			Position: mailerlite.EmailBlockPositionBefore,
			BlockID:  "other-block-id",
		},
		Etag: "etag-from-the-last-read-or-write",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Print(moved.Etag)
}
```

### Delete a block

The blocks nested in the block are deleted too. The response has only the new etag of the email.

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

	deleted, _, err := client.Email.DeleteBlock(ctx, "email-id", "block-id", "etag-from-the-last-read-or-write")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(deleted.Etag)
}
```

### Validate a block

This checks a block against the same rules as an insert, and saves nothing. Invalid content is a successful request, so check `Valid` and `Errors`.

If the content is valid but an insert would remove some keys, the result has `Rejected` with these keys.

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

	result, _, err := client.Email.ValidateBlock(ctx, &mailerlite.CreateEmailBlockOptions{
		Type: mailerlite.EmailBlockTypeParagraph,
		EditorContent: map[string]interface{}{
			"attrs": map[string]interface{}{"textAlign": "center"},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if !result.Valid {
		log.Print(result.Errors)
	}
}
```

### Update the design settings of an email

Set only the keys to change. The API stores every value that you send on the email, also a value that is the same as the account value.

- To unset a brand colour, send an empty string. The API stores the value empty, and the email then uses the account value.
- To unset a footer text, send an empty editor document: an empty `EditorContent` map, or a map with an empty `content` list.
- Other colours, sizes, `Button.BorderRadius` and `Brand.Style` cannot be empty. An empty string gives a `422` error.

If the API removes an unsafe value, the response has `Rejected` with the keys of the removed values.

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

	updated, _, err := client.Email.UpdateSettings(ctx, "email-id", &mailerlite.UpdateEmailSettingsOptions{
		EmailSettings: mailerlite.EmailSettings{
			Brand: &mailerlite.EmailBrandSettings{
				Mode:            mailerlite.String(mailerlite.EmailBrandModeCustom),
				BackgroundColor: mailerlite.String("#f4f4f4"),
			},
		},
		Etag: "etag-from-the-last-read-or-write",
	})
	if err != nil {
		log.Fatal(err)
	}

	if updated.Rejected != nil {
		log.Print(updated.Rejected.Keys)
	}
}
```

### Validate the design settings of an email

This checks a settings change against the same rules as an update, and saves nothing. No etag is necessary.

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

	result, _, err := client.Email.ValidateSettings(ctx, "email-id", &mailerlite.UpdateEmailSettingsOptions{
		EmailSettings: mailerlite.EmailSettings{
			Button: &mailerlite.EmailButtonSettings{BorderRadius: mailerlite.String("12")},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if !result.Valid {
		log.Print(result.Errors)
	}
}
```

### Save an email as a template

To start a new campaign from the saved template, convert `template.Data.ID` to an `int` with `strconv.Atoi`, and send it as `TemplateID` when you create the campaign.

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

	template, _, err := client.Email.CreateTemplate(ctx, "email-id", &mailerlite.CreateEmailTemplateOptions{
		Name: "Newsletter",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Print(template.Data.ID)
}
```

### Get a list of versions of an email

The list is newest first. `Current` marks the version that the email is on now.

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

	versions, _, err := client.Email.ListVersions(ctx, "email-id", &mailerlite.ListEmailVersionOptions{
		Page:  1,
		Limit: 10,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Print(len(versions.Data))
}
```

### Save a version of an email

Changes made through the API do not save a version. If nothing changed since the last version, the API returns that version with status `200` instead of `201`.

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

	version, _, err := client.Email.CreateVersion(ctx, "email-id")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(version.Data.ID)
}
```

### Restore a version of an email

The API saves the replaced state as a version first, so you can undo a restore. The response is the restored email.

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

	email, _, err := client.Email.RestoreVersion(ctx, "email-id", "version-id", "etag-from-the-last-read-or-write")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(email.Etag)
}
```

### Retry a write with an idempotency key

If a write fails with a network error, you do not know if the API did the write. Send the write with an idempotency key, and retry it with the same key and the same request. If the first request succeeded, the API returns the stored response again and does not do the write two times.

- Use a new key for each write. The API keeps a key for 24 hours.
- If the key is empty, the SDK sends no `Idempotency-Key` header.
- The SDK does not check the format of the key. The API returns a `400` error for an invalid key.
- `net/http` returns an error for a key with a control character, for example a new line. It then sends no request.
- For HTTP/1.1, `net/http` removes spaces at the start and the end of the key.
- The same key with a different request returns a `422` error. If the first request is still running, the API returns a `409` error without a code.
- The API uses the key only for the POST, PATCH and DELETE requests of `client.Email`, except the validate methods.

```go
package main

import (
	"context"
	"errors"
	"log"

	"github.com/mailerlite/mailerlite-go"
)

var APIToken = "Api Token Here"

func main() {
	client := mailerlite.NewClient(APIToken)

	ctx := context.TODO()

	block := &mailerlite.CreateEmailBlockOptions{
		Type:          mailerlite.EmailBlockTypeDivider,
		EditorContent: map[string]interface{}{},
		Etag:          "etag-from-the-last-read-or-write",
	}

	keyCtx := mailerlite.WithIdempotencyKey(ctx, "insert-divider-0001")

	created, res, err := client.Email.CreateBlock(keyCtx, "email-id", block)
	var errResp *mailerlite.ErrorResponse
	if err != nil && !errors.As(err, &errResp) {
		// Network error: retry with the same key and the same request.
		created, res, err = client.Email.CreateBlock(keyCtx, "email-id", block)
	}
	if err != nil {
		log.Fatal(err)
	}

	log.Print(created.Data.ID, res.IsIdempotentReplay())
}
```

## Brand styles

### Get the brand styles

The response has the account palette (`Brand`, `nil` if the account has none), the default palette and the presets. To apply a preset to an email, send its `Identifier` as `Brand.Style` in `Email.UpdateSettings`.

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

	styles, _, err := client.BrandStyle.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, preset := range styles.Data.Presets {
		log.Print(preset.Identifier)
	}
}
```

## Files

### Get a list of files

Use a file URL to put an image into an email block. You can filter by `name` (partial match) and `type` (`image`, `pdf` or `epub`).

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

	listOptions := &mailerlite.ListFileOptions{
		Filters: &[]mailerlite.Filter{{
			Name:  "type",
			Value: mailerlite.FileTypeImage,
		}},
		Sort:  mailerlite.SortByCreatedAtDescending,
		Page:  1,
		Limit: 10,
	}

	files, _, err := client.File.List(ctx, listOptions)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files.Data {
		log.Print(file.URL)
	}
}
```

## E-commerce

Every e-commerce call takes the shop ID. The examples below assume this setup:

```go
client := mailerlite.NewClient(APIToken)
ctx := context.TODO()
```

### Shops

```go
shops, _, err := client.Ecommerce.Shop.List(ctx, &mailerlite.ListEcommerceOptions{Limit: 25})
shop, _, err := client.Ecommerce.Shop.Get(ctx, "shop-id")
shop, _, err := client.Ecommerce.Shop.Create(ctx, &mailerlite.ShopParams{Name: "My shop", Url: "https://shop.example.com", Currency: "EUR"})
shop, _, err := client.Ecommerce.Shop.Update(ctx, "shop-id", &mailerlite.ShopParams{Name: "New name"})
_, err := client.Ecommerce.Shop.Delete(ctx, "shop-id")
```

### Products and variants

```go
products, _, err := client.Ecommerce.Product.List(ctx, "shop-id", &mailerlite.ListProductOptions{
	Filters: &[]mailerlite.Filter{{Name: "enabled", Value: true}},
	Sort:    "-created_at",
})
product, _, err := client.Ecommerce.Product.Get(ctx, "shop-id", "product-id")
product, _, err := client.Ecommerce.Product.Create(ctx, "shop-id", &mailerlite.ProductParams{Name: "Red T-shirt"})
product, _, err := client.Ecommerce.Product.Update(ctx, "shop-id", "product-id", &mailerlite.ProductParams{Name: "Blue T-shirt"})
_, err := client.Ecommerce.Product.Delete(ctx, "shop-id", "product-id")
product, _, err := client.Ecommerce.Product.ChangeStatus(ctx, "shop-id", "product-id", false)
product, _, err := client.Ecommerce.Product.AssignCategories(ctx, "shop-id", "product-id", []string{"category-id"}, true)

variants, _, err := client.Ecommerce.ProductVariant.List(ctx, "shop-id", "product-id")
product, _, err := client.Ecommerce.ProductVariant.Sync(ctx, "shop-id", "product-id", []mailerlite.ProductVariantParams{
	{Name: "Large", Attributes: map[string]string{"size": "L"}, ResourceID: "variant-l"},
})
product, _, err := client.Ecommerce.ProductVariant.Update(ctx, "shop-id", "product-id", "variant-id", &mailerlite.ProductVariantParams{Name: "Extra large", Attributes: map[string]string{"size": "XL"}})
_, err := client.Ecommerce.ProductVariant.Delete(ctx, "shop-id", "product-id", "variant-id")
```

### Categories

```go
categories, _, err := client.Ecommerce.Category.List(ctx, "shop-id", nil)
category, _, err := client.Ecommerce.Category.Get(ctx, "shop-id", "category-id")
category, _, err := client.Ecommerce.Category.Create(ctx, "shop-id", &mailerlite.CategoryParams{Name: "Hats"})
category, _, err := client.Ecommerce.Category.Update(ctx, "shop-id", "category-id", &mailerlite.CategoryParams{Name: "Caps"})
_, err := client.Ecommerce.Category.Delete(ctx, "shop-id", "category-id")

products, _, err := client.Ecommerce.Category.Products(ctx, "shop-id", "category-id", nil)
_, err := client.Ecommerce.Category.AssignProduct(ctx, "shop-id", "category-id", "product-id")
_, err := client.Ecommerce.Category.UnassignProduct(ctx, "shop-id", "category-id", "product-id")
```

### Customers, orders and carts

```go
customers, _, err := client.Ecommerce.Customer.List(ctx, "shop-id", &mailerlite.ListEcommerceOptions{
	Filters: &[]mailerlite.Filter{{Name: "email", Value: "customer@example.com"}},
})
customer, _, err := client.Ecommerce.Customer.Create(ctx, "shop-id", &mailerlite.CustomerParams{Email: "customer@example.com"})
customer, _, err := client.Ecommerce.Customer.Update(ctx, "shop-id", "customer-id", &mailerlite.CustomerParams{Name: "Jane"})
_, err := client.Ecommerce.Customer.Delete(ctx, "shop-id", "customer-id")

orders, _, err := client.Ecommerce.Order.List(ctx, "shop-id", &mailerlite.ListEcommerceOptions{
	Filters: &[]mailerlite.Filter{{Name: "attributed", Value: true}},
})
order, _, err := client.Ecommerce.Order.Create(ctx, "shop-id", &mailerlite.CreateOrderParams{
	Customer: mailerlite.CustomerParams{Email: "customer@example.com"},
	Cart:     mailerlite.OrderCartParams{Items: []mailerlite.CartItemParams{{EcommerceProductID: "product-id", Quantity: 1}}},
	Status:   "complete",
})
order, _, err := client.Ecommerce.Order.Update(ctx, "shop-id", "order-id", &mailerlite.UpdateOrderParams{Status: "complete"})
_, err := client.Ecommerce.Order.Delete(ctx, "shop-id", "order-id")

carts, _, err := client.Ecommerce.Cart.List(ctx, "shop-id", nil)
cart, _, err := client.Ecommerce.Cart.Update(ctx, "shop-id", "cart-id", &mailerlite.UpdateCartParams{CheckoutUrl: "https://shop.example.com/cart"})
items, _, err := client.Ecommerce.CartItem.List(ctx, "shop-id", "cart-id", nil)
item, _, err := client.Ecommerce.CartItem.Update(ctx, "shop-id", "cart-id", "item-id", &mailerlite.CartItemParams{Quantity: 2})
_, err := client.Ecommerce.CartItem.Delete(ctx, "shop-id", "cart-id", "item-id")
```

### Imports

```go
_, err := client.Ecommerce.Import.Products(ctx, "shop-id", []mailerlite.ProductParams{{Name: "Red T-shirt"}})
_, err := client.Ecommerce.Import.Categories(ctx, "shop-id", []mailerlite.CategoryParams{{Name: "Hats"}})
_, err := client.Ecommerce.Import.Customers(ctx, "shop-id", []mailerlite.CustomerParams{{Email: "customer@example.com"}})
_, err := client.Ecommerce.Import.Orders(ctx, "shop-id", []mailerlite.CreateOrderParams{{
	Customer: mailerlite.CustomerParams{Email: "customer@example.com"},
	Cart:     mailerlite.OrderCartParams{Items: []mailerlite.CartItemParams{{EcommerceProductID: "product-id", Quantity: 1}}},
}})
```

### Revenue reports and settings

```go
stats, _, err := client.Ecommerce.Shop.Stats(ctx, "shop-id", &mailerlite.ShopStatsOptions{StartDate: "2026-10-01", EndDate: "2026-10-31"})

stats, _, err := client.Campaign.EcommerceStats(ctx, "campaign-id", &mailerlite.EcommerceReportOptions{ShopID: "shop-id"})
orders, _, err := client.Campaign.EcommerceOrders(ctx, "campaign-id", &mailerlite.EcommerceReportOptions{ShopID: "shop-id"})

stats, _, err := client.Subscriber.EcommerceStats(ctx, "subscriber-id", &mailerlite.EcommerceReportOptions{ShopID: "shop-id"})
orders, _, err := client.Subscriber.EcommerceOrders(ctx, "subscriber-id", &mailerlite.EcommerceReportOptions{ShopID: "shop-id"})

stats, _, err := client.Automation.EcommerceStats(ctx, "automation-id", nil)
stats, _, err := client.Automation.EmailEcommerceStats(ctx, "automation-id", "email-id", nil)
orders, _, err := client.Automation.EcommerceOrders(ctx, "automation-id", &mailerlite.AutomationEcommerceOrderOptions{TimeFrame: "all_time"})

settings, _, err := client.Ecommerce.Setting.Get(ctx)
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
