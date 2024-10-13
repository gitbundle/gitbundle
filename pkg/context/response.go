// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package context

import (
	"bufio"
	"errors"
	"net"
	"net/http"
)

// ResponseWriter represents a response writer for HTTP
type ResponseWriter interface {
	http.ResponseWriter
	http.Hijacker
	Flush()
	Status() int
	Before(func(ResponseWriter))
	Size() int
}

var _ ResponseWriter = &Response{}

// Response represents a response
type Response struct {
	http.ResponseWriter
	written        int
	status         int
	befores        []func(ResponseWriter)
	beforeExecuted bool
}

// Size return written size
func (r *Response) Size() int {
	return r.written
}

// Write writes bytes to HTTP endpoint
func (r *Response) Write(bs []byte) (int, error) {
	if !r.beforeExecuted {
		for _, before := range r.befores {
			before(r)
		}
		r.beforeExecuted = true
	}
	size, err := r.ResponseWriter.Write(bs)
	r.written += size
	if err != nil {
		return size, err
	}
	if r.status == 0 {
		r.status = http.StatusOK
	}
	return size, nil
}

// WriteHeader write status code
func (r *Response) WriteHeader(statusCode int) {
	if !r.beforeExecuted {
		for _, before := range r.befores {
			before(r)
		}
		r.beforeExecuted = true
	}
	if r.status == 0 {
		r.status = statusCode
		r.ResponseWriter.WriteHeader(statusCode)
	}
}

// Flush flush cached data
func (r *Response) Flush() {
	if f, ok := r.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}
}

// Status returned status code written
func (r *Response) Status() int {
	return r.status
}

func (r *Response) Hijack() (rwc net.Conn, buf *bufio.ReadWriter, err error) {
	h, ok := r.ResponseWriter.(http.Hijacker)
	if ok {
		return h.Hijack()
	}

	return nil, nil, errors.New("services/context: Hijack method does not supported")
}

// Before allows for a function to be called before the ResponseWriter has been written to. This is
// useful for setting headers or any other operations that must happen before a response has been written.
func (r *Response) Before(f func(ResponseWriter)) {
	r.befores = append(r.befores, f)
}

// NewResponse creates a response
func NewResponse(resp http.ResponseWriter) *Response {
	if v, ok := resp.(*Response); ok {
		return v
	}
	return &Response{
		ResponseWriter: resp,
		status:         0,
		befores:        make([]func(ResponseWriter), 0),
	}
}
