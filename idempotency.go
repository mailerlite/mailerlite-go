package mailerlite

import (
	"context"
	"net/http"
	"strings"
)

type idempotencyKeyContextKey struct{}

// WithIdempotencyKey - return a copy of ctx that sends key in the Idempotency-Key header.
//
// Use a new key for each write, and the same key when you retry that write. If the first request succeeded,
// the API returns the stored response again and does not do the write two times. Response.IsIdempotentReplay
// then returns true. The API keeps a key for 24 hours.
//
// The API uses the key only for POST, PATCH and DELETE requests of EmailService, except the validate methods.
// If key is empty, the SDK sends no header. The SDK does not check the format of the key. The API checks it.
func WithIdempotencyKey(ctx context.Context, key string) context.Context {
	return context.WithValue(ctx, idempotencyKeyContextKey{}, key)
}

// IsIdempotentReplay - true if the API returned the stored response of an earlier request with the same idempotency key.
// The check ignores letter case and spaces at the start and the end of the header value.
func (r *Response) IsIdempotentReplay() bool {
	return r != nil && r.Response != nil && strings.EqualFold(strings.TrimSpace(r.Header.Get(HeaderIdempotentReplayed)), "true")
}

func setIdempotencyKey(req *http.Request) {
	key, _ := req.Context().Value(idempotencyKeyContextKey{}).(string)
	if key == "" {
		return
	}

	req.Header.Set(HeaderIdempotencyKey, key)
}
