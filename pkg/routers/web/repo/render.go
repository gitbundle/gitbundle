// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package repo

import (
	"bytes"
	"io"
	"net/http"
	"path"

	"github.com/gitbundle/modules/charset"
	"github.com/gitbundle/modules/git"
	"github.com/gitbundle/modules/markup"
	"github.com/gitbundle/modules/typesniffer"
	"github.com/gitbundle/modules/util"
	"github.com/gitbundle/server/pkg/context"
)

// RenderFile renders a file by repos path
func RenderFile(ctx *context.Context) {
	blob, err := ctx.Repo.Commit.GetBlobByPath(ctx.Repo.TreePath)
	if err != nil {
		if git.IsErrNotExist(err) {
			ctx.NotFound("GetBlobByPath", err)
		} else {
			ctx.ServerError("GetBlobByPath", err)
		}
		return
	}

	dataRc, err := blob.DataAsync()
	if err != nil {
		ctx.ServerError("DataAsync", err)
		return
	}
	defer dataRc.Close()

	buf := make([]byte, 1024)
	n, _ := util.ReadAtMost(dataRc, buf)
	buf = buf[:n]

	st := typesniffer.DetectContentType(buf)
	isTextFile := st.IsText()

	rd := charset.ToUTF8WithFallbackReader(io.MultiReader(bytes.NewReader(buf), dataRc))

	if markupType := markup.Type(blob.Name()); markupType == "" {
		if isTextFile {
			_, err = io.Copy(ctx.Resp, rd)
			if err != nil {
				ctx.ServerError("Copy", err)
			}
			return
		}
		ctx.Error(http.StatusInternalServerError, "Unsupported file type render")
		return
	}

	treeLink := ctx.Repo.RepoLink + "/src/" + ctx.Repo.BranchNameSubURL()
	if ctx.Repo.TreePath != "" {
		treeLink += "/" + util.PathEscapeSegments(ctx.Repo.TreePath)
	}

	ctx.Resp.Header().Add("Content-Security-Policy", "frame-src 'self'; sandbox allow-scripts")
	err = markup.Render(&markup.RenderContext{
		Ctx:              ctx,
		RelativePath:     ctx.Repo.TreePath,
		URLPrefix:        path.Dir(treeLink),
		Metas:            ctx.Repo.Repository.ComposeDocumentMetas(),
		GitRepo:          ctx.Repo.GitRepo,
		InStandalonePage: true,
	}, rd, ctx.Resp)
	if err != nil {
		ctx.ServerError("Render", err)
		return
	}
}
