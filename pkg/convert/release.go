// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package convert

import (
	api "github.com/gitbundle/api/pkg/structs"
	repo_model "github.com/gitbundle/server/pkg/repo"
	release_model "github.com/gitbundle/server/pkg/repo/release"
)

// ToRelease convert a release_model.Release to api.Release
func ToRelease(r *release_model.Release) *api.Release {
	assets := make([]*api.Attachment, 0)
	for _, att := range r.Attachments {
		assets = append(assets, ToReleaseAttachment(att))
	}
	return &api.Release{
		ID:           r.ID,
		TagName:      r.TagName,
		Target:       r.Target,
		Title:        r.Title,
		Note:         r.Note,
		URL:          r.APIURL(),
		HTMLURL:      r.HTMLURL(),
		TarURL:       r.TarURL(),
		ZipURL:       r.ZipURL(),
		IsDraft:      r.IsDraft,
		IsPrerelease: r.IsPrerelease,
		CreatedAt:    r.CreatedUnix.AsTime(),
		PublishedAt:  r.CreatedUnix.AsTime(),
		Publisher:    ToUser(r.Publisher, nil),
		Attachments:  assets,
	}
}

// ToReleaseAttachment converts repo_model.Attachment to api.Attachment
func ToReleaseAttachment(a *repo_model.Attachment) *api.Attachment {
	return &api.Attachment{
		ID:            a.ID,
		Name:          a.Name,
		Created:       a.CreatedUnix.AsTime(),
		DownloadCount: a.DownloadCount,
		Size:          a.Size,
		UUID:          a.UUID,
		DownloadURL:   a.DownloadURL(),
	}
}
