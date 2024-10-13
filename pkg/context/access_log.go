// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package context

import (
	"bytes"
	"context"
	"html/template"
	"net/http"
	"time"

	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/setting"
)

type routerLoggerOptions struct {
	req            *http.Request
	Identity       *string
	Start          *time.Time
	ResponseWriter http.ResponseWriter
	Ctx            map[string]interface{}
}

const SignedUserNameStringPointerKey = "signedUserNameStringPointerKey"

// AccessLogger returns a middleware to log access logger
func AccessLogger() func(http.Handler) http.Handler {
	logger := log.GetLogger("access")
	logTemplate, _ := template.New("log").Parse(setting.AccessLogTemplate)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			start := time.Now()
			identity := "-"
			r := req.WithContext(context.WithValue(req.Context(), SignedUserNameStringPointerKey, &identity))

			next.ServeHTTP(w, r)
			rw := w.(ResponseWriter)

			buf := bytes.NewBuffer([]byte{})
			err := logTemplate.Execute(buf, routerLoggerOptions{
				req:            req,
				Identity:       &identity,
				Start:          &start,
				ResponseWriter: rw,
				Ctx: map[string]interface{}{
					"RemoteAddr": req.RemoteAddr,
					"Req":        req,
				},
			})
			if err != nil {
				log.Error("Could not set up chi access logger: %v", err.Error())
			}

			err = logger.SendLog(log.INFO, "", "", 0, buf.String(), "")
			if err != nil {
				log.Error("Could not set up chi access logger: %v", err.Error())
			}
		})
	}
}
