// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package repo

import (
	"net/http"
	"strings"

	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/server/pkg/context"
	repo_model "github.com/gitbundle/server/pkg/repo"
)

// TopicsPost response for creating repository
func TopicsPost(ctx *context.Context) {
	if ctx.Doer == nil {
		ctx.JSON(http.StatusForbidden, map[string]interface{}{
			"message": "Only owners could change the topics.",
		})
		return
	}

	topics := make([]string, 0)
	topicsStr := ctx.FormTrim("topics")
	if len(topicsStr) > 0 {
		topics = strings.Split(topicsStr, ",")
	}

	validTopics, invalidTopics := repo_model.SanitizeAndValidateTopics(topics)

	if len(validTopics) > 25 {
		ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"invalidTopics": nil,
			"message":       ctx.Tr("repo.topic.count_prompt"),
		})
		return
	}

	if len(invalidTopics) > 0 {
		ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"invalidTopics": invalidTopics,
			"message":       ctx.Tr("repo.topic.format_prompt"),
		})
		return
	}

	err := repo_model.SaveTopics(ctx.Repo.Repository.ID, validTopics...)
	if err != nil {
		log.Error("SaveTopics failed: %v", err)
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Save topics failed.",
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
