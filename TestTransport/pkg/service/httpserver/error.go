package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/valyala/fasthttp"
)

type httpError struct {
	Code    int
	Message string
}

// Error returns a text message corresponding to the given error.
func (e *httpError) Error() string {
	return e.Message
}

// StatusCode returns an HTTP status code corresponding to the given error.
func (e *httpError) StatusCode() int {
	return e.Code
}

// ErrorProcessor ...
type ErrorProcessor interface {
	Encode(ctx context.Context, r *fasthttp.Response, err error)
	Decode(r *fasthttp.Response) error
}

type errProcessor struct {
	defaultCode    int
	defaultMessage string
}

//Encode writes a svc error to the given http.ResponseWriter.
func (e *errProcessor) Encode(ctx context.Context, r *fasthttp.Response, err error) {
	code := e.defaultCode
	message := e.defaultMessage
	if err, ok := err.(*httpError); ok {
		if err.Code != e.defaultCode {
			code = err.Code
			message = err.Message
		}
	}
	r.SetStatusCode(code)
	r.SetBodyString(message)
	return
}

// Decode reads a Service error from the given *http.Response.
func (e *errProcessor) Decode(r *fasthttp.Response) error {
	msgBytes := r.Body()
	msg := strings.TrimSpace(string(msgBytes))
	if msg == "" {
		msg = http.StatusText(r.StatusCode())
	}
	return &httpError{
		Code:    r.StatusCode(),
		Message: msg,
	}
}

// NewErrorProcessor ...
func NewErrorProcessor(defaultCode int, defaultMessage string) ErrorProcessor {
	return &errProcessor{
		defaultCode:    defaultCode,
		defaultMessage: defaultMessage,
	}
}

// NewError ...
func NewError(status int, format string, v ...interface{}) error {
	return &httpError{
		Code:    status,
		Message: fmt.Sprintf(format, v...),
	}
}