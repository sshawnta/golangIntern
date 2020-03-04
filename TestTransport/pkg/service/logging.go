//Package service logging wrapper
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package service

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	`github.com/sshawnta/TestTransport/models`
)

// loggingMiddleware wraps Service and logs request information to the provided logger
type loggingMiddleware struct {
	logger log.Logger
	svc    Service
}

func (s *loggingMiddleware) GetClientId(ctx context.Context, request *models.GetClientId) (response models.Response, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "GetClientId",
			"request", request,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.GetClientId(ctx, request)
}

func (s *loggingMiddleware) PostOrder(ctx context.Context, request *models.PostClientId) (response models.Response, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "PostOrder",
			"request", request,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.PostOrder(ctx, request)
}

func (s *loggingMiddleware) GetCount(ctx context.Context, request *models.GetClientId) (response models.Response, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "GetCount",
			"request", request,
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.GetCount(ctx, request)
}

func (s *loggingMiddleware) GetOrder(ctx context.Context,  ) (response models.Response, err error) {
	defer func(begin time.Time) {
		_ = s.wrap(err).Log(
			"method", "GetOrder",
			
			"err", err,
			"elapsed", time.Since(begin),
		)
	}(time.Now())
	return s.svc.GetOrder(ctx)
}

func (s *loggingMiddleware) wrap(err error) log.Logger {
	lvl := level.Debug
	if err != nil {
		lvl = level.Error
	}
	return lvl(s.logger)
}

// NewLoggingMiddleware ...
func NewLoggingMiddleware(logger log.Logger, svc Service) *loggingMiddleware {
	return &loggingMiddleware{
		logger: logger,
		svc:    svc,
	}
}
