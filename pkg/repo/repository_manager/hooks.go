// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package repository

import (
	"context"
	"fmt"

	"github.com/gitbundle/modules/git"
	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/server/pkg/db"
	repo_model "github.com/gitbundle/server/pkg/repo"
	repo_module "github.com/gitbundle/server/pkg/repo/repository"
	"github.com/gitbundle/server/pkg/webhook"

	"xorm.io/builder"
)

// SyncRepositoryHooks rewrites all repositories' pre-receive, update and post-receive hooks
// to make sure the binary and custom conf path are up-to-date.
func SyncRepositoryHooks(ctx context.Context) error {
	log.Trace("Doing: SyncRepositoryHooks")

	if err := db.Iterate(
		ctx,
		new(repo_model.Repository),
		builder.Gt{"id": 0},
		func(idx int, bean interface{}) error {
			repo := bean.(*repo_model.Repository)
			select {
			case <-ctx.Done():
				return db.ErrCancelledf("before sync repository hooks for %s", repo.FullName())
			default:
			}

			if err := repo_module.CreateDelegateHooks(repo.RepoPath()); err != nil {
				return fmt.Errorf("SyncRepositoryHook: %v", err)
			}
			if repo.HasWiki() {
				if err := repo_module.CreateDelegateHooks(repo.WikiPath()); err != nil {
					return fmt.Errorf("SyncRepositoryHook: %v", err)
				}
			}
			return nil
		},
	); err != nil {
		return err
	}

	log.Trace("Finished: SyncRepositoryHooks")
	return nil
}

// GenerateGitHooks generates git hooks from a template repository
func GenerateGitHooks(ctx context.Context, templateRepo, generateRepo *repo_model.Repository) error {
	generateGitRepo, err := git.OpenRepository(ctx, generateRepo.RepoPath())
	if err != nil {
		return err
	}
	defer generateGitRepo.Close()

	templateGitRepo, err := git.OpenRepository(ctx, templateRepo.RepoPath())
	if err != nil {
		return err
	}
	defer templateGitRepo.Close()

	templateHooks, err := templateGitRepo.Hooks()
	if err != nil {
		return err
	}

	for _, templateHook := range templateHooks {
		generateHook, err := generateGitRepo.GetHook(templateHook.Name())
		if err != nil {
			return err
		}

		generateHook.Content = templateHook.Content
		if err := generateHook.Update(); err != nil {
			return err
		}
	}
	return nil
}

// GenerateWebhooks generates webhooks from a template repository
func GenerateWebhooks(ctx context.Context, templateRepo, generateRepo *repo_model.Repository) error {
	templateWebhooks, err := webhook.ListWebhooksByOpts(ctx, &webhook.ListWebhookOptions{RepoID: templateRepo.ID})
	if err != nil {
		return err
	}

	ws := make([]*webhook.Webhook, 0, len(templateWebhooks))
	for _, templateWebhook := range templateWebhooks {
		ws = append(ws, &webhook.Webhook{
			RepoID:      generateRepo.ID,
			URL:         templateWebhook.URL,
			HTTPMethod:  templateWebhook.HTTPMethod,
			ContentType: templateWebhook.ContentType,
			Secret:      templateWebhook.Secret,
			HookEvent:   templateWebhook.HookEvent,
			IsActive:    templateWebhook.IsActive,
			Type:        templateWebhook.Type,
			OrgID:       templateWebhook.OrgID,
			Events:      templateWebhook.Events,
			Meta:        templateWebhook.Meta,
		})
	}
	return webhook.CreateWebhooks(ctx, ws)
}
