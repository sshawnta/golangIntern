//Package service http server
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package httpserver

import (
	"context"
	`net/http`

	`github.com/buaazp/fasthttprouter`
	"github.com/valyala/fasthttp"

	`github.com/sshawnta/TestTransport/models`
)

type errorProcessor interface {
	Encode(ctx context.Context, r *fasthttp.Response, err error)
}

type service interface {

	GetClientId(ctx context.Context, request *models.GetClientId) (response models.Response, err error)
	PostOrder(ctx context.Context, request *models.PostClientId) (response models.Response, err error)
	GetCount(ctx context.Context, request *models.GetClientId) (response models.Response, err error)
	GetOrder(ctx context.Context) (response models.Response, err error)
}


type getClientIdServer struct {
	transport      GetClientIdTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getClientIdServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	request, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err := s.service.GetClientId(ctx, &request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetClientIdServer the server creator
func NewGetClientIdServer(transport GetClientIdTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getClientIdServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}
type postOrderServer struct {
	transport      PostOrderTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *postOrderServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	request, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err := s.service.PostOrder(ctx, &request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewPostOrderServer the server creator
func NewPostOrderServer(transport PostOrderTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := postOrderServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}
type getCountServer struct {
	transport      GetCountTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getCountServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	request, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err := s.service.GetCount(ctx, &request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetCountServer the server creator
func NewGetCountServer(transport GetCountTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getCountServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}
type getOrderServer struct {
	transport      GetOrderTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getOrderServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	_, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err := s.service.GetOrder(ctx)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}
// NewGetOrderServer the server creator
func NewGetOrderServer(transport GetOrderTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getOrderServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

func NewPreparedServer(svc service) *fasthttprouter.Router {
	errorProcessor := NewErrorProcessor(http.StatusInternalServerError, "internal error")

	getClientIdTransport := NewGetClientIdTransport(NewError)
	postOrderTransport := NewPostOrderTransport(NewError)
	getCountTransport := NewGetCountTransport(NewError)
	getOrderTransport := NewGetOrderTransport(NewError)
	return MakeFastHTTPRouter(
		[]*HandlerSettings{
			{
				Path:    URIPathClientGetClientId,
				Method:  http.MethodGet,
				Handler: NewGetClientIdServer(getClientIdTransport, svc, errorProcessor),
			},
			{
				Path:    URIPathClientPostOrder,
				Method:  http.MethodPost,
				Handler: NewPostOrderServer(postOrderTransport, svc, errorProcessor),
			},
			{
				Path:    URIPathClientGetCount,
				Method:  http.MethodGet,
				Handler: NewGetCountServer(getCountTransport, svc, errorProcessor),
			},
			{
				Path:    URIPathClientGetOrder,
				Method:  http.MethodGet,
				Handler: NewGetOrderServer(getOrderTransport, svc, errorProcessor),
			},
		},
	)

}


