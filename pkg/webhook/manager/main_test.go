// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package webhook_test

import (
	"net/http"
	"net/url"
	"path/filepath"
	"testing"
	_ "unsafe"

	api "github.com/gitbundle/api/pkg/structs"
	"github.com/gitbundle/modules/setting"
	"github.com/gitbundle/server/pkg/unittest"
	webhook_model "github.com/gitbundle/server/pkg/webhook"
	webhook "github.com/gitbundle/server/pkg/webhook/manager"

	_ "github.com/gitbundle/server/pkg/access_token"
	_ "github.com/gitbundle/server/pkg/action"
	_ "github.com/gitbundle/server/pkg/perm/access"
	_ "github.com/gitbundle/server/pkg/repo/repoman"
)

func TestMain(m *testing.M) {
	setting.LoadForTest()
	setting.NewQueueService()
	unittest.MainTest(m, &unittest.TestOptions{
		GiteaRootPath: filepath.Join("..", ".."),
		SetUp:         webhook.Init,
	})
}

var (
	//go:linkname greenColor github.com/gitbundle/server/pkg/webhook/manager.greenColor
	greenColor int
	//go:linkname greenColorLight github.com/gitbundle/server/pkg/webhook/manager.greenColorLight
	greenColorLight int
	//go:linkname yellowColor github.com/gitbundle/server/pkg/webhook/manager.yellowColor
	yellowColor int
	//go:linkname greyColor github.com/gitbundle/server/pkg/webhook/manager.greyColor
	greyColor int
	//go:linkname purpleColor github.com/gitbundle/server/pkg/webhook/manager.purpleColor
	purpleColor int
	//go:linkname orangeColor github.com/gitbundle/server/pkg/webhook/manager.orangeColor
	orangeColor int
	//go:linkname orangeColorLight github.com/gitbundle/server/pkg/webhook/manager.orangeColorLight
	orangeColorLight int
	//go:linkname redColor github.com/gitbundle/server/pkg/webhook/manager.redColor
	redColor int
)

//go:linkname noneLinkFormatter github.com/gitbundle/server/pkg/webhook/manager.noneLinkFormatter
func noneLinkFormatter(url, text string) string

//go:linkname getIssuesPayloadInfo github.com/gitbundle/server/pkg/webhook/manager.getIssuesPayloadInfo
func getIssuesPayloadInfo(p *api.IssuePayload, linkFormatter linkFormatter, withSender bool) (string, string, string, int)

// //go:linkname linkFormatter github.com/gitbundle/server/pkg/webhook/manager.linkFormatter
type linkFormatter = func(string, string) string

//go:linkname getIssueCommentPayloadInfo github.com/gitbundle/server/pkg/webhook/manager.getIssueCommentPayloadInfo
func getIssueCommentPayloadInfo(p *api.IssueCommentPayload, linkFormatter linkFormatter, withSender bool) (string, string, int)

//go:linkname getPullRequestPayloadInfo github.com/gitbundle/server/pkg/webhook/manager.getPullRequestPayloadInfo
func getPullRequestPayloadInfo(p *api.PullRequestPayload, linkFormatter linkFormatter, withSender bool) (string, string, string, int)

//go:linkname getReleasePayloadInfo github.com/gitbundle/server/pkg/webhook/manager.getReleasePayloadInfo
func getReleasePayloadInfo(p *api.ReleasePayload, linkFormatter linkFormatter, withSender bool) (text string, color int)

//go:linkname webhookProxy github.com/gitbundle/server/pkg/webhook/manager.webhookProxy
func webhookProxy() func(req *http.Request) (*url.URL, error)

//go:linkname getMatrixTxnID github.com/gitbundle/server/pkg/webhook/manager.getMatrixTxnID
func getMatrixTxnID(payload []byte) (string, error)

//go:linkname getMatrixHookRequest github.com/gitbundle/server/pkg/webhook/manager.getMatrixHookRequest
func getMatrixHookRequest(w *webhook_model.Webhook, t *webhook_model.HookTask) (*http.Request, error)
