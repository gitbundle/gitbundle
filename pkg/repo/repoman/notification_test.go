// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package repoman_test

import (
	"testing"

	"github.com/gitbundle/server/pkg/db"
	issues_model "github.com/gitbundle/server/pkg/issues"
	"github.com/gitbundle/server/pkg/repo/repoman"
	"github.com/gitbundle/server/pkg/unittest"
	user_model "github.com/gitbundle/server/pkg/user"

	"github.com/stretchr/testify/assert"
)

func TestCreateOrUpdateIssueNotifications(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	issue := unittest.AssertExistsAndLoadBean(t, &issues_model.Issue{ID: 1}).(*issues_model.Issue)

	assert.NoError(t, repoman.CreateOrUpdateIssueNotifications(issue.ID, 0, 2, 0))

	// User 9 is inactive, thus notifications for user 1 and 4 are created
	notf := unittest.AssertExistsAndLoadBean(t, &repoman.Notification{UserID: 1, IssueID: issue.ID}).(*repoman.Notification)
	assert.Equal(t, repoman.NotificationStatusUnread, notf.Status)
	unittest.CheckConsistencyFor(t, &issues_model.Issue{ID: issue.ID})

	notf = unittest.AssertExistsAndLoadBean(t, &repoman.Notification{UserID: 4, IssueID: issue.ID}).(*repoman.Notification)
	assert.Equal(t, repoman.NotificationStatusUnread, notf.Status)
}

func TestNotificationsForUser(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	user := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2}).(*user_model.User)
	statuses := []repoman.NotificationStatus{repoman.NotificationStatusRead, repoman.NotificationStatusUnread}
	notfs, err := repoman.NotificationsForUser(db.DefaultContext, user, statuses, 1, 10)
	assert.NoError(t, err)
	if assert.Len(t, notfs, 3) {
		assert.EqualValues(t, 5, notfs[0].ID)
		assert.EqualValues(t, user.ID, notfs[0].UserID)
		assert.EqualValues(t, 4, notfs[1].ID)
		assert.EqualValues(t, user.ID, notfs[1].UserID)
		assert.EqualValues(t, 2, notfs[2].ID)
		assert.EqualValues(t, user.ID, notfs[2].UserID)
	}
}

func TestNotification_GetRepo(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	notf := unittest.AssertExistsAndLoadBean(t, &repoman.Notification{RepoID: 1}).(*repoman.Notification)
	repo, err := notf.GetRepo()
	assert.NoError(t, err)
	assert.Equal(t, repo, notf.Repository)
	assert.EqualValues(t, notf.RepoID, repo.ID)
}

func TestNotification_GetIssue(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	notf := unittest.AssertExistsAndLoadBean(t, &repoman.Notification{RepoID: 1}).(*repoman.Notification)
	issue, err := notf.GetIssue()
	assert.NoError(t, err)
	assert.Equal(t, issue, notf.Issue)
	assert.EqualValues(t, notf.IssueID, issue.ID)
}

func TestGetNotificationCount(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	user := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 1}).(*user_model.User)
	cnt, err := repoman.GetNotificationCount(db.DefaultContext, user, repoman.NotificationStatusRead)
	assert.NoError(t, err)
	assert.EqualValues(t, 0, cnt)

	cnt, err = repoman.GetNotificationCount(db.DefaultContext, user, repoman.NotificationStatusUnread)
	assert.NoError(t, err)
	assert.EqualValues(t, 1, cnt)
}

func TestSetNotificationStatus(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	user := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2}).(*user_model.User)
	notf := unittest.AssertExistsAndLoadBean(t,
		&repoman.Notification{UserID: user.ID, Status: repoman.NotificationStatusRead}).(*repoman.Notification)
	_, err := repoman.SetNotificationStatus(notf.ID, user, repoman.NotificationStatusPinned)
	assert.NoError(t, err)
	unittest.AssertExistsAndLoadBean(t,
		&repoman.Notification{ID: notf.ID, Status: repoman.NotificationStatusPinned})

	_, err = repoman.SetNotificationStatus(1, user, repoman.NotificationStatusRead)
	assert.Error(t, err)
	_, err = repoman.SetNotificationStatus(unittest.NonexistentID, user, repoman.NotificationStatusRead)
	assert.Error(t, err)
}

func TestUpdateNotificationStatuses(t *testing.T) {
	assert.NoError(t, unittest.PrepareTestDatabase())
	user := unittest.AssertExistsAndLoadBean(t, &user_model.User{ID: 2}).(*user_model.User)
	notfUnread := unittest.AssertExistsAndLoadBean(t,
		&repoman.Notification{UserID: user.ID, Status: repoman.NotificationStatusUnread}).(*repoman.Notification)
	notfRead := unittest.AssertExistsAndLoadBean(t,
		&repoman.Notification{UserID: user.ID, Status: repoman.NotificationStatusRead}).(*repoman.Notification)
	notfPinned := unittest.AssertExistsAndLoadBean(t,
		&repoman.Notification{UserID: user.ID, Status: repoman.NotificationStatusPinned}).(*repoman.Notification)
	assert.NoError(t, repoman.UpdateNotificationStatuses(user, repoman.NotificationStatusUnread, repoman.NotificationStatusRead))
	unittest.AssertExistsAndLoadBean(t,
		&repoman.Notification{ID: notfUnread.ID, Status: repoman.NotificationStatusRead})
	unittest.AssertExistsAndLoadBean(t,
		&repoman.Notification{ID: notfRead.ID, Status: repoman.NotificationStatusRead})
	unittest.AssertExistsAndLoadBean(t,
		&repoman.Notification{ID: notfPinned.ID, Status: repoman.NotificationStatusPinned})
}
