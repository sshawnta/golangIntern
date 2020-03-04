//Package service http client
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package httpclient

import (
	"context"

	"github.com/valyala/fasthttp"

	`github.com/sshawnta/TestTransport/models`
)

const (
	method = "http://"
	getClientId2 = "/api/v1/users"
	)
var (

	GetClientId = option{}
	PostOrder = option{}
	GetCount = option{}
	GetOrder = option{}
)

type option struct{}

// Option ...
type Option interface {
	Prepare(ctx context.Context, r *fasthttp.Request)
}

// Service implements Service interface
type Service interface {

	GetClientId(ctx context.Context, request *models.GetClientId) (response models.Response, err error)
	PostOrder(ctx context.Context, request *models.PostClientId) (response models.Response, err error)
	GetCount(ctx context.Context, request *models.GetClientId) (response models.Response, err error)
	GetOrder(ctx context.Context,  ) (response models.Response, err error)
}

type client struct {
	cli *fasthttp.HostClient

	transportGetClientId GetClientIdClientTransport
	transportPostOrder PostOrderClientTransport
	transportGetCount GetCountClientTransport
	transportGetOrder GetOrderClientTransport
	options map[interface{}]Option
}

// GetClientId ...
func (s *client) GetClientId(ctx context.Context, request *models.GetClientId) (response models.Response, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[GetClientId]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetClientId.EncodeRequest(ctx, req, request); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	//fmt.Println("122",req,response)
	return s.transportGetClientId.DecodeResponse(ctx, res)
}

// PostOrder ...
func (s *client) PostOrder(ctx context.Context, request *models.PostClientId) (response models.Response, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[PostOrder]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportPostOrder.EncodeRequest(ctx, req, request); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportPostOrder.DecodeResponse(ctx, res)
}

// GetCount ...
func (s *client) GetCount(ctx context.Context, request *models.GetClientId) (response models.Response, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()

	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[GetCount]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetCount.EncodeRequest(ctx, req, request); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportGetCount.DecodeResponse(ctx, res)
}

// GetOrder ...
func (s *client) GetOrder(ctx context.Context,  ) (response models.Response, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[GetOrder]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetOrder.EncodeRequest(ctx, req, ); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportGetOrder.DecodeResponse(ctx, res)
}

// NewClient the client creator
func NewClient(
	cli *fasthttp.HostClient,
	
	transportGetClientId GetClientIdClientTransport,
	transportPostOrder PostOrderClientTransport,
	transportGetCount GetCountClientTransport,
	transportGetOrder GetOrderClientTransport,
	options map[interface{}]Option,
) Service {
	return &client{
		cli: cli,
		
		transportGetClientId: transportGetClientId,
		transportPostOrder: transportPostOrder,
		transportGetCount: transportGetCount,
		transportGetOrder: transportGetOrder,
		options: options,
	}
}

// NewPreparedClient create and set up http client
func NewPreparedClient(
	serverURL string,
	serverHost string,
	maxConns int,
	options map[interface{}]Option,//nil
	errorProcessor errorProcessor,
	errorCreator errorCreator,

	uriPathGetClientId string,
	uriPathPostOrder string,
	uriPathGetCount string,
	uriPathGetOrder string,

	httpMethodGetClientId string,
	httpMethodPostOrder string,
	httpMethodGetCount string,
	httpMethodGetOrder string,
) Service {
	transportGetClientId := NewGetClientIdClientTransport(
		errorProcessor,
		errorCreator,
		method+serverURL+uriPathGetClientId,
		httpMethodGetClientId,
	)
	transportPostOrder := NewPostOrderClientTransport(
		errorProcessor,
		errorCreator,
		method+serverURL+uriPathPostOrder,
		httpMethodPostOrder,
	)
	transportGetCount := NewGetCountClientTransport(
		errorProcessor,
		errorCreator,
		method+serverURL+uriPathGetCount,
		httpMethodGetCount,
	)
	transportGetOrder := NewGetOrderClientTransport(
		errorProcessor,
		errorCreator,
		method+serverURL+uriPathGetOrder,
		httpMethodGetOrder,
	)

	return NewClient(
		&fasthttp.HostClient{
			Addr:     serverHost,
			MaxConns: maxConns,
		},

		transportGetClientId,
		transportPostOrder,
		transportGetCount,
		transportGetOrder,
		options,
	)
}
