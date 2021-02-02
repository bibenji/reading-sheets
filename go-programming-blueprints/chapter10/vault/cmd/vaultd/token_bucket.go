package main

import (
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/juju/ratelimit"
	"golang.org/x/net/context"
)

// ErrLimitExceed err for limit exceed
var ErrLimitExceed = errors.New("Rate Limit Exceed")

// NewTokenBucketLimiter create a token bucket limiter
func NewTokenBucketLimiter(tb *ratelimit.Bucket) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			if tb.TakeAvailable(1) == 0 {
				return nil, ErrLimitExceed
			}
			return next(ctx, request)
		}
	}
}
