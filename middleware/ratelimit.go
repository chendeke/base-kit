package middleware

import (
	"context"
	"github.com/chendeke/genproto/location-apis/mojo/core"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

func NewTokenBucketLimitMiddleware(bkt *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !bkt.Allow() {
				return nil, core.NewError(http.StatusTooManyRequests, "Rate limit exceed!")
			}
			return next(ctx, request)
		}
	}
}

func EveryRateLimiter(interval time.Duration, b int) endpoint.Middleware {
	limiter := rate.NewLimiter(rate.Every(interval), b)
	return NewTokenBucketLimitMiddleware(limiter)
}