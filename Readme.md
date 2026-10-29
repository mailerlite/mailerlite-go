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

Every write must send the `etag` that the last read or write of the email returned. If the etag is empty, the SDK returns `mailerlite.ErrEmptyEtag` and sends no request. If the email changed after that read or write, the API returns a `409` error with the code `mailerlite.ErrorCodeVersionConflict`. The `Etag` field of the `*mailerlite.ErrorResponse` contains the current etag of the email.

If an email, block or version ID is empty, the SDK returns `mailerlite.ErrEmptyID` and sends no request. If an ID contains a character that is not a digit, the SDK returns `mailerlite.ErrInvalidID` and sends no request.

The SDK also checks the `BlockID` of a placement when the position is `before` or `after`. For `start` and `end`, the API does not use `BlockID`. The SDK does not check it and does not send it. Your value does not change.

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

The position is `start`, `end`, `before` or `after`. `BlockID` is required for `before` and `after`, and must be a sibling of the block. For `start` and `end`, the SDK does not send `BlockID`.

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
- The key must be 1 to 255 printable ASCII characters. Otherwise the SDK returns `mailerlite.ErrInvalidIdempotencyKey` and sends no request.
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
