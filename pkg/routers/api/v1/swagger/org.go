// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package swagger

import (
	api "github.com/gitbundle/api/pkg/structs"
)

// Organization
// swagger:response Organization
type swaggerResponseOrganization struct {
	// in:body
	Body api.Organization `json:"body"`
}

// OrganizationList
// swagger:response OrganizationList
type swaggerResponseOrganizationList struct {
	// in:body
	Body []api.Organization `json:"body"`
}

// Team
// swagger:response Team
type swaggerResponseTeam struct {
	// in:body
	Body api.Team `json:"body"`
}

// TeamList
// swagger:response TeamList
type swaggerResponseTeamList struct {
	// in:body
	Body []api.Team `json:"body"`
}

// OrganizationPermissions
// swagger:response OrganizationPermissions
type swaggerResponseOrganizationPermissions struct {
	// in:body
	Body api.OrganizationPermissions `json:"body"`
}
