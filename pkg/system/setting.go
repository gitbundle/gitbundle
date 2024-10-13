// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package system

import (
	"github.com/gitbundle/modules/timeutil"
	"github.com/gitbundle/server/pkg/db"
)

type SystemSettings struct {
	ID    int64  `xorm:"pk autoincr" json:"-"`
	Key   string `xorm:"UNIQUE NOT NULL" json:"key"`
	Value string `xorm:"NOT NULL" json:"value"`

	CreatedUnix timeutil.TimeStamp `xorm:"created" json:"createdUnix"`
	UpdatedUnix timeutil.TimeStamp `xorm:"updated" json:"updatedUnix"`
}

func init() {
	db.RegisterModel(new(SystemSettings))
}

func (m *SystemSettings) CreateOrUpdateValueByKey() error {
	exists, err := db.GetEngine(db.DefaultContext).Where("key=?", m.Key).Exist(&SystemSettings{})
	if err != nil {
		return err
	}
	if exists {
		_, err = db.GetEngine(db.DefaultContext).Where("key=?", m.Key).Cols("value").Update(m)
	} else {
		_, err = db.GetEngine(db.DefaultContext).Insert(m)
	}
	return err
}

func GetSystemSettingsByKeys(keys ...string) (map[string]bool, error) {
	opts := SystemSettingsOptions{Keys: keys}
	return opts.BoolLists()
}

type SystemSettingsOptions struct {
	Keys []string
}

func (opts *SystemSettingsOptions) Lists() (map[string]string, error) {
	sess := db.GetEngine(db.DefaultContext)
	var (
		err error
		ss  = make([]*SystemSettings, 0, 100)
	)
	if len(opts.Keys) == 0 {
		err = sess.Find(&ss)
	} else {
		err = sess.In("key", opts.Keys).Find(&ss)
	}

	if err != nil {
		return nil, err
	}

	result := make(map[string]string, len(ss))
	for _, s := range ss {
		result[s.Key] = s.Value
	}
	return result, nil
}

func (opts *SystemSettingsOptions) BoolLists() (map[string]bool, error) {
	res, err := opts.Lists()
	if err != nil {
		return nil, err
	}

	result := make(map[string]bool)
	for _, key := range opts.Keys {
		var value bool
		val, ok := res[key]
		if ok {
			value = val == "yes"
		}
		result[key] = value
	}
	return result, nil
}

func (opts *SystemSettingsOptions) Delete() error {
	if len(opts.Keys) == 0 {
		return nil
	}
	_, err := db.GetEngine(db.DefaultContext).In("key", opts.Keys).Delete(&SystemSettings{})
	return err
}
