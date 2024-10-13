// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package routing

import (
	"net/http"
	"sync"
	"time"
)

type requestRecord struct {
	// index of the record in the records map
	index uint64

	// immutable fields
	startTime      time.Time
	request        *http.Request
	responseWriter http.ResponseWriter

	// mutex
	lock sync.RWMutex

	// mutable fields
	isLongPolling bool
	funcInfo      *FuncInfo
	panicError    interface{}
}
