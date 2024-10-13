// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package common

import (
	"fmt"
	"io"
	"net/url"
	"path"
	"path/filepath"
	"strings"
	"time"

	charsetModule "github.com/gitbundle/modules/charset"
	"github.com/gitbundle/modules/git"
	"github.com/gitbundle/modules/httpcache"
	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/modules/typesniffer"
	"github.com/gitbundle/modules/util"
	"github.com/gitbundle/server/pkg/context"
)

// ServeBlob download a git.Blob
func ServeBlob(ctx *context.Context, blob *git.Blob, lastModified time.Time) error {
	if httpcache.HandleGenericETagTimeCache(ctx.Req, ctx.Resp, `"`+blob.ID.String()+`"`, lastModified) {
		return nil
	}

	dataRc, err := blob.DataAsync()
	if err != nil {
		return err
	}
	defer func() {
		if err = dataRc.Close(); err != nil {
			log.Error("ServeBlob: Close: %v", err)
		}
	}()

	return ServeData(ctx, ctx.Repo.TreePath, blob.Size(), dataRc)
}

// ServeData download file from io.Reader
func ServeData(ctx *context.Context, filePath string, size int64, reader io.Reader) error {
	buf := make([]byte, 1024)
	n, err := util.ReadAtMost(reader, buf)
	if err != nil {
		return err
	}
	if n >= 0 {
		buf = buf[:n]
	}

	httpcache.AddCacheControlToHeader(ctx.Resp.Header(), 5*time.Minute)

	if size >= 0 {
		ctx.Resp.Header().Set("Content-Length", fmt.Sprintf("%d", size))
	} else {
		log.Error("ServeData called to serve data: %s with size < 0: %d", filePath, size)
	}

	fileName := path.Base(filePath)
	sniffedType := typesniffer.DetectContentType(buf)
	isPlain := sniffedType.IsText() || ctx.FormBool("render")
	mimeType := ""
	charset := ""

	if setting.MimeTypeMap.Enabled {
		fileExtension := strings.ToLower(filepath.Ext(fileName))
		mimeType = setting.MimeTypeMap.Map[fileExtension]
	}

	if mimeType == "" {
		if sniffedType.IsBrowsableBinaryType() {
			mimeType = sniffedType.GetMimeType()
		} else if isPlain {
			mimeType = "text/plain"
		} else {
			mimeType = typesniffer.ApplicationOctetStream
		}
	}

	if isPlain {
		charset, err = charsetModule.DetectEncoding(buf)
		if err != nil {
			log.Error("Detect raw file %s charset failed: %v, using by default utf-8", filePath, err)
			charset = "utf-8"
		}
	}

	if charset != "" {
		ctx.Resp.Header().Set("Content-Type", mimeType+"; charset="+strings.ToLower(charset))
	} else {
		ctx.Resp.Header().Set("Content-Type", mimeType)
	}
	ctx.Resp.Header().Set("X-Content-Type-Options", "nosniff")

	isSVG := sniffedType.IsSvgImage()

	// serve types that can present a security risk with CSP
	if isSVG {
		ctx.Resp.Header().Set("Content-Security-Policy", "default-src 'none'; style-src 'unsafe-inline'; sandbox")
	} else if sniffedType.IsPDF() {
		// no sandbox attribute for pdf as it breaks rendering in at least safari. this
		// should generally be safe as scripts inside PDF can not escape the PDF document
		// see https://bugs.chromium.org/p/chromium/issues/detail?id=413851 for more discussion
		ctx.Resp.Header().Set("Content-Security-Policy", "default-src 'none'; style-src 'unsafe-inline'")
	}

	disposition := "inline"
	if isSVG && !setting.UI.SVG.Enabled {
		disposition = "attachment"
	}

	// encode filename per https://datatracker.ietf.org/doc/html/rfc5987
	encodedFileName := `filename*=UTF-8''` + url.PathEscape(fileName)

	ctx.Resp.Header().Set("Content-Disposition", disposition+"; "+encodedFileName)
	ctx.Resp.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")

	_, err = ctx.Resp.Write(buf)
	if err != nil {
		return err
	}
	_, err = io.Copy(ctx.Resp, reader)
	return err
}
