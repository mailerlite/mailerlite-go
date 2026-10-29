package mailerlite

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

// ErrInvalidIdempotencyKey - the idempotency key is empty, longer than 255 characters, contains a character
// that is not printable ASCII, or starts or ends with a space. The SDK sends no request.
var ErrInvalidIdempotencyKey = errors.New("mailerlite: the idempotency key must be 1 to 255 printable ASCII characters, with no space at the start or the end")

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

// IsIdempotentReplay - true if the API returned the stored response of an earlier request with the same idempotency key.
// The check ignores letter case and spaces at the start and the end of the header value.
func (r *Response) IsIdempotentReplay() bool {
	return r != nil && r.Response != nil && strings.EqualFold(strings.TrimSpace(r.Header.Get(HeaderIdempotentReplayed)), "true")
}

func setIdempotencyKey(req *http.Request) error {
	key, ok := req.Context().Value(idempotencyKeyContextKey{}).(string)
	if !ok {
		return nil
	}

	if len(key) == 0 || len(key) > 255 {
		return ErrInvalidIdempotencyKey
	}
	// net/http removes spaces at the start and the end of a header value. The API then gets a different key.
	if key[0] == ' ' || key[len(key)-1] == ' ' {
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
