// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package swagger

import api "github.com/gitbundle/api/pkg/structs"

// GeneralRepoSettings
// swagger:response GeneralRepoSettings
type swaggerResponseGeneralRepoSettings struct {
	// in:body
	Body api.GeneralRepoSettings `json:"body"`
}

// GeneralUISettings
// swagger:response GeneralUISettings
type swaggerResponseGeneralUISettings struct {
	// in:body
	Body api.GeneralUISettings `json:"body"`
}

// GeneralAPISettings
// swagger:response GeneralAPISettings
type swaggerResponseGeneralAPISettings struct {
	// in:body
	Body api.GeneralAPISettings `json:"body"`
}

// GeneralAttachmentSettings
// swagger:response GeneralAttachmentSettings
type swaggerResponseGeneralAttachmentSettings struct {
	// in:body
	Body api.GeneralAttachmentSettings `json:"body"`
}
