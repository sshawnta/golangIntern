//Package service http client
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package httpclient

import (
	"bufio"
	"context"
	"encoding/json"
	`fmt`
	"net/http"
	`strconv`

	"github.com/valyala/fasthttp"

	`github.com/sshawnta/TestTransport/models`
)

type errorCreator func(status int, format string, v ...interface{}) error

type errorProcessor interface {
	Encode(ctx context.Context, r *fasthttp.Response, err error)
	Decode(r *fasthttp.Response) error
}


// GetClientIdClientTransport transport interface
type GetClientIdClientTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, request *models.GetClientId) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (response models.Response, err error)
}

type getClientIdClientTransport struct {
	errorProcessor errorProcessor
	errorCreator   errorCreator
	pathTemplate   string
	method         string
}

// EncodeRequest method for encoding requests on client side
func (t *getClientIdClientTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, request *models.GetClientId) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(t.pathTemplate)
	r.URI().QueryArgs().Set("id", strconv.Itoa(request.Id))
	fmt.Println(r.URI())
	return
}

// DecodeResponse method for decoding response on client side
func (t *getClientIdClientTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (response models.Response, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}
	err = response.UnmarshalJSON(r.Body())
	return
}

// NewGetClientIdClientTransport the transport creator for http requests
func NewGetClientIdClientTransport(
	errorProcessor errorProcessor,
	errorCreator   errorCreator,
	pathTemplate   string,
	method         string,
) GetClientIdClientTransport {
	return &getClientIdClientTransport{
		errorProcessor: errorProcessor,
		errorCreator:   errorCreator,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}
// PostOrderClientTransport transport interface
type PostOrderClientTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, request *models.PostClientId) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (response models.Response, err error)
}

type postOrderClientTransport struct {
	errorProcessor errorProcessor
	errorCreator   errorCreator
	pathTemplate   string
	method         string
}

// EncodeRequest method for encoding requests on client side
func (t *postOrderClientTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, request *models.PostClientId) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(t.pathTemplate)
	r.Header.Set("Content-Type", "application/json")
	r.SetBodyStreamWriter(func(w *bufio.Writer) {
		if err = json.NewEncoder(w).Encode(request); err != nil {
			return
		}
	})
	return
}

// DecodeResponse method for decoding response on client side
func (t *postOrderClientTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (response models.Response, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}
	err = response.UnmarshalJSON(r.Body())
	return
}

// NewPostOrderClientTransport the transport creator for http requests
func NewPostOrderClientTransport(
	errorProcessor errorProcessor,
	errorCreator   errorCreator,
	pathTemplate   string,
	method         string,
) PostOrderClientTransport {
	return &postOrderClientTransport{
		errorProcessor: errorProcessor,
		errorCreator:   errorCreator,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}
// GetCountClientTransport transport interface
type GetCountClientTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, request *models.GetClientId) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (response models.Response, err error)
}

type getCountClientTransport struct {
	errorProcessor errorProcessor
	errorCreator   errorCreator
	pathTemplate   string
	method         string
}

// EncodeRequest method for encoding requests on client side
func (t *getCountClientTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, request *models.GetClientId) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(fmt.Sprintf(t.pathTemplate, strconv.Itoa(request.Id)))


	//ctx.SetUserValue("id",strconv.Itoa(request.Id))
	//fmt.Println(ctx.UserValue("id").(string))

	//r.URI().QueryArgs().Set("id","12")
	//r.URI().QueryArgs().Set("id", strconv.Itoa(request.Id))
	return
}

// DecodeResponse method for decoding response on client side
func (t *getCountClientTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (response models.Response, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}
	err = response.UnmarshalJSON(r.Body())
	return
}

// NewGetCountClientTransport the transport creator for http requests
func NewGetCountClientTransport(
	errorProcessor errorProcessor,
	errorCreator   errorCreator,
	pathTemplate   string,
	method         string,
) GetCountClientTransport {
	return &getCountClientTransport{
		errorProcessor: errorProcessor,
		errorCreator:   errorCreator,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}
// GetOrderClientTransport transport interface
type GetOrderClientTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request,  ) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (response models.Response, err error)
}

type getOrderClientTransport struct {
	errorProcessor errorProcessor
	errorCreator   errorCreator
	pathTemplate   string
	method         string
}

// EncodeRequest method for encoding requests on client side
func (t *getOrderClientTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request,  ) (err error) {
	r.Header.SetMethod(t.method)
	r.SetRequestURI(t.pathTemplate)
	r.Header.Set("Content-Type", "application/json")
	r.SetBodyStreamWriter(func(w *bufio.Writer) {
		if err = json.NewEncoder(w).Encode(r); err != nil {
			return
		}
	})
	return
}

// DecodeResponse method for decoding response on client side
func (t *getOrderClientTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (response models.Response, err error) {
	if r.StatusCode() != http.StatusOK {
		err = t.errorProcessor.Decode(r)
		return
	}
	err = response.UnmarshalJSON(r.Body())
	return
}

// NewGetOrderClientTransport the transport creator for http requests
func NewGetOrderClientTransport(
	errorProcessor errorProcessor,
	errorCreator   errorCreator,
	pathTemplate   string,
	method         string,
) GetOrderClientTransport {
	return &getOrderClientTransport{
		errorProcessor: errorProcessor,
		errorCreator:   errorCreator,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}
