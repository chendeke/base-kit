package middleware

import (
	"context"
	"github.com/chendeke/base-kit/counter"
	//"git.aimap.io/location/genproto/location-apis/mojo/core"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/metrics"
	"time"
)



// InstrumentingMiddleware returns an endpoint middleware that records
// the duration of each invocation to the passed histogram. The middleware adds
// a single field: "success", which is "true" if no error is returned, and
// "false" otherwise.
func Instrumenting(duration metrics.Histogram, counter metrics.Counter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				errString := getErrorString(err)
				accessKey := getAccessKey(ctx)
				count := getCount(request)
				counter.With("error", errString, "access_key", accessKey).Add(count)
				duration.With("error", errString, "access_key", accessKey).Observe(time.Since(begin).Seconds())
			}(time.Now())
			return next(ctx, request)
		}
	}
}

func getAccessKey(ctx context.Context) string {
	var ret string
	if accessKey, ok := ctx.Value("access_key").(string); ok {
		ret = accessKey
	}
	return ret
}

func getCount(request interface{}) float64 {
	if r, ok := request.(counter.BatchRequest); ok {
		return float64(r.GetBatchCount())
	}
	return 1
}

func getErrorString(err error) string {
	var ret string
	if err != nil {
		if e, ok := err.(*core.Error); ok {
			if e.StatusCode() == 200 {
				ret = "ok"
			} else if e.StatusCode() >= 400 && e.StatusCode() < 500 {
				ret = "client_error"
			} else {
				ret = "server_error"
			}
		} else {
			ret = "server_error"
		}
	} else {
		ret = "ok"
	}
	return ret
}
