package mailerlite

import (
	"context"
	"errors"
	"net/http"
)

// ErrInvalidIdempotencyKey - the idempotency key is empty, longer than 255 characters, or contains a character
// that is not printable ASCII. The SDK sends no request.
var ErrInvalidIdempotencyKey = errors.New("mailerlite: the idempotency key must be 1 to 255 printable ASCII characters")

type idempotencyKeyContextKey struct{}

// WithIdempotencyKey - return a copy of ctx that sends key in the Idempotency-Key header.
//
// Use a new key for each write, and the same key when you retry that write. If the first request succeeded,
// the API returns the stored response again and does not do the write two times. Response.IsIdempotentReplay
// then returns true. The API keeps a key for 24 hours.
//
// The API uses the key only for POST, PATCH and DELETE requests of EmailService, except the validate methods.
func WithIdempotencyKey(ctx context.Context, key string) context.Context {
	return context.WithValue(ctx, idempotencyKeyContextKey{}, key)
}

// IsIdempotentReplay - true if the API returned the stored response of an earlier request with the same idempotency key
func (r *Response) IsIdempotentReplay() bool {
	return r != nil && r.Response != nil && r.Header.Get(HeaderIdempotentReplayed) == "true"
}

// setIdempotencyKey - set the Idempotency-Key header from the context of req, if the context has a key
func setIdempotencyKey(req *http.Request) error {
	key, ok := req.Context().Value(idempotencyKeyContextKey{}).(string)
	if !ok {
		return nil
	}

	if len(key) == 0 || len(key) > 255 {
		return ErrInvalidIdempotencyKey
	}
	for i := 0; i < len(key); i++ {
		if key[i] < 0x20 || key[i] > 0x7e {
			return ErrInvalidIdempotencyKey
		}
	}

	req.Header.Set(HeaderIdempotencyKey, key)

	return nil
}
