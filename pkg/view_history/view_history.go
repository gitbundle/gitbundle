// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package view_history

import (
	"context"
	"encoding/json"

	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/timeutil"
	"github.com/gitbundle/server/pkg/db"
)

type VisitType string

const (
	UserType  VisitType = "user"
	RepoType            = "repo"
	SavedType           = "saved"
)

type ViewHistory struct {
	ID         int64     `xorm:"pk autoincr" json:"-"`
	DoerID     int64     `xorm:"INDEX(uvt) NOT NULL" json:"-"`
	UserID     int64     `xorm:"INDEX(uvt) NOT NULL" json:"-"`
	RepoID     int64     `xorm:"INDEX(uvt) NOT NULL" json:"-"`
	VisitType  VisitType `xorm:"INDEX(uvt) NOT NULL" json:"visitType"`
	VisitCount int64     `xorm:"INDEX(uvt) NOT NULL DEFAULT 0" json:"visitCount"`
	Content    []byte    `xorm:"NOT NULL" json:"-"`
	Location   string

	CreatedUnix timeutil.TimeStamp `xorm:"INDEX created" json:"created"`
	UpdatedUnix timeutil.TimeStamp `xorm:"INDEX updated" json:"updated"`

	Object map[string]interface{} `xorm:"-" json:"content"`
}

func init() {
	db.RegisterModel(new(ViewHistory))
}

func (m *ViewHistory) AfterLoad() {
	err := json.Unmarshal(m.Content, &m.Object)
	if err != nil {
		log.Error("json.Unmarshal, error: %v", err)
	}
}

// TODO: if the user or repo change the name, what should we do
func (m *ViewHistory) CreateOrUpdate(ctx context.Context) (err error) {
	vh := ViewHistory{}
	exists, err := db.GetEngine(ctx).
		Where("doer_id=? and user_id=? and repo_id=?", m.DoerID, m.UserID, m.RepoID).
		Exist(&vh)
	if exists {
		_, err = db.GetEngine(ctx).
			Where("doer_id=? and user_id=? and repo_id=?", m.DoerID, m.UserID, m.RepoID).
			Update(&ViewHistory{VisitCount: vh.VisitCount + 1, Content: m.Content})
		if err != nil {
			log.Error("Updated view history, error: %v", err)
		}
	} else {
		_, err = db.GetEngine(ctx).Insert(m)
		if err != nil {
			log.Error("Inserted view history, error: %v", err)
		}
	}
	return err
}

func GetAllHistories(ctx context.Context, doerID int64) ([]*ViewHistory, error) {
	results := make([]*ViewHistory, 0, 10)
	err := db.GetEngine(ctx).Where("doer_id=?", doerID).OrderBy("updated_unix DESC").Find(&results)
	return results, err
}

func GetHistoriesByViewType(ctx context.Context, doerID int64, visitType VisitType, limit int) ([]*ViewHistory, error) {
	results := make([]*ViewHistory, 0, limit)
	err := db.GetEngine(ctx).Where("doer_id=? and visit_type=?", doerID, visitType).OrderBy("updated_unix DESC").Limit(limit).Find(&results)
	return results, err
}
