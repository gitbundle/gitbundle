// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package migrations

import (
	"fmt"

	"xorm.io/xorm"
)

func createUserSettingsTable(x *xorm.Engine) error {
	type UserSetting struct {
		ID           int64  `xorm:"pk autoincr"`
		UserID       int64  `xorm:"index unique(key_userid)"`              // to load all of someone's settings
		SettingKey   string `xorm:"varchar(255) index unique(key_userid)"` // ensure key is always lowercase
		SettingValue string `xorm:"text"`
	}
	if err := x.Sync2(new(UserSetting)); err != nil {
		return fmt.Errorf("sync2: %v", err)
	}
	return nil
}
