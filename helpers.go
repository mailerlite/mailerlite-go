package mailerlite

import (
	"net/url"
	"strconv"
)

var (
	SortByID                           = "id"
	SortByIDDescending                 = "-id"
	SortByName                         = "name"
	SortByNameDescending               = "-name"
	SortByType                         = "type"
	SortByTypeDescending               = "-type"
	SortByTotal                        = "total"
	SortByTotalDescending              = "-total"
	SortByOpenRate                     = "open_rate"
	SortByOpenRateDescending           = "-open_rate"
	SortByClickRate                    = "click_rate"
	SortByClickRateDescending          = "-click_rate"
	SortByConversionsCount             = "conversions_count"
	SortByConversionsCountDescending   = "-conversions_count"
	SortByConversionRate               = "conversion_rate"
	SortByConversionRateDescending     = "-conversion_rate"
	SortByClicksCount                  = "clicks_count"
	SortByClicksCountDescending        = "-clicks_count"
	SortByOpensCount                   = "opens_count"
	SortByOpensCountDescending         = "-opens_count"
	SortByVisitors                     = "visitors"
	SortByVisitorsDescending           = "-visitors"
	SortByLastRegistrationAt           = "last_registration_at"
	SortByLastRegistrationAtDescending = "-created_at"
	SortByCreatedAt                    = "created_at"
	SortByCreatedAtDescending          = "-created_at"
	SortByUpdatedAt                    = "updated_at"
	SortByUpdatedAtDescending          = "-updated_at"

	FormTypePopup     = "popup"
	FormTypeEmbedded  = "embedded"
	FormTypePromotion = "promotion"

	CampaignTypeRegular = "regular"
	CampaignTypeAB      = "ab"
	CampaignTypeResend  = "resend"

	CampaignScheduleTypeInstant   = "instant"
	CampaignScheduleTypeScheduled = "scheduled"
	CampaignScheduleTypeTimezone  = "timezone_based"
)

type Meta struct {
	// offset  based pagination
	CurrentPage int         `json:"current_page"`
	From        int         `json:"from"`
	LastPage    int         `json:"last_page"`
	Links       []MetaLinks `json:"links"`
	Path        string      `json:"path"`
	PerPage     int         `json:"per_page"`
	To          int         `json:"to"`

	*Aggregations `json:"aggregations,omitempty"`
	*Counts       `json:"counts,omitempty"`

	// cursor based pagination
	Count int `json:"count"`
	Last  int `json:"last"`

	Total           int `json:"total"`
	TotalUnfiltered int `json:"total_unfiltered,omitempty"`

	// cursor based string pagination (subscribers)
	NextCursor string `json:"next_cursor,omitempty"`
	PrevCursor string `json:"prev_cursor,omitempty"`
}

type Aggregations struct {
	Total int `json:"total"`
	Draft int `json:"draft"`
	Ready int `json:"ready"`
	Sent  int `json:"sent"`
}

type Counts struct {
	All          int `json:"all"`
	Opened       int `json:"opened"`
	Unopened     int `json:"unopened"`
	Clicked      int `json:"clicked"`
	Unsubscribed int `json:"unsubscribed"`
	Forwarded    int `json:"forwarded"`
	Hardbounced  int `json:"hardbounced"`
	Softbounced  int `json:"softbounced"`
	Junk         int `json:"junk"`
}

// Links manages links that are returned along with a List
type Links struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
}

type MetaLinks struct {
	URL    interface{} `json:"url"`
	Label  string      `json:"label"`
	Active bool        `json:"active"`
}

type OpenRate struct {
	Float  float64 `json:"float"`
	String string  `json:"string"`
}

type ClickRate struct {
	Float  float64 `json:"float"`
	String string  `json:"string"`
}

type BounceRate struct {
	Float  float64 `json:"float"`
	String string  `json:"string"`
}

type UnsubscribeRate struct {
	Float  float64 `json:"float"`
	String string  `json:"string"`
}

type SpamRate struct {
	Float  float64 `json:"float"`
	String string  `json:"string"`
}

type HardBounceRate struct {
	Float  float64 `json:"float"`
	String string  `json:"string"`
}

type SoftBounceRate struct {
	Float  float64 `json:"float"`
	String string  `json:"string"`
}

type ClickToOpenRate struct {
	Float  float64 `json:"float"`
	String string  `json:"string"`
}

// NextPageToken is the page token to request the next page of the list
func (l *Links) NextPageToken() (string, error) {
	return l.nextPageToken()
}

// PrevPageToken is the page token to request the previous page of the list
func (l *Links) PrevPageToken() (string, error) {
	return l.prevPageToken()
}

func (l *Links) nextPageToken() (string, error) {
	if l == nil || l.Next == "" {
		return "", nil
	}
	token, err := pageTokenFromURL(l.Next)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (l *Links) prevPageToken() (string, error) {
	if l == nil || l.Prev == "" {
		return "", nil
	}
	token, err := pageTokenFromURL(l.Prev)
	if err != nil {
		return "", err
	}
	return token, nil
}

// IsLastPage returns true if the current page is the last
func (l *Links) IsLastPage() bool {
	return l.isLast()
}

func (l *Links) isLast() bool {
	return l.Next == ""
}

func pageForURL(urlText string) (int, error) {
	u, err := url.ParseRequestURI(urlText)
	if err != nil {
		return 0, err
	}

	pageStr := u.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return 0, err
	}

	return page, nil
}

func pageTokenFromURL(urlText string) (string, error) {
	u, err := url.ParseRequestURI(urlText)
	if err != nil {
		return "", err
	}
	return u.Query().Get("page_token"), nil
}
