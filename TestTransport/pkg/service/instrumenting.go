//Package service instrumenting wrapper
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package service

import (
	"context"
	"strconv"
	"time"

	"github.com/go-kit/kit/metrics"

	`github.com/sshawnta/TestTransport/models`
)

// instrumentingMiddleware wraps Service and enables request metrics
type instrumentingMiddleware struct {
	reqCount    metrics.Counter
	reqDuration metrics.Histogram
	svc         Service
}

func (s *instrumentingMiddleware) GetClientId(ctx context.Context, request *models.GetClientId) (response models.Response, err error) {
	defer s.recordMetrics("GetClientId", time.Now(), err)
	return s.svc.GetClientId(ctx, request)
}

func (s *instrumentingMiddleware) PostOrder(ctx context.Context, request *models.PostClientId) (response models.Response, err error) {
	defer s.recordMetrics("PostOrder", time.Now(), err)
	return s.svc.PostOrder(ctx, request)
}

func (s *instrumentingMiddleware) GetCount(ctx context.Context, request *models.GetClientId) (response models.Response, err error) {
	defer s.recordMetrics("GetCount", time.Now(), err)
	return s.svc.GetCount(ctx, request)
}

func (s *instrumentingMiddleware) GetOrder(ctx context.Context,  ) (response models.Response, err error) {
	defer s.recordMetrics("GetOrder", time.Now(), err)
	return s.svc.GetOrder(ctx)
}

func (s *instrumentingMiddleware) recordMetrics(method string, startTime time.Time, err error) {
	labels := []string{
		"method", method,
		"error", strconv.FormatBool(err != nil),
	}
	s.reqCount.With(labels...).Add(1)
	s.reqDuration.With(labels...).Observe(time.Since(startTime).Seconds())
}

// NewInstrumentingMiddleware ...
func NewInstrumentingMiddleware(reqCount metrics.Counter, reqDuration metrics.Histogram, svc Service) *instrumentingMiddleware {
	return &instrumentingMiddleware{
		reqCount:    reqCount,
		reqDuration: reqDuration,
		svc:         svc,
	}
}
