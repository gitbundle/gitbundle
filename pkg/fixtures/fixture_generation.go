// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package fixtures

import (
	"fmt"
	"strings"

	"github.com/gitbundle/server/pkg/db"
	access_model "github.com/gitbundle/server/pkg/perm/access"
	repo_model "github.com/gitbundle/server/pkg/repo"
)

// GetYamlFixturesAccess returns a string containing the contents
// for the access table, as recalculated using repo.RecalculateAccesses()
func GetYamlFixturesAccess() (string, error) {
	repos := make([]*repo_model.Repository, 0, 50)
	if err := db.GetEngine(db.DefaultContext).Find(&repos); err != nil {
		return "", err
	}

	for _, repo := range repos {
		repo.MustOwner()
		if err := access_model.RecalculateAccesses(db.DefaultContext, repo); err != nil {
			return "", err
		}
	}

	var b strings.Builder

	accesses := make([]*access_model.Access, 0, 200)
	if err := db.GetEngine(db.DefaultContext).OrderBy("user_id, repo_id").Find(&accesses); err != nil {
		return "", err
	}

	for i, a := range accesses {
		fmt.Fprintf(&b, "-\n")
		fmt.Fprintf(&b, "  id: %d\n", i+1)
		fmt.Fprintf(&b, "  user_id: %d\n", a.UserID)
		fmt.Fprintf(&b, "  repo_id: %d\n", a.RepoID)
		fmt.Fprintf(&b, "  mode: %d\n", a.Mode)
		fmt.Fprintf(&b, "\n")
	}

	return b.String(), nil
}
